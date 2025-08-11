package room

type IRoomRepository interface {
	Create(name string, createdBy uint) error
	GetById(id int) (*RoomEntity, error)
	GetAll() ([]*RoomEntity, error)
}
