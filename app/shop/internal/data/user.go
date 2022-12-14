package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	userV1 "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
		sg:   &singleflight.Group{},
	}
}

func (r *userRepo) WXLogin(ctx context.Context, code string, encryptedData string, iv string, sessionIsValid bool) (string, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("wxlogin_with_code_%s", code), func() (interface{}, error) {
		reply, err := r.data.uc.WxLogin(ctx, &userV1.WxLoginRequest{
			Code:           code,
			EncryptedData:  encryptedData,
			Iv:             iv,
			SessionIsValid: sessionIsValid,
		})
		if err != nil {
			return nil, errors.Wrapf(biz.ErrTokenNotCreated, "shop_user")
		}
		return reply.Token, nil
	})
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (r *userRepo) CheckToken(ctx context.Context, token string) (*biz.UserCache, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("check_token_%s", token), func() (interface{}, error) {
		reply, err := r.data.uc.CheckToken(ctx, &userV1.CheckTokenRequest{
			Token: token,
		})
		if err != nil {
			return nil, errors.Wrapf(biz.ErrTokenNotFound, "shop_user")
		}
		return &biz.UserCache{
			UserId: reply.User.UserId,
			OpenId: reply.User.OpenId,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.UserCache), nil
}
