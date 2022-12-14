package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrAddCart         = errors.New(500, "CART_ADD_ERROR", "add cart item error")
	ErrGetCart         = errors.New(500, "CART_NOT_FOUND", "get cart by userId error")
	ErrUpdateCartNum   = errors.New(500, "CART_UPDATE_NUM_ERROR", "update cart num error")
	ErrRemoveCartItems = errors.New(500, "CART_REMOVE_ERROR", "remove cart items error")
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
	GetCart(ctx context.Context, userId int64) ([]*Cart, error)
	UpdateCartNum(ctx context.Context, userId int64, cartId int64, num int32) (int64, error)
	RemoveCartItems(ctx context.Context, userId int64, ids []int64) (int64, error)
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
	return uc.repo.AddCart(ctx, item)
}

func (uc *CartUsecase) GetCart(ctx context.Context, userId int64) ([]*Cart, error) {
	return uc.repo.GetCart(ctx, userId)
}

func (uc *CartUsecase) UpdateCartNum(ctx context.Context, userId int64, cartId int64, num int32) (int64, error) {
	return uc.repo.UpdateCartNum(ctx, userId, cartId, num)
}

func (uc *CartUsecase) RemoveCartItems(ctx context.Context, userId int64, ids []int64) (int64, error) {
	return uc.repo.RemoveCartItems(ctx, userId, ids)
}
