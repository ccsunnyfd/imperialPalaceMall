package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrListCategories = errors.New("categories not listed")
)

// Category is a Category model.
type Category struct {
	Id           int64
	CategoryName string
}

// CategoryRepo is a Category repo.
type CategoryRepo interface {
	ListCategories(ctx context.Context) ([]*Category, error)
}

// CategoryUsecase is a Category usecase.
type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

// NewCategoryUsecase new a Category usecase.
func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "shop/interface/usecase/category"))}
}

// ListCategories returns the Category list.
func (uc *CategoryUsecase) ListCategories(ctx context.Context) ([]*Category, error) {
	return uc.repo.ListCategories(ctx)
}
