package master

import "gitlab.com/upn-belajar-go/configs"

type MahasiswaService interface{}

type MahasiswaServiceImpl struct {
	MahasiswaRepository MahasiswaRepository
	Config              *configs.Config
}

func ProvideMahasiswaServiceImpl(repository MahasiswaRepository, config *configs.Config) *MahasiswaServiceImpl {
	s := new(MahasiswaServiceImpl)
	s.MahasiswaRepository = repository
	s.Config = config
	return s
}
