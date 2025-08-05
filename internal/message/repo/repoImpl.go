package repo

import (
	"gochatx/internal/message"

	"gorm.io/gorm"
)

type ChatRepoImpl struct {
	db gorm.DB
}

func (c *ChatRepoImpl) Get() message.MessageEntity {
	// c.db.

	return message.MessageEntity{}
}

func (c *ChatRepoImpl) Create(chat message.MessageEntity) message.MessageEntity {
	return message.MessageEntity{}
}
func (c *ChatRepoImpl) Update(id int, update message.MessageEntity) message.MessageEntity {
	return message.MessageEntity{}
}

func (c *ChatRepoImpl) Delete(id int) {
	///not implemented
}
