package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userPb "imperialPalaceMall/api/user/v1"
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
	GetByDetails(ctx context.Context, addr *Address) (*Address, error)
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
	_, err := u.repo.GetByDetails(ctx, addr)
	if err != nil {
		if userPb.IsAddressNotFound(err) {
			return u.repo.Save(ctx, addr)
		}
		return -1, err
	}

	return -1, userPb.ErrorAddressConflict("create address conflict") // 地址冲突不属于服务器错误
}
