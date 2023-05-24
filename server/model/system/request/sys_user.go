package request

type Register struct {
	Username string `json:"username" example:"用户名"`
	Password string `json:"password" example:"密码"`
	NickName string `json:"nickName" example:"昵称"`
	Phone    string `json:"phone" example:"手机号"`
	Email    string `json:"email" example:"邮箱"`
}

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
