package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/shop/internal/conf"
	"imperialPalaceMall/app/shop/internal/service"
)

// NewHTTPServer new a http server.
func NewHTTPServer(c *conf.Server, shopInterface *service.ShopInterface, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			validate.Validator(),
			recovery.Recovery(),
		),
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
