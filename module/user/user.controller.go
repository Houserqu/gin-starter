package user

import (
	"strconv"
	"time"

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

func InitUserRouter(r *gin.Engine) {
	// 查单个
	r.GET("/model/:id", func(c *gin.Context) {
		// 参数转换和校验
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.ResError(c, core.ErrParam, "")
			return
		}

		// 根据 ID 查找
		data, err := GetModelByID(id)
		if err != nil {
			core.ResError(c, core.ErrNotFound, "")
			return
		}

		core.ResSuccess(c, data)
	})

	// 查列表
	r.GET("/model/list", func(c *gin.Context) {
		data, err := GetModelAll()
		if err != nil {
			core.ResError(c, core.ErrNotFound, "")
			return
		}

		core.ResSuccess(c, data)
	})

	// 创建
	r.POST("/model/create", func(c *gin.Context) {
		var params ReqModelCreate
		if err := c.ShouldBindJSON(&params); err != nil {
			core.ResError(c, core.ErrParam, "")
			return
		}

		result := &User{Name: params.Name, Age: params.Age, Email: params.Email, Birthday: time.Now()}
		err := CreateModel(result)
		if err != nil {
			core.ResError(c, core.ErrCreateFail, "")
			return
		}

		core.ResSuccess(c, result)
	})

	// 更新
	r.POST("/model/update", func(c *gin.Context) {
		var params ReqModelUpdate
		if err := c.ShouldBindJSON(&params); err != nil {
			core.ResError(c, core.ErrParam, "")
			return
		}

		result, err := UpdateModel(params.ID, params.Name, params.Email)
		if err != nil {
			core.ResError(c, core.ErrUpdateFail, "")
			return
		}

		core.ResSuccess(c, result)
	})

	// 删除
	r.POST("/model/delete/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			core.ResError(c, core.ErrParam, "id invalid")
			return
		}

		err = DelModel(id)
		if err != nil {
			core.ResError(c, core.ErrDeleteFail, "id invalid")
			return
		}

		core.ResSuccess(c, id)
	})
}
