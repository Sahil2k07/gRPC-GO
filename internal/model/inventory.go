package model

import "gorm.io/gorm"

type (
	InventoryGroup struct {
		gorm.Model
		Code           string          `gorm:"not null; uniqueIndex"`
		Name           string          `gorm:"not null"`
		Description    string          `gorm:"not null"`
		InventoryItems []InventoryItem `gorm:"foreignKey:InventoryGroupID"`
	}

	InventoryItem struct {
		gorm.Model
		Code             string         `gorm:"not null; uniqueIndex"`
		Name             string         `gorm:"not null"`
		Description      string         `gorm:"not null"`
		Quantity         float32        `gorm:"not null"`
		Price            float32        `gorm:"not null"`
		InventoryGroupID uint           `gorm:"not null"`
		InventoryGroup   InventoryGroup `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
