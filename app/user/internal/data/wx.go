package data

import (
	"context"
	"github.com/pkg/errors"
	weixinPb "imperialPalaceMall/api/weixin/v1"
	"imperialPalaceMall/app/user/internal/biz"
)

type WxAuther struct {
	request          *biz.Code2sessionRequest
	weixinHTTPClient weixinPb.WeixinHTTPClient
}

func NewWxAuther(request *biz.Code2sessionRequest, weixinHTTPClient weixinPb.WeixinHTTPClient) *WxAuther {
	return &WxAuther{
		request:          request,
		weixinHTTPClient: weixinHTTPClient,
	}
}

func (r *userRepo) Code2Session(ctx context.Context, code string) (openid, sessionKey string, err error) {
	resp, err := r.data.wxLoginAuther.weixinHTTPClient.Code2Session(ctx, &weixinPb.Code2SessionRequest{
		Appid:     r.data.wxLoginAuther.request.Appid,
		Secret:    r.data.wxLoginAuther.request.Secret,
		JsCode:    code,
		GrantType: r.data.wxLoginAuther.request.GrantType,
	})

	if err != nil {
		return "", "", errors.Wrap(weixinPb.ErrorWxCode2sessionError("Code2Session_code_%s", code), "user")
	}

	return resp.Openid, resp.SessionKey, nil
}

func (r *userRepo) DecryptUserInfo(ctx context.Context, sessionKey string, encryptedData, iv string) (*biz.User, error) {
	decrypter := r.data.wxDecrypter.WithSessionKey(sessionKey)
	userInfo, err := decrypter.DecryptData(encryptedData, iv)
	if err != nil {
		return nil, errors.Wrap(weixinPb.ErrorWxDecryptDataError("DecryptUserInfo"), "user")
	}

	return &biz.User{
		OpenId:    userInfo.OpenId,
		NickName:  userInfo.NickName,
		Gender:    biz.Gender(userInfo.Gender),
		City:      userInfo.City,
		Province:  userInfo.Province,
		Country:   userInfo.Country,
		AvatarUrl: userInfo.AvatarUrl,
		UnionId:   userInfo.UnionId,
	}, nil
}
