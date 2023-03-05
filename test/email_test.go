package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"testing"
)

func Test_end(t *testing.T) {
	e := email.NewEmail()
	//发送者
	e.From = "小学生 <wx06594@gmail.com>"
	//接收者
	e.To = []string{"3096407768@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	//主题
	e.Subject = "测试"
	//文本
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "wx06594@gmail.com", "lei20001205@", "smtp.gmail.com"))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	err := e.SendWithStartTLS("smtp.gmail.com:465", smtp.PlainAuth("", "wx06594@gmail.com", "lei20001205@", "smtp.gmail.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		log.Println(err)
		return
	}
}
