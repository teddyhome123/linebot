package router

import (
	"linebotim/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/callback", controller.Callback)
	r.GET("/getusermes", controller.GetUserMes)
	r.Run(":8081")
	return r
}
