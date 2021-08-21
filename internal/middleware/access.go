package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"houserqu.com/gin-starter/internal"
	"time"
)

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()               // 开始时间
		c.Next()                              // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqUri := c.Request.RequestURI        // 请求路由
		statusCode := c.Writer.Status()       // 状态码
		clientIP := c.ClientIP()              // 请求IP

		internal.LogAccess.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
