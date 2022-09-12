package service

import (
	"context"
	pb "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/user/internal/biz"
)

type UserServiceService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUsecase
	ac *biz.AddressUsecase
}

func NewUserServiceService(uc *biz.UserUsecase, ac *biz.AddressUsecase) *UserServiceService {
	return &UserServiceService{
		uc: uc,
		ac: ac,
	}
}

func (s *UserServiceService) WxLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.WxLoginReply, error) {
	token, err := s.uc.WXLogin(ctx, req.Code, req.EncryptedData, req.Iv, req.SessionIsValid)
	if err != nil {
		return nil, err
	}
	return &pb.WxLoginReply{
		Token: token,
	}, nil
}

func (s *UserServiceService) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenReply, error) {
	user, err := s.uc.CheckToken(ctx, req.Token)
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

func (s *UserServiceService) GetAddressByUserId(ctx context.Context, req *pb.GetAddressByUserIdRequest) (*pb.GetAddressByUserIdReply, error) {
	addrs, err := s.ac.GetAddressByUserId(ctx, req.UserId.Value)
	if err != nil {
		return nil, err
	}

	rv := make([]*pb.Address, 0)
	for _, po := range addrs {
		rv = append(rv, &pb.Address{
			Id:         po.Id,
			UserId:     po.UserId,
			UserName:   po.UserName,
			Tel:        po.Tel,
			Region:     po.Region,
			DetailInfo: po.DetailInfo,
			PostCode:   po.PostCode,
		})
	}
	return &pb.GetAddressByUserIdReply{
		Addresses: rv,
	}, nil
}

func (s *UserServiceService) SaveAddress(ctx context.Context, req *pb.SaveAddressRequest) (*pb.SaveAddressReply, error) {
	id, err := s.ac.SaveAddress(ctx, &biz.Address{
		UserId:     req.UserId.Value,
		UserName:   req.UserName,
		Tel:        req.Tel,
		Region:     req.Region,
		DetailInfo: req.DetailInfo,
		PostCode:   req.PostCode,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SaveAddressReply{
		Id: id,
	}, nil
}

func (s *UserServiceService) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressReply, error) {
	affected, err := s.ac.UpdateAddress(ctx, req.UserId.Value,
		&biz.Address{
			Id:         req.Id.Value,
			UserName:   req.UserName,
			Tel:        req.Tel,
			Region:     req.Region,
			DetailInfo: req.DetailInfo,
			PostCode:   req.PostCode,
		})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAddressReply{
		Affected: affected,
	}, nil
}
