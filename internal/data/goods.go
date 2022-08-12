package data

import (
	"context"
	"imperialPalaceMall/internal/filters"

	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/internal/biz"
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

func (r *goodsRepo) FindByID(ctx context.Context, id int64) (*biz.GoodsDetail, error) {
	g, err := r.data.db.Goods.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	var infoList = make([]*biz.GoodsInfo, 0, len(g.Infos))

	for _, info := range g.Infos {
		infoList = append(infoList, &biz.GoodsInfo{
			Id:      info.Id,
			Kind:    info.Kind,
			Content: info.Content,
		})
	}
	return &biz.GoodsDetail{
		Id:         g.Id,
		SpuNo:      g.SpuNo,
		GoodsName:  g.GoodsName,
		StartPrice: g.StartPrice,
		CategoryId: g.CategoryId,
		BrandId:    g.BrandId,
		Infos:      infoList,
	}, nil
}

func (r *goodsRepo) List(ctx context.Context, categoryId int64, f filters.Filters) ([]*biz.GoodsSimplify, filters.Metadata, error) {
	gg, metadata, err := r.data.db.Goods.Query(ctx, categoryId, f)

	if err != nil {
		return nil, filters.Metadata{}, err
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
