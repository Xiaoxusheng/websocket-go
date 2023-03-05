package models

import (
	"Gin/db"
	"log"
)

type Message struct {
	Idently          string `json:"idently"`
	Message_id       int64  `json:"message_id"`
	Message          string `json:"message"`
	Room_idently     string `json:"room_idently"`
	Messagesend_time int64  `json:"messagesend_time"`
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
