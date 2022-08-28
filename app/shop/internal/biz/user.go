package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrTokenNotCreated = errors.New("token not created")
	ErrTokenNotFound   = errors.New("token not found")
)

type UserCache struct {
	UserId int64
	OpenId string
}

type User struct {
	Id        int64  `json:"id"`
	OpenId    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    Gender `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
}

type Gender int64

const (
	unknown Gender = 0
	male    Gender = 1
	female  Gender = 2
)

// UserRepo is a User repo.
type UserRepo interface {
	WXLogin(ctx context.Context, code string, encryptedData string, iv string, sessionIsValid bool) (string, error)
	CheckToken(ctx context.Context, token string) (*UserCache, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "shop/interface/usecase/user"))}
}

func (u *UserUsecase) WXLogin(ctx context.Context, code string, encryptedData string, iv string, sessionIsValid bool) (string, error) {
	return u.repo.WXLogin(ctx, code, encryptedData, iv, sessionIsValid)
}

func (u *UserUsecase) CheckToken(ctx context.Context, token string) (*UserCache, error) {
	return u.repo.CheckToken(ctx, token)
}
