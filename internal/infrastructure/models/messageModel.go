package models

import (
	"time"

	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	SenderID   uint
	ReceiverID *uint
	RoomID     *uint
	CreatedAt  time.Time
}
