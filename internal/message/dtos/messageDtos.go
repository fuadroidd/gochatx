package dtos

import "time"

type MessageRequestDTO struct {
	SenderID   uint
	ReceiverID *uint
	RoomID     *uint
	Text       string
}

type MessageResponseDTO struct {
	MessageId  uint
	SenderID   uint
	ReceiverID *uint
	RoomID     *uint
	Text       string
	CreatedAt  time.Time
}
