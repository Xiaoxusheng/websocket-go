package controltion

import (
	"Gin/sever"
	"github.com/gin-gonic/gin"
)

func Files(r *gin.Engine) {
	file := r.Group("/files")
	r.MaxMultipartMemory = 32
	{
		file.POST("/img", server.File)
	}

}
