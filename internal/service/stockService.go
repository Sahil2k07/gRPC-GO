package service

import (
	"context"
	"fmt"

	stock "github.com/Sahil2k07/gRPC-GO/internal/generated/stock/proto"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type stockService struct {
	stock.UnimplementedStockServiceServer
	repo interfaces.InventoryItemRepository
}

func NewStockService(r interfaces.InventoryItemRepository) interfaces.StockService {
	return &stockService{repo: r}
}

func (s *stockService) ListStockItems(ctx context.Context, req *stock.ListStockItemRequest) (*stock.StockItemResponse, error) {

	filters := view.ListInventoryItem{
		Code:               "",
		Name:               "",
		InventoryGroupCode: "",
		InventoryGroupName: "",
		InventoryGroupID:   nil,
		Page: view.PageFilter{
			PageSize: int(req.PageSize),
			PageNum:  int(req.PageNum),
		},
	}

	if req.Code != nil {
		filters.Code = *req.Code
	}
	if req.Name != nil {
		filters.Name = *req.Name
	}
	if req.GroupCode != nil {
		filters.InventoryGroupCode = *req.GroupCode
	}
	if req.GroupName != nil {
		filters.InventoryGroupName = *req.GroupName
	}

	models, count, err := s.repo.ListInventoryItems(filters)
	if err != nil {
		return &stock.StockItemResponse{}, fmt.Errorf("error while listing items: %v", err)
	}

	var items []*stock.StockItem

	for _, m := range models {
		item := stock.StockItem{
			Code:        m.Code,
			Name:        m.Name,
			Description: m.Description,
		}

		items = append(items, &item)
	}

	response := &stock.StockItemResponse{
		TotalRecords: count,
		Items:        items,
	}

	return response, nil
}

func (s *stockService) CheckStockAvailability(ctx context.Context, req *stock.StockAvailabilityRequest) (*stock.StockAvailabilityResponse, error) {
	var codes []string

	for _, r := range req.Items {
		codes = append(codes, r.Code)
	}

	models, err := s.repo.GetInventoryItemsFromCodes(codes)
	if err != nil {
		return &stock.StockAvailabilityResponse{}, fmt.Errorf("error while reading data")
	}

	mappings := make(map[string]model.InventoryItem, len(models))
	for _, md := range models {
		mappings[md.Code] = md
	}

	var items []*stock.StockAvailability
	for _, ri := range req.Items {
		item := stock.StockAvailability{
			Code:              ri.Code,
			Available:         false,
			AvailableQuantity: nil,
		}

		mp, exists := mappings[ri.Code]
		if exists {
			item.AvailableQuantity = &mp.Quantity

			if ri.Quantity <= mp.Quantity {
				item.Available = true
			}
		}

		items = append(items, &item)
	}

	return &stock.StockAvailabilityResponse{Items: items}, nil
}

func (s *stockService) ConsumeStock(ctx context.Context, req *stock.ConsumeStockRequest) (*stock.StockConsumptionResponse, error) {
	var codes []string
	for _, r := range req.Items {
		codes = append(codes, r.Code)
	}

	models, err := s.repo.GetInventoryItemsFromCodes(codes)
	if err != nil {
		return &stock.StockConsumptionResponse{}, fmt.Errorf("error while reading inventory data")
	}

	mappings := make(map[string]*model.InventoryItem, len(models))
	for i := range models {
		m := &models[i]
		mappings[m.Code] = m
	}

	for _, r := range req.Items {
		m, exist := mappings[r.Code]
		if !exist {
			return &stock.StockConsumptionResponse{
				Items: []*stock.StockConsumption{
					{
						Code:    r.Code,
						Success: false,
						Message: s.ptrString(fmt.Sprintf("Item with code %s not available", r.Code)),
					},
				},
			}, nil
		}
		if m.Quantity < r.Quantity {
			return &stock.StockConsumptionResponse{
				Items: []*stock.StockConsumption{
					{
						Code:    r.Code,
						Success: false,
						Message: s.ptrString(fmt.Sprintf("Not enough quantity for code %s", r.Code)),
					},
				},
			}, nil
		}
	}

	tx := util.NewTransactionScope()
	defer tx.Rollback()

	for _, r := range req.Items {
		m := mappings[r.Code]
		m.Quantity -= r.Quantity
		if err := s.repo.UpdateInventoryItem(m, tx.Tx); err != nil {
			return &stock.StockConsumptionResponse{
				Items: []*stock.StockConsumption{
					{
						Code:    r.Code,
						Success: false,
						Message: s.ptrString(fmt.Sprintf("Error while updating stock: %v", err)),
					},
				},
			}, nil
		}
	}

	if err := tx.Commit(); err != nil {
		return &stock.StockConsumptionResponse{}, fmt.Errorf("error while updating stock: %v", err)
	}

	var itemsResp []*stock.StockConsumption
	for _, r := range req.Items {
		itemsResp = append(itemsResp, &stock.StockConsumption{
			Code:    r.Code,
			Success: true,
			Message: s.ptrString("Stock updated successfully"),
		})
	}

	return &stock.StockConsumptionResponse{Items: itemsResp}, nil
}

func (s *stockService) ptrString(str string) *string {
	return &str
}
