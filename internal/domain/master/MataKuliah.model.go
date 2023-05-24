package master

import (
	"time"

	"github.com/gofrs/uuid"
)

type MataKuliah struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	KodeMatkul string     `db:"kode_matkul" json:"kodeMatkul"`
	NamaMatkul string     `db:"nama_matkul" json:"namaMatkul"`
	Sks        int        `db:"sks" json:"sks"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy  *string    `db:"created_by" json:"createdBy"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy  *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted  bool       `db:"is_deleted" json:"isDeleted"`
}

var ColumnMappMatkul = map[string]interface{}{
	"id":         "id",
	"kodeMatkul": "kode_matkul",
	"namaMatkul": "nama_matkul",
	"sks":        "sks",
	"createdBy":  "created_by",
	"createdAt":  "created_at",
	"updatedBy":  "updated_by",
	"updatedAt":  "updated_at",
	"isDeleted":  "is_deleted",
}
