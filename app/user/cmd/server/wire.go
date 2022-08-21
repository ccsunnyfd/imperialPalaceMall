//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"imperialPalaceMall/app/user/internal/biz"
	"imperialPalaceMall/app/user/internal/conf"
	"imperialPalaceMall/app/user/internal/data"
	"imperialPalaceMall/app/user/internal/server"
	"imperialPalaceMall/app/user/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
