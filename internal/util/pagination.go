package util

import (
	"strings"

	"github.com/Sahil2k07/gRPC-GO/internal/view"

	"gorm.io/gorm"
)

func AddPagination(db *gorm.DB, pf view.PageFilter, sf view.SortFilter) *gorm.DB {
	if sf.SortField != "" {
		order := "asc"
		if strings.ToLower(sf.SortOrder) == "desc" {
			order = "desc"
		}
		db = db.Order(sf.SortField + " " + order)
	}

	if !pf.AllPages {
		if pf.PageSize <= 0 {
			pf.PageSize = 10
		}
		if pf.PageNum <= 0 {
			pf.PageNum = 1
		}
		offset := (pf.PageNum - 1) * pf.PageSize
		db = db.Offset(offset).Limit(pf.PageSize)
	}

	return db
}
