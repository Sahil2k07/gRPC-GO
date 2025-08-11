package database

import (
	"log"
	"os"
	"time"

	"github.com/Sahil2k07/gRPC-GO/internal/config"

	charmbracelet "github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	isProduction := config.IsProduction()

	var gormLog gormLogger.Interface

	if isProduction {
		gormLog = gormLogger.Default.LogMode(gormLogger.Error)
	} else {
		stdLogger := log.New(os.Stdout, "\r\n", log.LstdFlags)
		gormLog = gormLogger.New(stdLogger, gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		})
	}

	postgresDSN := config.GetDBConfig()

	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		charmbracelet.Errorf("failed to connect to DB: %v", err)
		panic("Database was not found")
	}

	DB = db
}
