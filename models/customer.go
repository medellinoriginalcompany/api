package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	//* Entidade de cliente
	ID int32 `gorm:"primaryKey;"`

	FullName  string `gorm:"type:varchar(100);notnull"`
	Email     string `gorm:"type:varchar(50);unique;notnull"`
	Password  string `gorm:"type:varchar(255);notnull"`
	BirthDate string `gorm:"type:date;notnull"`
	CPF       string `gorm:"type:varchar(11);notnull"`
	Phone     string `gorm:"type:varchar(11);notnull"`
	Gender    int8   `gorm:"notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Customer) TableName() string {
	return "users.customers"
}
