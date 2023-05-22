package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/internal/domain/auth"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

// UserHandler the HTTP handler for User domain.
type UserHandler struct {
	UserService auth.UserService
	Config      *configs.Config
}

// ProvideUserHandler is the provider for this handler.
func ProvideUserHandler(userService auth.UserService, config *configs.Config) UserHandler {
	return UserHandler{
		UserService: userService,
		Config:      config,
	}
}

// Router sets up the router for this domain.
func (u *UserHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/login", u.Login)
	})
}

// Login sign in a user
// @Summary sign in a user
// @Description This endpoint sign in a user
// @Tags users
// @Param x-api-key header string false "Token api key"
// @Param users body auth.InputLogin true "The User to be sign in."
// @Produce json
// @Success 201 {object} response.Base{auth.ResponseLogin}
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/user/login [post]
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input auth.InputLogin
	// xApiKey := r.Header.Get("x-api-key")
	// cekToken := middleware.GetClaimsTokenApiKey(u.Config.App.TokenApiKey, xApiKey)
	// if !cekToken {
	// 	response.WithError(w, failure.Unauthorized("Token invalid."))
	// 	return
	// }

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err = shared.GetValidator().Struct(input)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	payloadAuth := auth.InputLogin{
		Username: input.Username,
		Password: input.Password,
	}

	resp, exist, err := u.UserService.Login(payloadAuth)
	if !exist {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	response.WithJSON(w, http.StatusOK, resp)
}
