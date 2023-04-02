package models

import (
	"Gin/db"
	"encoding/json"
	"log"
	"strconv"
)

type Message struct {
	Idently          string `json:"idently"`
	Message_id       int64  `json:"message_id"`
	Message          string `json:"message"`
	Room_idently     string `json:"room_idently"`
	Messagesend_time int64  `json:"messagesend_time"`
}

// 获取查询数据
type Messages struct {
	Id               int    `json:"id"`
	Idently          string `json:"idently"`
	Message_id       int64  `json:"message_id"`
	Message          string `json:"message"`
	Room_idently     string `json:"room_idently"`
	Messagesend_time int64  `json:"messagesend_time"`
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *Message) GetMessage() string {
	return "message"
}

func InsertMessage(message *Message) {
	_, err := db.DB.Exec("insert into message(idently,message_id,message,room_idently,messagesend_time) value (?,?,?,?,?)", message.Idently, message.Message_id, message.Message, message.Room_idently, message.Messagesend_time)
	if err != nil {
		log.Println("INSERT ERR:", err)
		return
	}
}

func GetMessages(id string, pageSize string, pageNumber int) (*[]Messages, error) {
	messages := []Messages{}
	pageSizes, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	err = db.DB.Select(&messages, "select * from message  where  room_idently=? order by messagesend_time  limit  ?,? ", id, pageSizes-1, pageNumber)
	if err != nil {
		log.Println("SELECT ERR:", err)
		return nil, err
	}
	return &messages, nil
}
