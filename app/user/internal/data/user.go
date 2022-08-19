package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/user/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}
