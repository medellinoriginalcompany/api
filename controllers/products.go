package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddProduct(c *gin.Context) {
	// Pegar info do produto do corpo da req
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		SKU         string `json:"sku"`
		Price       string `json:"price"`
		Active      bool   `json:"active"`
		Banner      string `json:"banner"`
		Type        string `json:"type"`
		Category    string `json:"category"`
		Size        string `json:"size"`
		Color       string `json:"color"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
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

	var productColor models.ProductColor
	database.DB.First(&productColor, "name = ?", body.Color)

	if productColor.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cor não cadastrada",
		})

		return
	}

	var productType models.ProductType
	database.DB.First(&productType, "name = ?", body.Type)

	if productType.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tipo não cadastrado",
		})

		return
	}

	var productCategory models.ProductCategory
	database.DB.First(&productCategory, "name = ?", body.Category)

	if productCategory.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Categoria não cadastrado",
		})

		return
	}

	var productSize models.ProductSize
	database.DB.First(&productSize, "name = ?", body.Size)

	if productSize.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tamanho não cadastrado",
		})

		return
	}

	// Converter string preço para float
	price, err := strconv.ParseFloat(body.Price, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse price",
		})

		return
	}

	// Criar produto
	product = models.Product{
		Name:        body.Name,
		Description: body.Description,
		SKU:         body.SKU,
		Price:       float32(price),
		Active:      body.Active,
		Banner:      body.Banner,
		TypeID:      productType.ID,
		CategoryID:  productCategory.ID,
		SizeID:      productSize.ID,
		ColorID:     productColor.ID,
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

func GetProducts(c *gin.Context) {
	var products []models.Product

	// Pegar produtos
	database.DB.Joins("Category").Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": &products,
	})
}

func GetProduct(c *gin.Context) {
	// Pegar id do produto
	id := c.Param("id")

	// Pegar produto
	var product models.Product
	database.DB.Joins("Type").Joins("Category").Joins("Size").Joins("Color").First(&product, "products.id = ?", &id)

	if product.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Produto não encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": &product,
	})
}

func DeleteProduct(c *gin.Context) {
	// Pegar id do produto
	id := c.Param("id")

	// Deletar produto
	response := database.DB.Delete(&models.Product{}, "id = ?", &id)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar produto",
			"error":   response.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto deletado com sucesso",
	})
}

func GetDeletedProducts(c *gin.Context) {
	var products []models.Product

	// Pegar produtos deletados
	response := database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&products)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao pegar produtos deletados",
			"error":   response.Error.Error(),
		})

		return
	}
	
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum produto deletado foi encontrado",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": &products,
	})
}

func PermaDeleteProdutct(c *gin.Context) {
	// Pegar id do produto
	id := c.Param("id")

	// Deletar permanentemente produto
	response := database.DB.Unscoped().Delete(&models.Product{}, "id = ?", &id)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar produto",
			"error":   response.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto deletado permanente com sucesso",
	})
}
