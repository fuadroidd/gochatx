package models

import (
	"time"

	"gorm.io/gorm"
)

type RoomModel struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedBy uint `gorm:"foreignKey:id"`
	CreatedAt time.Time
}
