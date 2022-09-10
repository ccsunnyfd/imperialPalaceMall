package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Address is a Address model.
type Address struct {
	Id         int64    `json:"id"`
	UserId     int64    `json:"userId"`
	UserName   string   `json:"userName"`
	Tel        string   `json:"tel"`
	Region     []string `json:"region"`
	DetailInfo string   `json:"detailInfo"`
	PostCode   string   `json:"postCode"`
}

// AddressRepo is a Address repo.
type AddressRepo interface {
	Save(context.Context, *Address) (int64, error)
	GetByUserId(ctx context.Context, userId int64) ([]*Address, error)
}

// AddressUsecase is a Address usecase.
type AddressUsecase struct {
	repo AddressRepo
	log  *log.Helper
}

// NewAddressUsecase new a Address usecase.
func NewAddressUsecase(repo AddressRepo, logger log.Logger) *AddressUsecase {
	return &AddressUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *AddressUsecase) GetAddressByUserId(ctx context.Context, userId int64) ([]*Address, error) {
	return u.repo.GetByUserId(ctx, userId)
}

func (u *AddressUsecase) SaveAddress(ctx context.Context, addr *Address) (int64, error) {
	return u.repo.Save(ctx, addr)
}
