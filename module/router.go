package module

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/module/example"
	"houserqu.com/gin-starter/module/user"
	"houserqu.com/gin-starter/module/view"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", view.IndexView)

	r.GET("/example", example.GetModel)
	r.GET("/error", example.ErrorExample)

	// 用户相关
	userRouter := r.Group("user")
	userRouter.GET("/:id", user.GetUser)            // 查单个
	userRouter.GET("/list", user.GetUserList)       // 查列表
	userRouter.POST("/create", user.CreateUser)     // 创建
	userRouter.POST("/update", user.UpdateUser)     // 更新
	userRouter.POST("/delete/:id", user.DeleteUser) // 删除
}
