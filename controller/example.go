package controller

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/internal"
	"houserqu.com/gin-starter/module/example"
	"strconv"
	"time"
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
	r.GET("/model/:id", func(c *gin.Context) {
		// 参数转换和校验
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(internal.ResParamError("id invalid"))
			return
		}

		// 根据 ID 查找
		data, err := example.GetModelByID(id)
		if err != nil {
			c.JSON(internal.ResNotFound(err.Error()))
			return
		}

		c.JSON(internal.ResSuccess(data, ""))
	})

	// 查列表
	r.GET("/model/list", func(c *gin.Context) {
		data, err := example.GetModelAll()
		if err != nil {
			c.JSON(internal.ResNotFound(""))
			return
		}

		c.JSON(internal.ResSuccess(data, ""))
	})

	// 创建
	r.POST("/model/create", func(c *gin.Context) {
		var params ReqModelCreate
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(internal.ResParamError(err.Error()))
			return
		}

		result := &example.Example{Name: params.Name, Age: params.Age, Email: params.Email, Birthday: time.Now()}
		err := example.CreateModel(result)
		if err != nil {
			c.JSON(internal.ResError(internal.ErrCreateFail, err.Error()))
			return
		}

		c.JSON(internal.ResSuccess(result, ""))
	})

	// 更新
	r.POST("/model/update", func(c *gin.Context) {
		var params ReqModelUpdate
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(internal.ResParamError(err.Error()))
			return
		}

		result, err := example.UpdateModel(params.ID, params.Name, params.Email)
		if err != nil {
			c.JSON(internal.ResError(internal.ErrUpdateFail, ""))
			return
		}

		c.JSON(internal.ResSuccess(result, ""))
	})

	// 删除
	r.POST("/model/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(internal.ResParamError("id invalid"))
			return
		}

		err = example.DelModel(id)
		if err != nil {
			c.JSON(internal.ResError(internal.ErrDeleteFail, err.Error()))
			return
		}

		c.JSON(internal.ResSuccess(id, ""))
	})
}
