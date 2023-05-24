package master

import (
	"time"

	"github.com/gofrs/uuid"
)

type JenisMbkm struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Nama      string     `db:"nama" json:"nama"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	CreatedBy *string    `db:"created_by" json:"createdBy"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string    `db:"updated_by" json:"updatedBy"`
	IsDeleted bool       `db:"is_deleted" json:"isDeleted"`
}

type RequestJenisMbkmFormat struct {
	ID   string `db:"id" json:"id"`
	Nama string `db:"nama" json:"nama"`
}

func (s *JenisMbkm) NewJenisMbkmFormat(reqFormat RequestJenisMbkmFormat, userID string) (newJenisMbkm JenisMbkm, err error) {
	newID, _ := uuid.NewV4()
	now := time.Now()

	if reqFormat.ID == "" {
		newJenisMbkm = JenisMbkm{
			ID:        newID,
			Nama:      reqFormat.Nama,
			CreatedAt: time.Now(),
			CreatedBy: &userID,
		}
	} else {
		id, _ := uuid.FromString(reqFormat.ID)
		newJenisMbkm = JenisMbkm{
			ID:        id,
			Nama:      reqFormat.Nama,
			UpdatedAt: &now,
			UpdatedBy: &userID,
		}
	}

	return
}

var ColumnMappJenisMbkm = map[string]interface{}{
	"id":        "id",
	"nama":      "nama",
	"createdBy": "created_by",
	"createdAt": "created_at",
	"updatedBy": "updated_by",
	"updatedAt": "updated_at",
	"isDeleted": "is_deleted",
}

func (jenis *JenisMbkm) SoftDelete(userId string) {
	now := time.Now()
	jenis.IsDeleted = true
	jenis.UpdatedBy = &userId
	jenis.UpdatedAt = &now
}
