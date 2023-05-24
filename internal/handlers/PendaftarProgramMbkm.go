package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type PendaftarProgramMbkmHandler struct {
	PendaftarProgramMbkmService master.PendaftarProgramMbkmService
}

func ProvidePendaftarProgramMbkmHandler(service master.PendaftarProgramMbkmService) PendaftarProgramMbkmHandler {
	return PendaftarProgramMbkmHandler{
		PendaftarProgramMbkmService: service,
	}
}

func (h *PendaftarProgramMbkmHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/mbkm/pendaftar-program-mbkm", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			r.Post("/", h.Create)
		})
	})
}

// create adalah untuk menambah data pendaftar program mbkm.
// @Summary menambahkan data pendaftar program mbkm.
// @Description Endpoint ini adalah untuk menambahkan data pendaftar program mbkm.
// @Tags pendaftar-program-mbkm
// @Produce json
// @Param PendaftarProgramMbkm body master.PendaftarProgramMbkmRequest true "Pendaftar program mbkm yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.PendaftarProgramMbkm}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/mbkm/pendaftar-program-mbkm [post]
func (h *PendaftarProgramMbkmHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.PendaftarProgramMbkmRequest
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.PendaftarProgramMbkmService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}
