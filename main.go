package main

import (
	"github.com/wermvasconcelos/api_go_gin/database"
	"github.com/wermvasconcelos/api_go_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
