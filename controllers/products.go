package controllers

import (
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func AddProduct(c *gin.Context) {
	// Pegar info do produto do corpo da req
	var body struct {
		Banner             string             `json:"banner"`
		Name               string             `json:"name"`
		Description        string             `json:"description"`
		Price              string             `json:"price"`
		DiscountPercentage string             `json:"discountPercentage"`
		Stock              string             `json:"stock"`
		Active             bool               `json:"active"`
		Manufacturer       string             `json:"manufacturer"`
		Print              string             `json:"print"`
		Category           string             `json:"category"`
		Type               string             `json:"type"`
		Color              string             `json:"color"`
		SKU                string             `json:"sku"`
		Sizes              map[string]float32 `json:"sizes"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"body":    body,
			"message": "Falha ao ler corpo da requisição",
			"warning": "Este é o corpo da requisição. Verifique se os campos estão corretos",
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

	// Verificar se as propriedades existem
	var productColor models.ProductColor
	database.DB.First(&productColor, "name = ?", body.Color)

	if productColor.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cor não cadastrada",
			"body":    body,
		})

		return
	}

	var productType models.ProductType
	database.DB.First(&productType, "name = ?", body.Type)

	if productType.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tipo não cadastrado",
			"body":    body,
		})

		return
	}

	var productCategory models.ProductCategory
	database.DB.First(&productCategory, "name = ?", body.Category)

	if productCategory.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Categoria não cadastrado",
			"body":    body,
		})

		return
	}

	// Converter strings para float/int
	price, err1 := strconv.ParseFloat(body.Price, 32)
	stock, err2 := strconv.ParseInt(body.Stock, 10, 32)

	if body.DiscountPercentage == "" {
		body.DiscountPercentage = "0"
	}

	discountPercentage, err3 := strconv.ParseFloat(body.DiscountPercentage, 32)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse price",
		})

		return
	}

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse stock",
		})

		return
	}

	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse discountPercentage",
			"body":    body,
		})

		return

	}

	// Criar produto
	product = models.Product{
		Banner:             body.Banner,
		Name:               body.Name,
		Description:        body.Description,
		Price:              float32(price),
		DiscountPercentage: float32(discountPercentage),
		Stock:              int32(stock),
		Active:             body.Active,
		Manufacturer:       body.Manufacturer,
		Print:              body.Print,
		CategoryID:         productCategory.ID,
		TypeID:             productType.ID,
		ColorID:            productColor.ID,
		SKU:                body.SKU,
	}

	result := database.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao criar produto",
		})

		return
	}

	// Criar relação entre produto e tamanho para cada tamanho recebido
	for _, size := range body.Sizes {
		var productSize models.ProductSize
		database.DB.First(&productSize, "id = ?", size)

		// Caso seja um tamanho inválido, pular
		if productSize.ID == 0 {
			continue
		}

		productSizeRelation := models.Products_product_size{
			ProductID: product.ID,
			SizeID:    productSize.ID,
		}

		result := database.DB.Create(&productSizeRelation)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Erro ao criar relação entre produto e tamanho",
			})

			return
		}
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

func GetActiveProducts(c *gin.Context) {
	var products []models.Product

	// Pegar produtos
	res := database.DB.Joins("Category").
		Where("active = ?", true).
		Find(&products)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao pegar produtos",
			"error":   res.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
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

func EditProduct(c *gin.Context) {
	id := c.Param("id")

	// Pegar info do produto do corpo da req
	var body struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		SKU             string `json:"sku"`
		Price           string `json:"price"`
		Stock           string `json:"stock"`
		Active          bool   `json:"active"`
		DiscountedPrice string `json:"discountedPrice"`
		Banner          string `json:"banner"`
		Type            string `json:"type"`
		Category        string `json:"category"`
		Size            string `json:"size"`
		Color           string `json:"color"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	// Converter strings para float/int
	price, err1 := strconv.ParseFloat(body.Price, 32)
	stock, err2 := strconv.ParseInt(body.Stock, 10, 32)

	if body.DiscountedPrice == "" {
		body.DiscountedPrice = "0"
	}

	discountedPrice, err3 := strconv.ParseFloat(body.DiscountedPrice, 32)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse price",
		})

		return
	}

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse stock",
		})

		return
	}

	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse discountedPrice",
		})

		return
	}

	// Necessário para atualizar valores nulos
	data := map[string]interface{}{
		"name":             body.Name,
		"description":      body.Description,
		"sku":              body.SKU,
		"price":            float32(price),
		"stock":            int32(stock),
		"active":           body.Active,
		"discounted_price": float32(discountedPrice),
		"banner":           body.Banner,
	}

	res := database.DB.Model(&models.Product{}).Where("id = ?", &id).Updates(data)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao editar produto",
			"error":   res.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto alterado com sucesso. Desconto removido.",
		"body":    body,
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

func PermaDeleteProduct(c *gin.Context) {
	// Pegar id do produto
	id := c.Param("id")
	cld, ctx := config.Credentials()
	var product models.Product

	// Pegar produto
	database.DB.Unscoped().First(&product, "id = ?", &id)

	productBanner := product.Banner

	// Deletar permanentemente produto
	response := database.DB.Unscoped().Delete(&product, "id = ?", &id)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar produto",
			"error":   response.Error.Error(),
		})

		return
	}

	// Deletar imagem do Cloudinary
	res, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs:    []string{productBanner},
		DeliveryType: "upload",
		AssetType:    "image",
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar imagem do Cloudinary. Produto deletado com sucesso",
			"error":   err.Error(),
		})

		return
	}

	if res.Deleted == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar imagem do Cloudinary. Produto deletado com sucesso",
			"error":   res.Error.Message,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto e imagem deletados permanentemente com sucesso",
	})
}

func RestoreProduct(c *gin.Context) {
	// Pegar id do produto
	id := c.Param("id")

	// Restaurar produto
	response := database.DB.Unscoped().Model(&models.Product{}).Where("id = ?", &id).Update("deleted_at", nil)

	if response.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao restaurar produto",
			"error":   response.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Produto restaurado com sucesso",
	})
}
