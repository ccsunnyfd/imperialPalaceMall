package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/wrapperspb"
	cartV1 "imperialPalaceMall/api/cart/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cartRepo) AddCart(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	reply, err := r.data.cc.AddCart(ctx, &cartV1.AddCartRequest{
		UserId:       wrapperspb.Int64(item.UserId),
		GoodsId:      wrapperspb.Int64(item.GoodsId),
		GoodsSkuId:   wrapperspb.Int64(item.GoodsSKUId),
		GoodsSkuDesc: wrapperspb.String(item.GoodsSKUDesc),
	})
	if err != nil {
		return nil, biz.ErrAddCart
	}
	return &biz.Cart{
		Id:           reply.Cart.Id,
		UserId:       reply.Cart.UserId,
		GoodsId:      reply.Cart.GoodsId,
		GoodsSKUId:   reply.Cart.GoodsSkuId,
		GoodsSKUDesc: reply.Cart.GoodsSkuDesc,
		Num:          reply.Cart.Num,
	}, nil
}

func (r *cartRepo) GetCart(ctx context.Context, userId int64) ([]*biz.Cart, error) {
	reply, err := r.data.cc.GetCartByUserId(ctx, &cartV1.GetCartByUserIdRequest{
		UserId: wrapperspb.Int64(userId),
	})
	if err != nil {
		return nil, biz.ErrGetCart
	}

	rv := make([]*biz.Cart, 0, len(reply.Items))
	for _, x := range reply.Items {
		rv = append(rv, &biz.Cart{
			Id:           x.CartId,
			UserId:       userId,
			GoodsId:      x.GoodsId,
			GoodsSKUId:   x.GoodsSkuId,
			GoodsSKUDesc: x.GoodsSkuDesc,
			Num:          x.Num,
		})
	}

	return rv, nil
}

func (r *cartRepo) UpdateCartNum(ctx context.Context, cartId int64, num int32) (*biz.Cart, error) {
	reply, err := r.data.cc.UpdateCartNum(ctx, &cartV1.UpdateCartNumRequest{
		CartId: wrapperspb.Int64(cartId),
		Num:    wrapperspb.Int32(num),
	})
	if err != nil {
		return nil, biz.ErrUpdateCartNum
	}

	return &biz.Cart{
		Id:           reply.Cart.Id,
		UserId:       reply.Cart.UserId,
		GoodsId:      reply.Cart.GoodsId,
		GoodsSKUId:   reply.Cart.GoodsSkuId,
		GoodsSKUDesc: reply.Cart.GoodsSkuDesc,
		Num:          reply.Cart.Num,
	}, nil
}
