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
	// JenisMbkmService interface and implementation
	master.ProvideJenisMbkmServiceImpl,
	wire.Bind(new(master.JenisMbkmService), new(*master.JenisMbkmServiceImpl)),
	// JenisMbkmRepository interface and implementation
	master.ProvideJenisMbkmRepositoryPostgreSQL,
	wire.Bind(new(master.JenisMbkmRepository), new(*master.JenisMbkmRepositoryPostgreSQL)),

	// MataKuliahService interface and implementation
	master.ProvideMataKuliahServiceImpl,
	wire.Bind(new(master.MataKuliahService), new(*master.MataKuliahServiceImpl)),
	// MataKuliahRepository interface and implementation
	master.ProvideMataKuliahRepositoryPostgreSQL,
	wire.Bind(new(master.MataKuliahRepository), new(*master.MataKuliahRepositoryPostgreSQL)),

	// PendaftarProgramMbkmService interface and implementation
	master.ProvidePendaftarProgramMbkmServiceImpl,
	wire.Bind(new(master.PendaftarProgramMbkmService), new(*master.PendaftarProgramMbkmServiceImpl)),
	// PendaftarProgramMbkmRepository interface and implementation
	master.ProvidePendaftarProgramMbkmRepositoryPostgreSQL,
	wire.Bind(new(master.PendaftarProgramMbkmRepository), new(*master.PendaftarProgramMbkmRepositoryPostgreSQL)),
)

// Wiring for all domains.
var domains = wire.NewSet(
	domainMaster,
)

// Wiring for HTTP routing.
var routing = wire.NewSet(
	wire.Struct(new(router.DomainHandlers), "*"),
	handlers.ProvideJenisMbkmHandler,
	handlers.ProvideMataKuliahHandler,
	handlers.ProvidePendaftarProgramMbkmHandler,
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
