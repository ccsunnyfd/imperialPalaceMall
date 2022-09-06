package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	mallPb "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/app/mall/internal/biz"
)

var _ biz.CategoryRepo = (*categoryRepo)(nil)

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) Save(ctx context.Context, g *biz.Category) (*biz.Category, error) {
	return g, nil
}

func (r *categoryRepo) Update(ctx context.Context, g *biz.Category) (*biz.Category, error) {
	return g, nil
}

func (r *categoryRepo) FindByID(context.Context, int64) (*biz.Category, error) {
	return nil, nil
}

func (r *categoryRepo) ListAll(ctx context.Context) ([]*biz.Category, error) {
	cc, err := r.data.db.Category.Query(ctx)

	if err != nil {
		return nil, errors.Wrap(mallPb.ErrorCategoryNotFound("ListAll"), "category")
	}
	rv := make([]*biz.Category, 0, len(cc))
	for _, po := range cc {
		rv = append(rv, &biz.Category{
			Id:           po.Id,
			CategoryName: po.CategoryName,
		})
	}
	return rv, nil
}
