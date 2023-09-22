package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID int32 `gorm:"primaryKey;"`

	ProductID int32 `gorm:"notnull"`
	URL       string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductImage) TableName() string {
	return "products.images"
}
