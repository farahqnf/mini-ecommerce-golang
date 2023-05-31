package config

import (
	"fmt"
	"log"

	"github.com/tugasmeilyanto/go-trial-class/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=24jan01 dbname=contoh_h8 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err.Error())
	} else {
		fmt.Println("DB connected")
		DB = db
	}

	db.AutoMigrate(&entity.Product{}, &entity.Order{}, entity.User{})
}

// func DBConnect() {
// 	dsn := ""
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("failed to connect to database", err.Error())
// 	} else {
// 		fmt.Println("connected to db")
// 		DB = db
// 	}
// }
