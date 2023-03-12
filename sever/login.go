package server

import (
	"Gin/db"
	"Gin/models"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
	"time"
)

// 业务实现层
/*登录*/
func Login(c *gin.Context) {
	fmt.Println(c.FullPath())
	var userfrom models.UserForm
	userfrom.Use = c.PostForm("username")
	userfrom.Pwd = c.PostForm("password")
	data := []byte(userfrom.Pwd)
	has := md5.Sum(data)
	md5pwd := fmt.Sprintf("%x", has)
	fmt.Println(userfrom.Use, userfrom.Pwd)
	if userfrom.Use == "" || userfrom.Pwd == "" {
		c.JSON(http.StatusBadRequest, Response{1, "", "用户登录失败，用户名或密码不能为空"})
	} else {
		var user []models.User

		err := db.DB.Select(&user, "select * from user where username=? and password=? ", userfrom.Use, md5pwd)
		if err != nil {
			log.Println("查询出错了:", err)
			c.JSON(http.StatusOK, gin.H{
				"msg":  err,
				"code": 1,
			})
		}
		if len(user) > 0 {
			//c.JSON(http.StatusOK, Response{0, user, "登陆成功"})
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登陆成功",
				"code": 0,
				"data": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "登陆失败，密码或用名错误",
				"code": 1,
			})
		}
	}

}

// 注册
func Register(c *gin.Context) {
	fmt.Println(c.FullPath())

	use := c.PostForm("username")
	pwd := c.PostForm("password")
	email := c.PostForm("email")
	data := []byte(pwd)
	has := md5.Sum(data)
	md5pwd := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if use == "" || pwd == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "用户名，密码，邮箱不能为空",
			"code": 1,
		})
	} else {
		var user []models.User

		err := db.DB.Select(&user, "select * from user where username=?  ", use)
		if err != nil {
			log.Println("查询出错了:", err)
		}
		if len(user) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "用户名已存在，请重新注册",
				"code": 1,
			})
		} else {
			r, err := db.DB.Exec("insert into user(username,password,status,time ,email)values(?,?,?,?,?)", use, md5pwd, 0, time.Now().Format("2006--01--02 15:03:05"), email)
			if err != nil {
				log.Println("注册出错:", err)
				return
			}
			id, err1 := r.LastInsertId()
			if err1 != nil {
				log.Println("出错了:, ", err)
				return
			}
			fmt.Println("insert succ:", id)
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "注册成功",
			})
		}

	}

}
