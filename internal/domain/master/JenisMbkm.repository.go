package master

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

var (
	jenisMbkmQuery = struct {
		Select string
		Insert string
		Update string
		Delete string
		Exist  string
		Count  string
	}{
		Select: `select id, nama, updated_at, created_at, created_by, updated_by, is_deleted from m_jenis_mbkm `,
		Insert: `INSERT INTO m_jenis_mbkm (id, nama, created_by, created_at) values(:id, :nama, :created_by, :created_at) `,
		Update: `UPDATE m_jenis_mbkm SET 
				nama=:nama,
				updated_at=:updated_at,
				updated_by=:updated_by, 
				is_deleted=:is_deleted`,
		Delete: `delete from m_jenis_mbkm `,
		Exist:  `select count(id)>0 from m_jenis_mbkm `,
		Count:  `select count(id) from m_jenis_mbkm `,
	}
)

type JenisMbkmRepository interface {
	Create(data JenisMbkm) error
	GetAllData() (data []JenisMbkm, errr error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data JenisMbkm, err error)
	Update(data JenisMbkm) error
}

type JenisMbkmRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideJenisMbkmRepositoryPostgreSQL(db *infras.PostgresqlConn) *JenisMbkmRepositoryPostgreSQL {
	s := new(JenisMbkmRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *JenisMbkmRepositoryPostgreSQL) Create(data JenisMbkm) error {
	stmt, err := r.DB.Read.PrepareNamed(jenisMbkmQuery.Insert)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(data)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}
	return nil
}

func (r *JenisMbkmRepositoryPostgreSQL) GetAllData() (data []JenisMbkm, errr error) {
	rows, err := r.DB.Read.Queryx(jenisMbkmQuery.Select + " where is_deleted = false ")
	if err == sql.ErrNoRows {
		_ = failure.NotFound("Jenis MBKM Not Found")
		return
	}

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	for rows.Next() {
		var items JenisMbkm
		err = rows.StructScan(&items)

		if err != nil {
			return
		}

		data = append(data, items)
	}
	return
}

func (r *JenisMbkmRepositoryPostgreSQL) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	var searchParams []interface{}
	var searchRoleBuff bytes.Buffer
	searchRoleBuff.WriteString(" WHERE coalesce(is_deleted, false) = false ")

	if req.Keyword != "" {
		searchRoleBuff.WriteString(" AND ")
		searchRoleBuff.WriteString(" concat(nama) ilike ? ")
		searchParams = append(searchParams, "%"+req.Keyword+"%")
	}

	// query := r.DB.Read.Rebind("select count(*) from (" + siswaQuery.SelectDTO + searchRoleBuff.String() + ")s")
	query := r.DB.Read.Rebind(jenisMbkmQuery.Count + searchRoleBuff.String())
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

	searchRoleBuff.WriteString("order by " + ColumnMappJenisMbkm[req.SortBy].(string) + " " + req.SortType + " ")

	offset := (req.PageNumber - 1) * req.PageSize
	searchRoleBuff.WriteString("limit ? offset ? ")
	searchParams = append(searchParams, req.PageSize)
	searchParams = append(searchParams, offset)

	searchJenisMbkmQuery := searchRoleBuff.String()
	searchJenisMbkmQuery = r.DB.Read.Rebind(jenisMbkmQuery.Select + searchJenisMbkmQuery)
	fmt.Println("query", searchJenisMbkmQuery)
	rows, err := r.DB.Read.Queryx(searchJenisMbkmQuery, searchParams...)
	if err != nil {
		return
	}
	for rows.Next() {
		var items JenisMbkm
		err = rows.StructScan(&items)
		if err != nil {
			return
		}

		data.Items = append(data.Items, items)
	}

	data.Meta = pagination.CreateMeta(totalData, req.PageSize, req.PageNumber)

	return
}

func (r *JenisMbkmRepositoryPostgreSQL) ResolveByID(id string) (data JenisMbkm, err error) {
	err = r.DB.Read.Get(&data, jenisMbkmQuery.Select+" WHERE id=$1 ", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *JenisMbkmRepositoryPostgreSQL) Update(data JenisMbkm) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := txUpdateJenisMbkm(tx, data); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func txUpdateJenisMbkm(tx *sqlx.Tx, data JenisMbkm) (err error) {
	stmt, err := tx.PrepareNamed(jenisMbkmQuery.Update + " WHERE id=:id")
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(data)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return
}
