package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddSize(c *gin.Context) {
	// Pegar info do tamanho do corpo da req
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

	// Verificar se o tamanho já existe
	var size models.ProductSize
	database.DB.First(&size, "name = ?", body.Name)

	if size.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tamanho já cadastrado",
		})

		return
	}

	// Criar tamanho
	size = models.ProductSize{
		Name:        body.Name,
		Description: body.Description,
	}

	database.DB.Create(&size)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tamanho cadastrado com sucesso",
	})
}

func GetSizes(c *gin.Context) {
	var sizes []models.ProductSize

	database.DB.Find(&sizes, "deleted_at IS NULL")

	c.JSON(http.StatusOK, gin.H{
		"sizes": sizes,
	})
}
