package user

import (
	"errors"
	"fmt"
)

type UserUsecase struct {
	Repo IUserRepository
}

func NewUsecase(r IUserRepository) *UserUsecase {
	return &UserUsecase{Repo: r}
}

func (u *UserUsecase) Register(user UserEntity, password string) error {
	_, err := u.Repo.GetByUsername(user.Username)
	if err == nil {
		return fmt.Errorf("user already exists %s", err)
	}
	if err := u.Repo.Create(user, password); err != nil {
		return errors.New("unable to create new user")
	}
	return nil
}
func (u *UserUsecase) GetByUsername(username string) (*UserEntity, error) {
	user, err := u.Repo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (u *UserUsecase) Authenticate(username string, password string) (*UserEntity, error) {
	if len(username) > 3 && len(password) >= 4 {

		user, err := u.Repo.Authenticate(username, password)
		if err != nil {
			return nil, fmt.Errorf("authentication failed for user: '%s':%s ", username, err)
		}
		return user, nil
	}
	return nil, fmt.Errorf("invalid username and password")
}
