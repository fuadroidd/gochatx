package roommember

import (
	"fmt"
	"gochatx/internal/infrastructure/models"

	"gorm.io/gorm"
)

type RoomRepositoryImpl struct {
	db *gorm.DB
}

func NewRoomMemberRepositoryImpl(db *gorm.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{db: db}
}

func (r RoomRepositoryImpl) GetByRoomId(roomID uint) ([]*RoomMemberEntity, error) {
	var roomMembers []models.RoomMemberModel
	if res := r.db.Find(&roomMembers, models.RoomMemberModel{RoomID: roomID}); res.Error != nil {
		return nil, fmt.Errorf("RoomRepositoryImpl.GetByRoomId(roomId uint) error: %s", res.Error)
	} else {
		var rooms []*RoomMemberEntity
		for i := 0; i < len(roomMembers); i++ {
			rooms = append(rooms, &RoomMemberEntity{
				ID:     roomMembers[i].ID,
				RoomID: roomMembers[i].RoomID,
				UserID: roomMembers[i].UserID,
			})
		}
		return rooms, nil
	}
}
