package view

import "github.com/Sahil2k07/gRPC-GO/internal/model"

type (
	AddInventoryItem struct {
		Code             string  `json:"code" validate:"required"`
		Name             string  `json:"name" validate:"required"`
		Description      string  `json:"description" validate:"required"`
		InventoryGroupID uint    `json:"inventoryGroupID" validate:"required"`
		Quantity         float32 `json:"quantity" validate:"required,min=1"`
		Price            float32 `json:"price" validate:"required,min=1"`
	}

	UpdateInventoryItem struct {
		ID uint `json:"id" validate:"required"`
		AddInventoryItem
	}

	ListInventoryItem struct {
		Code               string     `json:"code"`
		Name               string     `json:"name"`
		InventoryGroupCode string     `json:"inventoryGroupCode"`
		InventoryGroupName string     `json:"inventoryGroupName"`
		InventoryGroupID   *uint      `json:"inventoryGroupID"`
		Page               PageFilter `json:"page"`
		Sort               SortFilter `json:"sort"`
	}

	InventoryItemResponse struct {
		ID               uint    `json:"id"`
		GroupCode        string  `json:"groupCode"`
		InventoryGroupID uint    `json:"inventoryGroupID"`
		Code             string  `json:"code"`
		Name             string  `json:"name"`
		Description      string  `json:"description"`
		InventoryGroup   string  `json:"inventoryGroup"`
		Quantity         float32 `json:"quantity"`
		Price            float32 `json:"price"`
	}

	UpdateInventoryStock struct {
		Code     string  `json:"code" validate:"required"`
		Quantity float32 `json:"quantity" validate:"required"`
	}
)

func NewInventoryItemResponse(m model.InventoryItem) InventoryItemResponse {
	return InventoryItemResponse{
		ID:               m.ID,
		GroupCode:        m.InventoryGroup.Code,
		InventoryGroupID: m.InventoryGroupID,
		Code:             m.Code,
		Name:             m.Name,
		Description:      m.Description,
		InventoryGroup:   m.InventoryGroup.Name,
		Quantity:         m.Quantity,
		Price:            m.Price,
	}
}
