// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"imperialPalaceMall/app/mall/internal/biz"
	"imperialPalaceMall/app/mall/internal/conf"
	"imperialPalaceMall/app/mall/internal/data"
	"imperialPalaceMall/app/mall/internal/server"
	"imperialPalaceMall/app/mall/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client, cleanup := data.NewDB(confData, logger)
	dataData, cleanup2, err := data.NewData(client, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	categoryRepo := data.NewCategoryRepo(dataData, logger)
	categoryUsecase := biz.NewCategoryUsecase(categoryRepo, logger)
	categoryServiceService := service.NewCategoryServiceService(categoryUsecase)
	goodsRepo := data.NewGoodsRepo(dataData, logger)
	goodsUsecase := biz.NewGoodsUsecase(goodsRepo, logger)
	goodsServiceService := service.NewGoodsServiceService(goodsUsecase)
	httpServer := server.NewHTTPServer(confServer, categoryServiceService, goodsServiceService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
