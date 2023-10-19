package main

import (

	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/routes"
	// "github.com/medellinoriginalcompany/api/handlers"
)

func init() {
	config.Loadenv()      	// Carregar .env variables
	config.Credentials()		// Carregar credenciais do Cloudinary
	config.LoggerHandler(); // Configurar logger
	database.Connection() 	// Conectar a database
}

func main() {

	// Inicializa o roteador Gin e registra as rotas
	routes.HandleRequest()
}
