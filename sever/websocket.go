package server

import (
	"Gin/models"
	"Gin/utility"
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
}
var client = make(map[string]*websocket.Conn)

func Websecket(c *gin.Context) {
	message := &models.Message{}
	token := c.GetHeader("token")
	use, err := utility.ParseWithClaims(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "token不能为空！",
		})
		c.Abort()
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//client[i]
	client[use.Indently] = ws
	fmt.Println(client)
	fmt.Println(len(client))
	//fmt.Println(client)
	defer ws.Close()
	for {

		//获得数据
		err := ws.ReadJSON(message)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(message.Room_idently)
		//log.Printf("recv: %v   p:%v", messageType, string(p))
		//聊天记录记录
		models.InsertMessage(&models.Message{
			use.Indently, utility.GetMessageId(), message.Message, message.Room_idently, time.Now().Unix(),
		})
		//获取在线人数
		ws := models.GetUserbyIdentlyRoomId(message.Room_idently)
		//发送数据
		fmt.Println(client)
		for _, w := range ws {
			if cc, ok := client[w.Useridently]; ok {
				if err := cc.WriteMessage(websocket.TextMessage, []byte(message.Message)); err != nil {
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
