package login

const (
	WEIXIN_LOGIN_TYPE = 1
)

const (
	LoginInServiceName   = "gfriends.login.LoginIn"
	LoginAuthServiceName = "gfriends.login.LoginAuth"
)

// 登录请求
type LoginReq struct {
	Appid     int    `json:"appid"`
	Token     string `json:"token"`
	LoginType int    `json:"login_type"`
}

type LoginRsp struct {
	Code        int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	Ticket      string `json:"ticket"`
	Openid      string `json:"openid"`
	RegisterUid int64  `json:"uid"`
}

// 登录鉴权
type LoginAuthReq struct {
	Appid       int    `json:"appid"`
	Token       string `json:"token"`
	Openid      string `json:"openid"`
	RegisterUid int64  `json:"uid"`
	LoginType   int    `json:"login_type"`
}

type LoginAuthRsp struct {
	Code   int    `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}
