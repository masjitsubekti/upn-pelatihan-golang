//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/internal/handlers"
	"gitlab.com/upn-belajar-go/transport/http"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/router"
)

// Wiring for configurations.
var configurations = wire.NewSet(
	configs.Get,
)

// Wiring for persistences.
var persistences = wire.NewSet(
	infras.ProvidePostgreSQLConn,
)

// Wiring for domain Master
var domainMaster = wire.NewSet(
	// MahasiswaService interface and implementation
	master.ProvideMahasiswaServiceImpl,
	wire.Bind(new(master.MahasiswaService), new(*master.MahasiswaServiceImpl)),
	// MahasiswaRepository interface and implementation
	master.ProvideMahasiswaRepositoryPostgreSQL,
	wire.Bind(new(master.MahasiswaRepository), new(*master.MahasiswaRepositoryPostgreSQL)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainMaster,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	handlers.ProvideMahasiswaHandler,
	// jwt
	middleware.ProvideJWTMiddleware,
	router.ProvideRouter,
)

// Wiring for everything.
func InitializeService() *http.HTTP {
	wire.Build(
		// configurations
		configurations,
		// persistences
		persistences,
		// domains
		domains,
		// routing
		routing,
		// selected transport layer
		http.ProvideHTTP)
	return &http.HTTP{}
}
