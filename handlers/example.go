package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vnxcius/medellin-api/config"
)

func Example(c *gin.Context) {
	example := config.Getenv("EXAMPLE") // Carregar variável de ambiente

	c.JSON(200, gin.H{
		"API DIZ": example,
	})
}
