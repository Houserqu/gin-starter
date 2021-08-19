package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitViewRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
