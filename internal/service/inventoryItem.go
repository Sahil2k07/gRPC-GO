package service

import (
	"fmt"

	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
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
	item, err := s.repo.GetInventoryItem(req.ID)
	if err != nil {
		return err
	}

	if req.Code != item.Code {
		if err := s.repo.ValidateExistingItem(&req.ID, req.AddInventoryItem); err != nil {
			return err
		}
	}

	item.Code = req.Code
	item.Name = req.Name
	item.Description = req.Description
	item.Price = req.Price
	item.Quantity = req.Quantity

	return s.repo.UpdateInventoryItem(&item)
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

func (s *inventoryItemService) UpdateInventoryItemStock(req []view.UpdateInventoryStock) error {
	if len(req) == 0 {
		return nil
	}

	var codes []string
	for _, r := range req {
		codes = append(codes, r.Code)
	}

	items, err := s.repo.GetInventoryItemsFromCodes(codes)
	if err != nil {
		return err
	}

	itemMap := make(map[string]*model.InventoryItem)
	for i := range items {
		itemMap[items[i].Code] = &items[i]
	}

	for _, r := range req {
		item, exists := itemMap[r.Code]
		if !exists {
			return errz.NewNotFound(fmt.Sprintf("inventory item with Code %s not found", r.Code))
		}

		newQty := item.Quantity + r.Quantity
		if newQty < 0 {
			return errz.NewValidation(fmt.Sprintf("insufficient stock for item code %s", r.Code))
		}

		item.Quantity = newQty
	}

	for _, item := range itemMap {
		if err := s.repo.UpdateInventoryItem(item); err != nil {
			return err
		}
	}

	return nil
}
