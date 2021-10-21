package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"houserqu.com/gin-starter/core"
	"houserqu.com/gin-starter/middleware"
	"houserqu.com/gin-starter/module/example"
	"houserqu.com/gin-starter/module/view"
)

func main() {
	core.InitConfig()
	core.InitLogger()

	r := gin.New()

	// 注册中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Access())

	// 静态文件
	r.Static("/public", "./public")
	// html 模板文件
	r.LoadHTMLGlob("./module/**/*.html")

	// 注册路由
	view.InitViewRouter(r)
	example.InitExampleRouter(r)

	// 监听端口
	err := r.Run(viper.GetString("server.addr"))
	if err != nil {
		log.Fatal("Error listen")
	}
}
