package models

// 注册
type User struct {
	ID            int    `db:"id" `
	Username      string `db:"username"  `
	Password      string `db:"password" "`
	Use_status    int    `db:"use_status" `
	Register_time string `db:"register_time" `
	Email         string `db:"email" `
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
	return " LoginForm"
}
