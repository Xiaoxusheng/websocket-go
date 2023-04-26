package middleware

import (
	"Gin/models"
	"Gin/utility"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

func IPLimite() gin.HandlerFunc {

	return func(c *gin.Context) {
		var use *utility.User
		var wg sync.WaitGroup
		wg.Add(2)
		c.Set("ip", c.ClientIP())
		//访问IP
		ip := c.ClientIP()
		token := c.GetHeader("token")
		if token == "" {
			use = &utility.User{}
			use.Indently = "No"
		} else {
			use, _ = utility.ParseWithClaims(token)
			if use == nil {
				use = &utility.User{}
				use.Indently = "token 过期"
			}
		}

		//查询是否为封杀ip
		banIP := models.GetbanIp(ip)
		if banIP {
			//403(禁止)服务器拒绝请求
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "您的帐户请求过于频繁，已经被封禁，请联系管理员",
			})
			c.Abort()
			return
		}
		//插入ip数据
		go func() {
			defer wg.Done()
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
		}()
		go func() {
			defer wg.Done()
			//查看次数
			number, err := models.GetIPNumber(ip)
			if err != nil {
				log.Println("GetIPNumber", err)
				c.Abort()
				return
			}
			log.Println("IP请求次数:", number)

			if number > 1000 {
				//加入黑名单
				err := models.BanIP(&models.Bans{ip, time.Now().Format("2006-01-02 15:00:00")})
				if err != nil {
					log.Println("BanIP", err)
					c.Abort()
					return
				}

				//	封禁用户

			}
		}()
		wg.Wait()
		c.Next()

	}
}
