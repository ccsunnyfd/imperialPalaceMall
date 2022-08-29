package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/pkg/middleware"
)

var (
	ErrAddCart = errors.New("add cart item error")
)

// Cart is a Cart model.
type Cart struct {
	Id           int64
	UserId       int64
	GoodsId      int64
	GoodsSKUId   int64
	GoodsSKUDesc string
	Num          int32
}

// CartRepo is a Cart repo.
type CartRepo interface {
	AddCart(ctx context.Context, item *Cart) (*Cart, error)
}

// CartUsecase is a Cart usecase.
type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

// NewCartUsecase new a Cart usecase.
func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AddCart creates a Cart item if no same item has been added, otherwise update item num in Cart, and returns the new Cart.
func (uc *CartUsecase) AddCart(ctx context.Context, item *Cart) (*Cart, error) {
	userId := middleware.ContextGetUser(ctx).UserId
	item.UserId = userId
	return uc.repo.AddCart(ctx, item)
}
