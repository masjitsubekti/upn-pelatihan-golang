package master

import (
	"bytes"
	"database/sql"
	"fmt"

	"gitlab.com/upn-belajar-go/shared/model"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

var (
	jenismitraQuery = struct {
		Select string
		Insert string
		Update string
		Delete string
		Exist  string
		Count  string
	}{
		Select: `SELECT id, nama_jenis_mitra, created_at, updated_at, created_by, updated_by, is_deleted
		         FROM jenis_mitra`,
		Insert: `INSERT INTO jenis_mitra (id, nama_jenis_mitra, created_by, created_at) 
		                         values(:id, :nama_jenis_mitra,  :created_by, :created_at) `,
		Update: `UPDATE jenis_mitra SET 
				id=:id, 
				nama_jenis_mitra=:nama_jenis_mitra,
				updated_at=:updated_at,
				updated_by=:updated_by, 
				is_deleted=:is_deleted`,
		Delete: `delete from jenis_mitra `,
		Exist:  `select count(id)>0 from jenis_mitra `,
		Count:  `select count(id) from jenis_mitra `,
	}
)

type JenisMitraRepository interface {
	Create(data JenisMitra) error
	Update(data JenisMitra) error
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	GetAllData() (data []JenisMitra, err error)
	ResolveByID(id uuid.UUID) (data JenisMitra, err error)
	ExistByNama(nama string) (bool, error)
	ExistByNamaID(id uuid.UUID, nama string) (bool, error)
}

type JenisMitraRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideJenisMitraRepositoryPostgreSQL(db *infras.PostgresqlConn) *JenisMitraRepositoryPostgreSQL {
	s := new(JenisMitraRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *JenisMitraRepositoryPostgreSQL) Create(jenismitra JenisMitra) error {
	stmt, err := r.DB.Read.PrepareNamed(jenismitraQuery.Insert)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(jenismitra)
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}
	return nil
}

func (r *JenisMitraRepositoryPostgreSQL) Update(data JenisMitra) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		if err := txUpdateJenisMitra(tx, data); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func txUpdateJenisMitra(tx *sqlx.Tx, data JenisMitra) (err error) {
	stmt, err := tx.PrepareNamed(jenismitraQuery.Update + " WHERE id=:id")
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

func (r *JenisMitraRepositoryPostgreSQL) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	var searchParams []interface{}
	var searchRoleBuff bytes.Buffer
	searchRoleBuff.WriteString(" WHERE coalesce(is_deleted, false) = false ")

	if req.Keyword != "" {
		searchRoleBuff.WriteString(" AND ")
		searchRoleBuff.WriteString(" concat(nama_jenis_mitra) ilike ?  ")
		searchParams = append(searchParams, "%"+req.Keyword+"%")
	}

	query := r.DB.Read.Rebind(jenismitraQuery.Count + searchRoleBuff.String())

	var totalData int
	err = r.DB.Read.QueryRow(query, searchParams...).Scan(&totalData)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	if totalData < 1 {
		data.Items = make([]interface{}, 0)
		return
	}

	searchRoleBuff.WriteString("order by " + ColumnMappJenisMitra[req.SortBy].(string) + " " + req.SortType + " ")

	offset := (req.PageNumber - 1) * req.PageSize
	searchRoleBuff.WriteString("limit ? offset ? ")
	searchParams = append(searchParams, req.PageSize)
	searchParams = append(searchParams, offset)

	searchjenismitraQuery := searchRoleBuff.String()
	searchjenismitraQuery = r.DB.Read.Rebind(jenismitraQuery.Select + searchjenismitraQuery)
	fmt.Println("query", searchjenismitraQuery)
	rows, err := r.DB.Read.Queryx(searchjenismitraQuery, searchParams...)
	if err != nil {
		return
	}
	for rows.Next() {
		var jenismitra JenisMitra
		err = rows.StructScan(&jenismitra)
		if err != nil {
			return
		}

		data.Items = append(data.Items, jenismitra)
	}

	data.Meta = pagination.CreateMeta(totalData, req.PageSize, req.PageNumber)

	return
}

func (r *JenisMitraRepositoryPostgreSQL) GetAllData() (data []JenisMitra, errr error) {
	rows, err := r.DB.Read.Queryx(jenismitraQuery.Select + " where is_deleted = false ")
	if err == sql.ErrNoRows {
		_ = failure.NotFound("JenisMitra")
		return
	}

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	for rows.Next() {
		var master JenisMitra
		err = rows.StructScan(&master)

		if err != nil {
			return
		}

		data = append(data, master)
	}
	return
}

func (r *JenisMitraRepositoryPostgreSQL) ResolveByID(id uuid.UUID) (data JenisMitra, err error) {
	err = r.DB.Read.Get(&data, jenismitraQuery.Select+" WHERE id=$1 ", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *JenisMitraRepositoryPostgreSQL) ExistByNama(nama string) (bool, error) {
	var exist bool

	err := r.DB.Read.Get(&exist, jenismitraQuery.Exist+" where upper(nama_jenis_mitra)=upper($1) and coalesce(is_deleted, false)=false ", nama)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return exist, err
}

func (r *JenisMitraRepositoryPostgreSQL) ExistByNamaID(id uuid.UUID, nama string) (bool, error) {
	var exist bool

	err := r.DB.Read.Get(&exist, jenismitraQuery.Exist+" where id <> $1 and upper(nama_jenis_mitra)=upper($2) and coalesce(is_deleted, false)=false ", id, nama)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return exist, err
}
