package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wellrcosta/Api.Crud.Go.Gin/controllers"
)

func HandleRequests() error {
	r := gin.Default()
	r.GET("/", controllers.HealthCheck)
	r.GET("/students", controllers.ListAll)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}
