package main

import (
	"Gin/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := router.Router()
	//开启日志颜色
	gin.ForceConsoleColor()
	//静态文件
	r.Static("/img", "./img")
	//前端页面
	r.Static("/static", "./dist")
	// 记录到log文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Run(":8080") // 监听并在 0.0.0.0:80 上启动服务
}
