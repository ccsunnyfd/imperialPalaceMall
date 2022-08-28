package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/pkg/filters"
)

var (
	ErrListGoods     = errors.New("goods not listed")
	ErrGoodsNotFound = errors.New("goods not found")
	ErrSKUNotFound   = errors.New("goods sku not found")
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

// GoodsRepo is a Goods repo.
type GoodsRepo interface {
	List(ctx context.Context, categoryId int64, f filters.Filters) ([]*GoodsSimplify, filters.Metadata, error)
	GetGoodsDetail(ctx context.Context, goodsId int64) (*GoodsDetail, error)
	GetGoodsSKUs(ctx context.Context, goodsId int64) ([]*GoodsSKU, []*GoodsAttr, error)
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
