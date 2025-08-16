package service

import (
	"context"
	"fmt"

	"github.com/Sahil2k07/gRPC-GO/internal/database"
	stock "github.com/Sahil2k07/gRPC-GO/internal/generated/stock/proto"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type stockService struct {
	stock.UnimplementedStockServiceServer
}

func NewStockService() interfaces.StockService {
	return &stockService{}
}

func (s *stockService) ListStockItems(ctx context.Context, req *stock.ListStockItemRequest) (*stock.StockItemResponse, error) {
	q := database.DB.Model(&model.InventoryItem{}).Preload("InventoryGroup")

	q = q.Joins("JOIN inventory_groups AS ig ON inventory_items.inventory_group_id = ig.id")

	if req.GroupName != nil {
		q = q.Where("ig.name ILIKE ?", "%"+*req.GroupName+"%")
	}

	if req.GroupCode != nil {
		q = q.Where("ig.code ILIKE ?", "%"+*req.GroupCode+"%")
	}

	if req.Code != nil {
		q = q.Where("inventory_items.code ILIKE ?", "%"+*req.Code+"%")
	}

	if req.Name != nil {
		q = q.Where("inventory_items.name ILIKE ?", "%"+*req.Name+"%")
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return &stock.StockItemResponse{}, err
	}

	q = util.AddPagination(q, view.PageFilter{PageSize: int(req.PageSize), PageNum: int(req.PageNum)}, view.SortFilter{})

	var models []model.InventoryItem
	if err := q.Find(&models).Error; err != nil {
		return &stock.StockItemResponse{}, err
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

	var models []model.InventoryItem
	if err := database.DB.Where("code IN ?", codes).Find(&models).Error; err != nil {
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

	var models []model.InventoryItem
	if err := database.DB.Where("code IN ?", codes).Find(&models).Error; err != nil {
		return &stock.StockConsumptionResponse{}, fmt.Errorf("error while reading data")
	}

	mappings := make(map[string]model.InventoryItem, len(models))
	for _, m := range models {
		mappings[m.Code] = m
	}

	var items []*stock.StockConsumption
	for _, r := range req.Items {
		m, exist := mappings[r.Code]
		msg := fmt.Sprintf("Item with code %s not available", r.Code)
		success := exist

		if exist {
			newQty := m.Quantity - r.Quantity
			msg = "Stock updated successfully"

			if newQty < 0 {
				newQty = 0
				success = false
				msg = "Not enough quantity"
			} else {
				m.Quantity = newQty
				if err := database.DB.Save(&m).Error; err != nil {
					msg = fmt.Sprintf("Error while update stock: %v", err)
				}
			}
		}

		items = append(items, &stock.StockConsumption{
			Code:    r.Code,
			Success: success,
			Message: &msg,
		})
	}

	return &stock.StockConsumptionResponse{Items: items}, nil
}
