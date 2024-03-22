package config

import (
	"29_rest_api_basic/structs"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "developer:developer123@tcp(192.168.56.109:3306)/go_resful?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect to databases")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
