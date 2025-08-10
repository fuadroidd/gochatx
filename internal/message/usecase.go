package message

import (
	messageDTOs "gochatx/internal/message/dtos"
)

type MessageUsecase struct {
	Repo IMessageRepo
}

func (uc MessageUsecase) SendMessage(messageReq messageDTOs.MessageRequestDTO) error {
	newMessage := MessageEntity{
		SenderID: messageReq.SenderID,
		RoomID:   messageReq.ReceiverID,
		Text:     messageReq.Text,
	}
	if err := uc.Repo.Create(newMessage); err != nil {
		return nil
	} else {
		return err
	}

}

func (uc MessageUsecase) GetByRoomID(roomID uint) ([]*MessageEntity, error) {
	if res, err := uc.Repo.GetByRoomId(roomID); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (uc MessageUsecase) GetByReceiverId(receiverID uint) ([]*MessageEntity, error) {
	if res, err := uc.Repo.GetByReceiverId(receiverID); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (uc MessageUsecase) GetByID(id uint) (*MessageEntity, error) {
	if res, err := uc.Repo.GetById(id); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
