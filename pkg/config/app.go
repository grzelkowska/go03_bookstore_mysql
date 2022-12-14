package config

import (
	// "fmt"
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("Could not load .env file")
	// }
	// password := os.Getenv("USER_PASSWORD")
	d, err := gorm.Open("mysql", "user@tcp(127.0.0.1:3306)@/simplerest?charset=utf8&parseTime=True&loc=Loca")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
