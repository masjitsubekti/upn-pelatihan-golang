package master

import (
	"time"

	"github.com/gofrs/uuid"
)

type KelasSiswa struct {
	ID          uuid.UUID          `db:"id" json:"id"`
	IdKelas     string             `db:"id_kelas" json:"idKelas" validate:"required"`
	TahunAjaran string             `db:"tahun_ajaran" json:"tahunAjaran"`
	CreatedAt   time.Time          `db:"created_at" json:"createdAt"`
	CreatedBy   *string            `db:"created_by" json:"createdBy"`
	UpdatedAt   *time.Time         `db:"updated_at" json:"updatedAt"`
	UpdatedBy   *string            `db:"updated_by" json:"updatedBy"`
	IsDeleted   bool               `db:"is_deleted" json:"isDeleted"`
	Detail      []KelasSiswaDetail `db:"-" json:"detail"`
}

type KelasSiswaDetail struct {
	ID           uuid.UUID  `db:"id" json:"id"`
	IdKelasSiswa string     `db:"id_kelas_siswa" json:"idKelasSiswa" validate:"required"`
	IdSiswa      string     `db:"id_siswa" json:"idSiswa"`
	CreatedAt    time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy    *string    `db:"created_by" json:"createdBy"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy    *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted    bool       `db:"is_deleted" json:"isDeleted"`
}

type KelasSiswaDTO struct {
	ID          uuid.UUID             `db:"id" json:"id"`
	IdKelas     string                `db:"id_kelas" json:"idKelas"`
	Kode        string                `db:"kode" json:"kode"`
	NamaKelas   string                `db:"nama_kelas" json:"namaKelas"`
	TahunAjaran string                `db:"tahun_ajaran" json:"tahunAjaran"`
	Detail      []KelasSiswaDetailDTO `db:"-" json:"detail"`
}

type KelasSiswaDetailDTO struct {
	ID           uuid.UUID `db:"id" json:"id"`
	IdKelasSiswa string    `db:"id_kelas_siswa" json:"idKelasSiswa" validate:"required"`
	IdSiswa      string    `db:"id_siswa" json:"idSiswa"`
	NamaSiswa    string    `db:"nama_siswa" json:"namaSiswa"`
}
type KelasSiswaRequest struct {
	ID          string                    `db:"id" json:"id"`
	IdKelas     string                    `db:"id_kelas" json:"idKelas" validate:"required"`
	TahunAjaran string                    `db:"tahun_ajaran" json:"tahunAjaran"`
	Detail      []KelasSiswaDetailRequest `db:"-" json:"detail"`
}

type KelasSiswaDetailRequest struct {
	ID           string `db:"id" json:"id"`
	IdKelasSiswa string `db:"id_kelas_siswa" json:"idKelasSiswa" validate:"required"`
	IdSiswa      string `db:"id_siswa" json:"idSiswa"`
}

func (s *KelasSiswa) NewKelasSiswaFormat(reqFormat KelasSiswaRequest, userID string) (newKelas KelasSiswa, err error) {
	newID, _ := uuid.NewV4()
	now := time.Now()

	if reqFormat.ID == "" {
		newKelas = KelasSiswa{
			ID:          newID,
			IdKelas:     reqFormat.IdKelas,
			TahunAjaran: reqFormat.TahunAjaran,
			CreatedAt:   time.Now(),
			CreatedBy:   &userID,
		}
	} else {
		id, _ := uuid.FromString(reqFormat.ID)
		newKelas = KelasSiswa{
			ID:          id,
			IdKelas:     reqFormat.IdKelas,
			TahunAjaran: reqFormat.TahunAjaran,
			UpdatedAt:   &now,
			UpdatedBy:   &userID,
		}
	}

	details := make([]KelasSiswaDetail, 0)
	for _, d := range reqFormat.Detail {
		var detID uuid.UUID
		if d.ID == "" {
			detID, _ = uuid.NewV4()
		} else {
			detID, _ = uuid.FromString(d.ID)
		}

		newDetail := KelasSiswaDetail{
			ID:           detID,
			IdSiswa:      d.IdSiswa,
			IdKelasSiswa: newKelas.ID.String(),
			CreatedAt:    time.Now(),
			CreatedBy:    &userID,
		}

		details = append(details, newDetail)
	}

	newKelas.Detail = details

	return
}

var ColumnMappKelasSiswa = map[string]interface{}{
	"id":          "kl.id",
	"kode":        "k.kode",
	"namaKelas":   "k.nama",
	"id_kelas":    "kl.id_kelas",
	"tahunAjaran": "kl.tahun_ajaran",
	"createdBy":   "kl.created_by",
	"createdAt":   "kl.created_at",
	"updatedBy":   "kl.updated_by",
	"updatedAt":   "kl.updated_at",
	"isDeleted":   "kl.is_deleted",
}

func (kelas *KelasSiswa) SoftDelete(userId string) {
	now := time.Now()
	kelas.IsDeleted = true
	kelas.UpdatedBy = &userId
	kelas.UpdatedAt = &now
}
