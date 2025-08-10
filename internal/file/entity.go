package file

import "time"

type FileEntity struct {
	ID         uint
	UploaderID uint
	MessageID  uint
	FileURL    string
	FileType   string
	FileSize   string
	UploadedAt time.Time
}
