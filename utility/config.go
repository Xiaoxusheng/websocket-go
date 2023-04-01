package utility

var Key = []byte("suhusduhhfmifr56436$$#$#597**)*^$^*(s65+978t$67")

type Message struct {
	Message_id       int64  `json:"message_id"`
	Room_idently     string `json:"roomidently"`
	Room_type        string `json:"room_type"`
	Messagesend_time int64  `json:"messagesend_time"`
	Message          string `json:"message"`
	Useridently      string `json:"useridently"`
}

type Userinfo struct {
	Room_id   string      `json:"room_id"`
	Userinfo  interface{} `json:"userinfo"`
	Room_type string      `json:"room_type"`
}

/*{
   "room_idently":"123456",
   "message":"你好，我是张三",
   "room_type":"group"
}*/

/*
{
   "room_idently":"7172065",
   "message":"你好，我是张三",
   "room_type":"group"
}*/

//{
//"room_idently":"3725963",
//"message":"你好，我是leilong",
//"room_type":"group"
//}
