package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"houserqu.com/gin-starter/controller"
	"houserqu.com/gin-starter/internal/middleware"
	"log"
	"os"
)

func main() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Access())

	if os.Getenv("GIN_LOG") == "true" {
		r.Use(gin.Logger())
	}

	// 静态文件
	r.Static("/public", "./web/public")
	r.LoadHTMLGlob("./web/views/*")

	// 注册路由
	controller.InitViewRouter(r)
	controller.InitExampleRouter(r)
	controller.InitRedisRouter(r)

	// 监听端口
	err := r.Run(os.Getenv("SERVER_ADDR")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("Error listen")
	}
}
