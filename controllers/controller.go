package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wermvasconcelos/api_go_gin/database"
	"github.com/wermvasconcelos/api_go_gin/models"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai" + nome + ",tudo belez?",
	})
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
