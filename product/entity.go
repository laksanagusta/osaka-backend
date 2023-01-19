package product

import (
	"time"
)

type Product struct {
	ID          int    `gorm:"size:36;not null;uniqueIndex;primary"`
	Title       string `gorm:size:100`
	UnitPrice   int    `gorm:size:100`
	Description string `gorm:size:100`
	Code        string `gorm:size:100`
	Image       string `gorm:size:100`
	IsDeleted   int8   `gorm:size:1;default:0`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
