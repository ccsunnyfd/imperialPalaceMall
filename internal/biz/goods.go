package biz

import (
	"context"
	"imperialPalaceMall/internal/filters"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	mall "imperialPalaceMall/api/mall/v1"
)

var (
	// ErrGoodsNotFound is goods not found.
	ErrGoodsNotFound = errors.NotFound(mall.ErrorReason_GOODS_NOT_FOUND.String(), "goods not found")
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
	StartPrice float64
	CategoryId int64
	BrandId    int64
}

type GoodsDetail struct {
	Id         int64
	SpuNo      string
	GoodsName  string
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
	Save(context.Context, *Goods) (*Goods, error)
	Update(context.Context, *Goods) (*Goods, error)
	FindByID(context.Context, int64) (*GoodsDetail, error)
	List(context.Context, int64, filters.Filters) ([]*GoodsSimplify, filters.Metadata, error)
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

// Get returns the GoodsDetail by id.
func (uc *GoodsUsecase) Get(ctx context.Context, goodsId int64) (*GoodsDetail, error) {
	uc.log.WithContext(ctx).Infof("Get: ")

	g, err := uc.repo.FindByID(ctx, goodsId)
	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	return nil, ErrGoodsNotFound
		//}
		return nil, err
	}

	return g, nil
}
