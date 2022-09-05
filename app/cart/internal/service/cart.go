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

func (s *CartService) GetCartsByUserId(ctx context.Context, in *pb.GetCartsByUserIdRequest) (*pb.GetCartsByUserIdReply, error) {
	cc, err := s.uc.GetCartsByUserId(ctx, in.UserId.Value)
	if err != nil {
		return nil, err
	}
	rv := make([]*pb.GetCartsByUserIdReply_CartItem, 0, len(cc))
	for _, x := range cc {
		rv = append(rv, &pb.GetCartsByUserIdReply_CartItem{
			CartId:       x.Id,
			GoodsSkuId:   x.GoodsSKUId,
			GoodsId:      x.GoodsId,
			Num:          x.Num,
			GoodsSkuDesc: x.GoodsSKUDesc,
		})
	}

	return &pb.GetCartsByUserIdReply{
		Items: rv,
	}, nil
}

func (s *CartService) UpdateCartNum(ctx context.Context, in *pb.UpdateCartNumRequest) (*pb.UpdateCartNumReply, error) {
	cart, err := s.uc.UpdateCartNum(ctx, in.CartId.Value, in.Num.Value)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCartNumReply{
		Cart: &pb.Cart{
			Id:           cart.Id,
			UserId:       cart.UserId,
			GoodsId:      cart.GoodsId,
			GoodsSkuId:   cart.GoodsSKUId,
			GoodsSkuDesc: cart.GoodsSKUDesc,
			Num:          cart.Num,
		},
	}, nil
}
