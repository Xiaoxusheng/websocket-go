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

// 创建群聊
func CreateGroup(c *gin.Context) {
	//获取token里面的idently
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	room_type := c.PostForm("room_type")
	info := c.PostForm("info")
	if room_type == "" || info == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	create_uesr, err := models.GetUsername(use.Indently)
	if err != nil {
		log.Println("查询出错:", err)
		return
	}
	roomid := utility.GetRoomId()
	GroupNumber := models.SelectGroupNumber(create_uesr, "group")
	if GroupNumber == 5 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "创建群聊达到上限！",
		})
		return
	}
	f := models.CreateRoom(&models.Room_id{roomid, use.Indently, room_type, time.Now().Unix(), create_uesr, info})
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建群聊成功,群号为:" + roomid,
		})
		err := models.InsertUseridently(&models.User_room{use.Indently, roomid, time.Now().Unix(), time.Now().Unix(), "group"})
		if err != nil {
			log.Println("INSERT ERR", err)
			return
		}

	}

}

// 加入群聊
func JoinGroup(c *gin.Context) {
	room_id := c.Query("room_id")
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println(room_id)
	if room_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "房间号不能为空！",
		})
		return
	}
	value := models.GetRoomId(room_id)
	if value {
		err := models.InsertUseridently(&models.User_room{use.Indently, room_id, time.Now().Unix(), time.Now().Unix(), "group"})
		if err != nil {
			log.Println("INSERT ERR", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "加入群聊失败！",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "加入群聊成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "房间号错误！",
		})
		return
	}
}

//退出群聊
