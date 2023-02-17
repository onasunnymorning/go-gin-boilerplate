package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//DBConnection -> return db instance
func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	return db, nil
}
