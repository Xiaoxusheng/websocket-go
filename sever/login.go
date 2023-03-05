package server

import (
	"Gin/db"
	"Gin/models"
	"Gin/utility"
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
	var userfrom models.LoginForm
	userfrom.Username = c.PostForm("username")
	userfrom.Password = c.PostForm("password")
	//userfrom.Email = c.PostForm("email")
	//err := c.ShouldBind(&userfrom)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	log.Println(userfrom)
	if userfrom.Username == "" || userfrom.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "用户登录失败，用户名或密码不能为空",
		})
	} else {
		var user []models.User

		err := db.DB.Select(&user, "select * from user where username=? and password=? ", userfrom.Username, utility.Createmd5(userfrom.Password))
		if err != nil {
			log.Println("查询出错了:", err)
			c.JSON(http.StatusOK, gin.H{
				"msg":  "sql err" + err.Error(),
				"code": 1,
			})
		}
		if len(user) > 0 {
			//c.JSON(http.StatusOK, Response{0, user, "登陆成功"})
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登陆成功",
				"code": 200,
				"data": utility.Createtoken(userfrom.Username, userfrom.Password),
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
	var Register models.LoginForm
	err := c.ShouldBind(&Register)
	if err != nil {
		return
	}
	//use := c.PostForm("username")
	//pwd := c.PostForm("password")
	//email := c.PostForm("email")
	fmt.Println(Register)

	if Register.Username == "" || Register.Password == "" || Register.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "用户名，密码，邮箱不能为空",
			"code": 1,
		})
	} else {
		var user, emails []models.User
		err := db.DB.Select(&emails, "select * from user where email=?", Register.Email)
		if err != nil {
			return
		}
		fmt.Println(emails)
		if len(emails) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "此邮箱已经注册过,请更换新的邮箱",
				"code": 1,
			})
			return
		}
		err = db.DB.Select(&user, "select * from user where username=?  ", Register.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "sql err" + err.Error(),
				"code": 1,
			})
			log.Println("查询出错了:", err)
			return
		}
		fmt.Println(user)
		if len(user) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "用户名已存在，请重新注册",
				"code": 1,
			})
		} else {
			r, err := db.DB.Exec("insert into user(username,password,use_status,register_time ,email)values(?,?,?,?,?)", Register.Username, utility.Createmd5(Register.Password), 0, time.Now().Format("2006--01--02 15:03:05"), Register.Email)
			if err != nil {
				log.Println("注册出错:", err)
				return
			}
			id, err1 := r.LastInsertId()
			if err1 != nil {
				log.Println("出错了:, ", err)
				return
			}
			log.Println("insert succ:", id)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "注册成功",
			})
		}

	}

}
