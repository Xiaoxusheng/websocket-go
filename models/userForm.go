package models

import (
	"Gin/db"
	"Gin/utility"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// 注册
type User struct {
	Indently      string `db:"indently" `
	Username      string `db:"username"  `
	Password      string `db:"password" `
	Use_status    int    `db:"use_status" `
	Register_time string `db:"register_time" `
	Email         string `db:"email" `
	Account       string `json:"account"`
}

// 登录
type LoginForm struct {
	Username string `json:"username" form:"username" binding:"required" msg:"用户名不能为空" form:"username"`
	Password string `json:"password"  form:"password" binding:"min=3,max=7" msg:"密码长度不能小于3大于7"`
	Email    string `json:"email"   form:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

func (u User) User() string {
	return "User"
}

func (f LoginForm) UserForm() string {
	return " "
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
	r, err := db.DB.Exec("insert into user(indently,username,password,use_status,register_time ,email,account)values(?,?,?,?,?,?,?)", utility.Uuid(), Register.Username, utility.Createmd5(Register.Password), 0, time.Now().Format("2006--01--02 15:03:05"), Register.Email, account)
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
