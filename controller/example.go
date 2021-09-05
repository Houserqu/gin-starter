package controller

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/internal"
	"houserqu.com/gin-starter/module/example"
)

type ReqModelCreate struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
	Age   int    `form:"age" binding:"required"`
}

type ReqModelUpdate struct {
	ID    int    `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

func InitExampleRouter(r *gin.Engine) {
	// 查单个
	r.GET("/example", func(c *gin.Context) {
		// 根据 ID 查找
		data, err := example.GetModelByID()
		if err != nil {
			c.JSON(internal.ResNotFound(err.Error()))
			return
		}

		c.JSON(internal.ResSuccess(data, ""))
	})
}
