package master

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
	"gitlab.com/upn-belajar-go/transport/http/response"
)

type SiswaService interface {
	Create(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error)
	GetAllData() (data []Siswa, err error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data Siswa, err error)
	Update(reqFormat RequestSiswaFormat, userID string) (data Siswa, err error)
	DeleteByID(id string, userID string) error
	UploadFile(w http.ResponseWriter, r *http.Request, formValue string, path_file string) (path string, err error)
	DeleteBerkas(path string) (err error)
}

type SiswaServiceImpl struct {
	SiswaRepository SiswaRepository
	Config          *configs.Config
}

func ProvideSiswaServiceImpl(repository SiswaRepository, config *configs.Config) *SiswaServiceImpl {
	s := new(SiswaServiceImpl)
	s.SiswaRepository = repository
	s.Config = config
	return s
}

func (s *SiswaServiceImpl) Create(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error) {
	newSiswa, _ = newSiswa.NewSiswaFormat(reqFormat, userID)
	err = s.SiswaRepository.Create(newSiswa)
	if err != nil {
		return Siswa{}, err
	}
	return newSiswa, nil
}

func (s *SiswaServiceImpl) GetAllData() (data []Siswa, err error) {
	return s.SiswaRepository.GetAllData()
}

func (s *SiswaServiceImpl) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	return s.SiswaRepository.ResolveAll(req)
}

func (s *SiswaServiceImpl) ResolveByID(id string) (data Siswa, err error) {
	return s.SiswaRepository.ResolveByID(id)
}

func (s *SiswaServiceImpl) Update(reqFormat RequestSiswaFormat, userID string) (data Siswa, err error) {
	siswa, _ := data.NewSiswaFormat(reqFormat, userID)
	err = s.SiswaRepository.Update(siswa)
	if err != nil {
		log.Error().Msgf("service.UpdateSiswa error", err)
	}
	return siswa, nil
}

func (s *SiswaServiceImpl) DeleteByID(id string, userID string) error {
	siswa, err := s.SiswaRepository.ResolveByID(id)

	if err != nil || (Siswa{}) == siswa {
		return errors.New("Data siswa dengan ID :" + id + " tidak ditemukan")
	}

	siswa.SoftDelete(userID)
	err = s.SiswaRepository.Update(siswa)
	if err != nil {
		return errors.New("Ada kesalahan dalam menghapus data siswa dengan ID: " + id)
	}

	// Delete file
	// s.DeleteBerkas(siswa.Berkas)
	return nil
}

func (s *SiswaServiceImpl) UploadFile(w http.ResponseWriter, r *http.Request, formValue string, path_file string) (path string, err error) {
	if err = r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadedFile, handler, err := r.FormFile(formValue)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	newID, _ := uuid.NewV4()
	filename := fmt.Sprintf("%s%s", "berkas_siswa_"+newID.String(), filepath.Ext(handler.Filename))
	dir := s.Config.App.File.Dir
	DokumenBerkasSiswaDir := s.Config.App.File.BerkasSiswa

	if path_file == "" {
		path = filepath.Join(DokumenBerkasSiswaDir, filename)
	} else {
		path = path_file
	}
	fileLocation := filepath.Join(dir, path)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("ERROR FILE:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err = io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ERROR COPY FILE:", err)
		return
	}
	return
}

func (s *SiswaServiceImpl) DeleteBerkas(path string) (err error) {
	dir := s.Config.App.File.Dir
	DokumenBerkasDir := path
	fileLocation := filepath.Join(dir, DokumenBerkasDir)
	err = os.Remove(fileLocation)
	return
}
