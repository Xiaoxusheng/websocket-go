package router

import (
	"Gin/sever"
	"Gin/use"
	"github.com/gin-gonic/gin"
)

// 控制层
func Router() *gin.Engine {
	r := gin.Default()
	User := r.Group("/user")
	Administrator := r.Group("administrator")
	r.MaxMultipartMemory = 32
	{
		//公共方法
		User.POST("/login", use.Yanzheng, server.Login)
		User.POST("/register", server.Register)
		User.POST("/", server.Socket)
		User.POST("/img", server.File)
	}

	{ //管理员
		Administrator.POST("/")

	}
	return r
}
