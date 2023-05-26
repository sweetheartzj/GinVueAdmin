package request

type Register struct {
	Username string `json:"username" example:"用户名" binding:"required"`
	Password string `json:"password" example:"密码" binding:"required"`
	NickName string `json:"nickName" example:"昵称" binding:"required"`
	Phone    string `json:"phone" example:"手机号" binding:"required"`
	Email    string `json:"email" example:"邮箱" binding:"required"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
