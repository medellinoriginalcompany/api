package main

import (
	"fmt"

	"github.com/vnxcius/medellin-api/config"
	"github.com/vnxcius/medellin-api/database"
	"github.com/vnxcius/medellin-api/routes"
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