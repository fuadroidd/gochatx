package message

type CreateChatUsecase struct {
	repo IChatRepo
}

func (create *CreateChatUsecase) Execute(chat MessageEntity) MessageEntity {
	newChat := create.repo.Create(chat)
	return newChat

}

type UpdateChatUsecase struct {
	repo IChatRepo
}

func (update *UpdateChatUsecase) Execute(id int, updates MessageEntity) MessageEntity {
	updated := update.repo.Update(id, updates)
	return updated
}

type DeleteChatUsecase struct {
	repo IChatRepo
}

func (delete *DeleteChatUsecase) Execute(id int) {
	delete.repo.Delete(id)

}
