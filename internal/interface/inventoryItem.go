package interfaces

import (
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
	"gorm.io/gorm"
)

type (
	InventoryItemRepository interface {
		ValidateExistingItem(id *uint, req view.AddInventoryItem) error

		GetInventoryItem(id uint) (model.InventoryItem, error)

		AddInventoryItem(req view.AddInventoryItem) error

		UpdateInventoryItem(item *model.InventoryItem, tx ...*gorm.DB) error

		DeleteInventoryItem(id uint) error

		ValidateExistingGroup(id uint) error

		ListInventoryItems(req view.ListInventoryItem) ([]model.InventoryItem, int64, error)

		GetInventoryItemsFromCodes(ids []string) ([]model.InventoryItem, error)
	}

	InventoryItemService interface {
		GetInventoryItem(id string) (view.InventoryItemResponse, error)

		AddInventoryItem(req view.AddInventoryItem) error

		UpdateInventoryItem(req view.UpdateInventoryItem) error

		DeleteInventoryItem(id string) error

		ListInventoryItems(req view.ListInventoryItem) (view.ListResponse, error)

		UpdateInventoryItemStock(req []view.UpdateInventoryStock) error
	}
)
