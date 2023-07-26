package routes

import (
	"github.com/vnxcius/medellin-api/handlers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	//* Gerenciamento de rotas
	r := gin.Default()

	r.GET("/example", handlers.Example)

	// r.POST("/registro", controllers.Signup)
	// r.POST("/login", controllers.Login)
	// r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
