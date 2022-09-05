package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	cartPb "imperialPalaceMall/api/cart/v1"
)

var (
	ErrAddCartItem      = errors.New(500, cartPb.ErrorReason_CART_ITEM_ADD_ERROR.String(), "add cart item failed")
	ErrUpdateCartItem   = errors.New(500, cartPb.ErrorReason_CART_ITEM_UPDATE_ERROR.String(), "update cart item failed")
	ErrCartItemNotFound = errors.NotFound(cartPb.ErrorReason_CART_ITEM_NOT_FOUND.String(), "cart item not found")
	ErrGetCartItem      = errors.New(500, cartPb.ErrorReason_CART_ITEM_GET_ERROR.String(), "get cart item failed")
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
	FindOne(ctx context.Context, item *Cart) (*Cart, error)
	Create(ctx context.Context, item *Cart) (*Cart, error)
	UpdateNum(ctx context.Context, cartId int64, num int32) (*Cart, error)
	GetByUserId(ctx context.Context, userId int64) ([]*Cart, error)
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
	return uc.repo.UpdateNum(ctx, existItem.Id, existItem.Num+1)
}

func (uc *CartUsecase) GetCartsByUserId(ctx context.Context, userId int64) ([]*Cart, error) {
	uc.log.WithContext(ctx).Infof("GetCartsByUserId: %d", userId)
	return uc.repo.GetByUserId(ctx, userId)
}

func (uc *CartUsecase) UpdateCartNum(ctx context.Context, cartId int64, num int32) (*Cart, error) {
	uc.log.WithContext(ctx).Infof("UpdateCartNum_cartId_%d_num_%d", cartId, num)
	return uc.repo.UpdateNum(ctx, cartId, num)
}
