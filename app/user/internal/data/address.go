package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lib/pq"
	pkgErrors "github.com/pkg/errors"
	"gorm.io/gorm"
	userPb "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/user/internal/biz"
	"time"
)

var _ biz.AddressRepo = (*addressRepo)(nil)

type addressRepo struct {
	data *Data
	log  *log.Helper
}

type Address struct {
	gorm.Model
	Id         int64 `gorm:"primaryKey"`
	UserId     int64
	UserName   string
	Tel        string
	Region     pq.StringArray `gorm:"type:text[]"`
	DetailInfo string
	PostCode   string

	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAddressRepo .
func NewAddressRepo(data *Data, logger log.Logger) biz.AddressRepo {
	return &addressRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *addressRepo) GetByDetails(ctx context.Context, addr *biz.Address) (*biz.Address, error) {
	var a Address
	result := r.data.db.WithContext(ctx).
		Where(map[string]interface{}{
			"user_id":     addr.UserId,
			"user_name":   addr.UserName,
			"tel":         addr.Tel,
			"detail_info": addr.DetailInfo,
			"region":      pq.Array(addr.Region),
		}).
		First(&a)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, userPb.ErrorAddressNotFound("address by details not found")
		}
		return nil, pkgErrors.Wrap(userPb.ErrorAddressGetError("get address by details error"), "user/address")
	}
	return &biz.Address{
		Id:         a.Id,
		UserId:     a.UserId,
		UserName:   a.UserName,
		Tel:        a.Tel,
		Region:     a.Region,
		DetailInfo: a.DetailInfo,
		PostCode:   a.PostCode,
	}, nil
}

func (r *addressRepo) Save(ctx context.Context, addr *biz.Address) (int64, error) {
	a := Address{
		UserId:     addr.UserId,
		UserName:   addr.UserName,
		Tel:        addr.Tel,
		Region:     addr.Region,
		DetailInfo: addr.DetailInfo,
		PostCode:   addr.PostCode,
	}
	result := r.data.db.WithContext(ctx).Create(&a)
	if result.Error != nil || result.RowsAffected < 1 {
		return -1, pkgErrors.Wrap(userPb.ErrorAddressCreateError("create address error"), "user/address")
	}

	return a.Id, nil
}

func (r *addressRepo) GetByUserId(ctx context.Context, userId int64) ([]*biz.Address, error) {
	var addrs []Address
	result := r.data.db.WithContext(ctx).Where("user_id = ?", userId).Find(&addrs).Order("id")
	if result.Error != nil {
		return nil, pkgErrors.Wrap(userPb.ErrorAddressGetError("get address by userid error"), "user/address")
	}

	rv := make([]*biz.Address, 0)
	for _, po := range addrs {
		rv = append(rv, &biz.Address{
			Id:         po.Id,
			UserId:     po.UserId,
			UserName:   po.UserName,
			Tel:        po.Tel,
			Region:     po.Region,
			DetailInfo: po.DetailInfo,
			PostCode:   po.PostCode,
		})
	}
	return rv, nil
}

func (r *addressRepo) Update(ctx context.Context, userId int64, address *biz.Address) (int64, error) {
	result := r.data.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Updates(Address{
			Id:         address.Id,
			UserName:   address.UserName,
			Tel:        address.Tel,
			Region:     address.Region,
			DetailInfo: address.DetailInfo,
			PostCode:   address.PostCode,
		})

	if result.Error != nil {
		return -1, pkgErrors.Wrap(userPb.ErrorAddressEditError("edit address error"), "user/address")
	}

	if result.RowsAffected < 1 {
		return -1, userPb.ErrorAddressNotFound("edit address not found") // ????????????????????????????????????????????????????????????????????????????????????
	}

	return address.Id, nil
}

func (r *addressRepo) Delete(ctx context.Context, userId int64, id int64) (int64, error) {
	a := Address{Id: id}
	result := r.data.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&a)

	if result.Error != nil {
		return -1, pkgErrors.Wrap(userPb.ErrorAddressDeleteError("delete address error"), "user/address")
	}
	
	if result.RowsAffected < 1 {
		return -1, userPb.ErrorAddressDeleteNotFound("delete address not found")
	}

	return result.RowsAffected, nil
}
