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

func getUserIdFromContext(ctx context.Context) (int64, error) {
	ctxUser, err := middleware.ContextGetUser(ctx)
	if err != nil {
		return -1, pkgErrors.Wrap(errors.Unauthorized("user", "missing user value in request context"), "cartService_GetMyCart_error")
	}
	return ctxUser.UserId, nil
}
