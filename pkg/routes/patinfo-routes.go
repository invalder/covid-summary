package routes

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", HelloWorld)
	r.GET("/covid", GetAllPatients)

	return r
}
