package db

import (
	"log"

	"github.com/anucha-tk/go_todo/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Init() (*gorm.DB, error) {
	var err error
	dbConn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatically migrate the schema
	if err = dbConn.AutoMigrate(&domain.Todo{}); err != nil {
		return nil, err
	}
	log.Println("connect sqlite successful")

	return dbConn, nil
}
