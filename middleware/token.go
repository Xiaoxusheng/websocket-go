package middleware

import (
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TokenParse() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		log.Println("token", token == "null")
		fmt.Printf("%T", token)
		if token == "null" || token == "" {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "token不能为空！",
			})
			return
		}
		use, err := utility.ParseWithClaims(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "系统错误！" + err.Error(),
			})
			return
		}
		if use != nil {
			c.Set("use", use)
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "验证失败！",
			})
			return
		}
	}
}