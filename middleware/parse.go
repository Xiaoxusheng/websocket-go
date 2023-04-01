package middleware

import (
	"Gin/db"
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
	Use_status int64  `json:"use_status"`
	Username   string `json:"username"`
	Indently   string `json:"indently"`
}

func AuthontokenParse() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		use, err := utility.ParseWithClaims(token)
		if err != nil {
			c.Abort()
			return
		}
		if use.Username == "" || use.Indently == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "用户名不能为空！",
				"code": "1",
			})
			return
		}
		var admin Admin
		err = db.DB.Select(admin, "select * from user where username=?", use.Username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "系统内部出错！",
			})
			return
		}
		fmt.Println(admin)
		if admin.Use_status == 999 && admin.Username == use.Username {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "验证通过",
			})
			c.Set("indently", use.Indently)
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "验证失败，无权限!",
			})
			return
		}

	}

}
