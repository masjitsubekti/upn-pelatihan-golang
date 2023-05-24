package master

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
)

var (
	pendaftarMbkmQuery = struct {
		Select    string
		SelectDTO string
		Insert    string
		Update    string
	}{
		Select:    `select id, id_siamik_mahasiswa, id_mitra_terlibat_program, updated_at, created_at, created_by, updated_by, is_deleted from pendaftar_program_mbkm`,
		SelectDTO: `select id, id_siamik_mahasiswa, id_mitra_terlibat_program, updated_at, created_at, created_by, updated_by, is_deleted from pendaftar_program_mbkm`,
		Insert:    `INSERT INTO pendaftar_program_mbkm (id, id_siamik_mahasiswa, id_mitra_terlibat_program, created_by, created_at) values(:id, :id_siamik_mahasiswa, :id_mitra_terlibat_program, :created_by, :created_at) `,
		Update: `UPDATE pendaftar_program_mbkm SET 
			id=:id, 
			id_siamik_mahasiswa=:id_siamik_mahasiswa,
			id_mitra_terlibat_program=:id_mitra_terlibat_program,
			updated_at=:updated_at,
			updated_by=:updated_by, 
			is_deleted=:is_deleted`,
	}
)

var (
	paketKonversiQuery = struct {
		Select                string
		SelectDTO             string
		InsertBulk            string
		InsertBulkPlaceholder string
	}{
		Select: `SELECT id, id_paket_konversi, id_matkul, sks, nilai_angka, nilai_huruf, id_pendaftar_mbkm, created_at, updated_at, created_by, updated_by, is_deleted
		FROM public.matkul_paket_konversi_terpilih_awal`,
		SelectDTO: `SELECT id, id_paket_konversi, id_matkul, sks, nilai_angka, nilai_huruf, id_pendaftar_mbkm, created_at, updated_at, created_by, updated_by, is_deleted
		FROM public.matkul_paket_konversi_terpilih_awal`,
		InsertBulk:            `INSERT INTO public.matkul_paket_konversi_terpilih_awal(id, id_paket_konversi, id_matkul, sks, id_pendaftar_mbkm, created_at, created_by) values `,
		InsertBulkPlaceholder: ` (:id, :id_paket_konversi, :id_matkul, :sks, :id_pendaftar_mbkm, :created_at, :created_by) `,
	}
)

type PendaftarProgramMbkmRepository interface {
	Create(data PendaftarProgramMbkm) error
}

type PendaftarProgramMbkmRepositoryPostgreSQL struct {
	DB *infras.PostgresqlConn
}

func ProvidePendaftarProgramMbkmRepositoryPostgreSQL(db *infras.PostgresqlConn) *PendaftarProgramMbkmRepositoryPostgreSQL {
	s := new(PendaftarProgramMbkmRepositoryPostgreSQL)
	s.DB = db
	return s
}

// Function digunakan untuk create with transaction
func (r *PendaftarProgramMbkmRepositoryPostgreSQL) Create(data PendaftarProgramMbkm) error {
	return r.DB.WithTransaction(func(tx *sqlx.Tx, e chan error) {
		// Function create table pendaftar_program_mbkm
		if err := r.CreateTxPendaftarMbkm(tx, data); err != nil {
			e <- err
			return
		}

		// Function Insert Bulk table matkul_paket_konversi_terpilih_awal
		if err := txCreatePaketKonversi(tx, data.PaketKonversi); err != nil {
			e <- err
			return
		}
		e <- nil
	})
}

// fungsi ini digunakan untuk insert ke tabel pendaftar_program_mbkm
func (r *PendaftarProgramMbkmRepositoryPostgreSQL) CreateTxPendaftarMbkm(tx *sqlx.Tx, data PendaftarProgramMbkm) error {
	stmt, err := tx.PrepareNamed(pendaftarMbkmQuery.Insert)
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

// fungsi ini digunakan untuk insert ke tabel matkul_paket_konversi_terpilih_awal
func txCreatePaketKonversi(tx *sqlx.Tx, details []MatkulPaketKonversiTerpilihAwal) (err error) {
	if len(details) == 0 {
		return
	}
	query, args, err := composeBulkUpsertPaketKonversiQuery(details)
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

func composeBulkUpsertPaketKonversiQuery(details []MatkulPaketKonversiTerpilihAwal) (qResult string, params []interface{}, err error) {
	values := []string{}
	for _, d := range details {
		param := map[string]interface{}{
			"id":                d.ID,
			"id_paket_konversi": d.IdPaketKonversi,
			"id_matkul":         d.IdMatkul,
			"sks":               d.Sks,
			"id_pendaftar_mbkm": d.IdPendaftarMbkm,
			"created_at":        d.CreatedAt,
			"created_by":        d.CreatedBy,
		}
		q, args, err := sqlx.Named(paketKonversiQuery.InsertBulkPlaceholder, param)
		if err != nil {
			return qResult, params, err
		}
		values = append(values, q)
		params = append(params, args...)
	}
	qResult = fmt.Sprintf(`%v %v 
						ON CONFLICT (id) 
						DO UPDATE SET id_paket_konversi=EXCLUDED.id_paket_konversi, 
						id_matkul=EXCLUDED.id_matkul,
						sks=EXCLUDED.sks `, paketKonversiQuery.InsertBulk, strings.Join(values, ","))

	fmt.Println("query bulk", qResult)
	return
}
