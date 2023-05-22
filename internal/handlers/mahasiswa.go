package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type MahasiswaHandler struct {
	MahasiswaService master.MahasiswaService
}

func ProvideMahasiswaHandler(service master.MahasiswaService) MahasiswaHandler {
	return MahasiswaHandler{
		MahasiswaService: service,
	}
}

func (h *MahasiswaHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/mahasiswa", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			// r.Post("/", h.Create)
			// r.Post("/update", h.Update)
			r.Get("/all", h.GetAllData)
			// r.Get("/{id}", h.ResolveByID)
			// r.Put("/{id}", h.Update)
			// r.Delete("/{id}", h.Delete)
		})
	})
}

// GetDataAll list all.
// @Summary Get list all.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data sesuai dengan filter yang dikirimkan.
// @Tags mahasiswa
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/mahasiswa/all [get]
func (h *MahasiswaHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	// status, err := h.MahasiswaService.GetAllData()
	// if err != nil {
	// 	response.WithError(w, err)
	// 	return
	// }

	response.WithJSON(w, http.StatusOK, nil)
}
