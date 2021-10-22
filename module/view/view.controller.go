package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitViewRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
