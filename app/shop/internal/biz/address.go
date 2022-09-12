package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrAddressNotCreated = errors.New(500, "user/address", "address not created")
	ErrAddressNotFound   = errors.NotFound("user/address", "address not found")
	ErrAddressConflict   = errors.Conflict("user/address", "address conflict")
	ErrAddressEdit       = errors.Conflict("user/address", "address edit error")
	ErrAddressDelete     = errors.New(500, "user/address", "address delete error")
)

type Address struct {
	Id         int64    `json:"id"`
	UserId     int64    `json:"userId"`
	UserName   string   `json:"userName"`
	Tel        string   `json:"tel"`
	Region     []string `json:"region"`
	DetailInfo string   `json:"detailInfo"`
	PostCode   string   `json:"postCode"`
}

// AddressRepo is an Address repo.
type AddressRepo interface {
	Create(ctx context.Context, userId int64, address *Address) (int64, error)
	GetByUserId(ctx context.Context, userId int64) ([]*Address, error)
	Update(ctx context.Context, userId int64, address *Address) (int64, error)
	Delete(ctx context.Context, userId int64, id int64) (int64, error)
}

// AddressUsecase is an Address usecase.
type AddressUsecase struct {
	repo AddressRepo
	log  *log.Helper
}

// NewAddressUsecase new an Address usecase.
func NewAddressUsecase(repo AddressRepo, logger log.Logger) *AddressUsecase {
	return &AddressUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "shop/interface/usecase/address"))}
}

func (u *AddressUsecase) CreateAddress(ctx context.Context, userId int64, address *Address) (int64, error) {
	return u.repo.Create(ctx, userId, address)
}

func (u *AddressUsecase) GetAddressByUserId(ctx context.Context, userId int64) ([]*Address, error) {
	return u.repo.GetByUserId(ctx, userId)
}

func (u *AddressUsecase) UpdateAddress(ctx context.Context, userId int64, address *Address) (int64, error) {
	return u.repo.Update(ctx, userId, address)
}

func (u *AddressUsecase) DeleteAddress(ctx context.Context, userId int64, id int64) (int64, error) {
	return u.repo.Delete(ctx, userId, id)
}
