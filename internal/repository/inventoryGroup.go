package repository

import (
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	errz "github.com/Sahil2k07/gRPC-GO/internal/error"
	interfaces "github.com/Sahil2k07/gRPC-GO/internal/interface"
	"github.com/Sahil2k07/gRPC-GO/internal/model"
	"github.com/Sahil2k07/gRPC-GO/internal/util"
	"github.com/Sahil2k07/gRPC-GO/internal/view"
	"gorm.io/gorm"
)

type inventoryGroupRepository struct{}

func NewInventoryGroupRepository() interfaces.InventoryGroupRepository {
	return &inventoryGroupRepository{}
}

func (r *inventoryGroupRepository) GetInventoryGroup(id uint) (model.InventoryGroup, error) {
	var m model.InventoryGroup

	if err := database.DB.First(&m, id).Error; err != nil {
		return model.InventoryGroup{}, err
	}

	return m, nil
}

func (r *inventoryGroupRepository) ValidateExistingCode(id uint, req view.AddInventoryGroup) error {
	q := database.DB.Model(&model.InventoryGroup{}).Where("code = ?", req.Code)

	if id > 0 {
		q = q.Where("id != ?", id)
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errz.NewAlreadyExists("Item Group with same code already exist")
	}

	return nil
}

func (r *inventoryGroupRepository) AddInventoryGroup(req view.AddInventoryGroup) error {
	m := model.InventoryGroup{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := database.DB.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryGroupRepository) UpdateInventoryGroup(req view.InventoryGroupView) error {
	var model model.InventoryGroup

	if err := database.DB.First(&model, req.ID).Error; err != nil {
		return nil
	}

	model.Name = req.Name
	model.Description = req.Description
	model.Code = req.Code

	if err := database.DB.Save(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *inventoryGroupRepository) DeleteInventoryGroup(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		itemQuery := tx.Unscoped().Model(&model.InventoryItem{}).Where("inventory_group_id = ?", id)

		if err := itemQuery.Delete(&model.InventoryItem{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(&model.InventoryGroup{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *inventoryGroupRepository) ListInventoryGroup(req view.ListInventoryGroup) ([]model.InventoryGroup, int64, error) {
	q := database.DB.Model(&model.InventoryGroup{})

	if req.Code != "" {
		q = q.Where("code ILIKE ?", "%"+req.Code+"%")
	}

	if req.Name != "" {
		q = q.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	q = util.AddPagination(q, req.Page, req.Sort)

	var models []model.InventoryGroup
	if err := q.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	return models, count, nil
}
