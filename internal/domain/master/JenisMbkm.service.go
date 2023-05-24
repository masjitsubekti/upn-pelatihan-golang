package master

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

type JenisMbkmService interface {
	Create(reqFormat RequestJenisMbkmFormat, userID string) (newJenisMbkm JenisMbkm, err error)
	GetAllData() (data []JenisMbkm, err error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data JenisMbkm, err error)
	Update(reqFormat RequestJenisMbkmFormat, userID string) (data JenisMbkm, err error)
	DeleteByID(id string, userID string) error
}

type JenisMbkmServiceImpl struct {
	JenisMbkmRepository JenisMbkmRepository
	Config              *configs.Config
}

func ProvideJenisMbkmServiceImpl(repository JenisMbkmRepository, config *configs.Config) *JenisMbkmServiceImpl {
	s := new(JenisMbkmServiceImpl)
	s.JenisMbkmRepository = repository
	s.Config = config
	return s
}

func (s *JenisMbkmServiceImpl) Create(reqFormat RequestJenisMbkmFormat, userID string) (newJenisMbkm JenisMbkm, err error) {
	newJenisMbkm, _ = newJenisMbkm.NewJenisMbkmFormat(reqFormat, userID)
	err = s.JenisMbkmRepository.Create(newJenisMbkm)
	if err != nil {
		return JenisMbkm{}, err
	}
	return newJenisMbkm, nil
}

func (s *JenisMbkmServiceImpl) GetAllData() (data []JenisMbkm, err error) {
	return s.JenisMbkmRepository.GetAllData()
}

func (s *JenisMbkmServiceImpl) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	return s.JenisMbkmRepository.ResolveAll(req)
}

func (s *JenisMbkmServiceImpl) ResolveByID(id string) (data JenisMbkm, err error) {
	return s.JenisMbkmRepository.ResolveByID(id)
}

func (s *JenisMbkmServiceImpl) Update(reqFormat RequestJenisMbkmFormat, userID string) (data JenisMbkm, err error) {
	jenisMbkm, _ := data.NewJenisMbkmFormat(reqFormat, userID)
	err = s.JenisMbkmRepository.Update(jenisMbkm)
	if err != nil {
		log.Error().Msgf("service.UpdateJenisMbkm error", err)
	}
	return jenisMbkm, nil
}

func (s *JenisMbkmServiceImpl) DeleteByID(id string, userID string) error {
	jenisMbkm, err := s.JenisMbkmRepository.ResolveByID(id)

	if err != nil || (JenisMbkm{}) == jenisMbkm {
		return errors.New("Data Jenis MBKM dengan ID :" + id + " tidak ditemukan")
	}

	jenisMbkm.SoftDelete(userID)
	err = s.JenisMbkmRepository.Update(jenisMbkm)
	if err != nil {
		return errors.New("Ada kesalahan dalam menghapus data jenis mbkm dengan ID: " + id)
	}
	return nil
}
