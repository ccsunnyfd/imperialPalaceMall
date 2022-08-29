package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/pkg/middleware"
	"imperialPalaceMall/app/shop/internal/conf"
	"imperialPalaceMall/app/shop/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	// protobuf文件中 包名.服务名/方法名
	whiteList["/api.shop.interface.v1.ShopInterface/WxLogin"] = struct{}{}
	whiteList["/api.shop.interface.v1.ShopInterface/ListCategory"] = struct{}{}
	whiteList["/api.shop.interface.v1.ShopInterface/GetGoods"] = struct{}{}
	whiteList["/api.shop.interface.v1.ShopInterface/ListGoods"] = struct{}{}
	whiteList["/api.shop.interface.v1.ShopInterface/GetSKUs"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewTokenChecker(shopInterface *service.ShopInterface) middleware.CheckToken {
	return func(ctx context.Context, token string) (*middleware.User, error) {
		user, err := shopInterface.CheckToken(ctx, token)
		if err != nil {
			return nil, err
		}
		return &middleware.User{
			UserId: user.UserId,
		}, nil
	}
}

// NewHTTPServer new a http server.
func NewHTTPServer(c *conf.Server, shopInterface *service.ShopInterface, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(
				middleware.JWTAuth(NewTokenChecker(shopInterface)),
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterShopInterfaceHTTPServer(srv, shopInterface)
	return srv
}
