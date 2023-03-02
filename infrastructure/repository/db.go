package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_USERNAME = "hung"
const DB_PASSWORD = "12321221"
const DB_NAME = "seatmap"
const DB_HOST = "localhost"
const DB_PORT = "5432"

var Db *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
	}
	Db = db
}

func GetDB() *gorm.DB {
	return Db
}
