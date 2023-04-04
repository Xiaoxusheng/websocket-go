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

// PingExample godoc
// @Summary 添加好友接口
// @Param account query string true "账号"
// @Param token header string true "token"
// @Schemes
// @Description 账号 token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":200,"msg":"添加好友成功！"}"
// @Router  /user/join    [get]
func JoinPrivate(c *gin.Context) {
	//好友的account
	account := c.Query("account")
	user := c.MustGet("use")
	use := user.(*utility.User)
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
		log.Println(e)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "账号不存在！",
		})
		return
	}
	//是否已经为好友
	other := models.GetOther(u.Indently, use.Indently)
	fmt.Println(other)
	if other {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "已经互为好友！",
		})
		return
	}
	//房间号
	id := utility.GetRoomId()
	err := models.InsertUseridently(&models.User_room{use.Indently, id, time.Now().Unix(), time.Now().Unix(), "private", u.Indently})
	if err != nil {
		log.Println("1", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}

	err = models.InsertUseridently(&models.User_room{u.Indently, id, time.Now().Unix(), time.Now().Unix(), "private", use.Indently})
	if err != nil {
		log.Println("2", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}

	//
	f := models.CreateRoom(&models.Room_id{id, use.Indently, "private", time.Now().Unix(), use.Username, u.Username})
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加好友成功！",
		})
	}
}

// 删除好友
// PingExample godoc
// @Summary  删除好友接口
// @Param account query string true "账号"
// @Param token header string true "token"
// @Schemes
// @Description  token 账号为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":200,"msg":"删除成功！"}"
// @Router  /user/delete      [get]
func DelPrivate(c *gin.Context) {
	account := c.Query("account")
	user := c.MustGet("use")
	use := user.(*utility.User)
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	//验证账号是否存在
	mc, e := models.GetUserByaccount(account)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "账号不存在！",
		})
		return
	}
	//是否已经为好友
	other := models.GetOther(mc.Indently, use.Indently)
	fmt.Println(other)
	if !other {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "不是好友关系！",
		})
		return
	}
	//删除
	err := models.Del(use.Indently, mc.Indently)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "删除错误！",
		})
		return
	}
	room, err := models.GetRoom(use.Indently)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！",
		})
		return
	}

	err = models.DelGroup(room)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功！",
	})
}

// Friendlist
// PingExample godoc
// @Summary  好友列表接口
// @Param token header string true "token"
// @Schemes
// @Description  token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{ "code": 200, "data": {"data": [    {"Indently": "6a2a462c-a107-48ea-82e5-74e308327e6f", "Username": "admin", "Password": "21232f297a57a5a743894a0e4a801fc3", "Use_status": 0, "Register_time": "2023-03-13 17:05:08","Email": "3096407764@qq.com", "account": "3169387148"}]}, "msg": "获取数据成功！"}"
// @Router  /user/friend_list      [get]
func Friendlist(c *gin.Context) {
	userinfo := c.MustGet("use")
	use := userinfo.(*utility.User)
	frieendlist := models.GetFriendList(use.Indently)
	user := make([]*utility.Userinfo, 0)
	for _, userroom := range frieendlist {
		userinfos, err := models.GetUsername(userroom.Friendidently)
		if err != nil {
			return
		}
		user = append(user, &utility.Userinfo{
			userroom.Roomidently,
			userinfos,
			userroom.Room_type,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"data": user,
		},
	})
}

// 好友在线状态
// PingExample godoc
// @Summary  获取好友在线状态接口
// @Param account query string true "account"
// @Schemes
// @Description  account 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{ "code": 200, "msg": "获取用户状态成功！", "status": true }"
// @Router  /user/online      [get]
func GetUserOnline(c *gin.Context) {
	online := false
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	byaccount, err := models.GetUserByaccount(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误！" + err.Error(),
		})
		return
	}
	for i, _ := range Client {
		if i == byaccount.Indently {
			online = true
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"msg":    "获取用户状态成功！",
				"status": true,
			})
			break
		}
	}
	if !online {
		c.JSON(http.StatusOK, gin.H{
			"code":   1,
			"msg":    "获取用户状态成功！",
			"status": false,
		})
		return
	}
}
