package req

type LoginBody struct {
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}
