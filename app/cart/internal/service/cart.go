package service

import (
	"context"

	pb "imperialPalaceMall/api/cart/v1"
	"imperialPalaceMall/app/cart/internal/biz"
)

// CartService is a cart service.
type CartService struct {
	pb.UnimplementedCartServiceServer
	uc *biz.CartUsecase
}

// NewCartService new a cart service.
func NewCartService(uc *biz.CartUsecase) *CartService {
	return &CartService{uc: uc}
}

func (s *CartService) AddCart(ctx context.Context, in *pb.AddCartRequest) (*pb.AddCartReply, error) {
	c, err := s.uc.AddCart(ctx, &biz.Cart{
		UserId:       in.UserId.Value,
		GoodsId:      in.GoodsId.Value,
		GoodsSKUId:   in.GoodsSkuId.Value,
		GoodsSKUDesc: in.GoodsSkuDesc.Value,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddCartReply{
		Cart: &pb.Cart{
			Id:           c.Id,
			UserId:       c.UserId,
			GoodsId:      c.GoodsId,
			GoodsSkuId:   c.GoodsSKUId,
			GoodsSkuDesc: c.GoodsSKUDesc,
			Num:          c.Num,
		},
	}, nil
}
