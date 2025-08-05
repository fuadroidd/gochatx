package migrations

import (
	"gochatx/internal/infrastructure/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.UserModel{},
		//  &models.RoomModel{}, &models.MessageModel{}, &models.RoomMemberModel{}, &models.FileModel{},
	)
}
