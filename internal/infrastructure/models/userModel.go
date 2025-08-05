package models

import (
	"gorm.io/gorm"
)

// gorom model implementation for user with postgresql database connection
type UserModel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex"`
	Displayname string `gomr:"column:displayname"`
	Email       string `gorm:"uniqueIndex"`
	Password    string `gorm:"column:password"`
	//has rooms
}
