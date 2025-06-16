package database

import (
	"fmt"
	"toy-duman/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type db struct {
	db *gorm.DB
}

func NewDB() (*db, error) {
	d := &db{}
	if err := d.initDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// if err := initModels(); err != nil {
	// 	return nil, fmt.Errorf("failed to initialize models: %w", err)
	// }

	// if err := initUser(); err != nil {
	// 	return nil, fmt.Errorf("failed to initialize user: %w", err)
	// }

	return d, nil
}

func (d *db) initDB() error {
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

	d.db, err = gorm.Open(postgres.Open(d.getDSN()), c)
	if err != nil {
		return err
	}

	// if err := d.initModels(); err != nil {
	// 	return err
	// }

	return nil
}

func (d *db) GetDB() *gorm.DB {
	return d.db
}

func (d *db) CloseDB() error {
	if d.db != nil {
		sqlDB, err := d.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func (d *db) getDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "postgres"),
		config.GetEnv("DB_NAME", "toy_duman"),
	)
}
