package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderId      uint   `gorm:"primaryKey"`
	CostumerName string `json:"costumer_name" gorm:"not null;type:varchar(50)"`
	ItemId       []Item `json:"item_id" gorm:"foreignKey:ItemId"`
	OrderedAt    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     gorm.DeletedAt
}
