package models

import (
	"Gin/db"
	"Gin/utility"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type User struct {
	Indently      string `json:"indently" `
	Username      string `json:"username"  `
	Password      string `json:"password" `
	Use_status    int    `json:"use_status" `
	Register_time string `json:"register_time" `
	Email         string `json:"email" `
	Account       string `json:"account"`
	Headpicture   string `json:"headpicture"`
}

// 登录
type LoginForm struct {
	Username string `json:"username" form:"username" binding:"required" message:"用户名不能为空" form:"username"`
	Password string `json:"password"  form:"password" binding:"min=3,max=7" message:"密码长度不能小于3大于7"`
	Email    string `json:"email"   form:"email" binding:"email" message:"邮箱地址格式不正确"`
}

func (u User) User() string {
	return "User"
}

// 查询用户是否存在
func Getidently(username, password string) ([]User, error) {
	var user []User
	err := db.DB.Select(&user, "select * from user where username=? and password=? ", username, utility.Createmd5(password))
	if err != nil {
		log.Println("查询出错了:", err)
		return nil, err
	}
	return user, nil
}

// 插入数据
func InsetuserLoginForm(Register LoginForm, account string) (sql.Result, error) {
	r, err := db.DB.Exec("insert into user(indently,username,password,use_status,register_time ,email,account,headpicture)values(?,?,?,?,?,?,?,?)", utility.Uuid(), Register.Username, utility.Createmd5(Register.Password), 0, time.Now().Format("2006--01--02 15:03:05"), Register.Email, account, "")
	if err != nil {
		log.Println("注册出错:", err)
		return nil, err
	}
	//res, err := r.LastInsertId()
	return r, nil
}

// 根据useridently查询用户信息
func GetUsername(useridently string) (*User, error) {
	var use User
	//fmt.Println(useridently)
	err := db.DB.Get(&use, "select * from user where indently=?", useridently)
	if err != nil {
		//log.Println("查询出错:", err)
		return nil, err
	}
	return &use, nil
}

// 根据account获取用户信息
func GetUserByaccount(account string) (*User, error) {
	var use User
	err := db.DB.Get(&use, "select * from user where account=? ", account)
	if err != nil {
		return nil, err
	}
	return &use, nil
}

// 设置头像
func SetHeadPicture(url, indently string) error {
	_, err := db.DB.Exec("update user set headpicture=? where indently=?", url, indently)
	if err != nil {
		return err
	}
	return nil
}

// 根据用户名获取email
func GetEmail(username string) (*User, error) {
	var use User
	//fmt.Println(useridently)
	err := db.DB.Get(&use, "select * from user where username=?", username)
	if err != nil {
		//log.Println("查询出错:", err)
		return nil, err
	}
	return &use, nil
}
