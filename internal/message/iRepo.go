package message

type IChatRepo interface {
	Get() MessageEntity
	Create(newChat MessageEntity) MessageEntity
	Update(id int, update MessageEntity) MessageEntity
	Delete(id int)
}
