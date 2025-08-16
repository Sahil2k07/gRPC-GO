package util

import (
	"github.com/Sahil2k07/gRPC-GO/internal/database"
	"gorm.io/gorm"
)

type transaction struct {
	Tx *gorm.DB
}

func NewTransactionScope() *transaction {
	tx := database.DB.Begin()
	return &transaction{
		Tx: tx,
	}
}

func (u *transaction) Commit() error {
	return u.Tx.Commit().Error
}

func (u *transaction) Rollback() error {
	return u.Tx.Rollback().Error
}
