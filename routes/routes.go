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
		AllowOrigins: []string{
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
	r.POST("/validar", middleware.RequireAuth, controllers.Validate)

	r.POST("/admin/registro", middleware.RequireAdmin, controllers.AdminSignup)
	r.POST("/admin/login", controllers.AdminLogin)
	r.POST("/admin/logout", middleware.RequireAdmin, controllers.AdminLogout)
	r.GET("/admin/validar", middleware.RequireAdmin, controllers.Validate)

	r.GET("/produtos", controllers.GetProducts)
	r.GET("/produtos/:id", controllers.GetProduct)
	r.GET("/produtos/ativos", controllers.GetActiveProducts)

	r.POST("/admin/produtos/cadastrar-produto", middleware.RequireAdmin, controllers.AddProduct)
	r.POST("/admin/produtos/editar-produto/:id", middleware.RequireAdmin, controllers.EditProduct)
	r.POST("/admin/produtos/adicionar-propriedade/tamanhos", middleware.RequireAdmin, controllers.AddProductSize)
	r.POST("/admin/produtos/adicionar-propriedade/cores", middleware.RequireAdmin, controllers.AddProductColor)
	r.POST("/admin/produtos/adicionar-propriedade/categorias", middleware.RequireAdmin, controllers.AddProductCategory)
	r.POST("/admin/produtos/adicionar-propriedade/tipos", middleware.RequireAdmin, controllers.AddProductType)
	r.POST("/admin/produtos/editar-propriedade/tamanhos/:id", middleware.RequireAdmin, controllers.EditProductSize)
	r.POST("/admin/produtos/editar-propriedade/cores/:id", middleware.RequireAdmin, controllers.EditProductColor)
	r.POST("/admin/produtos/editar-propriedade/categorias/:id", middleware.RequireAdmin, controllers.EditProductCategory)
	r.POST("/admin/produtos/editar-propriedade/tipos/:id", middleware.RequireAdmin, controllers.EditProductType)
	r.POST("/admin/produtos/restaurar/:id", middleware.RequireAdmin, controllers.RestoreProduct)

	r.GET("/admin/produtos", middleware.RequireAdmin, controllers.GetProducts)
	r.GET("/admin/produto/:id", middleware.RequireAdmin, controllers.GetProduct)
	r.GET("/admin/produtos/propriedades", middleware.RequireAdmin, controllers.GetProductProperties)
	r.GET("/admin/produtos/lixeira", middleware.RequireAdmin, controllers.GetDeletedProducts)

	r.DELETE("/admin/deletar-produto/:id", middleware.RequireAdmin, controllers.DeleteProduct)
	r.DELETE("/admin/produtos/deletar-permanente/:id", middleware.RequireAdmin, controllers.PermaDeleteProduct)
	r.DELETE("/admin/deletar/:type/:id", middleware.RequireAdmin, controllers.DeleteProductProperty)

	r.GET("/example", handlers.Example)

	r.Run()
}
