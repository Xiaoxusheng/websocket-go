package utility

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Createmd5(s string) string {
	has := md5.Sum([]byte(s))
	md5pwd := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(len(md5pwd))
	return md5pwd
}

// 生成房间号
func GetRoomId() string {
	rand.Seed(time.Now().UnixMicro())
	//七位数
	rand.Int63n(10)
	str := ""
	for i := 0; i < 7; i++ {
		str += strconv.FormatInt(rand.Int63n(10), 10)
	}
	return str
}

// 生成message_id
func GetMessageId() int64 {
	rand.Seed(time.Now().UnixMicro())
	return rand.Int63n(100000000)
}

// 生成用户唯一账号
func GetUserAccount() string {
	rand.Seed(time.Now().UnixMicro())
	//七位数
	rand.Int63n(10)
	str := ""
	for i := 0; i < 10; i++ {
		str += strconv.FormatInt(rand.Int63n(10), 10)
	}
	return str
}
