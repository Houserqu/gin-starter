package controller

import (
	"github.com/gin-gonic/gin"
	"houserqu.com/gin-starter/internal"
)

func InitRedisRouter(r *gin.Engine) {
	// 根据 Key 查找
	r.GET("/redis/get", func(c *gin.Context) {
		res, err := internal.Redis.Get(internal.RedisCtx, "APP").Result()
		if err != nil {
			c.JSON(internal.ResNotFound(err.Error()))
			return
		}

		c.JSON(internal.ResSuccess(res, ""))
	})

	// 设置值
	r.GET("/redis/set", func(c *gin.Context) {
		err := internal.Redis.Set(internal.RedisCtx, "APP", "gin-starter", 0).Err()
		if err != nil {
			c.JSON(internal.ResNotFound(err.Error()))
			return
		}

		c.JSON(internal.ResSuccess("success", ""))
	})
}
