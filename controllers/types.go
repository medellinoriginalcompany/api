package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddType(c *gin.Context) {
	// Pegar info do tipo do corpo da req
	var body struct {
		Name        string
		Description string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Verificar se o tipo já existe
	var productType models.ProductType
	database.DB.First(&productType, "name = ?", body.Name)

	if productType.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tipo já cadastrado",
		})

		return
	}

	// Criar tipo
	productType = models.ProductType{
		Name:        body.Name,
		Description: body.Description,
	}

	database.DB.Create(&productType)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo cadastrado com sucesso",
	})
}

func GetTypes(c *gin.Context) {
	var productTypes []models.ProductType

	database.DB.Find(&productTypes, "deleted_at IS NULL")

	c.JSON(http.StatusOK, gin.H{
		"types": productTypes,
	})
}
