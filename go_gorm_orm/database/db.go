package database

import (
	"fmt"
	"go_gorm_orm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// variabel global
var (
	host     = "localhost"
	user     = "postgres"
	password = "1234"
	dbPort   = "5432"
	dbname   = "learning_gorm_golang"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
