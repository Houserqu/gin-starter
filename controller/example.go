package controller

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/lib"
)
import "houserqu.com/gin-starter/module/example"

func InitUserRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		tom := example.GetHello()
		lib.Log(tom)
		c.JSON(200, tom)
	})
}
