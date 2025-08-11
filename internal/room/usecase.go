package room

import (
	"fmt"
)

type RoomUsecase struct {
	repo IRoomRepository
}

func (r RoomUsecase) Create(name string, by uint) error {
	if err := r.repo.Create(name, by); err != nil {
		return fmt.Errorf("could create room %s error: %s", name, err)
	}
	return nil
}

func (r *RoomUsecase) GetById(id int) (*RoomEntity, error) {
	room, err := r.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("couldn't get room %s", err)
	}
	return room, nil

}
