package view

type (
	PageFilter struct {
		PageSize int  `json:"pageSize" validate:"required,min=10,max=100"`
		PageNum  int  `json:"pageNum" validate:"required,min=1"`
		AllPages bool `json:"allPages"`
	}

	SortFilter struct {
		SortField string `json:"sortField"`
		SortOrder string `json:"sortOrder"`
	}
)
