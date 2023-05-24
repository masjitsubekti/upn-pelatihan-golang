package handlers

import (
	"encoding/json"
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

type JenisMbkmHandler struct {
	JenisMbkmService master.JenisMbkmService
}

func ProvideJenisMbkmHandler(service master.JenisMbkmService) JenisMbkmHandler {
	return JenisMbkmHandler{
		JenisMbkmService: service,
	}
}

func (h *JenisMbkmHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/jenis-mbkm", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			r.Post("/", h.Create)
			r.Get("/", h.ResolveAll)
			r.Get("/all", h.GetAllData)
			r.Get("/{id}", h.ResolveByID)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}

// createJenisMbkm adalah untuk menambah data jenis mbkm.
// @Summary menambahkan data jenis mbkm.
// @Description Endpoint ini adalah untuk menambahkan data jenis mbkm.
// @Tags jenis-mbkm
// @Produce json
// @Param jenisMbkm body master.RequestJenisMbkmFormat true "Jenis mbkm yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.JenisMbkm}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mbkm [post]
func (h *JenisMbkmHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.RequestJenisMbkmFormat
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.JenisMbkmService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}

// GetDataAll list all.
// @Summary Get list all.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data sesuai dengan filter yang dikirimkan.
// @Tags jenis-mbkm
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mbkm/all [get]
func (h *JenisMbkmHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	data, err := h.JenisMbkmService.GetAllData()
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// ResolveAll list all jenis mbkm.
// @Summary Get list all jenis mbkm.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data jenis mbkm sesuai dengan filter yang dikirimkan.
// @Tags jenis-mbkm
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
// @Router /v1/master/jenis-mbkm [get]
func (h *JenisMbkmHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.JenisMbkmService.ResolveAll(req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// ResolveByID adalah untuk mendapatkan satu data jenis mbkm berdasarkan ID.
// @Summary Mendapatkan satu data jenis mbkm berdasarkan ID.
// @Description Endpoint ini adalah untuk mendapatkan jenis mbkm By ID.
// @Tags jenis-mbkm
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base{data=master.JenisMbkm}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mbkm/{id} [get]
func (h *JenisMbkmHandler) ResolveByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	data, err := h.JenisMbkmService.ResolveByID(ID)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}

// UpdateJenisMbkm adalah untuk mengubah data jenis mbkm.
// @Summary mengubah data jenis mbkm
// @Description Endpoint ini adalah untuk mengubah data jenis mbkm.
// @Tags jenis-mbkm
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Param siswa body master.RequestJenisMbkmFormat true "Jenis MBKM yang akan diubah"
// @Success 200 {object} response.Base{data=master.JenisMbkm}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mbkm/{id} [put]
func (h *JenisMbkmHandler) Update(w http.ResponseWriter, r *http.Request) {
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	id := chi.URLParam(r, "id")
	var reqFormat master.RequestJenisMbkmFormat
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	reqFormat.ID = id
	siswa, err := h.JenisMbkmService.Update(reqFormat, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, siswa)
}

// deleteJenisMbkm adalah untuk menghapus data JenisMbkm.
// @Summary hapus data JenisMbkm.
// @Description Endpoint ini adalah untuk menghapus data JenisMbkm.
// @Tags jenis-mbkm
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mbkm/{id} [delete]
func (h *JenisMbkmHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	err := h.JenisMbkmService.DeleteByID(id, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, "success")
}
