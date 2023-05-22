package master

import (
	"time"

	"github.com/gofrs/uuid"
	"gitlab.com/upn-belajar-go/shared"
)

type Siswa struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Nama      string     `db:"nama" json:"nama" validate:"required"`
	Kelas     string     `db:"kelas" json:"kelas"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy *string    `db:"created_by" json:"createdBy"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
	Berkas    *string    `db:"berkas" json:"berkas"`
	IdKelas   *string    `db:"id_kelas" json:"idKelas"`
}

type RequestSiswaFormat struct {
	ID      string `db:"id" json:"id"`
	Nama    string `db:"nama" json:"nama" validate:"required"`
	Kelas   string `db:"kelas" json:"kelas" validate:"required"`
	Berkas  string `db:"berkas" json:"berkas" validate:"required"`
	IdKelas string `db:"id_kelas" json:"idKelas"`
}

// Validate validates the entity.
func (f *Siswa) Validate() (err error) {
	validator := shared.GetValidator()
	return validator.Struct(f)
}

func (s *Siswa) NewSiswaFormat(reqFormat RequestSiswaFormat, userID string) (newSiswa Siswa, err error) {
	newID, _ := uuid.NewV4()
	now := time.Now()

	if reqFormat.ID == "" {
		newSiswa = Siswa{
			ID:        newID,
			Nama:      reqFormat.Nama,
			Kelas:     reqFormat.Kelas,
			Berkas:    &reqFormat.Berkas,
			IdKelas:   &reqFormat.IdKelas,
			CreatedAt: time.Now(),
			CreatedBy: &userID,
		}
	} else {
		id, _ := uuid.FromString(reqFormat.ID)
		newSiswa = Siswa{
			ID:        id,
			Nama:      reqFormat.Nama,
			Kelas:     reqFormat.Kelas,
			Berkas:    &reqFormat.Berkas,
			IdKelas:   &reqFormat.IdKelas,
			UpdatedAt: &now,
			UpdatedBy: &userID,
		}
	}
	err = newSiswa.Validate()
	return
}

var ColumnMappSiswa = map[string]interface{}{
	"id":        "s.id",
	"nama":      "s.nama",
	"kelas":     "s.kelas",
	"namaKelas": "k.nama",
	"createdBy": "s.created_by",
	"createdAt": "s.created_at",
	"updatedBy": "s.updated_by",
	"updatedAt": "s.updated_at",
	"isDeleted": "s.is_deleted",
}

func (siswa *Siswa) SoftDelete(userId string) {
	now := time.Now()
	siswa.IsDeleted = true
	siswa.UpdatedBy = &userId
	siswa.UpdatedAt = &now
}

type SiswaDTO struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Nama      string     `db:"nama" json:"nama" validate:"required"`
	Kelas     string     `db:"kelas" json:"kelas"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy *string    `db:"created_by" json:"createdBy"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
	Berkas    *string    `db:"berkas" json:"berkas"`
	IdKelas   *string    `db:"id_kelas" json:"idKelas"`
	NamaKelas *string    `db:"nama_kelas" json:"namaKelas"`
}
