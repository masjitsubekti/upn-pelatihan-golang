package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/upn-belajar-go/internal/domain/orm"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type KelasHandler struct {
	KelasService orm.KelasService
}

func ProvideKelasHandler(kelasService orm.KelasService) KelasHandler {
	return KelasHandler{
		KelasService: kelasService,
	}
}

// Router untuk setup dari router untuk domain ini
func (h *KelasHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/kelas", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// Diaktifkan jika menggunakan authorization jwt
			// r.Use(middleware.VerifyToken)
			r.Get("/", h.ResolvePagination)
			r.Get("/all", h.ResolveAll)
			r.Get("/{id}", h.ResolveByID)
			r.Post("/", h.Create)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}

// ResolvePagination adalah endpoint untuk mendapatkan data semua master Kelas
// @Summary Mendapatkan semua master Kelas
// @Schemes
// @Description Endpoint ini digunakan untuk mendapatkan list semua master Kelas
// @Tags master/kelas
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param q query string false "Keyword search"
// @Param pageSize query int true "Set pageSize data"
// @Param pageNumber query int true "Set page number"
// @Param sortBy query string false "Set sortBy parameter diisi dengan salah satu dari [id, nama, keterangan]"
// @Param sortType query string false "Set sortType with asc or desc"
// @Success 200 {object} response.Base{data=[]orm.Kelas}
// @Router /v1/master/kelas [get]
func (h *KelasHandler) ResolvePagination(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageNumberStr := r.URL.Query().Get("pageNumber")
	sortBy := r.URL.Query().Get("sortBy")
	sortType := r.URL.Query().Get("sortType")

	if sortBy == "" {
		sortBy = "nama"
	}

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

	dataPage, err := h.KelasService.ResolvePagination(req)
	if err != nil {
		response.WithJSON(w, http.StatusInternalServerError, nil)
	}

	response.WithJSON(w, http.StatusOK, dataPage)
}

// ResolveAll adalah endpoint untuk mendapatkan data Kelas
// @Summary Mendapatkan data Kelas
// @Schemes
// @Description Endpoint ini digunakan untuk mendapatkan Kelas
// @Tags master/kelas
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Success 200 {object} response.Base{data=orm.Kelas}
// @Router /v1/master/kelas/all [get]
func (h *KelasHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.KelasService.ResolveAll()
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// ResolveByID adalah endpoint untuk mendapatkan data Kelas berdasarikan ID-nya
// @Summary Mendapatkan data Kelas berdasarikan ID-nya
// @Schemes
// @Description Endpoint ini digunakan untuk mendapatkan data Kelas berdasarikan ID-nya
// @Tags master/kelas
// @Accept json
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base{data=orm.Kelas}
// @Router /v1/master/kelas/{id} [get]
func (h *KelasHandler) ResolveByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.KelasService.ResolveByID(id)
	if err != nil {
		if err.Error() == model.RECORD_NOT_FOUND {
			response.WithError(w, failure.NotFound("Kelas tidak ditemukan"))
		} else {
			response.WithError(w, failure.BadRequest(err))
		}
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}

// Create adalah endpoint untuk menambahkan master Kelas
// @Summary Menambahkan Kelas baru
// @Schemes
// @Description Endpoint ini digunakan untuk menambahkan Kelas
// @Tags master/kelas
// @Accept json
// @Param Authorization header string false "Bearer <token>"
// @Produce json
// @Param role body orm.KelasRequest true "Kelas yang akan ditambahkan"
// @Success 200 {object} response.Base{data=orm.Kelas}
// @Router /v1/master/kelas [post]
func (h *KelasHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req orm.KelasRequest
	// Diaktifkan jika menggunakan authorization jwt
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err = shared.GetValidator().Struct(req)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	data, err := h.KelasService.Create(req, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// Update adalah endpoint untuk update master Kelas
// @Summary Update master Kelas
// @Schemes
// @Description Endpoint ini digunakan untuk update master Kelas
// @Tags master/kelas
// @Accept json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Produce json
// @Param role body orm.KelasRequest true "Kelas yang akan diupdate"
// @Success 200 {object} response.Base{data=orm.Kelas}
// @Router /v1/master/kelas/{id} [put]
func (h *KelasHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req orm.KelasRequest
	id := chi.URLParam(r, "id")
	// Diaktifkan jika menggunakan authorization jwt
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	err = shared.GetValidator().Struct(req)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	data, err := h.KelasService.Update(req, id, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// Delete adalah endpoint untuk hapus master kelas
// @Summary Hapus master kelas
// @Schemes
// @Description Endpoint ini digunakan untuk hapus master kelas
// @Tags master/kelas
// @Accept json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base{data=orm.Kelas}
// @Router /v1/master/kelas/{id} [delete]
func (h *KelasHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Diaktifkan jika menggunakan authorization jwt
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""

	err := h.KelasService.SoftDelete(id, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, "success")
}
