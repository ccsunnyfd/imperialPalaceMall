package data

import (
	"context"
	pb "imperialPalaceMall/api/user/v1"
)

func (r *userRepo) Login(ctx context.Context, token, sessionKey string) error {
	r.data.wxLoginClient.WxLogin(ctx, &pb.WxLoginRequest{
		Code: "",
	})
	return nil
}