package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddColor(c *gin.Context) {
	// Pegar info da cor do corpo da req
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

	// Verificar se a cor já existe
	var color models.ProductColor
	database.DB.First(&color, "name = ?", body.Name)

	if color.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cor já cadastrada",
		})

		return
	}

	// Criar cor
	color = models.ProductColor{
		Name: body.Name,
		Description: body.Description,
	}

	database.DB.Create(&color)

	c.JSON(http.StatusOK, gin.H{
		"message": "Cor cadastrada com sucesso",
	})
}

func GetColors(c *gin.Context) {
	var colors []models.ProductColor

	database.DB.Find(&colors, "deleted_at IS NULL")

	c.JSON(http.StatusOK, gin.H{
		"colors": colors,
	})
}
