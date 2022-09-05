package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	pkgErrors "github.com/pkg/errors"
	"gorm.io/gorm"
	"imperialPalaceMall/app/user/internal/biz"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Id        int64 `gorm:"primaryKey"`
	OpenId    string
	NickName  string
	Gender    biz.Gender
	City      string
	Province  string
	Country   string
	AvatarUrl string
	UnionId   string
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{}
	result := r.data.db.WithContext(ctx).Where("open_id = ?", user.OpenId).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 新建
			u = User{
				OpenId:    user.OpenId,
				NickName:  user.NickName,
				Gender:    user.Gender,
				City:      user.City,
				Province:  user.Province,
				Country:   user.Country,
				AvatarUrl: user.AvatarUrl,
				UnionId:   user.UnionId,
			}
			result = r.data.db.WithContext(ctx).Create(&u)
			if result.Error != nil {
				return nil, pkgErrors.Wrapf(biz.ErrUserCreate, "Save: create_user")
			}
			return &biz.User{
				Id: u.Id,
			}, nil
		}
		return nil, pkgErrors.Wrapf(biz.ErrUserGet, "Save: get_by_open_id_%s", user.OpenId)
	}

	u.NickName = user.NickName
	u.Gender = user.Gender
	u.AvatarUrl = user.AvatarUrl
	u.City = user.City
	u.Country = user.Country
	u.Province = user.Province
	u.UnionId = user.UnionId

	result = r.data.db.WithContext(ctx).Save(&u)
	if result.Error != nil {
		return nil, pkgErrors.Wrapf(biz.ErrUserUpdate, "Save: update")
	}

	return &biz.User{
		Id: u.Id,
	}, nil
}
