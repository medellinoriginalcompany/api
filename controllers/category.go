package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddCategory(c *gin.Context) {
	// Pegar info da categoria do corpo da req
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

	// Verificar se a categoria já existe
	var category models.ProductCategory
	database.DB.First(&category, "name = ?", body.Name)

	if category.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Categoria já cadastrada",
		})

		return
	}

	// Criar categoria
	category = models.ProductCategory{
		Name:        body.Name,
		Description: body.Description,
	}

	database.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria cadastrada com sucesso",
	})
}

func GetCategories(c *gin.Context) {
	var categories []models.ProductCategory

	database.DB.Find(&categories, "deleted_at IS NULL")

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
