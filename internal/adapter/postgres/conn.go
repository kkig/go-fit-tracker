package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	pass := "someRandomStr"
	dns := fmt.Sprintf("password=%v", pass)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)	
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to server!")

	return db
}