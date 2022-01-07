package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("leads.db"))
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic("failed to connect to database")
	}

	fmt.Println("Connection to database established")

	return db
}

func Close() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		panic("cannot close database connection")
	}

	sqlDB.Close()
	fmt.Println("Connection to database closed")
}
