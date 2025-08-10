package repo

import (
	"fmt"
	msgpkg "gochatx/internal/message"
	"strconv"

	"gorm.io/gorm"
)

type MessageRepoImpl struct {
	DB gorm.DB
}

func (c MessageRepoImpl) GetById(messageID uint) (*msgpkg.MessageEntity, error) {
	var message MessageModel
	res := c.DB.Where(&MessageModel{ID: messageID}).First(&message)
	if res.Error != nil {
		return nil, fmt.Errorf("MessageRepo.GetByID(%s) error: %s", strconv.FormatUint(uint64(messageID), 10), res.Error)
	} else {
		return &msgpkg.MessageEntity{
			ID:        message.ID,
			SenderID:  message.SenderID,
			RoomID:    message.RoomId,
			Text:      message.Text,
			TimeStamp: message.CreatedAt,
		}, nil
	}

}
func (c MessageRepoImpl) GetByReceiverId(ReceiverID uint) ([]*msgpkg.MessageEntity, error) {
	var messages []MessageModel
	if res := c.DB.Where(&MessageModel{ReceiverID: &ReceiverID}).First(&messages); res.Error != nil {
		return nil, fmt.Errorf("MessageRepoImpl.GetByReciverID(%s) error:%s", strconv.FormatUint(uint64(ReceiverID), 10), res.Error)
	} else {
		var eMessage []*msgpkg.MessageEntity
		for i := 0; i < len(messages); i++ {
			eMessage = append(eMessage, &msgpkg.MessageEntity{
				ID:         messages[i].ID,
				SenderID:   messages[i].SenderID,
				ReceiverID: messages[i].ReceiverID,
				TimeStamp:  messages[i].CreatedAt,
			})
		}
		return eMessage, nil
	}

}
func (c MessageRepoImpl) GetByRoomId(roomID uint) ([]*msgpkg.MessageEntity, error) {
	var messages []MessageModel
	if res := c.DB.Where(&MessageModel{RoomId: &roomID}).First(&messages); res.Error != nil {
		return nil, fmt.Errorf("MessageRepoImpl.GetByRoomID(%s) error:%s", strconv.FormatUint(uint64(roomID), 10), res.Error)
	} else {
		var eMessage []*msgpkg.MessageEntity
		for i := 0; i < len(messages); i++ {
			eMessage = append(eMessage, &msgpkg.MessageEntity{
				ID:        messages[i].ID,
				SenderID:  messages[i].SenderID,
				RoomID:    messages[i].RoomId,
				TimeStamp: messages[i].CreatedAt,
			})
		}
		return eMessage, nil
	}

}
func (c MessageRepoImpl) Create(newMessage msgpkg.MessageEntity) error {
	message := MessageModel{
		SenderID:   newMessage.SenderID,
		RoomId:     newMessage.RoomID,
		ReceiverID: newMessage.ReceiverID,
		Text:       newMessage.Text,
	}
	if res := c.DB.Create(&message); res.Error != nil {
		return fmt.Errorf("MessageRepoIml.Create(%s) error %s", newMessage.Text, res.Error)
	} else {
		return nil
	}
}
