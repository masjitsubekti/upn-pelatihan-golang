package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type JenisMitraHandler struct {
	JenisMitraService master.JenisMitraService
}

func ProvideJenisMitraHandler(service master.JenisMitraService) JenisMitraHandler {
	return JenisMitraHandler{
		JenisMitraService: service,
	}
}

func (h *JenisMitraHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/jenis-mitra", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middleware.VerifyToken)
			r.Post("/", h.Create)
			r.Put("/{id}", h.Update)
			r.Get("/", h.ResolveAll)
			r.Get("/all", h.GetAllData)
			r.Delete("/{id}", h.Delete)
			r.Get("/{id}", h.ResolveJenisMitraByID)
		})
	})
}

// createJenisMitra adalah untuk menambah data jenis mitra.
// @Summary menambahkan data jenis mitra.
// @Description Endpoint ini adalah untuk menambahkan data jenis mitra.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param jenis-mitra body master.RequestJenisMitraFormat true "Jenis Mitra yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.JenisMitra}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra [post]
func (h *JenisMitraHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.RequestJenisMitraFormat
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}
	userID, err := uuid.FromString(middleware.GetClaimsValue(r.Context(), "userId").(string))
	if err != nil {
		fmt.Print("error user id")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	newMitra, err := h.JenisMitraService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error response")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, newMitra)
}

// UpdateJenisMitra adalah untuk mengubah data JenisMitra.
// @Summary mengubah data JenisMitra
// @Description Endpoint ini adalah untuk mengubah data JenisMitra.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "ID"
// @Param jenis-mitra body master.RequestJenisMitraFormat true "Jenis Mitra yang akan diubah"
// @Success 200 {object} response.Base{data=[]master.JenisMitra}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra/{id} [put]
func (h *JenisMitraHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.FromString(middleware.GetClaimsValue(r.Context(), "userId").(string))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.WithError(w, err)
		return
	}
	var reqFormat master.RequestJenisMitraFormat
	err = json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	newMitra, err := h.JenisMitraService.Update(id, reqFormat, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, newMitra)
}

// ResolveAll list all jenis mitra.
// @Summary Get list all jenis mitra.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data jenis mitra sesuai dengan filter yang dikirimkan.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param q query string false "Keyword search"
// @Param pageSize query int true "Set pageSize data"
// @Param pageNumber query int true "Set page number"
// @Param sortBy query string false "Set sortBy parameter is one of [ nama_jenis_mitra ]"
// @Param sortType query string false "Set sortType with asc or desc"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra [get]
func (h *JenisMitraHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	fmt.Println("pageSizeStr", pageSizeStr)
	fmt.Println("pageNumberStr", pageNumberStr)
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "namaJenisMitra"
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

	status, err := h.JenisMitraService.ResolveAll(req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, status)
}

// GetDataAll list all jenis mitra.
// @Summary Get list all jenis mitra.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data jenis mitra sesuai dengan filter yang dikirimkan.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra/all [get]
func (h *JenisMitraHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	status, err := h.JenisMitraService.GetAllData()
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, status)
}

// deleteJenisMitra adalah untuk menghapus data jenis mitra.
// @Summary hapus data jenis mitra.
// @Description Endpoint ini adalah untuk menghapus data jenis mitra.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra/{id} [delete]
func (h *JenisMitraHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	newID, err := uuid.FromString(id)

	userID, err := uuid.FromString(middleware.GetClaimsValue(r.Context(), "userId").(string))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	err = h.JenisMitraService.DeleteByID(newID, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, "success")
}

// ResolveJenisMitraByID adalah untuk mendapatkan satu data Jenis Mitra berdasarkan ID.
// @Summary Mendapatkan satu data Jenis Mitra berdasarkan ID.
// @Description Endpoint ini adalah untuk mendapatkan Jenis Mitra By ID.
// @Tags jenis-mitra
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base{data=master.JenisMitra}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/jenis-mitra/{id} [get]
func (h *JenisMitraHandler) ResolveJenisMitraByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.WithError(w, err)
		return
	}
	jenisMitra, err := h.JenisMitraService.ResolveByID(ID)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, jenisMitra)
}
