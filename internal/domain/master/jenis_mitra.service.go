package master

import (
	"errors"

	"gitlab.com/upn-belajar-go/shared/model"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

type JenisMitraService interface {
	Create(reqFormat RequestJenisMitraFormat, userID uuid.UUID) (newData JenisMitra, error error)
	Update(id uuid.UUID, newData RequestJenisMitraFormat, userID uuid.UUID) (data JenisMitra, err error)
	GetAllData() ([]JenisMitra, error)
	ResolveAll(request model.StandardRequest) (orders pagination.Response, err error)
	DeleteByID(id uuid.UUID, userID uuid.UUID) error
	ResolveByID(id uuid.UUID) (data JenisMitra, err error)
}

type JenisMitraServiceImpl struct {
	JenisMitraRepository JenisMitraRepository
	Config               *configs.Config
}

func ProvideJenisMitraServiceImpl(repository JenisMitraRepository, config *configs.Config) *JenisMitraServiceImpl {
	s := new(JenisMitraServiceImpl)
	s.JenisMitraRepository = repository
	s.Config = config
	return s
}

func (s *JenisMitraServiceImpl) Create(reqFormat RequestJenisMitraFormat, userID uuid.UUID) (newJenisMitra JenisMitra, err error) {
	exist, err := s.JenisMitraRepository.ExistByNama(reqFormat.NamaJenisMitra)
	if exist {
		x := errors.New("Nama Jenis Mitra sudah dipakai")
		return JenisMitra{}, x
	}
	if err != nil {
		return JenisMitra{}, err
	}
	newJenisMitra, _ = newJenisMitra.NewJenisMitraFormat(reqFormat, userID)
	err = s.JenisMitraRepository.Create(newJenisMitra)
	if err != nil {
		return JenisMitra{}, err
	}
	return newJenisMitra, nil
}

func (s *JenisMitraServiceImpl) Update(id uuid.UUID, reqFormat RequestJenisMitraFormat, userID uuid.UUID) (data JenisMitra, err error) {
	exist, err := s.JenisMitraRepository.ExistByNamaID(reqFormat.ID, reqFormat.NamaJenisMitra)
	if exist {
		x := errors.New("Nama Jenis Mitra sudah dipakai")
		return JenisMitra{}, x
	}
	if err != nil {
		return JenisMitra{}, err
	}
	np, _ := data.NewJenisMitraFormat(reqFormat, userID)
	err = s.JenisMitraRepository.Update(np)
	if err != nil {
		log.Error().Msgf("service.UpdateJenisMitra error", err)
	}
	return np, nil
}

func (s *JenisMitraServiceImpl) GetAllData() (data []JenisMitra, err error) {
	return s.JenisMitraRepository.GetAllData()
}

func (s *JenisMitraServiceImpl) ResolveAll(request model.StandardRequest) (orders pagination.Response, err error) {
	return s.JenisMitraRepository.ResolveAll(request)
}

func (s *JenisMitraServiceImpl) ResolveByID(id uuid.UUID) (jenisMitra JenisMitra, err error) {
	return s.JenisMitraRepository.ResolveByID(id)
}

func (s *JenisMitraServiceImpl) DeleteByID(id uuid.UUID, userID uuid.UUID) error {
	jenismitra, err := s.JenisMitraRepository.ResolveByID(id)

	if err != nil || (JenisMitra{}) == jenismitra {
		return errors.New("Data Jenis Mitra dengan ID :" + id.String() + " tidak ditemukan")
	}

	jenismitra.SoftDelete(userID)
	err = s.JenisMitraRepository.Update(jenismitra)
	if err != nil {
		return errors.New("Ada kesalahan dalam menghapus data Jenis Mitra dengan ID: " + id.String())
	}
	return nil
}
