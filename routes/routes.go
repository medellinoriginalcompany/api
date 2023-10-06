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
		AllowOrigins:     []string{
			"http://localhost:5173",
			"http://localhost:5174",
		},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-Requested-With"},
		AllowMethods:     []string{"GET, POST, PUT, DELETE, OPTIONS"},
	}))

	r.POST("/registro", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuth, controllers.Logout)

	r.POST("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/admin/validate", middleware.RequireAdmin, controllers.Validate)

	r.POST("/admin/registro", middleware.RequireAdmin, controllers.AdminSignup)
	r.POST("/admin/login", controllers.AdminLogin)
	r.POST("/admin/logout", middleware.RequireAdmin, controllers.AdminLogout)
	r.POST("/admin/cadastrar-produto", middleware.RequireAdmin, controllers.AddProduct)
	
	r.GET("/admin/products", middleware.RequireAdmin, controllers.GetProducts)
	r.GET("/admin/products/properties", middleware.RequireAdmin, controllers.GetProductProperties)

	r.POST("/admin/products/add-property/tamanhos", middleware.RequireAdmin, controllers.AddProductSize)
	r.POST("/admin/products/add-property/cores", middleware.RequireAdmin, controllers.AddProductColor)
	r.POST("/admin/products/add-property/categorias", middleware.RequireAdmin, controllers.AddProductCategory)
	r.POST("/admin/products/add-property/tipos", middleware.RequireAdmin, controllers.AddProductType)
	
	r.DELETE("/admin/delete-product/:id", middleware.RequireAdmin, controllers.DeleteProduct)
	r.DELETE(("/admin/delete/:type/:id"), middleware.RequireAdmin, controllers.DeleteProductProperty)

	r.GET("/example", handlers.Example)


	r.Run()
}
