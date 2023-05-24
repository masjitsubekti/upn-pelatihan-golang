package master

import (
	"gitlab.com/upn-belajar-go/configs"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

type MataKuliahService interface {
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
}

type MataKuliahServiceImpl struct {
	MataKuliahRepository MataKuliahRepository
	Config               *configs.Config
}

func ProvideMataKuliahServiceImpl(repository MataKuliahRepository, config *configs.Config) *MataKuliahServiceImpl {
	s := new(MataKuliahServiceImpl)
	s.MataKuliahRepository = repository
	s.Config = config
	return s
}

func (s *MataKuliahServiceImpl) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	return s.MataKuliahRepository.ResolveAll(req)
}
