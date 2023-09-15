package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	//* Entidade de admin
	ID int32 `gorm:"primaryKey;"`

	FirstName string `gorm:"type:varchar(100);notnull"`
	LastName  string `gorm:"type:varchar(100);notnull"`
	Email     string `gorm:"type:varchar(50);unique;notnull"`
	Password  string `gorm:"type:varchar(255);notnull"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Admin) TableName() string {
	return "users.admins"
}