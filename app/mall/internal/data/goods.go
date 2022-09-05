package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"imperialPalaceMall/app/mall/internal/biz"
	"imperialPalaceMall/app/pkg/filters"
)

var _ biz.GoodsRepo = (*goodsRepo)(nil)

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *goodsRepo) Save(ctx context.Context, g *biz.Goods) (*biz.Goods, error) {
	return g, nil
}

func (r *goodsRepo) Update(ctx context.Context, g *biz.Goods) (*biz.Goods, error) {
	return g, nil
}

func (r *goodsRepo) GetGoodsDetail(ctx context.Context, id int64) (*biz.GoodsDetail, error) {
	g, err := r.data.db.Goods.Get(ctx, id)

	if err != nil {
		return nil, errors.Wrapf(biz.ErrGoodsNotFound, "GetGoodsDetail: id_%d", id)
	}

	infos, err := r.data.db.Goods.GetGoodsInfo(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(biz.ErrGoodsInfoNotFound, "GetGoodsDetail: id_%d", id)
	}

	var infoList = make([]*biz.GoodsInfo, 0, len(infos))
	for _, f := range infos {
		infoList = append(infoList, &biz.GoodsInfo{
			Id:      f.Id,
			Kind:    f.Kind,
			Content: f.Content,
		})
	}

	return &biz.GoodsDetail{
		Id:         g.Id,
		SpuNo:      g.SpuNo,
		GoodsName:  g.GoodsName,
		GoodsDesc:  g.GoodsDesc,
		StartPrice: g.StartPrice,
		CategoryId: g.CategoryId,
		BrandId:    g.BrandId,
		Infos:      infoList,
	}, nil
}

func (r *goodsRepo) GetSKUs(ctx context.Context, goodsId int64) ([]*biz.GoodsSKU, error) {
	skus, err := r.data.db.Goods.GetSKUs(ctx, goodsId)

	if err != nil {
		return nil, errors.Wrapf(biz.ErrSkuNotFound, "GetSKUs: goodsId_%d", goodsId)
	}
	rv := make([]*biz.GoodsSKU, 0, len(skus))
	for _, sku := range skus {
		rv = append(rv, &biz.GoodsSKU{
			Id:            sku.Id,
			GoodsId:       sku.GoodsId,
			GoodsAttrPath: sku.GoodsAttrPath,
			Price:         sku.Price,
			Stock:         sku.Stock,
		})
	}
	return rv, nil
}

func (r *goodsRepo) GetBySkuId(ctx context.Context, skuId int64) (*biz.GoodsSKU, error) {
	sku, err := r.data.db.Goods.GetSku(ctx, skuId)

	if err != nil {
		return nil, errors.Wrapf(biz.ErrSkuNotFound, "GetBySkuId: skuId_%d", skuId)
	}
	return &biz.GoodsSKU{
		Id:            sku.Id,
		GoodsId:       sku.GoodsId,
		GoodsAttrPath: sku.GoodsAttrPath,
		Price:         sku.Price,
		Stock:         sku.Stock,
	}, nil
}

func (r *goodsRepo) GetAttrs(ctx context.Context, goodsId int64) ([]*biz.GoodsAttr, error) {
	kk, err := r.data.db.Goods.GetGoodsAttrKeys(ctx, goodsId)
	if err != nil {
		return nil, errors.Wrapf(biz.ErrAttrsNotFound, "GetAttrs: goodsId_%d", goodsId)
	}

	attrList := make([]*biz.GoodsAttr, 0, len(kk))
	for _, k := range kk {
		attrKeyId := k.Id
		vv, err := r.data.db.Goods.GetGoodsAttrValuesByKeyId(ctx, attrKeyId)
		if err != nil {
			return nil, errors.Wrapf(biz.ErrAttrValuesNotFound, "GetAttrs: goodsId_%d", goodsId)
		}

		attrValueList := make([]*biz.GoodsAttrValue, 0, len(vv))
		for _, v := range vv {
			attrValueList = append(attrValueList, &biz.GoodsAttrValue{
				Id:        v.Id,
				AttrKeyId: v.AttrKeyId,
				AttrValue: v.AttrValue,
			})
		}

		attrList = append(attrList, &biz.GoodsAttr{
			Id:         k.Id,
			GoodsId:    goodsId,
			AttrKey:    k.AttrKey,
			AttrValues: attrValueList,
		})
	}

	return attrList, nil
}

func (r *goodsRepo) List(ctx context.Context, categoryId int64, f filters.Filters) ([]*biz.GoodsSimplify, filters.Metadata, error) {
	gg, metadata, err := r.data.db.Goods.Query(ctx, categoryId, f)

	if err != nil {
		return nil, filters.Metadata{}, errors.Wrapf(biz.ErrGoodsList, "List: categoryId_%d", categoryId)
	}
	rv := make([]*biz.GoodsSimplify, 0, len(gg))
	for _, po := range gg {
		rv = append(rv, &biz.GoodsSimplify{
			Id:         po.Id,
			SpuNo:      po.SpuNo,
			GoodsName:  po.GoodsName,
			StartPrice: po.StartPrice,
			CategoryId: po.CategoryId,
			BrandId:    po.BrandId,
			Info: biz.GoodsInfo{
				Id:      po.Info.Id,
				Kind:    po.Info.Kind,
				Content: po.Info.Content,
			},
		})
	}
	return rv, metadata, nil
}
