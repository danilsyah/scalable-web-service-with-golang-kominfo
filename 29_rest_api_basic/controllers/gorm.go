package controllers

import "github.com/jinzhu/gorm"

type InDB struct {
	DB *gorm.DB
}

type Person struct {
	gorm.Model
	First_Name string
	Last_Name  string
}
