package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/medellinoriginalcompany/api/controllers"
	"github.com/medellinoriginalcompany/api/handlers"
	"github.com/medellinoriginalcompany/api/middleware"
)

func HandleRequest() {
	//* Gerenciamento de rotas
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
	}))

	r.POST("/registro", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)

	r.POST("/admin/registro", middleware.RequireAdmin, controllers.AdminSignup)
	r.POST("/admin/login", controllers.AdminLogin)
	r.POST("/admin/logout", middleware.RequireAdmin, controllers.Logout)

	r.GET("/example", handlers.Example)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
