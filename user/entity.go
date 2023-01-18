package user

import (
	"time"
)

type User struct {
	ID             int    `gorm:"size:36;not null;uniqueIndex;primary"`
	Username       string `gorm:size:100`
	Name           string `gorm:size:100`
	Occupation     string `gorm:size:100`
	Email          string `gorm:size:100;not null`
	PasswordHash   string `gorm:size:100;not null`
	Role           string `gorm:size:100;not null`
	OrganizationId int    `gorm:size:100`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
