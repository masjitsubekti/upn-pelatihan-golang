package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type MataKuliahHandler struct {
	MataKuliahService master.MataKuliahService
}

func ProvideMataKuliahHandler(service master.MataKuliahService) MataKuliahHandler {
	return MataKuliahHandler{
		MataKuliahService: service,
	}
}

func (h *MataKuliahHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/mata-kuliah", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			r.Get("/", h.ResolveAll)
		})
	})
}

// ResolveAll list all mata kuliah.
// @Summary Get list all mata kuliah.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data mata kuliah sesuai dengan filter yang dikirimkan.
// @Tags mata-kuliah
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param q query string false "Keyword search"
// @Param pageSize query int true "Set pageSize data"
// @Param pageNumber query int true "Set page number"
// @Param sortBy query string false "Set sortBy parameter is one of [ nama ]"
// @Param sortType query string false "Set sortType with asc or desc"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/mata-kuliah [get]
func (h *MataKuliahHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	fmt.Println("pageSizeStr", pageSizeStr)
	fmt.Println("pageNumberStr", pageNumberStr)
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "createdAt"
	}

	sortType := r.URL.Query().Get("sortType")
	if sortType == "" {
		sortType = "DESC"
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	req := model.StandardRequest{
		Keyword:    keyword,
		PageSize:   pageSize,
		PageNumber: pageNumber,
		SortBy:     sortBy,
		SortType:   sortType,
	}

	err = shared.GetValidator().Struct(req)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	data, err := h.MataKuliahService.ResolveAll(req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}
