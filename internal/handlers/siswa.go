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

type SiswaHandler struct {
	SiswaService master.SiswaService
}

func ProvideSiswaHandler(service master.SiswaService) SiswaHandler {
	return SiswaHandler{
		SiswaService: service,
	}
}

func (h *SiswaHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/siswa", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			r.Post("/", h.Create)
			r.Post("/update", h.UpdateSiswa)
			r.Get("/all", h.GetAllData)
			r.Get("/", h.ResolveAll)
			r.Get("/{id}", h.ResolveByID)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})
	})
}

// createSiswa adalah untuk menambah data siswa.
// @Summary menambahkan data siswa.
// @Description Endpoint ini adalah untuk menambahkan data siswa.
// @Tags siswa
// @Produce json
// @Param nama formData string true "Nama siswa"
// @Param kelas formData string false "Kelas"
// @Param berkas formData file false "Berkas"
// @Success 200 {object} response.Base{data=master.Siswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa [post]
func (h *SiswaHandler) Create(w http.ResponseWriter, r *http.Request) {
	// @Param siswa body master.RequestSiswaFormat true "Siswa yang akan ditambahkan"
	// var reqFormat master.RequestSiswaFormat
	// err := json.NewDecoder(r.Body).Decode(&reqFormat)
	// if err != nil {
	// 	fmt.Print("error jsondecoder")
	// 	response.WithError(w, failure.BadRequest(err))
	// 	return
	// }

	nama := r.FormValue("nama")
	kelas := r.FormValue("kelas")
	idKelas := r.FormValue("idKelas")

	var path string
	path, err := h.SiswaService.UploadFile(w, r, "berkas", "")
	if err != nil {
		fmt.Println("ERROR:", err)
		response.WithError(w, failure.BadRequest(err))
		return
	}

	var reqFormat = master.RequestSiswaFormat{
		Nama:    nama,
		Kelas:   kelas,
		Berkas:  path,
		IdKelas: idKelas,
	}

	// Validasi required
	err = shared.GetValidator().Struct(reqFormat)
	if err != nil {
		response.WithStatusMessage(w, http.StatusCreated, false, "Nama, Kelas Wajib diisi")
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.SiswaService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}

// GetDataAll list all siswa.
// @Summary Get list all siswa.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data siswa sesuai dengan filter yang dikirimkan.
// @Tags siswa
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa/all [get]
func (h *SiswaHandler) GetAllData(w http.ResponseWriter, r *http.Request) {
	status, err := h.SiswaService.GetAllData()
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, status)
}

// ResolveAll list all siswa.
// @Summary Get list all siswa.
// @Description endpoint ini digunakan untuk mendapatkan seluruh data siswa sesuai dengan filter yang dikirimkan.
// @Tags siswa
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
// @Router /v1/master/siswa [get]
func (h *SiswaHandler) ResolveAll(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.SiswaService.ResolveAll(req)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, data)
}

// ResolveByID adalah untuk mendapatkan satu data siswa berdasarkan ID.
// @Summary Mendapatkan satu data siswa berdasarkan ID.
// @Description Endpoint ini adalah untuk mendapatkan siswa By ID.
// @Tags siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base{data=master.Siswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa/{id} [get]
func (h *SiswaHandler) ResolveByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	data, err := h.SiswaService.ResolveByID(ID)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, data)
}

// UpdateSiswa adalah untuk mengubah data Siswa.
// @Summary mengubah data Siswa
// @Description Endpoint ini adalah untuk mengubah data Siswa.
// @Tags siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Param siswa body master.RequestSiswaFormat true "Siswa yang akan diubah"
// @Success 200 {object} response.Base{data=master.Siswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa/{id} [put]
func (h *SiswaHandler) Update(w http.ResponseWriter, r *http.Request) {
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	id := chi.URLParam(r, "id")
	var reqFormat master.RequestSiswaFormat
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	reqFormat.ID = id
	siswa, err := h.SiswaService.Update(reqFormat, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, siswa)
}

// deleteSiswa adalah untuk menghapus data siswa.
// @Summary hapus data siswa.
// @Description Endpoint ini adalah untuk menghapus data siswa.
// @Tags siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id path string true "ID"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa/{id} [delete]
func (h *SiswaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	err := h.SiswaService.DeleteByID(id, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, "success")
}

// UpdateSiswa adalah untuk mengubah data Siswa.
// @Summary mengubah data Siswa
// @Description Endpoint ini adalah untuk mengubah data Siswa.
// @Tags siswa
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param id formData string true "ID"
// @Param nama formData string true "Nama siswa"
// @Param kelas formData string false "Kelas"
// @Param idKelas formData string false "ID Kelas"
// @Param berkas formData file false "Berkas"
// @Success 200 {object} response.Base{data=master.Siswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa/update [post]
func (h *SiswaHandler) UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""

	id := r.FormValue("id")
	nama := r.FormValue("nama")
	kelas := r.FormValue("kelas")
	idKelas := r.FormValue("idKelas")

	getSiswa, err := h.SiswaService.ResolveByID(id)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	var reqFormat = master.RequestSiswaFormat{
		ID:      id,
		Nama:    nama,
		Kelas:   kelas,
		IdKelas: idKelas,
	}

	uploadedFile, _, _ := r.FormFile("berkas")
	var path string
	fmt.Println("image path:", getSiswa.Berkas)

	fmt.Println("berkas", uploadedFile)
	if uploadedFile != nil {
		filepath, err := h.SiswaService.UploadFile(w, r, "berkas", "")
		if err != nil {
			response.WithError(w, failure.BadRequest(err))
			return
		}
		path = filepath
		reqFormat.Berkas = path

		// Delete berkas
		if getSiswa.Berkas != nil {
			h.SiswaService.DeleteBerkas(*getSiswa.Berkas)
		}
	} else {
		reqFormat.Berkas = *getSiswa.Berkas
	}

	siswa, err := h.SiswaService.Update(reqFormat, userID)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusOK, siswa)
}
