package main

import (
	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/routes"
)

func init() {
	config.Loadenv()      // Carregar .env variables
	config.Credentials()	// Carregar credenciais do Cloudinary
	database.Connection() // Conectar a database
}

func main() {

	// Inicializa o roteador Gin e registra as rotas
	routes.HandleRequest()
}
