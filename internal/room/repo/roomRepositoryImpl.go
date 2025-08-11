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

func (r *RoomRepositoryImpl) Create(roomName string, by uint) error {

	if err := r.DB.Create(&model.RoomModel{Name: roomName, CreatedBy: by}); err != nil {
		return fmt.Errorf("RoomRepositoryImpl.Create error: %s ", err.Error)
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

func (r *RoomRepositoryImpl) GetAll() ([]*entity.RoomEntity, error) {
	var rooms []model.RoomModel
	if err := r.DB.Find(&rooms); err != nil {
		return nil, fmt.Errorf("RoomRepositoryImpl.GetAll() error: %s", err.Error)
	} else {
		var entities []*entity.RoomEntity
		for i := 0; i < len(rooms); i++ {
			entities = append(entities, &entity.RoomEntity{
				ID:        rooms[i].ID,
				Name:      rooms[i].Name,
				CreatedBy: rooms[i].CreatedBy,
			})
		}
		return entities, nil
	}
}
