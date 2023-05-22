package master

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"gitlab.com/upn-belajar-go/shared/nuuid"
)

type JenisMitra struct {
	ID             uuid.UUID   `db:"id" json:"id"`
	NamaJenisMitra string      `db:"nama_jenis_mitra" json:"namaJenisMitra"`
	CreatedAt      time.Time   `db:"created_at" json:"createdAt"`
	CreatedBy      nuuid.NUUID `db:"created_by" json:"createdBy"`
	UpdatedAt      null.Time   `db:"updated_at" json:"updatedAt"`
	UpdatedBy      nuuid.NUUID `db:"updated_by" json:"updatedBy"`
	IsDeleted      bool        `db:"is_deleted" json:"isDeleted"`
}

type RequestJenisMitraFormat struct {
	ID             uuid.UUID `db:"id" json:"id"`
	NamaJenisMitra string    `db:"nama_jenis_mitra" json:"namaJenisMitra"`
}

var ColumnMappJenisMitra = map[string]interface{}{
	"id":             "id",
	"namaJenisMitra": "nama_jenis_mitra",
	"createdBy":      "created_by",
	"createdAt":      "created_at",
	"updatedBy":      "updated_by",
	"updatedAt":      "updated_at",
	"isDeleted":      "is_deleted",
}

func (jenismitra *JenisMitra) NewJenisMitraFormat(reqFormat RequestJenisMitraFormat, userID uuid.UUID) (newJenisMitra JenisMitra, err error) {
	newID, _ := uuid.NewV4()

	if reqFormat.ID == uuid.Nil {
		newJenisMitra = JenisMitra{
			ID:             newID,
			NamaJenisMitra: reqFormat.NamaJenisMitra,
			CreatedAt:      time.Now(),
			CreatedBy:      nuuid.From(userID),
		}
	} else {
		newJenisMitra = JenisMitra{
			ID:             reqFormat.ID,
			NamaJenisMitra: reqFormat.NamaJenisMitra,
			UpdatedAt:      null.TimeFrom(time.Now()),
			UpdatedBy:      nuuid.From(userID),
		}
	}
	return
}

func (jenismitra *JenisMitra) SoftDelete(userId uuid.UUID) {
	jenismitra.IsDeleted = true
	jenismitra.UpdatedBy = nuuid.From(userId)
	jenismitra.UpdatedAt = null.TimeFrom(time.Now())
}
