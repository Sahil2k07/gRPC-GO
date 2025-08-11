package view

type (
	PageFilter struct {
		PageSize    int  `json:"pageSize"`
		CurrentPage int  `json:"currentPage"`
		AllPages    bool `json:"allPages"`
	}

	SortFilter struct {
		SortField string `json:"sortField"`
		SortOrder string `json:"sortOrder"`
	}
)
