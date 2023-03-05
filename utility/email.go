package utility

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

func Sendemails() {
	e := email.NewEmail()
	//发送者
	e.From = "小学生 <wx06594@gmail.com>"
	//接收者
	e.To = []string{"3096407768@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	//主题
	e.Subject = "登录验证码"

	//文本
	e.Text = []byte(Getcode())
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")

	err := e.SendWithStartTLS("smtp.gmail.com:465", smtp.PlainAuth("", "wx06594@gmail.com", "lei20001205@", "smtp.gmail.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		log.Println(err)

	}

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
