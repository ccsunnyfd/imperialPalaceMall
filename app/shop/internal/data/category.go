package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	mallV1 "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

var _ biz.CategoryRepo = (*categoryRepo)(nil)

type categoryRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

// NewCategoryRepo .
func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/category")),
		sg:   &singleflight.Group{},
	}
}

func (r *categoryRepo) ListCategories(ctx context.Context) ([]*biz.Category, error) {
	result, err, _ := r.sg.Do("list_categories", func() (interface{}, error) {
		reply, err := r.data.catc.ListCategory(ctx, &mallV1.ListCategoryRequest{})
		if err != nil {
			return nil, errors.Wrapf(biz.ErrListCategories, "shop_category")
		}
		rv := make([]*biz.Category, 0)
		for _, x := range reply.Categories {
			rv = append(rv, &biz.Category{
				Id:           x.Id,
				CategoryName: x.CategoryName,
			})
		}
		return rv, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]*biz.Category), nil
}
