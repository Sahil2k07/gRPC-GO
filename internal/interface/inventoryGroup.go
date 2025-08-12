package interfaces

import (
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type (
	InventoryGroupRepository interface {
		GetInventoryGroup(id uint) (model.InventoryGroup, error)

		AddInventoryGroup(req view.AddInventoryGroup) error

		UpdateInventoryGroup(req view.InventoryGroupView) error

		DeleteInventoryGroup(id uint) error

		ListInventoryGroup(req view.ListInventoryGroup) ([]model.InventoryGroup, int64, error)

		ValidateExistingCode(id uint, req view.AddInventoryGroup) error
	}

	InventoryGroupService interface {
		GetInventoryGroup(id string) (view.InventoryGroupView, error)

		AddInventoryGroup(req view.AddInventoryGroup) error

		UpdateInventoryGroup(req view.InventoryGroupView) error

		DeleteInventoryGroup(id string) error

		ListInventoryGroup(req view.ListInventoryGroup) (view.ListResponse, error)
	}
)
