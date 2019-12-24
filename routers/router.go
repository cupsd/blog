package routers

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.Static("/statics/", "./statics")

	r.GET("/index", controller.IndexHandler)

	return r
}
