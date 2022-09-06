package data

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	userPb "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/user/internal/biz"
	"time"
)

func (r *userRepo) Get(ctx context.Context, openidOrToken string) (*biz.UserCache, error) {
	result, err := r.data.rdb.Get(ctx, openidOrToken).Result()

	if err != nil {
		return nil, errors.Wrap(userPb.ErrorUserCacheNotFound("Get_by_openidOrToken_%s", openidOrToken), "user")
	}
	var cacheUser = &biz.UserCache{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, errors.Wrap(userPb.ErrorUserCacheUnmarshalError("Get_by_openidOrToken_%s", openidOrToken), "user")
	}
	return cacheUser, nil
}

func (r *userRepo) SetUserCache(ctx context.Context, user *biz.UserCache, key string, ttl time.Duration) error {
	marshal, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(userPb.ErrorUserCacheMarshalError("fail to set user cache:json.Marshal(%v) error(%v)", user, err), "user")
	}
	err = r.data.rdb.Set(ctx, key, string(marshal), ttl).Err()
	if err != nil {
		return errors.Wrap(userPb.ErrorUserCacheSetError("fail to set user cache:redis.Set(%v) error(%v)", user, err), "user")
	}
}

func (r *userRepo) DeleteOldToken(ctx context.Context, oldTokenKey string) {
	r.data.rdb.Del(ctx, oldTokenKey)
}
