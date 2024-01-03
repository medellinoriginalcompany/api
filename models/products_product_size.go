package models

type Products_product_size struct {
	ID int32 `gorm:"primaryKey;"`

	ProductID int32   `gorm:"notnull"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE;"`

	SizeID int32       `gorm:"notnull"`
	Size   ProductSize `gorm:"foreignKey:SizeID;references:ID;constraint:OnDelete:CASCADE;"`

	Stock int32 `gorm:"notnull;default:0"`
}

func (Products_product_size) TableName() string {
	return "products.product_size"
}
