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
	Appid     int    `json:"appid"`      //appid
	Token     string `json:"token"`      //授权码
	LoginType int    `json:"login_type"` //登录类型 WEIXIN_LOGIN_TYPE=1
}

type LoginRsp struct {
	Code        int    `json:"errcode"` //错误码
	ErrMsg      string `json:"errmsg"`  //错误信息
	Ticket      string `json:"ticket"`  //登录票据
	Openid      string `json:"openid"`  //账号openid
	RegisterUid int64  `json:"uid"`     //uid
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
