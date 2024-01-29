package database

import (
	"currency_api/internal/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DatabaseInterface interface {
	CurrencyDataBaseInterface
}

type DataBase struct {
	db *gorm.DB
}

func InitDB(conf models.DBConfig) (DatabaseInterface, error) {
	conn := sqlite.Open(fmt.Sprintf("./db/%s.db", conf.DataBase))
	db, err := gorm.Open(conn)
	if err != nil {
		log.Printf("Failed to initialize GORM with SQLite dialect: %v", err)
		return nil, err
	}

	if !db.Migrator().HasTable(&models.CurrencyDB{}) {
		db.Migrator().CreateTable(&models.CurrencyDB{})
	}
	sl := &DataBase{db: db}

	return sl, nil
}
