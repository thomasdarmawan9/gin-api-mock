package database

import (
	"fmt"
	"gin-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "containers-us-west-42.railway.app"
	user     = "postgres"
	password = "PZD7xkuhWUiBgJTPf2lV"
	dbPort   = "5433"
	dbname   = "railway"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.AutoMigrate(&models.Car{})
}

func GetDB() *gorm.DB {
	
	return db
}