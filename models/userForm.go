package models

// 注册
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"`
	Time     string `json:"time"`
	Email    string `json:"email"`
}
type UserForm struct {
	Use   string `json:"username" binding:"required" msg:"用户名不能为空"`
	Pwd   string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email string `json:"email"   binding:"email" msg:"邮箱地址格式不正确"`
}
