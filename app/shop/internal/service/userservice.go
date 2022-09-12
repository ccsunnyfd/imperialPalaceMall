package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	pkgErrors "github.com/pkg/errors"
	pb "imperialPalaceMall/api/shop/interface/v1"
	"imperialPalaceMall/app/pkg/middleware"
	"imperialPalaceMall/app/shop/internal/biz"
)

func (s *ShopInterface) WxLogin(ctx context.Context, req *pb.WxLoginRequest) (*pb.WxLoginReply, error) {
	token, err := s.uc.WXLogin(ctx, req.Code, req.EncryptedData, req.Iv, req.SessionIsValid)
	if err != nil {
		return nil, err
	}
	return &pb.WxLoginReply{
		Token: token,
	}, nil
}

func (s *ShopInterface) SaveAddress(ctx context.Context, req *pb.SaveAddressRequest) (*pb.SaveAddressReply, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	createdId, err := s.ac.CreateAddress(ctx, userId, &biz.Address{
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
		Id: createdId,
	}, nil
}

func (s *ShopInterface) GetAddress(ctx context.Context, req *pb.GetAddressRequest) (*pb.GetAddressReply, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	po, err := s.ac.GetAddressByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}

	addrs := make([]*pb.Address, 0, len(po))
	for _, x := range po {
		addrs = append(addrs, &pb.Address{
			Id:         x.Id,
			UserId:     x.UserId,
			UserName:   x.UserName,
			Tel:        x.Tel,
			Region:     x.Region,
			DetailInfo: x.DetailInfo,
			PostCode:   x.PostCode,
		})
	}

	return &pb.GetAddressReply{
		Addresses: addrs,
	}, nil
}

func (s *ShopInterface) UpdateAddress(ctx context.Context, req *pb.UpdateAddressRequest) (*pb.UpdateAddressReply, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := s.ac.UpdateAddress(ctx, userId, &biz.Address{
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
		Id: id,
	}, nil
}

func (s *ShopInterface) DeleteAddress(ctx context.Context, req *pb.DeleteAddressRequest) (*pb.DeleteAddressReply, error) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	affected, err := s.ac.DeleteAddress(ctx, userId, req.Id.Value)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAddressReply{
		Affected: affected,
	}, nil
}

func (s *ShopInterface) CheckToken(ctx context.Context, token string) (*biz.UserCache, error) {
	return s.uc.CheckToken(ctx, token)
}

func getUserIdFromContext(ctx context.Context) (int64, error) {
	ctxUser, err := middleware.ContextGetUser(ctx)
	if err != nil {
		return -1, pkgErrors.Wrap(errors.Unauthorized("user", "missing user value in request context"), "cartService_GetMyCart_error")
	}
	return ctxUser.UserId, nil
}
