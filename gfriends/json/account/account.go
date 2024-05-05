package account

const (
	WEIXIN_OPEN_TYPE = 1
)

// 登录请求
type GetUidReq struct {
	Appid    int    `json:"appid"`
	OpenType int    `json:"open_type"`
	OpenId   string `json:"openid"`
	UnionId  string `json:"unionid"`
}

type GetUidRsp struct {
	Code        int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	RegisterUid int64  `json:"uid"`
}
