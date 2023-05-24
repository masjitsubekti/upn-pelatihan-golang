package master

import (
	"bytes"
	"fmt"

	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

var (
	matkulQuery = struct {
		Select string
		Count  string
	}{
		Select: `select id, kode_matkul, nama_matkul, sks, updated_at, created_at, created_by, updated_by, is_deleted from siamik_mata_kuliah `,
		Count:  `select count(id) from siamik_mata_kuliah `,
	}
)

type MataKuliahRepository interface {
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
}

type MataKuliahRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideMataKuliahRepositoryPostgreSQL(db *infras.PostgresqlConn) *MataKuliahRepositoryPostgreSQL {
	s := new(MataKuliahRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *MataKuliahRepositoryPostgreSQL) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	var searchParams []interface{}
	var searchRoleBuff bytes.Buffer
	searchRoleBuff.WriteString(" WHERE coalesce(is_deleted, false) = false ")

	if req.Keyword != "" {
		searchRoleBuff.WriteString(" AND ")
		searchRoleBuff.WriteString(" concat(kode_matkul, nama_matkul) ilike ? ")
		searchParams = append(searchParams, "%"+req.Keyword+"%")
	}

	query := r.DB.Read.Rebind(matkulQuery.Count + searchRoleBuff.String())
	fmt.Println("query", searchRoleBuff.String())

	var totalData int
	err = r.DB.Read.QueryRow(query, searchParams...).Scan(&totalData)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	fmt.Println(totalData)

	if totalData < 1 {
		data.Items = make([]interface{}, 0)
		return
	}

	searchRoleBuff.WriteString("order by " + ColumnMappMatkul[req.SortBy].(string) + " " + req.SortType + " ")

	offset := (req.PageNumber - 1) * req.PageSize
	searchRoleBuff.WriteString("limit ? offset ? ")
	searchParams = append(searchParams, req.PageSize)
	searchParams = append(searchParams, offset)

	searchMatkulQuery := searchRoleBuff.String()
	searchMatkulQuery = r.DB.Read.Rebind(matkulQuery.Select + searchMatkulQuery)
	fmt.Println("query", searchMatkulQuery)
	rows, err := r.DB.Read.Queryx(searchMatkulQuery, searchParams...)
	if err != nil {
		return
	}
	for rows.Next() {
		var items MataKuliah
		err = rows.StructScan(&items)
		if err != nil {
			return
		}

		data.Items = append(data.Items, items)
	}

	data.Meta = pagination.CreateMeta(totalData, req.PageSize, req.PageNumber)

	return
}
