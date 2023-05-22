package master

import "gitlab.com/upn-belajar-go/infras"

type MahasiswaRepository interface{}

type MahasiswaRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideMahasiswaRepositoryPostgreSQL(db *infras.PostgresqlConn) *MahasiswaRepositoryPostgreSQL {
	s := new(MahasiswaRepositoryPostgreSQL)
	s.DB = db
	return s
}
