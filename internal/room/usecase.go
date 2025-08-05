package room

import (
	"fmt"
)

type RoomUsecase struct {
	repo IRoomRepository
}

// func (r *RoomUsecase)AddUser(user usr.UserEntity) error{
// 	if err:=r.repo.addUser(user.Id); err!=nil{
// 		fmt.Errorf("unable to add %s to room %s : ", user.Username,r.repo.Create().Name)
// 	}

// }

func (r RoomUsecase) Create(name string, by int) error {
	if err := r.repo.Create(name, by); err != nil {
		return fmt.Errorf("could create room %s error: %s", name, err)
	}
	return nil
}
func (r RoomUsecase) Delete(id int) error {
	if _, err := r.repo.GetById(id); err != nil {
		return fmt.Errorf("room not found: %s", err)
	}
	if err := r.repo.Delete(id); err != nil {
		return fmt.Errorf("couldn't delete room %s ", err)
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
