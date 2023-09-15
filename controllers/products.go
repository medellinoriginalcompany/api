package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddProduct(c *gin.Context) {
	// Pegar info do produto do corpo da req
	var body struct {
		Name        string
		Description string
		SKU         string
		Price       float32
		Stock       int32
		Active      bool
		Discount    float32
		Banner      string
		CategoryID  int32
		SizeID      int32
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Verificar se o produto já existe
	var product models.Product
	database.DB.First(&product, "sku = ?", body.SKU)

	if product.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "SKU já cadastrado",
		})

		return
	}

	// Criar produto
	product = models.Product{
		Name:        body.Name,
		Description: body.Description,
		SKU:         body.SKU,
		Price:       body.Price,
		Stock:       body.Stock,
		Active:      body.Active,
		Discount:    body.Discount,
		Banner:      body.Banner,
		CategoryID:  body.CategoryID,
		SizeID:      body.SizeID,
	}

	result := database.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao criar produto",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Produto criado com sucesso",
	})
}
