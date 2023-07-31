package main

import (
	"fmt"

	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/routes"
)

func init() {
	config.Loadenv()      // Load .env variables
	database.Connection() // Connect to database
}

func main() {

	fmt.Println("Testando .env")
	// Inicializa o roteador Gin e registra as rotas
	routes.HandleRequest()
}