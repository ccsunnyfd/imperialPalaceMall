package weixin

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type WXBizDataCrypt struct {
	AppId      string
	SessionKey string
}

type WxData struct {
	OpenId    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int64     `json:"gender"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatarUrl"`
	UnionId   string    `json:"unionId"`
	WaterMark WaterMark `json:"watermark"`
}

type WaterMark struct {
	Timestamp int    `json:"timestamp"`
	AppId     string `json:"appid"`
}

func NewWXBizDataCrypt(appId, sessionKey string) *WXBizDataCrypt {
	return &WXBizDataCrypt{
		appId,
		sessionKey,
	}
}

func (crypt *WXBizDataCrypt) WithSessionKey(sessionKey string) *WXBizDataCrypt {
	crypt.SessionKey = sessionKey
	return crypt
}

func (crypt *WXBizDataCrypt) DecryptData(encryptedData string, iv string) (*WxData, error) {
	log.Infof("sessionKey: %s\n", crypt.SessionKey)
	sessionKey, _ := base64.StdEncoding.DecodeString(crypt.SessionKey)
	_iv, _ := base64.StdEncoding.DecodeString(iv)
	_encryptedData, _ := base64.StdEncoding.DecodeString(encryptedData)

	block, err := aes.NewCipher(sessionKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_encryptedData, _encryptedData)
	data := WxData{}
	err = json.Unmarshal(PKCS7UnPadding(_encryptedData), &data)
	if err != nil {
		return nil, err
	}

	if data.WaterMark.AppId != crypt.AppId {
		return nil, errors.New("Invalid Buffer")
	}

	return &data, nil

}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
