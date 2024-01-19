package main

import (
	"github.com/wermvasconcelos/api_go_gin/database"
	"github.com/wermvasconcelos/api_go_gin/models"
	"github.com/wermvasconcelos/api_go_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Werm", CPF: "00000000000", RG: "100000000"},
		{Nome: "Lizz", CPF: "11111111111", RG: "299999000"},
	}
	routes.HandleRequests()
}
