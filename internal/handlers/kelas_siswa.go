package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"

	"github.com/go-chi/chi"
)

type KelasSiswaHandler struct {
	KelasSiswaService master.KelasSiswaService
}

func ProvideKelasSiswaHandler(service master.KelasSiswaService) KelasSiswaHandler {
	return KelasSiswaHandler{
		KelasSiswaService: service,
	}
}

func (h *KelasSiswaHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/kelas-siswa", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", h.ResolveAll)
			r.Post("/", h.Create)
			r.Put("/", h.Update)
			r.Get("/cek-siswa", h.ExistByIDSiswa)
			r.Get("/{id}", h.ResolveByIDDTO)
			r.Delete("/{id}", h.Delete)
		})
	})
}

// ResolveAll list all kelas siswa.
// @Summary Get list all kelas siswa.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data kelas siswa sesuai dengan filter yang dikirimkan.
// @Tags kelas-siswa
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
// @Router /v1/master/kelas-siswa [get]
func (h *KelasSiswaHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.KelasSiswaService.ResolveAll(req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// createSiswa adalah untuk menambah data kelas siswa.
// @Summary menambahkan data kelas siswa.
// @Description Endpoint ini adalah untuk menambahkan data kelas siswa.
// @Tags kelas-siswa
// @Produce json
// @Param kelasSiswa body master.KelasSiswaRequest true "Kelas Siswa yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.KelasSiswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa [post]
func (h *KelasSiswaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.KelasSiswaRequest
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.KelasSiswaService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}

// updateKelasSiswa adalah untuk menambah data kelas siswa.
// @Summary menambahkan data kelas siswa.
// @Description Endpoint ini adalah untuk menambahkan data kelas siswa.
// @Tags kelas-siswa
// @Produce json
// @Param kelasSiswa body master.KelasSiswaRequest true "Kelas Siswa yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.KelasSiswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa [put]
func (h *KelasSiswaHandler) Update(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.KelasSiswaRequest
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.KelasSiswaService.Update(reqFormat, userID)
	if err != nil {
		fmt.Print("error update")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}

// ExistByIDSiswa adalah untuk mendapatkan satu data kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Summary Mendapatkan satu data kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Description Endpoint ini adalah untuk mendapatkan kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Tags kelas-siswa
// @Produce json
// @Param idSiswa query string true "Set idSiswa"
// @Param idKelasSiswa query string true "Set idKelasSiswa"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa/cek-siswa [get]
func (h *KelasSiswaHandler) ExistByIDSiswa(w http.ResponseWriter, r *http.Request) {
	idSiswa := r.URL.Query().Get("idSiswa")
	idKelasSiswa := r.URL.Query().Get("idKelasSiswa")
	siswa, err := h.KelasSiswaService.ExistByIdSiswa(idSiswa, idKelasSiswa)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, siswa)
}

// ResolveByIDDTO adalah untuk mendapatkan satu data kelas siswa berdasarkan ID.
// @Summary Mendapatkan satu data kelas siswa berdasarkan ID.
// @Description Endpoint ini adalah untuk mendapatkan kelas siswa By ID.
// @Tags kelas-siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa/{id} [get]
func (h *KelasSiswaHandler) ResolveByIDDTO(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	data, err := h.KelasSiswaService.ResolveByIDDTO(ID)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}

// deleteKelasSiswa adalah untuk menghapus data kelas siswa.
// @Summary hapus data kelas siswa.
// @Description Endpoint ini adalah untuk menghapus data kelas siswa.
// @Tags kelas-siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa/{id} [delete]
func (h *KelasSiswaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	err := h.KelasSiswaService.DeleteByID(id, userID)
	if err != nil {
		fmt.Println(err)
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, "success")
}
