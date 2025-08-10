package message

type IMessageRepo interface {
	GetById(MessageID uint) (*MessageEntity, error)
	GetByReceiverId(ReceiverID uint) ([]*MessageEntity, error)
	GetByRoomId(roomID uint) ([]*MessageEntity, error)
	Create(newMessage MessageEntity) error
}
