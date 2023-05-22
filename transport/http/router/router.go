package router

import (
	"gitlab.com/upn-belajar-go/internal/handlers"
	"gitlab.com/upn-belajar-go/transport/http/middleware"

	"github.com/go-chi/chi"
)

// DomainHandlers is a struct that contains all domain-specific handlers.
type DomainHandlers struct {
	JenisMitraHandler handlers.JenisMitraHandler
	SiswaHandler      handlers.SiswaHandler
	KelasHandler      handlers.KelasHandler
	KelasSiswaHandler handlers.KelasSiswaHandler
	FileHandler       handlers.FileHandler
	UserHandler       handlers.UserHandler
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
		r.DomainHandlers.JenisMitraHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.SiswaHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.KelasHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.KelasSiswaHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.FileHandler.Router(rc, r.JwtMiddleware)
		r.DomainHandlers.UserHandler.Router(rc, r.JwtMiddleware)
	})
}
