package utility

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func Sendemails(code string) {
	e := email.NewEmail()
	//发送者
	e.From = "小学生 <2673893724@qq.com>"
	//接收者
	e.To = []string{"3096407768@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	//主题
	e.Subject = "登录验证码"
	//文本
	e.Text = []byte("【IM在线聊天系统】您的验证码为：" + code)
	//e.HTML = []byte("<h1>g</h1>")

	err := e.SendWithStartTLS("smtp.qq.com:587", smtp.PlainAuth("", "2673893724@qq.com", "myucgbfyfcnodjch", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		log.Println("stmp:", err)

	}
	fmt.Println("发送成功！")
	//myucgbfyfcnodjch
}

// 生成验证码
func Getcode() string {
	var str string
	rand.Seed(time.Now().UnixNano())
	rand.Int63()
	for i := 0; i < 6; i++ {
		str += strconv.FormatInt(rand.Int63n(10), 10)
	}
	return str
}
