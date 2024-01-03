package models

type Products_product_size struct {
	ID int32 `gorm:"primaryKey;"`

	ProductID int32   `gorm:"notnull"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;"`

	SizeID int32       `gorm:"notnull"`
	Size   ProductSize `gorm:"foreignKey:SizeID;constraint:OnDelete:CASCADE;"`
}

func (Products_product_size) TableName() string {
	return "products.product_size"
}
