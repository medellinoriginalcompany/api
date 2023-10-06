package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductSize struct {
	ID int32 `gorm:"primaryKey;"`

	Name        string `gorm:"type:varchar(100);notnull;unique;"`
	Description string `gorm:"type:varchar(255);"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductSize) TableName() string {
	return "products.sizes"
}
