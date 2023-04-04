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

// PingExample godoc
// @Summary 登录接口
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param code  formData string true "验证码"
// @Schemes
// @Description 用户名 密码 验证码为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}    "{"code": 200, "msg": "登陆成功", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmRlbnRseSI6IjZhMmE0NjJjLWExMDctNDhlYS04MmU1LTc0ZTMwODMyN2U2ZiIsInVzZXJuYW1lIjoiYWRtaW4iLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjc4Nzg2NTM1fQ.P4dJ_f2UGhKbpiIqHxTxghRKwKIlCpF2XjryHCSnKKk"}"
// @Router /user/login [post]
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
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户登录失败，用户名或密码不能为空",
		})
	} else {
		getidently, err := models.Getidently(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "系统错误！",
			})
			return
		}
		//fmt.Println(c.GetInt("status"))
		//fmt.Print(len(getidently))
		if len(getidently) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登陆失败，用户不存在！",
				"code": 1,
			})
			return
		}
		if len(getidently) > 0 && c.GetInt("status") == 200 {
			c.JSON(http.StatusOK, gin.H{
				"msg":   "登陆成功",
				"code":  200,
				"token": utility.CreateToken(getidently[0].Indently, username),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "登陆失败，密码或验证码错误",
				"code": 1,
			})

		}

	}

}

// 注册
// PingExample godoc
// @Summary 注册接口
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param email  formData string true "邮箱"
// @Schemes
// @Description 用户名 密码 邮箱为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 { string } "{"account": "3169387148", "code": 200, "msg": "注册成功"}"
// @Router /user/register  [post]
func Register(c *gin.Context) {
	fmt.Println(c.FullPath())
	var Register models.LoginForm
	//usename := c.PostForm("username")
	err := c.ShouldBind(&Register)
	fmt.Println(Register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "验证失败！",
			"code": 1,
		})
		return
	}

	//pwd := c.PostForm("password")
	//email := c.PostForm("email")
	fmt.Println(Register)
	//fmt.Println(usename)

	if Register.Username == "" || Register.Password == "" || Register.Email == "" {
		c.JSON(http.StatusOK, gin.H{
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
			c.JSON(http.StatusOK, gin.H{
				"msg":  "此邮箱已经注册过,请更换新的邮箱",
				"code": 1,
			})
			return
		}
		err = db.DB.Select(&user, "select * from user where username=?  ", Register.Username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "sql err" + err.Error(),
				"code": 1,
			})
			log.Println("查询出错了:", err)
			return
		}
		fmt.Println(user)
		if len(user) > 0 {
			c.JSON(http.StatusOK, gin.H{
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

// Userinfo
// PingExample godoc
// 聊天记录
// @Summary 获取用户信息接口
// @Param token header string true "token"
// @Schemes
// @Description token 为必填
// @Tags 公共方法
// @Accept multipart/form-data
// @Produce json
// @Success 200 {string}   "{"code":200,"data":{"data":{"Indently":"949c770f-0b1e-4ca0-a15a-927ee5396c3b","Username":"aadmin","Password":"21232f297a57a5a743894a0e4a801fc3","Use_status":0,"Register_time":"2023-03-12 19:07:30","Email":"3096407769@qq.com","account":"8356511203"}},"msg":"获取数据成功！"}"
// @Router   /user/userinfo   [get]
func Userinfo(c *gin.Context) {
	user := c.MustGet("use")
	use := user.(*utility.User)
	username, err := models.GetUsername(use.Indently)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误或用户不存在，" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功！",
		"data": gin.H{
			"data": username,
			"ip":   c.MustGet("ip"),
		},
	})

}
