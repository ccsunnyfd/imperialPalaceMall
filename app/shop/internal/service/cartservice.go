package service

import (
	"context"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

func (s *ShopInterface) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartReply, error) {
	cart, err := s.cc.AddCart(ctx, &biz.Cart{
		GoodsId:  req.GoodsId.Value,
		GoodsSKU: req.GoodsSku.Value,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddCartReply{
		Cart: &pb.Cart{
			Id:       cart.Id,
			UserId:   cart.UserId,
			GoodsId:  cart.GoodsId,
			GoodsSku: cart.GoodsSKU,
			Num:      cart.Num,
		},
	}, nil
}
