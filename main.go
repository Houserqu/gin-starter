package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"houserqu.com/gin-starter/internal/middleware"
	"houserqu.com/gin-starter/module/example"
	"houserqu.com/gin-starter/module/view"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Access())

	if os.Getenv("GIN_LOG") == "true" {
		r.Use(gin.Logger())
	}

	// 静态文件
	r.Static("/public", "./public")
	r.LoadHTMLGlob("./module/**/*.html")

	// 注册路由
	view.InitViewRouter(r)
	example.InitExampleRouter(r)

	// 监听端口
	err := r.Run(os.Getenv("SERVER_ADDR")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("Error listen")
	}
}
