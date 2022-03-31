package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var AppDatabase *gorm.DB

func InitDatabase() {

	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresIp := os.Getenv("POSTGRES_IP")
	PostgresPort := os.Getenv("POSTGRES_PORT")
	PostgresDb := os.Getenv("POSTGRES_DB")

	dsn := "host=" + PostgresIp + " user=" + PostgresUser + " password=" + PostgresPassword + " dbname=" + PostgresDb + " port=" + PostgresPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	AppDatabase = db
}
