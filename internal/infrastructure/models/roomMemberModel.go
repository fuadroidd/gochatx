package models

import (
	"gorm.io/gorm"
)

type RoomMemberModel struct {
	gorm.Model
	ID     uint `gorm:"foreignKey"`
	UserID uint
	RoomID uint
}
