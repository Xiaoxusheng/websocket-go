package models

import (
	"Gin/db"
	"fmt"
	"log"
)

type User_room struct {
	Useridently   string `json:"useridently"`
	Roomidently   string `json:"roomidently"`
	Create_time   int64  `json:"create_time"`
	Update_time   int64  `json:"update_time"`
	Room_type     string `json:"room_type"`
	Friendidently string `json:"friendidently"`
}
type Userroom struct {
	Id            int    `json:"id"`
	Useridently   string `json:"useridently"`
	Roomidently   string `json:"roomidently"`
	Create_time   int64  `json:"create_time"`
	Update_time   int64  `json:"update_time"`
	Room_type     string `json:"room_type"`
	Friendidently string `json:"friendidently"`
}

// InsertUseridently 加入群聊
func InsertUseridently(user_room *User_room) error {
	_, err := db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type,friendidently) value (?,?,?,?,?,?)", user_room.Useridently, user_room.Roomidently, user_room.Create_time, user_room.Update_time, user_room.Room_type, user_room.Friendidently)
	if err != nil {
		return err
	}
	return nil
}

// GetUserbyIdentlyRoomId 查
func GetUserbyIdentlyRoomId(roomidently string) []Userroom {
	user_room := []Userroom{}
	err := db.DB.Select(&user_room, "select * from user_room where  roomidently=?", roomidently)
	if err != nil {
		log.Println(err)
		return nil
	}
	//fmt.Println(user_room)
	return user_room
}

// PrivateInsertUseridently 加好友
func PrivateInsertUseridently(use1 User_room) error {
	_, err := db.DB.Exec("insert into user_room (useridently,roomidently,create_time,update_time,room_type) value (?,?,?,?,?)", use1.Useridently, use1.Roomidently, use1.Create_time, use1.Room_type)
	if err != nil {
		return err
	}
	return nil
}

// 是否已经互为好友
func GetOther(indently1, indently2 string) bool {
	user_room1 := []Userroom{}
	user_room2 := []Userroom{}
	err := db.DB.Select(&user_room1, "select * from user_room where  Useridently=? and room_type=? ", indently1, "private")
	if err != nil {
		//log.Println("err1", err)
		return false
	}

	err = db.DB.Select(&user_room2, "select * from user_room where  Useridently=? and room_type=?", indently2, "private")
	if err != nil {
		//log.Println("err2", err)
		return false
	}
	//fmt.Println(user_room1, user_room2)

	for _, userroom := range user_room1 {
		for _, u := range user_room2 {
			if userroom.Roomidently == u.Roomidently {
				return true
			}
		}
	}

	return false
}

// 删除
func Del(indently1, indently2 string) error {
	exec, err := db.DB.Exec("delete from user_room where useridently=?", indently1)
	if err != nil {
		return err
	}
	fmt.Println(exec)
	exec, err = db.DB.Exec("delete from user_room where useridently=?", indently2)
	if err != nil {
		return err
	}
	return nil
}

// 查询是否已经加入房间
func GetGroup(indently, room_id string) error {
	user_room := Userroom{}
	err := db.DB.Get(&user_room, "select * from user_room where useridently=? and roomidently=?", indently, room_id)
	if err != nil {
		return err
	}
	return nil
}

// 退出群聊
func ExitGroupUser(indently, room_id string) error {
	_, err := db.DB.Exec("delete from user_room where useridently=? and room_type=? and roomidently=?", indently, "group", room_id)
	if err != nil {
		return err
	}

	return nil
}

// 好友列表
func GetFriendList(indently string) []Userroom {
	user_room := []Userroom{}
	err := db.DB.Select(&user_room, "select * from user_room  where useridently=? and room_type=? ", indently, "private")
	if err != nil {
		log.Println(err)
		return nil
	}
	//fmt.Println(user_room)
	return user_room
}

// 判断是否为群主
func GetGroupList(room_id string) []Userroom {
	user_room := []Userroom{}
	err := db.DB.Select(&user_room, "select * from user_room  where roomidently=? and room_type=? ", room_id, "group")
	if err != nil {
		log.Println(err)
		return nil
	}
	//fmt.Println(user_room)
	return user_room
}

// 解散群聊
func DissolveGroup(room_id string) error {
	_, err := db.DB.Exec("delete from user_room where roomidently=? ", room_id)
	if err != nil {
		return err
	}
	return nil

}

// 房间成员
func GetUserByUserindently(useridently string) *[]Userroom {
	room := []Userroom{}
	err := db.DB.Select(&room, "select * from user_room where useridently=?", useridently)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &room
}
