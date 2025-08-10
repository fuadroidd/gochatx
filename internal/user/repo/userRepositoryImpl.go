package repo

import (
	"fmt"
	"gochatx/internal/infrastructure/models"
	"gochatx/internal/user"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (ur *UserRepositoryImpl) Create(user user.UserEntity, password string) error {

	if err := ur.Db.Create(&models.UserModel{
		// CreatedAt:time.Now(),UpdatedAt:time.Now(),
		Username:    user.Username,
		Displayname: user.Displayname,
		Email:       user.Email,
		Password:    password,
	}); err.Error != nil {
		return fmt.Errorf("failed to create user: %s", err.Error)
	}
	return nil

}

func (ur *UserRepositoryImpl) Update(id int, displayName string) error {
	if err := ur.Db.Model(&models.UserModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"display_name": displayName,
	}).Error; err != nil {
		return fmt.Errorf("failed to update %s", err)
	}
	return nil
}

func (ur UserRepositoryImpl) GetById(id int) (*user.UserEntity, error) {
	var usr models.UserModel
	if err := ur.Db.First(&usr, id); err != nil {
		return nil, fmt.Errorf("failed to get user %s", err.Error)
	}
	userEntity := user.UserEntity{

		ID:          usr.ID,
		Username:    usr.Username,
		Displayname: usr.Displayname,
		Email:       usr.Email,
	}
	return &userEntity, nil
}

func (ur *UserRepositoryImpl) Delete(id int) error {
	var usr models.UserModel
	if err := ur.Db.First(&usr, id); err != nil {
		return fmt.Errorf("user not found %s", err.Error)
	}
	ur.Db.Where("id=?", id).First(&usr).Delete(&usr)
	return nil
}

func (ur UserRepositoryImpl) GetByUsername(username string) (*user.UserEntity, error) {
	var usr models.UserModel
	if err := ur.Db.Where("username = ?", username).First(&usr); err.Error != nil {

		return nil, fmt.Errorf("/nuser by %s no found: %s/n", username, err.Error)
	}
	fmt.Printf("log %s", usr.Displayname)
	ue := user.UserEntity{
		ID:          usr.ID,
		Username:    usr.Username,
		Displayname: usr.Displayname,
		Email:       usr.Email,
	}
	return &ue, nil

}

func (ur *UserRepositoryImpl) Authenticate(username string, passwdHash string) (*user.UserEntity, error) {
	var usr models.UserModel
	if err := ur.Db.First(&usr, username); err != nil {
		return nil, fmt.Errorf("user %s not found error: %s", username, err.Error)
	}
	if usr.Password == passwdHash {
		return &user.UserEntity{
			ID:          usr.ID,
			Displayname: usr.Displayname,
			Username:    usr.Username,
			Email:       usr.Email,
		}, nil
	}
	return nil, fmt.Errorf("incorect password for user %s ", username)

}

func (ur UserRepositoryImpl) GetAll() ([]*user.UserEntity, error) {
	var users []models.UserModel
	result := ur.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	} else {
		var ues []*user.UserEntity
		for i := 0; i < int(result.RowsAffected); i++ {
			ues = append(ues, &user.UserEntity{
				ID:          users[i].ID,
				Username:    users[i].Username,
				Displayname: users[i].Displayname,
				Email:       users[i].Email,
			})
		}
		return ues, nil
	}
}
