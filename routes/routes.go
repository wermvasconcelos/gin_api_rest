package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wermvasconcelos/api_go_gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
