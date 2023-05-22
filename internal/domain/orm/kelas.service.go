package orm

import (
	"database/sql"

	"gitlab.com/upn-belajar-go/infras"
	"gitlab.com/upn-belajar-go/shared/logger"
	"gitlab.com/upn-belajar-go/shared/model"
	"gitlab.com/upn-belajar-go/shared/pagination"
	"gorm.io/gorm"
)

type KelasServiceImpl struct {
	DB *infras.PostgresqlConn
}

func ProvideKelasServiceImpl(db *infras.PostgresqlConn) *KelasServiceImpl {
	s := new(KelasServiceImpl)
	s.DB = db
	return s
}

type KelasService interface {
	ResolvePagination(req model.StandardRequest) (result pagination.Response, err error)
	ResolveAll() (res []Kelas, err error)
	ResolveByID(id string) (res Kelas, err error)
	Create(req KelasRequest, userID string) (res Kelas, err error)
	Update(req KelasRequest, id string, userID string) (res Kelas, err error)
	SoftDelete(id, userID string) (err error)
}

func (s *KelasServiceImpl) ResolvePagination(req model.StandardRequest) (result pagination.Response, err error) {
	req.SortBy = ColumnMappKelas[req.SortBy].(string)
	var count int64
	criteria := " is_deleted = false AND (nama || coalesce(kode,'')) ilike ? "
	err = s.DB.GormRead.Model(&Kelas{}).Where(criteria, "%"+req.Keyword+"%").Count(&count).Error
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	result.Items = make([]interface{}, 0)
	if count > 0 {
		var items []Kelas
		if find := s.DB.GormRead.
			Scopes(pagination.Paginate(req)).
			Where(criteria, "%"+req.Keyword+"%").
			Find(&items); find.Error == nil {
			for _, d := range items {
				result.Items = append(result.Items, d)
			}
		}
	} else {
		result.Items = make([]interface{}, 0)
	}

	result.Meta = pagination.CreateMeta(int(count), req.PageSize, req.PageNumber)

	return
}

// ResolveAll mendapatkan all data id, nama kelas untuk combobox
func (s *KelasServiceImpl) ResolveAll() (res []Kelas, err error) {
	err = s.DB.GormRead.Table(tableNameKelas).
		Select("id", "kode", "nama").
		Where("is_deleted = false").
		Order("nama asc").Scan(&res).Error
	if res == nil {
		return make([]Kelas, 0), nil
	}
	return
}

// ResolveByID mendapatkan satu data kelas berdasarkan ID nya
func (s *KelasServiceImpl) ResolveByID(id string) (res Kelas, err error) {
	err = s.DB.GormRead.First(&res, "is_deleted = false AND id=?", id).Error
	return
}

// Create menambahkan kelas baru
func (s *KelasServiceImpl) Create(req KelasRequest, userID string) (res Kelas, err error) {
	res.BindFromRequest(req, "", userID)
	if err := s.DB.GormWrite.Create(&res).Error; err != nil {
		logger.ErrorWithStack(err)
		return Kelas{}, err
	}

	return
}

// Update untuk merubah data kelas
func (s *KelasServiceImpl) Update(req KelasRequest, id string, userID string) (res Kelas, err error) {
	var current Kelas
	err = s.DB.GormRead.First(&current, "is_deleted = false AND id=?", id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.ErrorWithStack(err)
		return
	}
	current.BindFromRequest(req, id, userID)
	err = s.DB.GormWrite.Save(&current).Error

	return current, err
}

// SoftDelete untuk menghapus kelas berdasarkan id
func (s *KelasServiceImpl) SoftDelete(id, userID string) (err error) {
	var current Kelas
	err = s.DB.GormRead.First(&current, "is_deleted = false AND id=?", id).Error
	if err != nil {
		if err == sql.ErrNoRows {
			logger.ErrorWithStack(err)
			return err
		}
		logger.ErrorWithStack(err)
		return err
	}

	current.SoftDelete(userID)
	err = s.DB.GormWrite.Save(&current).Error

	return err
}
