package login

const (
	WEIXIN_LOGIN_TYPE = 1
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
	RegisterUid int64  `json:"uid"`
}
