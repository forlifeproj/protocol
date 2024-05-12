package account

const (
	WEIXIN_OPEN_TYPE = 1
)

const (
	GetUidServiceName = "gfriends.account.GetUid"
)

// 登录注册getuid请求
type GetUidReq struct {
	Appid        int    `json:"appid"`
	OpenType     int    `json:"open_type"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionid"`
	AutoRegister int    `json:"auto_register"`
}

type GetUidRsp struct {
	Code        int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	RegisterUid int64  `json:"uid"`
}
