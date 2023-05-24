package master

import (
	"gitlab.com/upn-belajar-go/configs"
)

type PendaftarProgramMbkmService interface {
	Create(reqFormat PendaftarProgramMbkmRequest, userID string) (newPendaftar PendaftarProgramMbkm, err error)
}

type PendaftarProgramMbkmServiceImpl struct {
	PendaftarProgramMbkmRepository PendaftarProgramMbkmRepository
	Config                         *configs.Config
}

func ProvidePendaftarProgramMbkmServiceImpl(repository PendaftarProgramMbkmRepository, config *configs.Config) *PendaftarProgramMbkmServiceImpl {
	s := new(PendaftarProgramMbkmServiceImpl)
	s.PendaftarProgramMbkmRepository = repository
	s.Config = config
	return s
}

func (s *PendaftarProgramMbkmServiceImpl) Create(reqFormat PendaftarProgramMbkmRequest, userID string) (newPendaftar PendaftarProgramMbkm, err error) {
	newPendaftar, _ = newPendaftar.NewPendaftarProgramMbkmFormat(reqFormat, userID)
	err = s.PendaftarProgramMbkmRepository.Create(newPendaftar)
	if err != nil {
		return PendaftarProgramMbkm{}, err
	}
	return newPendaftar, nil
}
