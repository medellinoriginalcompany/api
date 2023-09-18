package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID int32 `gorm:"primaryKey;"`

	Name        string `gorm:"type:varchar(255);notnull"`
	Description string `gorm:"type:varchar(255);notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductType) TableName() string {
	return "products.types"
}
