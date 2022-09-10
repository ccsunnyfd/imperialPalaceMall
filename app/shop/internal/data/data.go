package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	cartV1 "imperialPalaceMall/api/cart/v1"
	mallV1 "imperialPalaceMall/api/mall/v1"
	userV1 "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/shop/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserClient, NewCategoryClient, NewGoodsClient, NewCartClient, NewCategoryRepo, NewGoodsRepo, NewUserRepo, NewCartRepo, NewAddressRepo)

// Data .
type Data struct {
	log  *log.Helper
	uc   userV1.UserClient
	catc mallV1.CategoryServiceClient
	gc   mallV1.GoodsServiceClient
	cc   cartV1.CartServiceClient
}

// NewData .
func NewData(logger log.Logger, uc userV1.UserClient, catc mallV1.CategoryServiceClient, gc mallV1.GoodsServiceClient, cc cartV1.CartServiceClient) (*Data, error) {
	logH := log.NewHelper(log.With(logger, "module", "shop/interface/data"))

	return &Data{
		log:  logH,
		uc:   uc,
		catc: catc,
		gc:   gc,
		cc:   cc,
	}, nil
}

func NewUserClient(conf *conf.Data, logger log.Logger) userV1.UserClient {
	logH := log.NewHelper(log.With(logger, "module", "shop/interface/data/userClient"))

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Service.User.Addr),
		grpc.WithTimeout(conf.Service.User.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logH.Fatalf("failed opening connection to user service: %v", err)
	}
	return userV1.NewUserClient(conn)
}

func NewCategoryClient(conf *conf.Data, logger log.Logger) mallV1.CategoryServiceClient {
	logH := log.NewHelper(log.With(logger, "module", "shop/interface/data/categoryClient"))

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Service.Mall.Addr),
		grpc.WithTimeout(conf.Service.Mall.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logH.Fatalf("failed opening connection to category service: %v", err)
	}
	return mallV1.NewCategoryServiceClient(conn)
}

func NewGoodsClient(conf *conf.Data, logger log.Logger) mallV1.GoodsServiceClient {
	logH := log.NewHelper(log.With(logger, "module", "shop/interface/data/goodsClient"))

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Service.Mall.Addr),
		grpc.WithTimeout(conf.Service.Mall.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logH.Fatalf("failed opening connection to goods service: %v", err)
	}
	return mallV1.NewGoodsServiceClient(conn)
}

func NewCartClient(conf *conf.Data, logger log.Logger) cartV1.CartServiceClient {
	logH := log.NewHelper(log.With(logger, "module", "shop/interface/data/cartClient"))

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Service.Cart.Addr),
		grpc.WithTimeout(conf.Service.Cart.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logH.Fatalf("failed opening connection to cart service: %v", err)
	}
	return cartV1.NewCartServiceClient(conn)
}
