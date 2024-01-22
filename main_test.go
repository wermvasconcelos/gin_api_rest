package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wermvasconcelos/api_go_gin/controllers"
	"github.com/wermvasconcelos/api_go_gin/database"
	"github.com/wermvasconcelos/api_go_gin/models"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678911", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/werm", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	mockDaResposta := `{"API diz:":"E ai werm, Tudo belez?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
