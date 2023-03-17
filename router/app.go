package router

import (
	docs "Gin/docs"
	"Gin/sever"
	"Gin/use"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// 控制层
func Router() *gin.Engine {
	r := gin.Default()
	// 公共方法
	User := r.Group("/user", use.IPLimite())
	docs.SwaggerInfo.BasePath = ""
	r.MaxMultipartMemory = 32
	{
		User.POST("/login", use.VerifyEmail(), server.Login)
		User.POST("/register", server.Register)
		User.GET("/send_code", server.Send_email)
		User.POST("/img", server.File)
		User.GET("/websocket", server.Websecket)
		User.GET("/join", server.JoinPrivate)
		User.GET("/delete", server.DelPrivate)
		User.GET("friend_list", server.Friendlist)
		User.GET("/get_message", server.ChatRecord)
	}
	// 群聊
	Group := r.Group("/group")
	{
		Group.POST("/group", server.CreateGroup)
		Group.GET("/join", server.JoinGroup)
		Group.GET("/exit", server.ExitGroup)
		Group.GET("/grouplist", server.GetGroupList)
	}

	// 管理员,
	Administrator := r.Group("administrator", use.AuthontokenParse())
	{
		Administrator.POST("/")

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
