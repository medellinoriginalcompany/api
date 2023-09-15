package database

import (
	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	//* Fazer a conexão com banco de dados postgres

	dsn := config.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&models.Admin{},
		&models.Customer{},
		&models.Product{},
	)

}