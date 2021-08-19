package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"houserqu.com/gin-starter/controller"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	// 注册路由
	controller.InitUserRouter(r)

	// 监听端口
	err := r.Run(os.Getenv("SERVER_ADDR")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("Error listen")
	}
}
