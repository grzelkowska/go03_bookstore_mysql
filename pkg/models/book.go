package models

import (
	// "github.com/grzelkowska/11projects/03go_bookstore_mysql/pkg/config"
	"github.com/grzelkowska/go03_bookstore_mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
