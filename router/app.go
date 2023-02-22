package router

import (
	"Gin/sever"
	"github.com/gin-gonic/gin"
)

// 控制层
func Router() *gin.Engine {
	r := gin.Default()
	User := r.Group("/user")
	r.MaxMultipartMemory = 32
	{
		User.POST("/login", server.Login)
		User.POST("/register", server.Register)
		User.POST("/", server.Socket)
		User.POST("/img", server.File)
	}

	return r
}
