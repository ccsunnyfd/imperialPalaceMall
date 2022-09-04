package service

import (
	"context"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/pkg/middleware"
	"imperialPalaceMall/app/shop/internal/biz"
)

func (s *ShopInterface) AddCart(ctx context.Context, req *pb.AddCartRequest) (*pb.AddCartReply, error) {
	userId := middleware.ContextGetUser(ctx).UserId

	cart, err := s.cc.AddCart(ctx, &biz.Cart{
		UserId:       userId,
		GoodsId:      req.GoodsId.Value,
		GoodsSKUId:   req.GoodsSkuId.Value,
		GoodsSKUDesc: req.GoodsSkuDesc.Value,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddCartReply{
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

func (s *ShopInterface) GetMyCart(ctx context.Context, req *pb.GetMyCartRequest) (*pb.GetMyCartReply, error) {
	userId := middleware.ContextGetUser(ctx).UserId

	cartItems, err := s.cc.GetCart(ctx, userId)
	if err != nil {
		return nil, err
	}
	rv := make([]*pb.GetMyCartReply_CartItem, 0, len(cartItems))
	for _, x := range cartItems {
		goodsId := x.GoodsId
		skuId := x.GoodsSKUId

		goodsAndSku, err := s.gc.GetGoodsAndSkuDetail(ctx, goodsId, skuId)
		if err != nil {
			return nil, err
		}

		rv = append(rv, &pb.GetMyCartReply_CartItem{
			CartId:        x.Id,
			GoodsSkuId:    skuId,
			GoodsId:       goodsId,
			Num:           x.Num,
			GoodsSkuDesc:  x.GoodsSKUDesc,
			GoodsName:     goodsAndSku.GoodsName,
			GoodsDesc:     goodsAndSku.GoodsDesc,
			GoodsImage:    goodsAndSku.GoodsImage,
			Price:         goodsAndSku.Price,
			Stock:         goodsAndSku.Stock,
			GoodsAttrPath: goodsAndSku.GoodsAttrPath,
		})
	}

	return &pb.GetMyCartReply{
		Items: rv,
	}, nil
}

func (s *ShopInterface) UpdateCartNum(ctx context.Context, req *pb.UpdateCartNumRequest) (*pb.UpdateCartNumReply, error) {
	cart, err := s.cc.UpdateCartNum(ctx, req.Id.Value, req.Num.Value)
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
