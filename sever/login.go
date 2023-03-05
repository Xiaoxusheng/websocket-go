package server

import (
	"Gin/db"
	"Gin/models"
	"Gin/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 业务实现层
/*登录*/
func Login(c *gin.Context) {
	fmt.Println(c.FullPath())
	username := c.PostForm("username")
	password := c.PostForm("password")
	//userfrom.Email = c.PostForm("email")
	//err := c.ShouldBind(&userfrom)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(userfrom)
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "用户登录失败，用户名或密码不能为空",
		})
	} else {
		getidently, err := models.Getidently(username, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 1,
				"msg":  "系统错误！",
			})
			return
		}
		//fmt.Println(c.GetInt("status"))
		//fmt.Print(len(getidently))
		if len(getidently) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "登陆失败，用户不存在！",
				"code": 1,
			})
			return
		}
		if len(getidently) > 0 && c.GetInt("status") == 200 {
			c.JSON(http.StatusOK, gin.H{
				"msg":   "登陆成功",
				"code":  200,
				"token": utility.Createtoken(getidently[0].Indently, username),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "登陆失败，密码或验证码错误",
				"code": 1,
			})

		}

	}

}

// 注册
func Register(c *gin.Context) {
	fmt.Println(c.FullPath())
	var Register models.LoginForm
	use := c.PostForm("username")
	err := c.ShouldBind(&Register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "验证失败！",
			"code": 1,
		})
		return
	}

	//pwd := c.PostForm("password")
	//email := c.PostForm("email")
	fmt.Println(Register)
	fmt.Println(use)

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
		fmt.Println("email", emails)
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
			account := utility.GetUserAccount()
			r, err := models.InsetuserLoginForm(Register, account)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  "系统错误！",
				})
				return
			}
			fmt.Print(r.LastInsertId())
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"msg":     "注册成功",
				"account": account,
			})
		}

	}

}
