package service

import (
	"context"
	pb "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/user/internal/biz"
)

type UserServiceService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
}

func NewUserServiceService(uc *biz.UserUsecase) *UserServiceService {
	return &UserServiceService{
		uc: uc,
	}
}

func (s *UserServiceService) WxLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.WxLoginReply, error) {
	token, err := s.uc.WXLogin(ctx, req.Code, req.EncryptedData, req.Iv, req.SessionIsValid)
	if err != nil {
		return &pb.WxLoginReply{}, err
	}
	return &pb.WxLoginReply{
		Token: token,
	}, nil
}