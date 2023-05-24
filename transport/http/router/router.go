package router

import (
	"gitlab.com/upn-belajar-go/internal/handlers"
	"gitlab.com/upn-belajar-go/transport/http/middleware"

	"github.com/go-chi/chi"
)

// DomainHandlers is a struct that contains all domain-specific handlers.
type DomainHandlers struct {
	JenisMbkmHandler            handlers.JenisMbkmHandler
	MataKuliahHandler           handlers.MataKuliahHandler
	PendaftarProgramMbkmHandler handlers.PendaftarProgramMbkmHandler
}

// Router is the router struct containing handlers.
type Router struct {
	JwtMiddleware  *middleware.JWT
	DomainHandlers DomainHandlers
}

// ProvideRouter is the provider function for this router.
func ProvideRouter(domainHandlers DomainHandlers, jwtMiddleware *middleware.JWT) Router {
	return Router{
		DomainHandlers: domainHandlers,
		JwtMiddleware:  jwtMiddleware,
	}
}

// SetupRoutes sets up all routing for this server.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v1", func(rc chi.Router) {
		r.DomainHandlers.JenisMbkmHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.MataKuliahHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.PendaftarProgramMbkmHandler.Router(rc, r.JwtMiddleware)
	})
}
