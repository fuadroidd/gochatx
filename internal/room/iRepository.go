package room

type IRoomRepository interface {
	Create(name string, createdBy int) error

	Delete(id int) error
	GetById(id int) (*RoomEntity, error)
}
