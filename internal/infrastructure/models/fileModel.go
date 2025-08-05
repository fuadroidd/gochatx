package models

import (
	"time"

	"gorm.io/gorm"
)

type FileModel struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	UploaderID uint
	MessageID  uint
	FileURL    string
	FileType   string
	FileSize   string
	UploadedAt time.Time
}
