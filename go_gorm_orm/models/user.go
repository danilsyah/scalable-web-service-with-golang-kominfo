package models

import "time"

type User struct {
	ID       uint   `gorm:"primarykey"`
	Email    string `gorm:"not null;unique;type:varchar(191)"`
	Products []Product
	CreateAt time.Time
	UpdateAt time.Time
}
