//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/asynccnu/ccnu-service/internal/biz"
	"github.com/asynccnu/ccnu-service/internal/conf"
	"github.com/asynccnu/ccnu-service/internal/data"
	"github.com/asynccnu/ccnu-service/internal/registry"
	"github.com/asynccnu/ccnu-service/internal/server"
	"github.com/asynccnu/ccnu-service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		registry.ProviderSet,
		wire.Bind(new(biz.UserRepo), new(*data.UserRepo)),
		newApp))
}
