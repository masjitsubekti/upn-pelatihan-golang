package pagination

import (
	"math"

	"gitlab.com/upn-belajar-go/shared/model"
	"gorm.io/gorm"
)

// Response is a standard list data
type Response struct {
	Items []interface{} `json:"items"`
	Meta  Metadata      `json:"meta"`
}

// Metadata is a additional info for list data
type Metadata struct {
	TotalItems   int `json:"totalItems"`
	TotalPage    int `json:"totalPage"`
	PreviousPage int `json:"previousPage"`
	CurrentPage  int `json:"currentPage"`
	NextPage     int `json:"nextPage"`
	LimitPerPage int `json:"limitPerPage"`
}

// CreateMeta is a metadata creator
func CreateMeta(totalItems int, dataPerPage int, pageNumber int) (meta Metadata) {
	totalPageRaw := float64(totalItems) / float64(dataPerPage)
	maxPage := int(math.Ceil(totalPageRaw))
	minPage := 1

	if maxPage < minPage {
		maxPage = minPage
	}

	nextPage := pageNumber + 1
	if nextPage > maxPage {
		nextPage = maxPage
	}

	prevPage := pageNumber - 1
	if prevPage < minPage {
		prevPage = minPage
	}

	return Metadata{
		TotalItems:   totalItems,
		TotalPage:    maxPage,
		PreviousPage: prevPage,
		CurrentPage:  pageNumber,
		NextPage:     nextPage,
		LimitPerPage: dataPerPage,
	}
}

func Paginate(req model.StandardRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := req.PageNumber
		if req.PageNumber == 0 {
			page = 1
		}

		pageSize := req.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize).Order(req.SortBy + " " + req.SortType)
	}
}
