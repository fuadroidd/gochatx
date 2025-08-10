package user

type IUserRepository interface {
	GetById(id int) (*UserEntity, error)
	GetByUsername(username string) (*UserEntity, error)

	Create(user UserEntity, password string) error
	Update(id int, displayname string) error
	Delete(id int) error
	Authenticate(username string, password string) (*UserEntity, error)
	GetAll() ([]*UserEntity, error)
}
