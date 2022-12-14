// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/user/internal/biz"
	"imperialPalaceMall/app/user/internal/conf"
	"imperialPalaceMall/app/user/internal/data"
	"imperialPalaceMall/app/user/internal/server"
	"imperialPalaceMall/app/user/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	client, cleanup := data.NewRedis(confData, logger)
	wxAuther, cleanup2 := data.NewWxLoginAuther(confData, logger)
	wxBizDataCrypt := data.NewWxDecrypter(confData)
	dataData, cleanup3, err := data.NewData(db, client, wxAuther, wxBizDataCrypt, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	addressRepo := data.NewAddressRepo(dataData, logger)
	addressUsecase := biz.NewAddressUsecase(addressRepo, logger)
	userServiceService := service.NewUserServiceService(userUsecase, addressUsecase)
	grpcServer := server.NewGRPCServer(confServer, userServiceService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
