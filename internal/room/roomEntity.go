package room

import (
	"time"
)

type RoomEntity struct {
	ID        uint
	Name      string
	CreatedBy uint
	CreatedAt time.Time
}
