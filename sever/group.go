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

// CreateGroup
// 创建群聊
// @Summary 创建群聊接口
// @Param token header string true "token"
// @Param info formData string true "群名称"
// @Schemes
// @Description info token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}  "{"code": 200,"msg": "创建群聊成功,群号为:8660920"}"
// @Router  /group/group  [post]
func CreateGroup(c *gin.Context) {
	info := c.PostForm("info")
	//获取token里面的idently
	user := c.MustGet("use")
	use := user.(*utility.User)
	if info == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	//根据indently获取信息
	create_uesr, err := models.GetUsername(use.Indently)
	if err != nil {
		log.Println("查询出错:", err)
		return
	}
	roomid := utility.GetRoomId()
	GroupNumber := models.SelectGroupNumber(create_uesr.Username, "group")
	if GroupNumber == 5 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "创建群聊达到上限！",
		})
		return
	}
	f := models.CreateRoom(&models.Room_id{roomid, use.Indently, "group", time.Now().Unix(), create_uesr.Username, info})
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "创建群聊成功,群号为:" + roomid,
		})
		err := models.InsertUseridently(&models.User_room{use.Indently, roomid, time.Now().Unix(), time.Now().Unix(), "group", ""})
		if err != nil {
			log.Println("INSERT ERR", err)
			return
		}

	}

}

// JoinGroup
// 加入群聊
// @Summary 加入群聊接口
// @Param room_id query string true "群号"
// @Param token header string true "token"
// @Schemes
// @Description room_id token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}  "{"code": 200, "msg": "加入群聊成功！"}"
// @Router  /group/join  [get]
func JoinGroup(c *gin.Context) {
	room_id := c.Query("room_id")
	user := c.MustGet("use")
	use := user.(*utility.User)
	fmt.Println(room_id)
	if room_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "房间号不能为空！",
		})
		return
	}
	value := models.GetRoomId(room_id)
	err := models.GetGroup(use.Indently, room_id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经加入群聊！",
		})
		return
	}
	if value {
		err := models.InsertUseridently(&models.User_room{use.Indently, room_id, time.Now().Unix(), time.Now().Unix(), "group", ""})
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

// ExitGroup
// PingExample godoc
// 退出群聊
// @Summary 退出群聊接口
// @Param room_id query string true "群号"
// @Param token header string true "token"
// @Schemes
// @Description room_id token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}  "{"code":200,"msg":"退出成功！"}"
// @Router  /group/exit  [get]
func ExitGroup(c *gin.Context) {
	room_id := c.Query("room_id")
	user := c.MustGet("use")
	use := user.(*utility.User)
	fmt.Println(room_id)
	if room_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	f := models.GetRoomId(room_id)
	if !f {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "群号不存在！，",
		})
		return
	}

	err := models.ExitGroupUser(use.Indently, room_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误，" + err.Error(),
		})
		return
	}
	room := models.GetGroupLord(use.Indently)
	if len(room) != 0 {
		err := models.DissolveGroup(room_id)
		if err != nil {
			return
		}
		models.DelGroup(room_id)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "解散群聊成功！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功！",
	})
}

// ExitGroup
// PingExample godoc
// 获取群成员列表
// @Summary 获取群成员列表接口
// @Param room_id query string true "群号"
// @Schemes
// @Description room_id  为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
//
//	@Success 200 {string}  "{"code": 200,"data": {	"data": [	{	"Indently": "6a2a462c-a107-48ea-82e5-74e308327e6f",	"Username": "admin", "Password": "21232f297a57a5a743894a0e4a801fc3",	"Use_status": 0,	"Register_time": "2023-03-13 17:05:08",	"Email": "3096407764@qq.com",	"account": "3169387148"	},	{		"Indently": "cacda2d3-4a77-4afa-94b5-6ff2c036d126", "Username": "leilong",		   "Password": "e10adc3949ba59abbe56e057f20f883e", "Use_status": 0, "Register_time": "2023-03-12 19:07:16","Email": "3096407768@qq.com", "account": "9546270155"}]	}, "msg": "获取数据成功！"}"
//
// @Router  /group/grouplist  [get]
func GetGroupList(c *gin.Context) {
	room_id := c.Query("room_id")
	if room_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	f := models.GetRoomId(room_id)
	if !f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "群号不存在！",
		})
		return
	}
	Gouplist := models.GetGroupList(room_id)
	user := make([]*models.User, 0)
	for _, userroom := range Gouplist {
		username, err := models.GetUsername(userroom.Useridently)
		if err != nil {
			return
		}
		user = append(user, username)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"data": user,
		},
	})
}

// GetGrouperList
// PingExample godoc
// 群聊
// @Summary 群列表
// @Param token header string true "token"
// @Schemes
// @Description  token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
//
//	@Success 200 {string}  "{	"code": 200,	"data": [ { "roomidently": "3725963", "useridently": "", "room_type": "group", "creaet_time": 0, "create_uesr": "admin", "info": "小学生" } ], "msg": "获取数据成功！" }"
//
// @Router  /group/grouperlist   [get]
func GetGrouperList(c *gin.Context) {
	user := c.MustGet("use")
	use := user.(*utility.User)
	//fmt.Println(use.Indently)
	grouplist := models.GetGrouperList(use.Indently)
	if grouplist == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"data": grouplist,
		},
	})
}
