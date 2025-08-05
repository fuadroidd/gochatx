package repo

import (
	"fmt"
	model "gochatx/internal/infrastructure/models"
	entity "gochatx/internal/room"

	"gorm.io/gorm"
)

type RoomRepositoryImpl struct {
	DB *gorm.DB
}

func (r *RoomRepositoryImpl) Create(room entity.RoomEntity, password string) error {
	roomModel := model.RoomModel{Name: room.Name, CreatedBy: room.CreatedBy, CreatedAt: room.CreatedAt}
	if err := r.DB.Create(roomModel); err != nil {
		return fmt.Errorf("couldn't add user: %s ", err.Error)
	}
	return nil
}

func (r *RoomRepositoryImpl) GetById(roomID int) (*entity.RoomEntity, error) {
	var room model.RoomModel
	if err := r.DB.First(&room, roomID).Error; err != nil {
		return nil, fmt.Errorf("room not found %s", err)
	}
	roomEntity := entity.RoomEntity{ID: room.ID, Name: room.Name, CreatedBy: room.CreatedBy, CreatedAt: room.CreatedAt}
	return &roomEntity, nil
}

func (r *RoomRepositoryImpl) Delete(roomID int) error {
	if _, err := r.GetById(roomID); err != nil {
		return fmt.Errorf("room doesn't exist: %s", err)
	}
	var room model.RoomModel
	if err := r.DB.Delete(room, roomID).Error; err != nil {
		return fmt.Errorf("couldn't delete, %s", err)
	}
	return nil
}
