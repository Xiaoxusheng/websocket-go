package use

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func IPLimite() gin.HandlerFunc {
	return func(c *gin.Context) {
		//访问IP
		ip := c.ClientIP()

		body := c.Request.GetBody
		fmt.Println(ip, body)
	}
}
