package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/pkg/filters"
)

type GoodsInfo struct {
	Id      int64
	Kind    int32
	Content string
}

// Goods is a Goods model.
type Goods struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	GoodsDesc  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
}

type GoodsAttrValue struct {
	Id        int64
	AttrKeyId int64
	AttrValue string
}

type GoodsAttr struct {
	Id         int64
	GoodsId    int64
	AttrKey    string
	AttrValues []*GoodsAttrValue
}

type GoodsSKU struct {
	Id            int64
	GoodsId       int64
	GoodsAttrPath []int64
	Price         int64
	Stock         int64
}

type GoodsDetail struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	GoodsDesc  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
	Infos      []*GoodsInfo
}

type GoodsSimplify struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
	Info       GoodsInfo
}

type GoodsAndSkuDetail struct {
	GoodsAttrPath []int64
	Price         int64
	Stock         int64
	GoodsName     string
	GoodsDesc     string
	GoodsImage    string
}

// GoodsRepo is a Goods repo.
type GoodsRepo interface {
	Save(context.Context, *Goods) (*Goods, error)
	Update(context.Context, *Goods) (*Goods, error)
	GetGoodsDetail(context.Context, int64) (*GoodsDetail, error)
	List(context.Context, int64, filters.Filters) ([]*GoodsSimplify, filters.Metadata, error)
	GetAttrs(context.Context, int64) ([]*GoodsAttr, error)
	GetSKUs(context.Context, int64) ([]*GoodsSKU, error)
	GetBySkuId(context.Context, int64) (*GoodsSKU, error)
}

// GoodsUsecase is a Goods usecase.
type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

// NewGoodsUsecase new a Goods usecase.
func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

// List returns the Goods list.
func (uc *GoodsUsecase) List(ctx context.Context, categoryId int64, f filters.Filters) ([]*GoodsSimplify, filters.Metadata, error) {
	uc.log.WithContext(ctx).Infof("List: ")
	return uc.repo.List(ctx, categoryId, f)
}

// GetGoodsDetail returns the GoodsDetail by goods_id.
func (uc *GoodsUsecase) GetGoodsDetail(ctx context.Context, goodsId int64) (*GoodsDetail, error) {
	uc.log.WithContext(ctx).Infof("Get goodsDetail by goods_id: %d", goodsId)

	g, err := uc.repo.GetGoodsDetail(ctx, goodsId)
	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	return nil, ErrGoodsNotFound
		//}
		return nil, err
	}

	return g, nil
}

// GetGoodsSKUs returns the GoodsSKU and GoodsAttr list by goods_id.
func (uc *GoodsUsecase) GetGoodsSKUs(ctx context.Context, goodsId int64) ([]*GoodsSKU, []*GoodsAttr, error) {
	uc.log.WithContext(ctx).Infof("GetSKUs by goods_id: %d", goodsId)

	skus, err := uc.repo.GetSKUs(ctx, goodsId)
	if err != nil {
		return nil, nil, err
	}

	attrs, err := uc.repo.GetAttrs(ctx, goodsId)
	if err != nil {
		return nil, nil, err
	}

	return skus, attrs, nil
}

func (uc *GoodsUsecase) GetGoodsAndSkuDetails(ctx context.Context, goodsId int64, skuId int64) (*GoodsAndSkuDetail, error) {
	goodsDetail, err := uc.GetGoodsDetail(ctx, goodsId)
	if err != nil {
		return nil, err
	}
	sku, err := uc.repo.GetBySkuId(ctx, skuId)
	if err != nil {
		return nil, err
	}

	var imagePath string
	for _, x := range goodsDetail.Infos {
		if x.Kind == 0 {
			imagePath = x.Content
			break
		}
	}

	return &GoodsAndSkuDetail{
		GoodsAttrPath: sku.GoodsAttrPath,
		Price:         sku.Price,
		Stock:         sku.Stock,
		GoodsName:     goodsDetail.GoodsName,
		GoodsDesc:     goodsDetail.GoodsDesc,
		GoodsImage:    imagePath,
	}, nil
}
