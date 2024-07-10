package database

import (
	"log"

	"github.com/anucha-tk/go_todo/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Automatically migrate the schema
	if err = DB.AutoMigrate(&domain.Todo{}); err != nil {
		return err
	}
	log.Println("connect sqlite successful")

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
