package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemId      uint   `gorm:"primaryKey"`
	ItemCode    string `json:"item_code" gorm:"not null;type:varchar(50)"`
	Descryption string `json:"descryption_item"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
