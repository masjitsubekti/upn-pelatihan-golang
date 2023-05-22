package master

import (
	"errors"
	"time"

	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

type KelasSiswaService interface {
	Create(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error)
	Update(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error)
	ExistByIdSiswa(idSiswa string, idKelasSiswa string) (exist bool, err error)
	ResolveByIDDTO(id string) (data KelasSiswaDTO, err error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	DeleteByID(id string, userID string) error
}

type KelasSiswaServiceImpl struct {
	KelasSiswaRepository KelasSiswaRepository
	Config               *configs.Config
}

func ProvideKelasSiswaServiceImpl(repository KelasSiswaRepository, config *configs.Config) *KelasSiswaServiceImpl {
	s := new(KelasSiswaServiceImpl)
	s.KelasSiswaRepository = repository
	s.Config = config
	return s
}

func (s *KelasSiswaServiceImpl) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	return s.KelasSiswaRepository.ResolveAll(req)
}

func (s *KelasSiswaServiceImpl) Create(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error) {
	newKelas, _ = newKelas.NewKelasSiswaFormat(reqFormat, userID)
	err = s.KelasSiswaRepository.Create(newKelas)
	if err != nil {
		return KelasSiswa{}, err
	}
	return newKelas, nil
}

func (s *KelasSiswaServiceImpl) Update(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error) {
	newKelas, _ = newKelas.NewKelasSiswaFormat(reqFormat, userID)
	err = s.KelasSiswaRepository.UpdateKelasSiswa(newKelas)
	if err != nil {
		return KelasSiswa{}, err
	}
	return newKelas, nil
}

func (s *KelasSiswaServiceImpl) ExistByIdSiswa(idSiswa string, idKelasSiswa string) (exist bool, err error) {
	exist, err = s.KelasSiswaRepository.ExistByIdSiswa(idSiswa, idKelasSiswa)
	if err != nil {
		return false, err
	}
	return
}

func (s *KelasSiswaServiceImpl) ResolveByIDDTO(id string) (data KelasSiswaDTO, err error) {
	return s.KelasSiswaRepository.ResolveByIDDTO(id)
}

func (s *KelasSiswaServiceImpl) DeleteByID(id string, userID string) error {
	kelas, err := s.KelasSiswaRepository.ResolveByID(id)

	if err != nil {
		return errors.New("Data kelas siswa dengan ID :" + id + " tidak ditemukan")
	}

	now := time.Now()
	kelas.IsDeleted = true
	kelas.UpdatedBy = &userID
	kelas.UpdatedAt = &now
	err = s.KelasSiswaRepository.Update(kelas)
	if err != nil {
		return errors.New("Ada kesalahan dalam menghapus data kelas siswa dengan ID: " + id)
	}

	return nil
}
