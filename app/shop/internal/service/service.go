package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	pb.UnimplementedShopInterfaceServer
	uc   *biz.UserUsecase
	catc *biz.CategoryUsecase
	gc   *biz.GoodsUsecase
	cc   *biz.CartUsecase

	log *log.Helper
}

func NewShopInterface(uc *biz.UserUsecase, catc *biz.CategoryUsecase, gc *biz.GoodsUsecase, cc *biz.CartUsecase, logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log:  log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:   uc,
		catc: catc,
		gc:   gc,
		cc:   cc,
	}
}
