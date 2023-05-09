package server

import (
	"Gin/models"
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ChatRecord
// PingExample godoc
// 获取聊天记录
// @Summary 获取聊天记录接口
// @Param room_id query string true "房间号"
// @Param token header string true "token"
// @Param pageSize query string false "pageSize"
// @Schemes
// @Description room_id token 为必填 pageSize默认值为1 选填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{"code":1,"data":{"data":[{"id":29,"idently":"6a2a462c-a107-48ea-82e5-74e308327e6f","message_id":6858759,"message":"你好，我是张三","room_idently":"0820018","messagesend_time":1679052325},{"id":30,"idently":"cacda2d3-4a77-4afa-94b5-6ff2c036d126","message_id":72843315,"message":"你好，我是李四","room_idently":"0820018","messagesend_time":1679052347},{"id":31,"idently":"cacda2d3-4a77-4afa-94b5-6ff2c036d126","message_id":91900639,"message":"你好，我是李四","room_idently":"0820018","messagesend_time":1679054157},{"id":32,"idently":"cacda2d3-4a77-4afa-94b5-6ff2c036d126","message_id":63367923,"message":"你好，我是李四","room_idently":"0820018","messagesend_time":1679054157},{"id":33,"idently":"cacda2d3-4a77-4afa-94b5-6ff2c036d126","message_id":30613339,"message":"你好，我是李四","room_idently":"0820018","messagesend_time":1679054158}]},"msg":"获取数据成功！"}"
// @Router  /user/get_message   [get]
func ChatRecord(c *gin.Context) {
	roomId := c.Query("room_id")
	pageSize := c.DefaultQuery("pageSize", "1")
	pageNumber := 20
	user := c.MustGet("use")
	use := user.(*utility.User)
	if roomId == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误！",
		})
		return
	}
	mc := models.GetUserByUserindently(use.Indently)
	//房间成员
	fmt.Println(mc)
	f := false
	for _, res := range *mc {
		if res.Roomidently == roomId {
			f = true
		}
	}
	if !f {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "违法操作！",
		})
		return
	}
	messages, err := models.GetMessages(roomId, pageSize, pageNumber)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "1系统错误，" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"data": messages,
		},
	})

}

// ChatRecord
// PingExample godoc
// 撤回消息记录
// @Summary 消息撤回接口
// @Param message_id query string true "房间号"
// @Param token header string true "token"
// @Schemes
// @Description message_id token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string}  "{}"
// @Router   /user/recallchatrecord    [get]
func RecallChatRecord(c *gin.Context) {
	message_id := c.Query("message_id")
	if message_id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	//判断是否非法数据
	err := models.GetMess(message_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "非法操作！",
		})
		return
	}
	//撤回消息
	err = models.DelMessage(message_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "撤回成功！",
	})
}

func Html(c *gin.Context) {
	//times := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	//number, err := models.GetIPNumber(times)
	//if err != nil {
	//	return
	//}
	//c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//	"iplist": number,
	//})
}
