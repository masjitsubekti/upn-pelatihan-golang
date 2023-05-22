package master

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
)

var (
	kelasSiswaQuery = struct {
		Select    string
		SelectDTO string
		Insert    string
		Update    string
	}{
		Select: `select id, id_kelas, tahun_ajaran, updated_at, created_at, created_by, updated_by, is_deleted from kelas_siswa`,
		SelectDTO: `select kl.id, kl.id_kelas, kl.tahun_ajaran, k.kode, k.nama nama_kelas from public.kelas_siswa kl
		left join m_kelas k on kl.id_kelas = k.id`,
		Insert: `INSERT INTO kelas_siswa (id, id_kelas, tahun_ajaran, created_by, created_at) values(:id, :id_kelas, :tahun_ajaran, :created_by, :created_at) `,
		Update: `UPDATE kelas_siswa SET 
			id=:id, 
			id_kelas=:id_kelas,
			tahun_ajaran=:tahun_ajaran,
			updated_at=:updated_at,
			updated_by=:updated_by, 
			is_deleted=:is_deleted`,
	}
)

var (
	kelasSiswaDetailQuery = struct {
		Select                string
		SelectDTO             string
		InsertBulk            string
		InsertBulkPlaceholder string
		Exist                 string
	}{
		Select: `select id, id_kelas_siswa, id_siswa, updated_at, created_at, created_by, updated_by, is_deleted from kelas_siswa_detail`,
		SelectDTO: `select kl.id, kl.id_kelas_siswa, kl.id_siswa, s.nama nama_siswa from public.kelas_siswa_detail kl
		left join m_siswa s on kl.id_siswa = s.id`,
		InsertBulk:            `INSERT INTO public.kelas_siswa_detail(id, id_kelas_siswa, id_siswa, created_at, created_by) values `,
		InsertBulkPlaceholder: ` (:id, :id_kelas_siswa, :id_siswa, :created_at, :created_by) `,
		Exist: ` select count(id_siswa)>0 from kelas_siswa_detail s
			left join kelas_siswa ks on ks.id = s.id_kelas_siswa
		`,
	}
)

type KelasSiswaRepository interface {
	Create(data KelasSiswa) error
	UpdateKelasSiswa(data KelasSiswa) error
	Update(data KelasSiswa) error
	ExistByIdSiswa(idSiswa string, idKelasSiswa string) (bool, error)
	ResolveByIDDTO(id string) (data KelasSiswaDTO, err error)
	ResolveAll(req model.StandardRequest) (data pagination.Response, err error)
	ResolveByID(id string) (data KelasSiswa, err error)
}

type KelasSiswaRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvideKelasSiswaRepositoryPostgreSQL(db *infras.PostgresqlConn) *KelasSiswaRepositoryPostgreSQL {
	s := new(KelasSiswaRepositoryPostgreSQL)
	s.DB = db
	return s
}

func (r *KelasSiswaRepositoryPostgreSQL) ResolveAll(req model.StandardRequest) (data pagination.Response, err error) {
	var searchParams []interface{}
	var searchRoleBuff bytes.Buffer
	searchRoleBuff.WriteString(" WHERE coalesce(kl.is_deleted, false) = false ")

	if req.Keyword != "" {
		searchRoleBuff.WriteString(" AND ")
		searchRoleBuff.WriteString(" concat(k.nama, kl.tahun_ajaran) ilike ? ")
		searchParams = append(searchParams, "%"+req.Keyword+"%")
	}

	query := r.DB.Read.Rebind("select count(*) from (" + kelasSiswaQuery.SelectDTO + searchRoleBuff.String() + ")s")

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

	searchRoleBuff.WriteString("order by " + ColumnMappKelasSiswa[req.SortBy].(string) + " " + req.SortType + " ")

	offset := (req.PageNumber - 1) * req.PageSize
	searchRoleBuff.WriteString("limit ? offset ? ")
	searchParams = append(searchParams, req.PageSize)
	searchParams = append(searchParams, offset)

	searchSiswaQuery := searchRoleBuff.String()
	searchSiswaQuery = r.DB.Read.Rebind(kelasSiswaQuery.SelectDTO + searchSiswaQuery)
	fmt.Println("query", searchSiswaQuery)
	rows, err := r.DB.Read.Queryx(searchSiswaQuery, searchParams...)
	if err != nil {
		return
	}
	for rows.Next() {
		var siswa KelasSiswaDTO
		err = rows.StructScan(&siswa)
		if err != nil {
			return
		}

		data.Items = append(data.Items, siswa)
	}

	data.Meta = pagination.CreateMeta(totalData, req.PageSize, req.PageNumber)

	return
}

// Function digunakan untuk create with transaction
func (r *KelasSiswaRepositoryPostgreSQL) Create(data KelasSiswa) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		// Function create table kelas_siswa
		if err := r.CreateTxKelasSiswa(tx, data); err != nil {
			e <- err
			return
		}

		// Function Insert Bulk table kelas_siswa_detail
		if err := txCreateKelasSiswaDetail(tx, data.Detail); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

// Function digunakan untuk update with transaction
func (r *KelasSiswaRepositoryPostgreSQL) UpdateKelasSiswa(data KelasSiswa) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		// Function update table kelas_siswa
		if err := r.UpdateTxKelasSiswa(tx, data); err != nil {
			e <- err
			return
		}

		// Function delete not in table kelas_siswa_detail
		ids := make([]string, 0)
		for _, d := range data.Detail {
			ids = append(ids, d.ID.String())
		}
		fmt.Println(ids)
		if err := r.txDeleteDetailNotIn(tx, data.ID.String(), ids); err != nil {
			e <- err
			return
		}

		// Function Insert Bulk table kelas_siswa_detail
		if err := txCreateKelasSiswaDetail(tx, data.Detail); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func (r *KelasSiswaRepositoryPostgreSQL) CreateTxKelasSiswa(tx *sqlx.Tx, data KelasSiswa) error {
	stmt, err := tx.PrepareNamed(kelasSiswaQuery.Insert)
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

func (r *KelasSiswaRepositoryPostgreSQL) Update(data KelasSiswa) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		// Function update table kelas_siswa
		if err := r.UpdateTxKelasSiswa(tx, data); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

func (r *KelasSiswaRepositoryPostgreSQL) UpdateTxKelasSiswa(tx *sqlx.Tx, data KelasSiswa) error {
	stmt, err := tx.PrepareNamed(kelasSiswaQuery.Update + " WHERE id=:id")
	if err != nil {
		logger.ErrorWithStack(err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(data)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return nil
}

func txCreateKelasSiswaDetail(tx *sqlx.Tx, details []KelasSiswaDetail) (err error) {
	if len(details) == 0 {
		return
	}
	query, args, err := composeBulkUpsertKelasSiswaDetailQuery(details)
	if err != nil {
		return
	}

	query = tx.Rebind(query)
	stmt, err := tx.Preparex(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Stmt.Exec(args...)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func composeBulkUpsertKelasSiswaDetailQuery(details []KelasSiswaDetail) (qResult string, params []interface{}, err error) {
	values := []string{}
	for _, d := range details {
		param := map[string]interface{}{
			"id":             d.ID,
			"id_kelas_siswa": d.IdKelasSiswa,
			"id_siswa":       d.IdSiswa,
			"created_at":     d.CreatedAt,
			"created_by":     d.CreatedBy,
		}
		q, args, err := sqlx.Named(kelasSiswaDetailQuery.InsertBulkPlaceholder, param)
		if err != nil {
			return qResult, params, err
		}
		values = append(values, q)
		params = append(params, args...)
	}
	qResult = fmt.Sprintf(`%v %v 
						ON CONFLICT (id) 
						DO UPDATE SET id_kelas_siswa=EXCLUDED.id_kelas_siswa, id_siswa=EXCLUDED.id_siswa `, kelasSiswaDetailQuery.InsertBulk, strings.Join(values, ","))

	fmt.Println("tes", qResult)
	return
}

func (r *KelasSiswaRepositoryPostgreSQL) ExistByIdSiswa(idSiswa string, idKelasSiswa string) (bool, error) {
	var exist bool
	err := r.DB.Read.Get(&exist, kelasSiswaDetailQuery.Exist+" where coalesce(s.is_deleted) = false and s.id_siswa = $1 and ks.id = $2 ", idSiswa, idKelasSiswa)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	return exist, err
}

func (r *KelasSiswaRepositoryPostgreSQL) txDeleteDetailNotIn(tx *sqlx.Tx, idKelasSiswa string, ids []string) (err error) {
	query, args, err := sqlx.In("delete from kelas_siswa_detail where id_kelas_siswa = ? AND id NOT IN (?)", idKelasSiswa, ids)
	query = tx.Rebind(query)

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	res, err := r.DB.Write.Exec(query, args...)
	_, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}

func (r *KelasSiswaRepositoryPostgreSQL) ResolveByID(id string) (data KelasSiswa, err error) {
	err = r.DB.Read.Get(&data, kelasSiswaQuery.Select+" WHERE id=$1 ", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	return
}

func (r *KelasSiswaRepositoryPostgreSQL) ResolveByIDDTO(id string) (data KelasSiswaDTO, err error) {
	err = r.DB.Read.Get(&data, kelasSiswaQuery.SelectDTO+" WHERE kl.id=$1 ", id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	siswaDetail, err := r.GetAllSiswaByID(id)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	data.Detail = siswaDetail
	return
}

func (r *KelasSiswaRepositoryPostgreSQL) GetAllSiswaByID(idKelasSiswa string) (data []KelasSiswaDetailDTO, err error) {
	rows, err := r.DB.Read.Queryx(kelasSiswaDetailQuery.SelectDTO+" where kl.id_kelas_siswa=$1 and kl.is_deleted = false ", idKelasSiswa)
	if err == sql.ErrNoRows {
		_ = failure.NotFound("Siswa Not Found")
		return
	}

	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	for rows.Next() {
		var master KelasSiswaDetailDTO
		err = rows.StructScan(&master)

		if err != nil {
			return
		}

		data = append(data, master)
	}
	return
}
