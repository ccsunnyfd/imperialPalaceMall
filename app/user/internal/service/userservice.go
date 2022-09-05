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
	token, err := s.uc.WXLogin(ctx, req.Code.Value, req.EncryptedData.Value, req.Iv.Value, req.SessionIsValid)
	if err != nil {
		return nil, err
	}
	return &pb.WxLoginReply{
		Token: token,
	}, nil
}

func (s *UserServiceService) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenReply, error) {
	user, err := s.uc.CheckToken(ctx, req.Token.Value)
	if err != nil {
		return nil, err
	}
	return &pb.CheckTokenReply{
		User: &pb.CheckTokenReply_UserRet{
			UserId: user.UserId,
			OpenId: user.OpenId,
		},
	}, nil
}
