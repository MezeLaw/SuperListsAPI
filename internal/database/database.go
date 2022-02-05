package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var AppDatabase *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=meze password=meze dbname=superlists port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	AppDatabase = db
}
