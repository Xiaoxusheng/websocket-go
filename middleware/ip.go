package middleware

import (
	"github.com/gin-gonic/gin"
)

func IPLimite() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ip", c.ClientIP())
		//访问IP
		//ip := c.ClientIP()
		//token := c.GetHeader("token")
		//uses, err := utility.ParseWithClaims(token)
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 1,
		//		"msg":  "系统错误，" + err.Error(),
		//	})
		//	return
		//}
		//err := models.InsertIpbyUser(&models.IPs{ip, time.Now().Unix(), c.Request.Header["User-Agent"][0]})
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 1,
		//		"msg":  "系统错误！",
		//	})
		//	log.Println("insert err", err)
		//	c.Abort()
		//}
		//time := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
		//fmt.Println(time)
		//number, err := models.GetIPNumber(1679203619)
		//if err != nil {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 1,
		//		"msg":  "系统错误！",
		//	})
		//	log.Println(err)
		//	return
		//}
		//fmt.Println("number:", number)
		//body := c.Request.Header["User-Agent"][0]
		//fmt.Println(ip, body)
		//

	}
}
