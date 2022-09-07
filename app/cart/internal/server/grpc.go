package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "imperialPalaceMall/api/cart/v1"
	"imperialPalaceMall/app/cart/internal/conf"
	"imperialPalaceMall/app/cart/internal/service"
	"imperialPalaceMall/app/pkg/middleware"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, cartService *service.CartService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			middleware.LogServer(logger),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterCartServiceServer(srv, cartService)
	return srv
}
