package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/app/mall/internal/conf"
	service2 "imperialPalaceMall/app/mall/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, categoryService *service2.CategoryServiceService, goodsService *service2.GoodsServiceService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
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
	pb.RegisterCategoryServiceHTTPServer(srv, categoryService)
	pb.RegisterGoodsServiceHTTPServer(srv, goodsService)
	return srv
}
