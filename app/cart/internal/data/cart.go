package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(biz.ErrAddCartItem, "Create: userId_%d, goodsId_%d, goodsSKUId_%d, goodsSKUDesc_%s", item.UserId, item.GoodsId, item.GoodsSKUId, item.GoodsSKUDesc)
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

func (r *cartRepo) UpdateNum(ctx context.Context, cartId int64, num int32) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		UpdateOneID(cartId).
		SetNum(num).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrapf(biz.ErrUpdateCartItem, "UpdateNum to %d by cartId_%d", num, cartId)
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

func (r *cartRepo) FindOne(ctx context.Context, item *biz.Cart) (*biz.Cart, error) {
	po, err := r.data.db.Cart.
		Query().
		Where(
			cart.UserIDEQ(item.UserId),
			cart.GoodsIDEQ(item.GoodsId),
			cart.GoodsSkuIDEQ(item.GoodsSKUId)).
		First(ctx)

	if err != nil {
		return nil, errors.Wrapf(biz.ErrCartItemNotFound, "FindOneBy userId_%d, goodsId_%d, goodsSKUId_%d", item.UserId, item.GoodsId, item.GoodsSKUId)
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
			cart.UserIDEQ(userId)).
		All(ctx)
	// 如果查询为空列表，ent不认为是错误，只会返回空列表
	if err != nil {
		return nil, errors.Wrapf(biz.ErrGetCartItem, "GetByUserId_%d", userId)
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
	affected, err := r.data.db.Cart.Delete().Where(cart.IDIn(ids...), cart.UserIDEQ(userId)).Exec(ctx)
	if err != nil {
		return 0, errors.Wrapf(biz.ErrDeleteCartItems, "DeleteByIds_userId_%d_ids_%v", userId, ids)
	}
	return int64(affected), nil
}
