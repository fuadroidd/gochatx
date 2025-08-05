package message

import (
	"time"

	"gorm.io/gorm"
)

type MessageEntity struct {
	gorm.Model
	Id        int
	RoomId    int
	SenderId  int
	Text      string
	TimeStamp time.Time
	FileLink  string
}
