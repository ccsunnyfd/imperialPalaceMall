package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	cartPb "imperialPalaceMall/api/cart/v1"
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
		SetGoodsSkuID(item.GoodsSKUId).
		SetGoodsSkuDesc(item.GoodsSKUDesc).
		SetNum(1).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(cartPb.ErrorCartItemAddError("create cart item error"), "cart")
	}
	return &biz.Cart{
		Id:           po.ID,
		UserId:       po.UserID,
		GoodsId:      po.GoodsID,
		GoodsSKUId:   po.GoodsSkuID,
		GoodsSKUDesc: po.GoodsSkuDesc,
		Num:          po.Num,
	}, nil
}

func (r *cartRepo) UpdateNum(ctx context.Context, userId int64, cartId int64, num int32) (int64, error) {
	affected, err := r.data.db.Cart.
		Update().
		Where(
			cart.UserID(userId),
			cart.ID(cartId),
		).
		SetNum(num).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(cartPb.ErrorCartItemUpdateError("update cart num to %d error", num), "cart")
	}
	return int64(affected), nil
}

func (r *cartRepo) FindOne(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		Query().
		Where(
			cart.UserID(item.UserId),
			cart.GoodsID(item.GoodsId),
			cart.GoodsSkuID(item.GoodsSKUId)).
		First(ctx)

	if err != nil {
		return nil, errors.Wrap(cartPb.ErrorCartItemNotFound("find cart item error"), "cart")
	}
	return &biz.Cart{
		Id:           po.ID,
		UserId:       po.UserID,
		GoodsId:      po.GoodsID,
		GoodsSKUId:   po.GoodsSkuID,
		GoodsSKUDesc: po.GoodsSkuDesc,
		Num:          po.Num,
	}, nil
}

func (r *cartRepo) GetByUserId(ctx context.Context, userId int64) ([]*biz.Cart, error) {
	po, err := r.data.db.Cart.
		Query().
		Where(
			cart.UserID(userId)).
		All(ctx)
	// ???????????????????????????ent??????????????????????????????????????????
	if err != nil {
		return nil, errors.Wrap(cartPb.ErrorCartItemGetError("get user by id error"), "cart")
	}

	rv := make([]*biz.Cart, 0, len(po))
	for _, x := range po {
		rv = append(rv, &biz.Cart{
			Id:           x.ID,
			UserId:       x.UserID,
			GoodsId:      x.GoodsID,
			GoodsSKUId:   x.GoodsSkuID,
			GoodsSKUDesc: x.GoodsSkuDesc,
			Num:          x.Num,
		})
	}
	return rv, nil
}

func (r *cartRepo) DeleteByIds(ctx context.Context, userId int64, ids []int64) (int64, error) {
	affected, err := r.data.db.Cart.Delete().Where(cart.IDIn(ids...), cart.UserID(userId)).Exec(ctx)
	if err != nil {
		return 0, errors.Wrap(cartPb.ErrorCartItemDeleteError("delete cart item by ids_%v error", ids), "cart")
	}
	return int64(affected), nil
}
