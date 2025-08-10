package message

import (
	"time"

	"gorm.io/gorm"
)

type MessageEntity struct {
	gorm.Model
	ID         uint
	SenderID   uint
	ReceiverID *uint
	RoomID     *uint
	Text       string
	TimeStamp  time.Time
}
