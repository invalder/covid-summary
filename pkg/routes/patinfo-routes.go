package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/invalder/covid-summary/pkg/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.HelloWorld)
	r.GET("/covid", controllers.GetAllPatients)
	r.GET("/covid/summary", controllers.GetSummary)

	return r
}
