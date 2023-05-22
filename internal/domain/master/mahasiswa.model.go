package master

import (
	"time"

	"github.com/gofrs/uuid"
)

type Mahasiswa struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	NPM       string     `db:"npm" json:"npm"`
	Nama      string     `db:"nama" json:"nama"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy *string    `db:"created_by" json:"createdBy"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}
