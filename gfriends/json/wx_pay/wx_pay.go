package wx_pay

// 微信预下单
type WxPrePayReq struct {
	OpenId  string `json:"openid"`   // 付款人openid
	Amount  int64  `json:"amount"`   // 付款金额, 分
	BizInfo string `json:"biz_info"` // 业务信息
}

type WxPrePayRsp struct {
	PrepayId string `json:"prepay_id"` // 业务信息
}
