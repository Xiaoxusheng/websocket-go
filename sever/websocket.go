package server

import (
	"Gin/middleware"
	"Gin/models"
	"Gin/utility"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins
	},
	Subprotocols: []string{"token"},
}
var Client = make(map[string]*websocket.Conn)

// websocket
// PingExample godoc
// @Summary  websocket连接接口
// @Param token header string true "token"
// @Schemes
// @Description  token 为必填
// @Tags 公共方法
// @Accept json
// @Produce json
// @Router   /user/websocket       [get]
func Websecket(c *gin.Context) {
	message := &models.Message{}
	token := c.Query("token")
	//fmt.Println(token)
	//Authorization := c.GetHeader("Authorization")
	//fmt.Println("Authorization", Authorization)
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "token不能为空！",
		})
		log.Println(err)
		c.Abort()
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//client[i]
	Client[use.Indently] = ws
	fmt.Println(Client)
	fmt.Println(len(Client))
	//fmt.Println(client)
	ws.SetCloseHandler(func(code int, text string) error {
		fmt.Printf("WebSocket connection closed with code %d and reason: %s\n", code, text)
		return nil
	})
	defer ws.Close()
	defer middleware.DelConsumerGroup(message.Room_idently)
	for {
		//获得数据
		err := ws.ReadJSON(message)
		if err != nil {
			log.Println("1", err)
			return
		}
		//消息进入队列
		//err = middleware.Producer(message)
		//if err != nil {
		//	log.Println("生产者", err)
		//	return
		//}
		//fmt.Println("房间", message.Room_idently)
		//创建消费者群组
		//middleware.CreateConsumerGroup(message.Room_idently)
		//log.Printf("recv: %v   p:%v", messageType, string(p))
		//聊天记录记录
		models.InsertMessage(&models.Message{
			use.Indently, utility.GetMessageId(), message.Message, message.Room_idently, time.Now().Unix(),
		})
		//获取在线人数
		ws := models.GetUserbyIdentlyRoomId(message.Room_idently)
		//发送数据
		//fmt.Println(ws)
		//取出消息队列消息
		//log.Println("开始取出消息")
		//messagelist := middleware.Consumer(message.Room_idently)
		////确认消息
		//log.Println("开始确认消息")
		//for _, xMessage := range messagelist {
		//	fmt.Println(message.Room_idently, xMessage.ID)
		//	middleware.Confirmationmessage(message.Room_idently, xMessage.ID)
		//	log.Println("确认完毕", xMessage.ID)
		//}
		//
		//for i, xMessage := range messagelist {
		//	fmt.Println(i, xMessage.Values[message.Room_idently])
		//}
		//for i, i2 := range collection {
		//
		//}
		messages, _ := json.Marshal(message)
		for _, w := range ws {
			if cc, ok := Client[w.Useridently]; ok {
				if err := cc.WriteMessage(websocket.TextMessage, []byte(messages)); err != nil {
					log.Println(err)
					return
				}
			}
		}

		//for _, ws := range client {
		//	if err := ws.WriteMessage(websocket.TextMessage, []byte(message.Message)); err != nil {
		//		log.Println(err)
		//		return
		//	}
		//}
	}

}
