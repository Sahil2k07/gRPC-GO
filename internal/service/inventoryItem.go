package service

import (
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type inventoryItemService struct {
	repo interfaces.InventoryItemRepository
}

func NewInventoryItemService(r interfaces.InventoryItemRepository) interfaces.InventoryItemService {
	return &inventoryItemService{repo: r}
}

func (s *inventoryItemService) GetInventoryItem(id string) (view.InventoryItemResponse, error) {
	intID, err := util.StringToUint(id)
	if err != nil {
		return view.InventoryItemResponse{}, err
	}

	item, err := s.repo.GetInventoryItem(intID)
	if err != nil {
		return view.InventoryItemResponse{}, err
	}

	return view.NewInventoryItemResponse(item), nil
}

func (s *inventoryItemService) AddInventoryItem(req view.AddInventoryItem) error {
	if err := s.repo.ValidateExistingItem(nil, req); err != nil {
		return err
	}

	return s.repo.AddInventoryItem(req)
}

func (s *inventoryItemService) UpdateInventoryItem(req view.UpdateInventoryItem) error {
	if err := s.repo.ValidateExistingItem(&req.ID, req.AddInventoryItem); err != nil {
		return err
	}

	return s.repo.UpdateInventoryItem(req)
}

func (s *inventoryItemService) DeleteInventoryItem(id string) error {
	intID, err := util.StringToUint(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteInventoryItem(intID)
}

func (s *inventoryItemService) ListInventoryItems(req view.ListInventoryItem) (view.ListResponse, error) {
	items, count, err := s.repo.ListInventoryItems(req)
	if err != nil {
		return view.ListResponse{}, err
	}

	records := util.Map(items, func(m model.InventoryItem) view.InventoryItemResponse {
		return view.NewInventoryItemResponse(m)
	})

	response := view.ListResponse{
		TotalRecords: count,
		Data:         records,
	}

	return response, nil
}
