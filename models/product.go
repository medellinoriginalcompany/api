package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID int32 `gorm:"primaryKey;"`

	Name        string  `gorm:"type:varchar(255);notnull"`
	Description string  `gorm:"type:varchar(255);notnull"`
	SKU         string  `gorm:"type:varchar(50);unique;notnull"`
	Price       float32 `gorm:"notnull"`
	Stock       int32   `gorm:"notnull"`
	Active      bool    `gorm:"notnull;default:true"`
	Discount    float32 `gorm:"notnull"`
	Banner      string  `gorm:"type:varchar(255);notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CategoryID int32    `gorm:"notnull"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID"`

	SizeID int32 `gorm:"notnull"`
	Size   ProductSize  `gorm:"foreignKey:SizeID"`
}

func (Product) TableName() string {
	return "products.products"
}
