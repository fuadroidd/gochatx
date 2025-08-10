package router

import (
	msgPkg "gochatx/internal/message"
	dtos "gochatx/internal/message/dtos"
)

type MessageRoutes struct {
	MessageUseCase msgPkg.MessageUsecase
}

func (routes MessageRoutes) SendMessage(msgReq dtos.MessageRequestDTO) error {
	if err := routes.MessageUseCase.Repo.Create(
		msgPkg.MessageEntity{
			SenderID:   msgReq.SenderID,
			RoomID:     msgReq.RoomID,
			ReceiverID: msgReq.ReceiverID,
			Text:       msgReq.Text,
		}); err != nil {
		return err
	} else {
		return nil
	}

}

func (routes MessageRoutes) GetMessageByID(id uint) (*dtos.MessageResponseDTO, error) {
	if message, err := routes.MessageUseCase.Repo.GetById(id); err != nil {
		return nil, err
	} else {
		return &dtos.MessageResponseDTO{
			MessageId:  message.ID,
			SenderID:   message.SenderID,
			ReceiverID: message.ReceiverID,
			Text:       message.Text,
		}, nil
	}
}
