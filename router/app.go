package router

import (
	"Gin/sever"
	"Gin/use"
	"github.com/gin-gonic/gin"
)

// 控制层
func Router() *gin.Engine {
	r := gin.Default()
	//公共方法
	User := r.Group("/user")
	r.MaxMultipartMemory = 32
	{
		User.POST("/login", use.VerifyEmail(), server.Login)
		User.POST("/register", server.Register)
		User.POST("/send_code", server.Send_email)
		User.POST("/img", server.File)
		User.GET("/websocket", server.Websecket)
		User.GET("/join", server.JoinPrivate)
	}
	//群聊
	Group := r.Group("/group")
	{
		Group.POST("/group", server.CreateGroup)
		Group.GET("/join", server.JoinGroup)
		Group.GET("/delete")
	}

	//管理员,
	Administrator := r.Group("administrator", use.AuthontokenParse())
	{
		Administrator.POST("/")

	}
	return r
}
