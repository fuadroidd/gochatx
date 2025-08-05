package handler

import (
	"fmt"
	"gochatx/internal/user"
	"gochatx/internal/user/dtos"
)

type UserHandler struct {
	Uc *user.UserUsecase
}

func (u *UserHandler) Register(userReq dtos.UserRequestDTO) error {
	newUser := user.UserEntity{
		Username:    userReq.Username,
		Displayname: userReq.Displayname,
		Email:       userReq.Email,
	}
	print("display name:", userReq.Displayname)
	password := userReq.Password
	// registered, _ := u.Uc.Register(newUser)

	if err := u.Uc.Register(newUser, password); err != nil {
		return fmt.Errorf("failed to create user: %s", err.Error())
	}

	return nil
}

func (u UserHandler) Authenticate(username string, password string) error {
	if _, err := u.Uc.Authenticate(username, password); err != nil {
		return fmt.Errorf("error: %s", err)
	}
	return nil
}

func (u UserHandler) GetByUsername(username string) (*dtos.UserResponseDTO, error) {
	if usr, err := u.Uc.GetByUsername(username); err == nil {
		return &dtos.UserResponseDTO{
			ID:          usr.ID,
			Username:    usr.Username,
			Displayname: usr.Displayname,
			Email:       usr.Email,
		}, nil
	}
	return nil, fmt.Errorf("not found")
}
