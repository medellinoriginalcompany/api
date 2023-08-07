package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//* Entidade usuário
	ID     int32 `gorm:"primaryKey;"`

	FirstName string `gorm:"type:varchar(50);notnull"`
	LastName  string `gorm:"type:varchar(50);notnull"`
	Email     string `gorm:"type:varchar(50);unique;notnull"`
	Password  string `gorm:"type:varchar(255);notnull"`
	Birth     string `gorm:"type:date;notnull"`
	CPF       string `gorm:"type:varchar(50);notnull"`
	NumeroCel string `gorm:"type:varchar(20);notnull"`
	Gender    int8   `gorm:"notnull"`

	Profile_pic string `gorm:"type:varchar(100);default:default_profile.webp;notnull"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
