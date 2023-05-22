package orm

import (
	"time"

	"github.com/gofrs/uuid"
)

const tableNameKelas = "m_kelas"

type Kelas struct {
	ID        uuid.UUID  `gorm:"column:id" db:"id" json:"id"`
	Kode      string     `gorm:"column:kode" db:"kode" json:"kode"`
	Nama      string     `gorm:"column:nama" db:"nama" json:"nama"`
	CreatedBy *string    `gorm:"column:created_by" json:"createdBy" db:"created_by" validate:"required"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt" db:"created_at" validate:"required"`
	UpdatedBy *string    `gorm:"column:updated_by" json:"updatedBy,omitempty" db:"updated_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updatedAt,omitempty" db:"updated_at"`
	IsDeleted bool       `gorm:"column:is_deleted" json:"isDeleted" db:"is_deleted"`
}

func (*Kelas) TableName() string {
	return tableNameKelas
}

type KelasRequest struct {
	ID   string `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

func (k *Kelas) BindFromRequest(req KelasRequest, id string, userID string) {
	var now = time.Now()
	if id == "" {
		newID, _ := uuid.NewV4()
		k.ID = newID
		k.CreatedAt = now
		k.CreatedBy = &userID
	} else {
		id, _ := uuid.FromString(id)
		k.ID = id
		k.UpdatedAt = &now
		k.UpdatedBy = &userID
	}
	k.Kode = req.Kode
	k.Nama = req.Nama
}

var ColumnMappKelas = map[string]interface{}{
	"id":   "id",
	"kode": "kode",
	"nama": "nama",
}

func (item *Kelas) SoftDelete(userID string) {
	var now = time.Now()
	item.UpdatedAt = &now
	item.UpdatedBy = &userID
	item.IsDeleted = true
}
