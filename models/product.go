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
	Stock       int32   `gorm:"notnull;default:0"`
	Active      bool    `gorm:"notnull;default:true"`
	Discount    float32 `gorm:"notnull"`
	Banner      string  `gorm:"type:varchar(255);notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TypeID     int32
	CategoryID int32
	SizeID     int32
	ColorID    int32
	Type       ProductType     `gorm:"foreignKey:TypeID;constraint:OnDelete:SET NULL;"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL;"`
	Size       ProductSize     `gorm:"foreignKey:SizeID;constraint:OnDelete:SET NULL;"`
	Color      ProductColor    `gorm:"foreignKey:ColorID;constraint:OnDelete:SET NULL;"`

	Images []ProductImage `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "products.products"
}
