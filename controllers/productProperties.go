package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
)

func GetProductProperties(c *gin.Context) {
	// Obter todas as propriedades e retornar em json
	var categories []models.ProductCategory
	var colors []models.ProductColor
	var sizes []models.ProductSize
	var types []models.ProductType

	database.DB.Order("name ASC").Find(&categories)
	database.DB.Order("name ASC").Find(&colors)
	database.DB.Order("name ASC").Find(&types)
	database.DB.Order("name ASC").Find(&sizes)

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"colors":     colors,
		"sizes":      sizes,
		"types":      types,
	})
}

func DeleteProductProperty(c *gin.Context) {
	var productType = c.Param("type")
	var id = c.Param("id")

	var property interface{}

	switch productType {
	case "tamanhos":
		property = models.ProductSize{}
	case "cores":
		property = models.ProductColor{}
	case "categorias":
		property = models.ProductCategory{}
	case "tipos":
		property = models.ProductType{}

	default:
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Tipo não encontrado",
		})

		return
	}

	result := database.DB.Unscoped().Delete(&property, "id = ?", &id)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao deletar propriedade",
			"error":   result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Propriedade deletada com sucesso",
	})
}

func AddProductSize(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var size models.ProductSize
	database.DB.First(&size, "name = ?", body.Name)

	if size.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tamanho já cadastrado",
		})

		return
	}

	result := database.DB.Create(&models.ProductSize{
		Name: body.Name,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao cadastrar tamanho",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tamanho cadastrado com sucesso",
	})
}

func AddProductColor(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var color models.ProductColor
	database.DB.First(&color, "name = ?", body.Name)

	if color.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cor já cadastrada",
		})

		return
	}

	result := database.DB.Create(&models.ProductColor{
		Name: body.Name,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao cadastrar tamanho",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cor cadastrada com sucesso",
	})
}

func AddProductCategory(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var category models.ProductCategory
	database.DB.First(&category, "name = ?", body.Name)

	if category.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Categoria já cadastrada",
		})

		return
	}

	result := database.DB.Create(&models.ProductCategory{
		Name: body.Name,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao cadastrar tamanho",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria cadastrada com sucesso",
	})
}

func AddProductType(c *gin.Context) {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var productType models.ProductType
	database.DB.First(&productType, "name = ?", body.Name)

	if productType.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Tipo já cadastrado",
		})

		return
	}

	result := database.DB.Create(&models.ProductType{
		Name: body.Name,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao cadastrar tamanho",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo cadastrado com sucesso",
	})
}

func EditProductSize(c *gin.Context) {
	var id = c.Param("id")

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})
		fmt.Println("Failed to read body")

		return
	}

	result := database.DB.Model(models.ProductSize{}).Where("id = ?", &id).Updates(models.ProductSize{
		Name: body.Name,
		Description: body.Description,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao editar tamanho",
		})
		fmt.Println(result.Error.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tamanho editado com sucesso",
	})
}

func EditProductColor(c *gin.Context) {
	var id = c.Param("id")

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var color models.ProductColor

	result := database.DB.Model(&color).Where("id = ?", &id).Updates(models.ProductColor{
		Name: body.Name,
		Description: body.Description,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao editar cor",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cor editada com sucesso",
	})
}

func EditProductCategory(c *gin.Context) {
	var id = c.Param("id")

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var category models.ProductCategory

	result := database.DB.Model(&category).Where("id = ?", &id).Updates(models.ProductCategory{
		Name: body.Name,
		Description: body.Description,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao editar categoria",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categoria editada com sucesso",
	})
}

func EditProductType(c *gin.Context) {
	var id = c.Param("id")

	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to read body",
			"body":    body,
		})

		return
	}

	var productType models.ProductType

	result := database.DB.Model(&productType).Where("id = ?", &id).Updates(models.ProductType{
		Name: body.Name,
		Description: body.Description,
	})

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Erro ao editar tipo",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tipo editado com sucesso",
	})
}
