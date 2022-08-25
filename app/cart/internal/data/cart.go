package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/cart/internal/biz"
	"imperialPalaceMall/app/cart/internal/data/ent/cart"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *cartRepo) Create(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		Create().
		SetUserID(item.UserId).
		SetGoodsID(item.GoodsId).
		SetGoodsSku(item.GoodsSKU).
		SetNum(1).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cart{
		Id:       po.ID,
		UserId:   po.UserID,
		GoodsId:  po.GoodsID,
		GoodsSKU: po.GoodsSku,
		Num:      po.Num,
	}, nil
}

func (r *cartRepo) Update(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		UpdateOneID(item.Id).
		SetNum(item.Num).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Cart{
		Id:       po.ID,
		UserId:   po.UserID,
		GoodsId:  po.GoodsID,
		GoodsSKU: po.GoodsSku,
		Num:      po.Num,
	}, nil
}

func (r *cartRepo) FindOne(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		Query().
		Where(
			cart.UserIDEQ(item.UserId),
			cart.GoodsIDEQ(item.GoodsId),
			cart.GoodsSkuEQ(item.GoodsSKU)).
		First(ctx)

	if err != nil {
		return nil, biz.ErrCartItemNotFound
	}
	return &biz.Cart{
		Id:       po.ID,
		UserId:   po.UserID,
		GoodsId:  po.GoodsID,
		GoodsSKU: po.GoodsSku,
		Num:      po.Num,
	}, nil
}
