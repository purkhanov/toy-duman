package database

import (
	"fmt"
	"log"
	"toy-duman/config"
	"toy-duman/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func initModels() error {
	models := []any{
		&model.User{},
		&model.Status{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("Error auto migrating model: %v", err)
			return err
		}
	}

	return nil
}

// func NewDB() (*db, error) {
// 	d := &db{}
// 	if err := d.initDB(); err != nil {
// 		return nil, fmt.Errorf("failed to initialize database: %w", err)
// 	}

// 	// if err := initModels(); err != nil {
// 	// 	return nil, fmt.Errorf("failed to initialize models: %w", err)
// 	// }

// 	// if err := initUser(); err != nil {
// 	// 	return nil, fmt.Errorf("failed to initialize user: %w", err)
// 	// }

// 	return d, nil
// }

func InitDB() error {
	var err error
	var gormLogger logger.Interface

	// if config.IsDebug() {
	// 	gormLogger = logger.Default
	// } else {
	// 	gormLogger = logger.Discard
	// }

	c := &gorm.Config{
		Logger: gormLogger,
	}

	db, err = gorm.Open(postgres.Open(getDSN()), c)
	if err != nil {
		return err
	}

	if err := initModels(); err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func getDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "postgres"),
		config.GetEnv("DB_NAME", "toy_duman"),
	)
}
