package example

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/core"
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
		core.Log(c).Info("example")
		// 根据 ID 查找
		data, err := GetModelByID()
		if err != nil {
			c.JSON(core.ResNotFound(err.Error()))
			return
		}

		c.JSON(core.ResSuccess(data, ""))
	})
}
