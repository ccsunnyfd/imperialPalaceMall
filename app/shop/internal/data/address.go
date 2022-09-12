package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/wrapperspb"
	userV1 "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

var _ biz.AddressRepo = (*addressRepo)(nil)

type addressRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

// NewAddressRepo .
func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return &addressRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user/address")),
		sg:   &singleflight.Group{},
	}
}

func (r *addressRepo) Create(ctx context.Context, userId int64, address *biz.Address) (int64, error) {
	reply, err := r.data.uc.SaveAddress(ctx, &userV1.SaveAddressRequest{
		UserId:     wrapperspb.Int64(userId),
		UserName:   address.UserName,
		Tel:        address.Tel,
		Region:     address.Region,
		DetailInfo: address.DetailInfo,
		PostCode:   address.PostCode,
	})
	if err != nil {
		if userV1.IsAddressConflict(err) {
			return -1, biz.ErrAddressConflict // 地址冲突不是业务错误，是客户端问题，不需要打堆栈
		}
		return -1, errors.Wrapf(biz.ErrAddressNotCreated, "shop_user_address")
	}
	return reply.Id, nil
}

func (r *addressRepo) GetByUserId(ctx context.Context, userId int64) ([]*biz.Address, error) {
	result, err, _ := r.sg.Do(fmt.Sprintf("get_address_by_userid_%d", userId), func() (interface{}, error) {
		reply, err := r.data.uc.GetAddressByUserId(ctx, &userV1.GetAddressByUserIdRequest{
			UserId: wrapperspb.Int64(userId),
		})
		if err != nil {
			return nil, errors.Wrapf(biz.ErrAddressNotFound, "shop_user_address")
		}
		rv := make([]*biz.Address, 0, len(reply.Addresses))
		for _, x := range reply.Addresses {
			rv = append(rv, &biz.Address{
				Id:         x.Id,
				UserId:     x.UserId,
				UserName:   x.UserName,
				Tel:        x.Tel,
				Region:     x.Region,
				DetailInfo: x.DetailInfo,
				PostCode:   x.PostCode,
			})
		}
		return rv, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]*biz.Address), nil
}
