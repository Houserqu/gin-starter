package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"houserqu.com/gin-starter/controller"
	"log"
	"os"
)

func main() {
	// 加载配置文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// 注册路由
	controller.InitUserRouter(r)

	// 监听端口
	err = r.Run(os.Getenv("SERVER_ADDR")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("Error listen")
	}
}
