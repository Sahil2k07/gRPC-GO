package view

import "github.com/Sahil2k07/gRPC-GO/internal/model"

type (
	AddInventoryGroup struct {
		Code        string `json:"code" validate:"required"`
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	InventoryGroupView struct {
		ID uint `json:"id" validate:""`
		AddInventoryGroup
	}

	ListInventoryGroup struct {
		Code string     `json:"code"`
		Name string     `json:"name"`
		Page PageFilter `json:"page"`
		Sort SortFilter `json:"sort"`
	}
)

func NewInventoryGroupView(m model.InventoryGroup) InventoryGroupView {
	return InventoryGroupView{
		ID: m.ID,
		AddInventoryGroup: AddInventoryGroup{
			Code:        m.Code,
			Name:        m.Name,
			Description: m.Description,
		},
	}
}
