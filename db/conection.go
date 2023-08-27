package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	DSN = "host=localhost user=misa password=mysecret dbname=p5w1 port=5432"
)

var DB *gorm.DB

func Connection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect:", err)
	} else {
		log.Println("Database connected")
	}
}
