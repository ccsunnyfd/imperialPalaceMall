package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/wrapperspb"
	mallV1 "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/app/pkg/filters"
	"imperialPalaceMall/app/shop/internal/biz"
)

var _ biz.GoodsRepo = (*goodsRepo)(nil)

type goodsRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/goods")),
		sg:   &singleflight.Group{},
	}
}

func (r *goodsRepo) List(ctx context.Context, categoryId int64, f filters.Filters) ([]*biz.GoodsSimplify, filters.Metadata, error) {
	type ret struct {
		goods []*biz.GoodsSimplify
		metas filters.Metadata
	}
	result, err, _ := r.sg.Do(fmt.Sprintf("list_categories_categoryid_%d_page_%d_pagesize_%d", categoryId, f.Page, f.PageSize), func() (interface{}, error) {
		reply, err := r.data.gc.ListGoods(ctx, &mallV1.ListGoodsRequest{
			CategoryId: categoryId,
			Page:       f.Page,
			PageSize:   f.PageSize,
		})
		if err != nil {
			return nil, biz.ErrListGoods
		}
		rv := make([]*biz.GoodsSimplify, 0)
		for _, x := range reply.GoodsList {
			rv = append(rv, &biz.GoodsSimplify{
				Id:         x.Id,
				SpuNo:      x.SpuNo,
				GoodsName:  x.GoodsName,
				StartPrice: x.StartPrice,
				CategoryId: x.CategoryId,
				BrandId:    x.BrandId,
				Info: biz.GoodsInfo{
					Id:      x.Info.Id,
					Kind:    x.Info.Kind,
					Content: x.Info.Content,
				},
			})
		}
		metas := filters.Metadata{
			CurrentPage:  reply.Metadata.CurrentPage,
			PageSize:     reply.Metadata.PageSize,
			FirstPage:    reply.Metadata.FirstPage,
			LastPage:     reply.Metadata.LastPage,
			TotalRecords: reply.Metadata.TotalRecords,
		}
		return ret{
			goods: rv,
			metas: metas,
		}, nil
	})
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	retV := result.(ret)
	return retV.goods, retV.metas, nil
}

func (r *goodsRepo) GetGoodsDetail(ctx context.Context, goodsId int64) (*biz.GoodsDetail, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_goods_detail_%d", goodsId), func() (interface{}, error) {
		reply, err := r.data.gc.GetGoods(ctx, &mallV1.GetGoodsRequest{
			Id: goodsId,
		})
		if err != nil {
			return nil, biz.ErrGoodsNotFound
		}

		infos := make([]*biz.GoodsInfo, 0)
		for _, x := range reply.Goods.Infos {
			infos = append(infos, &biz.GoodsInfo{
				Id:      x.Id,
				Kind:    x.Kind,
				Content: x.Content,
			})
		}

		return &biz.GoodsDetail{
			Id:         reply.Goods.Id,
			SpuNo:      reply.Goods.SpuNo,
			GoodsName:  reply.Goods.GoodsName,
			GoodsDesc:  reply.Goods.GoodsDesc,
			StartPrice: reply.Goods.StartPrice,
			CategoryId: reply.Goods.CategoryId,
			BrandId:    reply.Goods.BrandId,
			Infos:      infos,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return result.(*biz.GoodsDetail), nil
}

func (r *goodsRepo) GetGoodsSKUs(ctx context.Context, goodsId int64) ([]*biz.GoodsSKU, []*biz.GoodsAttr, error) {
	type ret struct {
		skus  []*biz.GoodsSKU
		attrs []*biz.GoodsAttr
	}
	result, err, _ := r.sg.Do(fmt.Sprintf("get_goods_sku_%d", goodsId), func() (interface{}, error) {
		reply, err := r.data.gc.GetSKUs(ctx, &mallV1.GetSKUsRequest{
			Id: wrapperspb.Int64(goodsId),
		})
		if err != nil {
			return nil, biz.ErrSKUNotFound
		}

		skus, attrs := make([]*biz.GoodsSKU, 0), make([]*biz.GoodsAttr, 0)
		for _, x := range reply.Skus {
			skus = append(skus, &biz.GoodsSKU{
				Id:            x.Id,
				GoodsId:       x.GoodsId,
				GoodsAttrPath: x.GoodsAttrPath,
				Price:         x.Price,
				Stock:         x.Stock,
			})
		}
		for _, y := range reply.Attrs {
			attrVals := make([]*biz.GoodsAttrValue, 0)
			for _, z := range y.AttrValues {
				attrVals = append(attrVals, &biz.GoodsAttrValue{
					Id:        z.Id,
					AttrKeyId: z.AttrKeyId,
					AttrValue: z.AttrValue,
				})
			}
			attrs = append(attrs, &biz.GoodsAttr{
				Id:         y.Id,
				GoodsId:    y.GoodsId,
				AttrKey:    y.AttrKey,
				AttrValues: attrVals,
			})
		}

		return ret{skus, attrs}, nil
	})

	if err != nil {
		return nil, nil, err
	}

	retV := result.(ret)
	return retV.skus, retV.attrs, nil
}

func (r *goodsRepo) GetGoodsAndSkuDetail(ctx context.Context, goodsId int64, skuId int64) (*biz.GoodsAndSkuDetailItem, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_goods_and_sku_detail_goodsId:%d_skuId:%d", goodsId, skuId), func() (interface{}, error) {
		reply, err := r.data.gc.GetGoodsAndSkuDetail(ctx, &mallV1.GetGoodsAndSkuDetailRequest{
			GoodsId: wrapperspb.Int64(goodsId),
			SkuId:   wrapperspb.Int64(skuId),
		})
		if err != nil {
			r.log.Errorf("get goods and sku detail error: %v", err)
			return nil, biz.ErrGoodsAndSkuDetailNotFound
		}

		return &biz.GoodsAndSkuDetailItem{
			GoodsName:     reply.GoodsName,
			GoodsDesc:     reply.GoodsDesc,
			GoodsImage:    reply.GoodsImage,
			Price:         reply.Price,
			Stock:         reply.Stock,
			GoodsAttrPath: reply.GoodsAttrPath,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return result.(*biz.GoodsAndSkuDetailItem), nil
}
