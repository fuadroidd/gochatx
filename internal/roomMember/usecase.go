package roommember

import "fmt"

type RoomMemberUsecase struct {
	repo RoomRepositoryI
}

func NewRoomUsecase(r RoomRepositoryI) *RoomMemberUsecase {
	return &RoomMemberUsecase{repo: r}
}

func (u RoomMemberUsecase) GetRoomMembersByRoomId(roomID uint) ([]*RoomMemberEntity, error) {
	if rooms, err := u.repo.GetByRoomId(roomID); err != nil {
		return nil, fmt.Errorf("RoomMemberUsecase.GetRoomMembersByRoomId(roomId uint) error: %s", err)
	} else {
		return rooms, nil
	}
}
func (u RoomMemberUsecase) GetRoomsByMemberId(userID uint) ([]*RoomMemberEntity, error) {
	if rooms, err := u.repo.GetUserId(userID); err != nil {
		return nil, fmt.Errorf("RoomMemberUsecase.GetRoomsByMemberId(userID uint) error: %s", err)
	} else {
		return rooms, nil
	}
}
