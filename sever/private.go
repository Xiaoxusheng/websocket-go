package server

import (
	"Gin/models"
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 添加好友
func JoinPrivate(c *gin.Context) {
	//好友的account
	account := c.Query("account")
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	//验证添加的好友是否存在
	u, e := models.GetUserByaccount(account)
	if e != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "账号不存在！",
		})
		return
	}
	//是否已经为好友
	other, err := models.GetOther(u.Indently, use.Indently)
	if other && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经互为好友！",
		})
		return
	}
	//房间号
	id := utility.GetRoomId()
	err = models.InsertUseridently(&models.User_room{use.Indently, id, time.Now().Unix(), time.Now().Unix(), "private"})
	if err != nil {
		log.Println("1", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	mc, err := models.GetUserByaccount(account)
	fmt.Println(mc.Indently)
	log.Println("err", err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	err = models.InsertUseridently(&models.User_room{mc.Indently, id, time.Now().Unix(), time.Now().Unix(), "private"})
	if err != nil {
		log.Println("2", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}
	//
	f := models.CreateRoom(&models.Room_id{id, use.Indently, "private", time.Now().Unix(), use.Username, mc.Username})
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加好友成功！",
		})
	}
}
