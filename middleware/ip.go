package middleware

import (
	"Gin/models"
	"Gin/utility"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func IPLimite() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ip", c.ClientIP())
		//访问IP
		ip := c.ClientIP()
		user := c.MustGet("use")
		use := user.(*utility.User)
		//查询是否为封杀ip
		banIP := models.GetbanIp(ip)
		if banIP {
			c.Abort()
			return
		}
		//插入ip数据
		err := models.InsertIpbyUser(&models.IPs{ip, time.Now().Unix(), use.Indently})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "系统错误！",
			})
			log.Println("insert err", err)
			c.Abort()
			return
		}
		number, err := models.GetIPNumber(ip)
		if err != nil {
			c.Abort()
			return
		}

		if number > 1500 {
			//加入黑名单
			err := models.BanIP(&models.Bans{ip, time.Now().Format("2006-01-02 15:00:00")})
			if err != nil {
				c.Abort()
				return
			}
			//	封禁用户

		}

	}
}
