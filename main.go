package main

import (
	"Gin/controltion"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//开启日志颜色
	gin.ForceConsoleColor()

	//静态文件
	r.Static("/img", "./img")
	//允许跨域
	r.Use(cors.Default())
	//用户
	controltion.User(r)
	//文件
	controltion.Files(r)

	r.Run(":5701") // 监听并在 0.0.0.0:80 上启动服务
}
