package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	cart "imperialPalaceMall/api/cart/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrCartItemNotFound is cart item not found.
	ErrCartItemNotFound = errors.NotFound(cart.ErrorReason_CART_ITEM_NOT_FOUND.String(), "cart item not found")
)

// Cart is a Cart model.
type Cart struct {
	Id       int64
	UserId   int64
	GoodsId  int64
	GoodsSKU int64
	Num      int32
}

// CartRepo is a Cart repo.
type CartRepo interface {
	FindOne(ctx context.Context, item *Cart) (*Cart, error)
	Create(ctx context.Context, item *Cart) (*Cart, error)
	Update(ctx context.Context, item *Cart) (*Cart, error)
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
	uc.log.WithContext(ctx).Infof("AddCart: %+v", item)

	existItem, err := uc.repo.FindOne(ctx, item)
	// 不存在则新建
	if err != nil {
		if errors.Is(err, ErrCartItemNotFound) {
			return uc.repo.Create(ctx, item)
		}
		return nil, err
	}
	// 存在则更新
	existItem.Num++
	return uc.repo.Update(ctx, existItem)
}
