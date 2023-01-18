package organization

import (
	"time"
	"tokokecilkita-go/user"
)

type Organization struct {
	ID        int    `gorm:"size:36;not null;uniqueIndex;primary"`
	Name      string `gorm:size:100`
	Status    int8   `gorm:size:8`
	UserID    int    `gorm:size:32`
	User      user.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
