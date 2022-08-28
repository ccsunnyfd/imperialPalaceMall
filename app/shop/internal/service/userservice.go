package service

import (
	"context"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/shop/internal/biz"
)

func (s *ShopInterface) WxLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.WxLoginReply, error) {
	token, err := s.uc.WXLogin(ctx, req.Code.Value, req.EncryptedData.Value, req.Iv.Value, req.SessionIsValid)
	if err != nil {
		return nil, err
	}
	return &pb.WxLoginReply{
		Token: token,
	}, nil
}

func (s *ShopInterface) CheckToken(ctx context.Context, token string) (*biz.UserCache, error) {
	return s.uc.CheckToken(ctx, token)
}