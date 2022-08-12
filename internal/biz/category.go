package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	mall "imperialPalaceMall/api/mall/v1"
)

var (
	// ErrCategoryNotFound is category not found.
	ErrCategoryNotFound = errors.NotFound(mall.ErrorReason_CATEGORY_NOT_FOUND.String(), "category not found")
)

// Category is a Category model.
type Category struct {
	Id           int64
	CategoryName string
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	Save(context.Context, *Category) (*Category, error)
	Update(context.Context, *Category) (*Category, error)
	FindByID(context.Context, int64) (*Category, error)
	ListAll(context.Context) ([]*Category, error)
}

// CategoryUsecase is a Category usecase.
type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCategoryUsecase new a Category usecase.
func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

// ListCategories returns the Category list.
func (uc *CategoryUsecase) ListCategories(ctx context.Context) ([]*Category, error) {
	uc.log.WithContext(ctx).Infof("ListCategories: ")
	return uc.repo.ListAll(ctx)
}
