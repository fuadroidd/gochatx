package repo

import (
	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	SenderID   uint
	ReceiverID *uint
	RoomId     *uint
	Text       string
}
