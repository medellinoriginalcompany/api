package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/controllers"
	"github.com/medellinoriginalcompany/api/handlers"
)

func HandleRequest() {
	//* Gerenciamento de rotas
	r := gin.Default()

	r.GET("/example", handlers.Example)
	r.POST("/registro", controllers.Signup)
	r.POST("/login", controllers.Login)
	// r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
