package controller

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/internal"
	"houserqu.com/gin-starter/module/example"
)

type FindReq struct {
	ID int `form:"id" binding:"required"`
}

func InitExampleRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		tom := example.GetHello()
		internal.Log.Debug(tom)
		c.JSON(internal.Success(tom, "success"))
	})

	// 查单个
	r.GET("/find", func(c *gin.Context) {
		var query FindReq
		if err := c.ShouldBind(&query); err != nil {
			c.JSON(internal.Error(internal.ErrParam, err.Error()))
			return
		}

		data := example.GetModelByID(query.ID)
		c.JSON(200, data)
	})

	// 查列表
	r.GET("/findAll", func(c *gin.Context) {
		data, err := example.GetModelAll()
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(200, data)
	})
}
