package router

import (
	docs "Gin/docs"
	"Gin/middleware"
	"Gin/sever"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Router() *gin.Engine {
	r := gin.Default()

	//允许跨域
	//r.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "token"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	//html模板
	r.LoadHTMLGlob("view/*")

	// 公共方法
	User := r.Group("/user", middleware.IPLimite())
	docs.SwaggerInfo.BasePath = ""
	r.MaxMultipartMemory = 32
	{
		User.POST("/login", middleware.VerifyEmail(), server.Login)
		User.GET("/userinfo", middleware.TokenParse(), server.Userinfo)
		User.POST("/register", server.Register)
		User.GET("/send_code", server.Send_email)
		User.POST("/file", middleware.TokenParse(), server.File)
		User.GET("/websocket", server.Websecket)
		User.GET("/join", middleware.TokenParse(), server.JoinPrivate)
		User.GET("/delete", middleware.TokenParse(), server.DelPrivate)
		User.GET("friend_list", middleware.TokenParse(), server.Friendlist)
		User.GET("/get_message", middleware.TokenParse(), server.ChatRecord)
		User.GET("/html", server.Html)
		User.GET("/online", middleware.TokenParse(), server.GetUserOnline)
		User.POST("/SetHeadPicture", middleware.TokenParse(), server.SetHeadPicture)
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
	Administrator := r.Group("administrator", middleware.AuthontokenParse())
	{
		Administrator.POST("/")

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
