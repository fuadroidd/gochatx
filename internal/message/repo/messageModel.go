package repo

import (
	"time"

	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	SenderId  uint
	RoomId    uint
	Text      string
	TimeStamp time.Time
}
