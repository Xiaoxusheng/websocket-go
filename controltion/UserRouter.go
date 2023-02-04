package controltion

import (
	"Gin/sever"
	"github.com/gin-gonic/gin"
)

// 控制层
func User(r *gin.Engine) {
	User := r.Group("/user")
	r.MaxMultipartMemory = 32
	{
		User.POST("/login", server.Login)
		User.POST("/register", server.Register)
		User.POST("/", server.Socket)
	}

}
