package models

import (
	"Gin/db"
	"fmt"
	"log"
)

type User_room struct {
	Useridently string `json:"useridently"`
	Roomidently string `json:"roomidently"`
	Create_time int64  `json:"create_time"`
	Update_time int64  `json:"update_time"`
	Room_type   string `json:"room_type"`
}
type Userroom struct {
	Id          int    `json:"id"`
	Useridently string `json:"useridently"`
	Roomidently string `json:"roomidently"`
	Create_time int64  `json:"create_time"`
	Update_time int64  `json:"update_time"`
	Room_type   string `json:"room_type"`
}

// 加入群聊
func InsertUseridently(user_room *User_room) error {
	_, err := db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type) value (?,?,?,?,?)", user_room.Useridently, user_room.Roomidently, user_room.Create_time, user_room.Update_time, user_room.Room_type)
	if err != nil {
		return err
	}
	return nil
}

// 查
func GetUserbyIdentlyRoomId(roomidently string) []User_room {
	fmt.Println(roomidently)
	user_room := []User_room{}
	err := db.DB.Select(&user_room, "select * from user_room where  roomidently=?", roomidently)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println(user_room)
	return user_room
}

// 加好友
func PrivateInsertUseridently(idently1, idently2 string) {
	db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type) value (?,?,?,?,?)")

}

// 是否已经互为好友
func GetOther(indently1, indently2 string) (bool, error) {
	user_room1 := Userroom{}
	user_room2 := Userroom{}
	err := db.DB.Get(&user_room1, "select * from user_room where  Useridently=?", indently1)
	if err != nil {
		return true, err
	}

	err = db.DB.Get(&user_room2, "select * from user_room where  Useridently=?", indently2)
	if err != nil {
		return true, err
	}

	if user_room1.Roomidently == user_room2.Roomidently {
		return true, nil
	}
	return false, nil
}