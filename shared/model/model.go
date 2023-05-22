package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

const RECORD_NOT_FOUND = "record not found"

// StandardRequest is a standard query string request
type StandardRequest struct {
	Keyword      string `json:"q" validate:"omitempty"`
	StartDate    string `json:"startDate" validate:"omitempty"`
	EndDate      string `json:"endDate" validate:"omitempty"`
	PageNumber   int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize     int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy       string `json:"sortBy" validate:"required"`
	SortType     string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status       string `json:"status" validate:"omitempty"`
	IgnorePaging bool   `json:"ignorePaging" validate:"omitempty"`
}
type StandardRequestPegawai struct {
	Keyword         string `json:"q" validate:"omitempty"`
	StartDate       string `json:"startDate" validate:"omitempty"`
	EndDate         string `json:"endDate" validate:"omitempty"`
	PageNumber      int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize        int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy          string `json:"sortBy" validate:"required"`
	SortType        string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status          string `json:"status" validate:"omitempty"`
	IdStatusPegawai string `json:"idStatusPegawai" validate:"omitempty"`
	IgnorePaging    bool   `json:"ignorePaging" validate:"omitempty"`
}
type StandardModel struct {
	ID     string `json:"id" db:"id"`
	Nama   string `json:"nama" db:"nama"`
	Alamat string `json:"alamat" db:"alamat"`
}
type ReservasiRequest struct {
	Keyword    string `json:"q" validate:"omitempty"`
	StartDate  string `json:"startDate" validate:"omitempty"`
	EndDate    string `json:"endDate" validate:"omitempty"`
	IDCabang   string `json:"idCabang" validate:"omitempty"`
	IDPegawai  string `json:"idPegawai" validate:"omitempty"`
	PageNumber int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize   int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy     string `json:"sortBy" validate:"required"`
	SortType   string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}
type ReportRequestParams struct {
	IDOpd     uuid.UUID `json:"id" db:"id"`
	IdItem    string    `json:"idItem" validate:"omitempty"`
	IdBidang  string    `json:"idBidang" validate:"omitempty"`
	StartDate string    `json:"startDate" validate:"omitempty"`
	EndDate   string    `json:"endDate" validate:"omitempty"`
	Status    string    `json:"status"`
}

type ScanBarcodeRequestParams struct {
	ID            string          `json:"id" validate:"omitempty"`
	ProduksiID    string          `json:"produksiId" validate:"omitempty"`
	Nomor         string          `json:"nomor" validate:"omitempty"`
	Step          int             `json:"step" validate:"omitempty"`
	KelompokID    string          `json:"kelompokId" validate:"omitempty"`
	KelompokDetID string          `json:"kelompokDetId" validate:"omitempty"`
	PlastikKe     int             `json:"plastikKe" validate:"omitempty"`
	BeratTimbang  decimal.Decimal `json:"beratTimbang" validate:"omitempty"`
	BeratBahan    decimal.Decimal `json:"beratBahan" validate:"omitempty"`
	Toleransi     decimal.Decimal `json:"toleransi" validate:"omitempty"`
}

type StandardRequestJabatan struct {
	Keyword       string `json:"q" validate:"omitempty"`
	StartDate     string `json:"startDate" validate:"omitempty"`
	EndDate       string `json:"endDate" validate:"omitempty"`
	PageNumber    int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize      int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy        string `json:"sortBy" validate:"required"`
	SortType      string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status        string `json:"status" validate:"omitempty"`
	IgnorePaging  bool   `json:"ignorePaging" validate:"omitempty"`
	IdTipeJabatan string `json:"idTipeJabatan" validate:"omitempty"`
}
type RequestMatriksPeran struct {
	Keyword      string `json:"q" validate:"omitempty"`
	StartDate    string `json:"startDate" validate:"omitempty"`
	EndDate      string `json:"endDate" validate:"omitempty"`
	PageNumber   int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize     int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy       string `json:"sortBy" validate:"required"`
	SortType     string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	IdIndikator  string `json:"idIndikator" validate:"required"`
	Status       string `json:"status" validate:"omitempty"`
	IgnorePaging bool   `json:"ignorePaging" validate:"omitempty"`
}

// JSONRaw ...
type StandardRequestAktivitasHarian struct {
	Keyword      string    `json:"q" validate:"omitempty"`
	StartDate    string    `json:"startDate" validate:"omitempty"`
	EndDate      string    `json:"endDate" validate:"omitempty"`
	PageNumber   int       `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize     int       `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy       string    `json:"sortBy" validate:"required"`
	SortType     string    `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status       string    `json:"status" validate:"omitempty"`
	IgnorePaging bool      `json:"ignorePaging" validate:"omitempty"`
	IdPegawai    uuid.UUID `json:"idPegawai" validate:"omitempty"`
	IdIki        string    `json:"idIki" validate:"omitempty"`
}

type StandardRequestProgram struct {
	Keyword     string `json:"q" validate:"omitempty"`
	IdJenisMbkm string `json:"idJenisMbkm" validate:"omitempty"`
	PageNumber  int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize    int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy      string `json:"sortBy" validate:"required"`
	SortType    string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}
type StandardRequestLookupProgram struct {
	Keyword     string `json:"q" validate:"omitempty"`
	IdProgram   string `json:"idProgram" validate:"omitempty"`
	IdJenisMbkm string `json:"idJenisMbkm" validate:"omitempty"`
	PageNumber  int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize    int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy      string `json:"sortBy" validate:"required"`
	SortType    string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}

type StandardRequestSyaratLuaran struct {
	Keyword        string `json:"q" validate:"omitempty"`
	IdJenisLuaran  string `json:"idJenisLuaran" validate:"omitempty"`
	IdProgram      string `json:"idProgram" validate:"omitempty"`
	IdProgramProdi string `json:"idProgramProdi" validate:"omitempty"`
	PageNumber     int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize       int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy         string `json:"sortBy" validate:"required"`
	SortType       string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}

type StandardRequestPosisiKegiatanTopik struct {
	Keyword    string `json:"q" validate:"omitempty"`
	IdProgram  string `json:"idProgram" validate:"omitempty"`
	IdMitra    string `json:"idMitra" validate:"omitempty"`
	PageNumber int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize   int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy     string `json:"sortBy" validate:"required"`
	SortType   string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}
type StandardRequestProgramDitawarkan struct {
	Keyword        string `json:"q" validate:"omitempty"`
	IdProgram      string `json:"idProgram" validate:"omitempty"`
	IdProgramProdi string `json:"idProgramProdi" validate:"omitempty"`
	PageNumber     int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize       int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy         string `json:"sortBy" validate:"required"`
	SortType       string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}

type StandardRequestPicMitra struct {
	Keyword    string `json:"q" validate:"omitempty"`
	IdMitra    string `json:"idMitra" validate:"omitempty"`
	PageNumber int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize   int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy     string `json:"sortBy" validate:"required"`
	SortType   string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
}

type StandardRequestLogbookMagang struct {
	Keyword      string `json:"q" validate:"omitempty"`
	StartDate    string `json:"startDate" validate:"omitempty"`
	EndDate      string `json:"endDate" validate:"omitempty"`
	PageNumber   int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize     int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy       string `json:"sortBy" validate:"required"`
	SortType     string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status       string `json:"status" validate:"omitempty"`
	IgnorePaging bool   `json:"ignorePaging" validate:"omitempty"`
	IdPendaftar  string `json:"idPendaftar" validate:"omitempty"`
}

type StandardRequestMagang struct {
	Keyword        string `json:"q" validate:"omitempty"`
	StartDate      string `json:"startDate" validate:"omitempty"`
	EndDate        string `json:"endDate" validate:"omitempty"`
	PageNumber     int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize       int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy         string `json:"sortBy" validate:"required"`
	SortType       string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	Status         string `json:"status" validate:"omitempty"`
	IgnorePaging   bool   `json:"ignorePaging" validate:"omitempty"`
	IdMahasiswa    string `json:"idMahasiswa" validate:"omitempty"`
	IdPegawai      string `json:"idPegawai" validate:"omitempty"`
	IdRole         string `json:"idRole" validate:"omitempty"`
	IdProgramProdi string `json:"idProgramProdi" validate:"omitempty"`
	IdProgram      string `json:"idProgram" validate:"omitempty"`
	IdJenisMbkm    string `json:"idJenisMbkm" validate:"omitempty"`
	StatusMbkm     string `json:"statusMbkm" validate:"omitempty"`
	ProgramId      string `json:"ProgramId" validate:"omitempty"`
	ProgramProdiId string `json:"programProdiId" validate:"omitempty"`
	IdPeriode      string `json:"idPeriode" validate:"omitempty"`
}

type StandardRequestUser struct {
	Keyword        string `json:"q" validate:"omitempty"`
	StartDate      string `json:"startDate" validate:"omitempty"`
	EndDate        string `json:"endDate" validate:"omitempty"`
	PageNumber     int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize       int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy         string `json:"sortBy" validate:"required"`
	SortType       string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	IdRole         string `json:"idRole" validate:"omitempty"`
	IdProgramProdi string `json:"idProgramProdi" validate:"omitempty"`
}

type StandardRequestLuaranMagang struct {
	Keyword            string `json:"q" validate:"omitempty"`
	StartDate          string `json:"startDate" validate:"omitempty"`
	EndDate            string `json:"endDate" validate:"omitempty"`
	PageNumber         int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize           int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy             string `json:"sortBy" validate:"required"`
	SortType           string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	IdPendaftarProgram string `json:"idPendaftarProgram" validate:"omitempty"`
}

type StandardRequestMitraTerlibatProgram struct {
	Keyword    string `json:"q" validate:"omitempty"`
	StartDate  string `json:"startDate" validate:"omitempty"`
	EndDate    string `json:"endDate" validate:"omitempty"`
	PageNumber int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize   int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy     string `json:"sortBy" validate:"required"`
	SortType   string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	IdProgram  string `json:"idProgram" validate:"omitempty"`
	IdMitra    string `json:"idMitra" validate:"omitempty"`
}

type StandardRequestMenu struct {
	Keyword    string `json:"q" validate:"omitempty"`
	PageNumber int    `json:"pageNumber" validate:"omitempty,gte=0"`
	PageSize   int    `json:"pageSize" validate:"omitempty,gte=0"`
	SortBy     string `json:"sortBy" validate:"required"`
	SortType   string `json:"sortType" validate:"required,oneof=asc ASC desc DESC"`
	App        string `json:"app" validate:"omitempty"`
}

type ResponseApi struct {
	Success bool         `json:"success"`
	Data    *interface{} `json:"data"`
	Message *string      `json:"message,omitempty"`
}

type ResponseStandartApi struct {
	Success   bool         `json:"success"`
	Data      *interface{} `json:"data"`
	Message   *string      `json:"message,omitempty"`
	Page      int          `json:"page,omitempty"`
	TotalData int          `json:"total_data,omitempty"`
}

// JSONRaw ...
type JSONRaw json.RawMessage

// Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)

	return driver.Value(byteArr), nil
}

// Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}

	return nil
}

// MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

// UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
