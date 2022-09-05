package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/pkg/filters"
)

var (
	ErrListGoods                 = errors.New(500, "goods", "list goods error")
	ErrGoodsNotFound             = errors.New(500, "goods", "get goods error")
	ErrSKUNotFound               = errors.New(500, "goods", "get sku error")
	ErrGoodsAndSkuDetailNotFound = errors.New(500, "goods", "get goods and sku detail error")
)

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

type GoodsInfo struct {
	Id      int64
	Kind    int32
	Content string
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

type GoodsAndSkuDetailItem struct {
	// goods
	GoodsName  string
	GoodsDesc  string
	GoodsImage string
	// goods_sku
	Price         int64
	Stock         int64
	GoodsAttrPath []int64
}

// GoodsRepo is a Goods repo.
type GoodsRepo interface {
	List(ctx context.Context, categoryId int64, f filters.Filters) ([]*GoodsSimplify, filters.Metadata, error)
	GetGoodsDetail(ctx context.Context, goodsId int64) (*GoodsDetail, error)
	GetGoodsSKUs(ctx context.Context, goodsId int64) ([]*GoodsSKU, []*GoodsAttr, error)
	GetGoodsAndSkuDetail(ctx context.Context, goodsId int64, skuId int64) (*GoodsAndSkuDetailItem, error)
}

// GoodsUsecase is a Goods usecase.
type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

// NewGoodsUsecase new a Goods usecase.
func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "shop/interface/usecase/goods"))}
}

// List returns the Goods list.
func (uc *GoodsUsecase) List(ctx context.Context, categoryId int64, f filters.Filters) ([]*GoodsSimplify, filters.Metadata, error) {
	return uc.repo.List(ctx, categoryId, f)
}

func (uc *GoodsUsecase) GetGoodsDetail(ctx context.Context, goodsId int64) (*GoodsDetail, error) {
	return uc.repo.GetGoodsDetail(ctx, goodsId)
}

func (uc *GoodsUsecase) GetGoodsSKUs(ctx context.Context, goodsId int64) ([]*GoodsSKU, []*GoodsAttr, error) {
	return uc.repo.GetGoodsSKUs(ctx, goodsId)
}

func (uc *GoodsUsecase) GetGoodsAndSkuDetail(ctx context.Context, goodsId, skuId int64) (*GoodsAndSkuDetailItem, error) {
	return uc.repo.GetGoodsAndSkuDetail(ctx, goodsId, skuId)
}
