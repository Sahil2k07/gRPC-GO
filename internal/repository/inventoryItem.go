package repository

import (
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
)

type inventoryItemRepository struct{}

func NewInventoryItemRepository() interfaces.InventoryItemRepository {
	return &inventoryItemRepository{}
}

func (r *inventoryItemRepository) ValidateExistingItem(id *uint, req view.AddInventoryItem) error {
	q := database.DB.Model(&model.InventoryItem{})

	q = q.Where("code = ?", req.Code)

	if id != nil {
		q = q.Where("id != ?", *id)
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errz.NewAlreadyExists("Item with same code already exist")
	}

	return nil
}

func (r *inventoryItemRepository) GetInventoryItem(id uint) (model.InventoryItem, error) {
	var item model.InventoryItem

	if err := database.DB.Preload("InventoryGroup").First(&item, id).Error; err != nil {
		return model.InventoryItem{}, err
	}

	return item, nil
}

func (r *inventoryItemRepository) ValidateExistingGroup(id uint) error {
	var count int64
	err := database.DB.Model(&model.InventoryGroup{}).
		Where("id = ?", id).
		Count(&count).Error

	if err != nil {
		return nil
	}

	if count < 1 {
		return errz.NewNotFound("Group does not exist")
	}

	return nil
}

func (r *inventoryItemRepository) AddInventoryItem(req view.AddInventoryItem) error {
	if err := r.ValidateExistingGroup(req.InventoryGroupID); err != nil {
		return err
	}

	item := model.InventoryItem{
		Code:             req.Code,
		Name:             req.Name,
		Description:      req.Description,
		Quantity:         req.Quantity,
		Price:            req.Price,
		InventoryGroupID: req.InventoryGroupID,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryItemRepository) UpdateInventoryItem(req view.UpdateInventoryItem) error {
	item, err := r.GetInventoryItem(req.ID)
	if err != nil {
		return err
	}

	item.Code = req.Code
	item.Name = req.Name
	item.Description = req.Description
	item.Quantity = req.Quantity
	item.Price = req.Price

	if err := database.DB.Save(&item).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryItemRepository) DeleteInventoryItem(id uint) error {
	if err := database.DB.Delete(&model.InventoryItem{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryItemRepository) ListInventoryItems(req view.ListInventoryItem) ([]model.InventoryItem, int64, error) {
	q := database.DB.Model(&model.InventoryItem{}).Preload("InventoryGroup")

	q = q.Joins("JOIN inventory_groups AS ig ON inventory_items.inventory_group_id = ig.id")

	if req.InventoryGroupID != nil {
		q = q.Where("inventory_items.inventory_group_id = ?", *req.InventoryGroupID)
	}

	if req.InventoryGroupName != "" {
		q = q.Where("ig.name ILIKE ?", "%"+req.InventoryGroupName+"%")
	}

	if req.InventoryGroupCode != "" {
		q = q.Where("ig.code ILIKE ?", "%"+req.InventoryGroupCode+"%")
	}

	if req.Code != "" {
		q = q.Where("inventory_items.code ILIKE ?", "%"+req.Code+"%")
	}

	if req.Name != "" {
		q = q.Where("inventory_items.name ILIKE ?", "%"+req.Name+"%")
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	q = util.AddPagination(q, req.Page, req.Sort)

	var items []model.InventoryItem
	if err := q.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, count, nil
}
