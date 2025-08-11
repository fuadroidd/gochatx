package roommember

type RoomRepositoryI interface {
	GetUserId(userID uint) ([]*RoomMemberEntity, error)
	GetByRoomId(roomID uint) ([]*RoomMemberEntity, error)
}
