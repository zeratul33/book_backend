package req

type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password" v:"required#请输入密码"`
	Nickname string `json:"nickname" v:"required#请输入昵称"`
}
