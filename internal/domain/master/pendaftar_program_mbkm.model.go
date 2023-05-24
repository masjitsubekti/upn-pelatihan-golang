package master

import (
	"time"

	"github.com/gofrs/uuid"
)

type PendaftarProgramMbkm struct {
	ID                     uuid.UUID                         `db:"id" json:"id"`
	IdSiamikMahasiswa      string                            `db:"id_siamik_mahasiswa" json:"idSiamikMahasiswa"`
	IdMitraTerlibatProgram *string                           `db:"id_mitra_terlibat_program" json:"idMitraTerlibatProgram"`
	CreatedAt              time.Time                         `db:"created_at" json:"createdAt"`
	CreatedBy              *string                           `db:"created_by" json:"createdBy"`
	UpdatedAt              *time.Time                        `db:"updated_at" json:"updatedAt"`
	UpdatedBy              *string                           `db:"updated_by" json:"updatedBy"`
	IsDeleted              bool                              `db:"is_deleted" json:"isDeleted"`
	PaketKonversi          []MatkulPaketKonversiTerpilihAwal `db:"-" json:"paketKonversi"`
}

type PendaftarProgramMbkmRequest struct {
	ID                     string                                   `db:"id" json:"id"`
	IdSiamikMahasiswa      string                                   `db:"id_siamik_mahasiswa" json:"idSiamikMahasiswa"`
	IdMitraTerlibatProgram *string                                  `db:"id_mitra_terlibat_program" json:"idMitraTerlibatProgram"`
	PaketKonversi          []MatkulPaketKonversiTerpilihAwalRequest `db:"-" json:"paketKonversi"`
}

type MatkulPaketKonversiTerpilihAwal struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	IdPaketKonversi *string    `db:"id_paket_konversi" json:"idPaketKonversi"`
	IdMatkul        string     `db:"id_matkul" json:"idMatkul"`
	IdPendaftarMbkm string     `db:"id_pendaftar_mbkm" json:"idPendaftarMbkm"`
	Sks             string     `db:"sks" json:"sks"`
	NilaiAngka      *string    `db:"nilai_angka" json:"nilaiAngka"`
	NilaiHuruf      *string    `db:"nilai_huruf" json:"nilaiHuruf"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy       *string    `db:"created_by" json:"createdBy"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy       *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted       bool       `db:"is_deleted" json:"isDeleted"`
}

type MatkulPaketKonversiTerpilihAwalRequest struct {
	ID              string  `db:"id" json:"id"`
	IdPaketKonversi *string `db:"id_paket_konversi" json:"idPaketKonversi"`
	IdMatkul        string  `db:"id_matkul" json:"idMatkul"`
	Sks             string  `db:"sks" json:"sks"`
}

func (s *PendaftarProgramMbkm) NewPendaftarProgramMbkmFormat(reqFormat PendaftarProgramMbkmRequest, userID string) (newPendaftar PendaftarProgramMbkm, err error) {
	newID, _ := uuid.NewV4()
	now := time.Now()

	if reqFormat.ID == "" {
		newPendaftar = PendaftarProgramMbkm{
			ID:                     newID,
			IdSiamikMahasiswa:      reqFormat.IdSiamikMahasiswa,
			IdMitraTerlibatProgram: reqFormat.IdMitraTerlibatProgram,
			CreatedAt:              time.Now(),
			CreatedBy:              &userID,
		}
	} else {
		id, _ := uuid.FromString(reqFormat.ID)
		newPendaftar = PendaftarProgramMbkm{
			ID:                     id,
			IdSiamikMahasiswa:      reqFormat.IdSiamikMahasiswa,
			IdMitraTerlibatProgram: reqFormat.IdMitraTerlibatProgram,
			UpdatedAt:              &now,
			UpdatedBy:              &userID,
		}
	}

	details := make([]MatkulPaketKonversiTerpilihAwal, 0)
	for _, d := range reqFormat.PaketKonversi {
		var detID uuid.UUID
		if d.ID == "" {
			detID, _ = uuid.NewV4()
		} else {
			detID, _ = uuid.FromString(d.ID)
		}

		newDetail := MatkulPaketKonversiTerpilihAwal{
			ID:              detID,
			IdPaketKonversi: d.IdPaketKonversi,
			IdMatkul:        d.IdMatkul,
			Sks:             d.Sks,
			IdPendaftarMbkm: newPendaftar.ID.String(),
			CreatedAt:       time.Now(),
			CreatedBy:       &userID,
		}

		details = append(details, newDetail)
	}

	newPendaftar.PaketKonversi = details

	return
}
