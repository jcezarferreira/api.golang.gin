package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcezarferreira/api.gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.FindAllStudents)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.AddStudent)
	r.DELETE("/students/:id", controllers.RemoveStudentById)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.GET("/:name", controllers.Greetings)
	r.Run()
}
