package view

type ListResponse struct {
	TotalRecords int64 `json:"totalRecords"`
	Data         any   `json:"data"`
}
