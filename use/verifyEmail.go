package use

import (
	"Gin/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 邮件验证
func VerifyEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		code := c.PostForm("code")
		ctx := context.Background()
		result, err := db.Rdb.HMGet(ctx, username, "username", "time", "randString").Result()
		fmt.Println(db.Rdb.TTL(ctx, username).Result())
		fmt.Println(result)
		if err != nil {
			return
		}
		hset := make([]string, 0)
		for _, v := range result {
			if v == nil {
				return
			}
			//fmt.Println(i, "v", v)
			hset = append(hset, v.(string))
		}
		log.Println(hset)
		if username == "" || code == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "用户名或验证码不能为空！",
				"code": 1,
			})
			c.Abort()
		}
		if hset[2] == code {
			c.Set("status", 200)
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "验证不通过！",
				"code": 1,
			})
			c.Abort()
			return
		}
	}
}
