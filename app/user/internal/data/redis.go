package data

import (
	"context"
	"encoding/json"
	"imperialPalaceMall/app/user/internal/biz"
	"time"
)

func (r *userRepo) GetByOpenId(ctx context.Context, openid string) (*biz.UserCache, error) {
	result, err := r.data.rdb.Get(ctx, openid).Result()
	r.log.Infof("getByOpenId: openid %s result %s err %v\n", openid, result, err)

	if err != nil {
		return nil, err
	}
	var cacheUser = &biz.UserCache{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *userRepo) SetUserCache(ctx context.Context, user *biz.UserCache, key string, ttl time.Duration) {
	marshal, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("fail to set user cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.rdb.Set(ctx, key, string(marshal), ttl).Err()
	if err != nil {
		r.log.Errorf("fail to set user cache:redis.Set(%v) error(%v)", user, err)
	}
}
