package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/pkg/tokens"
	"time"
)

// User is a User model.
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

type UserCache struct {
	Token      string `json:"token"`
	OpenId     string `json:"open_id"`
	UserId     int64  `json:"user_id"`
	SessionKey string `json:"session_key"`
}

type Gender int64

const (
	unknown Gender = 0
	male    Gender = 1
	female  Gender = 2
)

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Code2Session(context.Context, string) (string, string, error)
	DecryptUserInfo(context.Context, string, string, string) (*User, error)
	Get(context.Context, string) (*UserCache, error)
	SetUserCache(context.Context, *UserCache, string, time.Duration) error
	DeleteOldToken(context.Context, string)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) WXLogin(ctx context.Context, code string, encryptedData string, iv string, sessionIsValid bool) (string, error) {
	openid, newSessionKey, err := u.repo.Code2Session(ctx, code)
	if err != nil {
		return "", err
	}

	sessionKeyForDecrypt := newSessionKey

	// 老的token
	var oldToken string
	if sessionIsValid {
		// 从缓存中拿老的session_key用来解密用户敏感数据
		userCache, err := u.repo.Get(ctx, openid)
		if err != nil {
			return "", err
		}
		sessionKeyForDecrypt = userCache.SessionKey
		oldToken = userCache.Token
	}

	userInfo, err := u.repo.DecryptUserInfo(ctx, sessionKeyForDecrypt, encryptedData, iv)
	if err != nil {
		// 解密不成功有可能是因为Redis缓存失效，不应该作为致命失败，只保存必要的openId等信息进数据库也可以的，应该让流程继续
		userInfo = &User{
			OpenId: openid,
		}
	}

	u.log.Infof("openid: %s, userInfo: %v", openid, userInfo)

	// 更新或新建用户数据到数据库
	retUser, err := u.repo.Save(ctx, userInfo)

	// 取得数据库中的userId
	userId := retUser.Id

	// 生成token
	tokenStr, err := tokens.GenerateToken(userId, 72*time.Hour, tokens.ScopeAuthentication)
	if err != nil {
		return "", err
	}

	// 缓存。由于微信最新的CheckSession状态是自动维护的，可以一直不过期，所以这里需要设置缓存永不过期0，新生成token时先把老的token记录删除
	if oldToken != "" {
		u.repo.DeleteOldToken(ctx, oldToken)
	}

	err = u.repo.SetUserCache(ctx, &UserCache{
		Token:      tokenStr,
		OpenId:     openid,
		UserId:     userId,
		SessionKey: newSessionKey,
	}, tokenStr, 0)

	if err != nil {
		return tokenStr, err
	}

	err = u.repo.SetUserCache(ctx, &UserCache{
		Token:      tokenStr,
		OpenId:     openid,
		UserId:     userId,
		SessionKey: newSessionKey,
	}, openid, 0)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

func (u *UserUsecase) CheckToken(ctx context.Context, token string) (*UserCache, error) {
	return u.repo.Get(ctx, token)
}

type Code2sessionRequest struct {
	Appid     string
	Secret    string
	GrantType string
}
