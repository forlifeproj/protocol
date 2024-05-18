package wxpay

import "time"

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type Amount struct {
	PayerTotal    int    `json:"payer_total"`
	Total         int    `json:"total"`
	Currency      string `json:"currency"`
	PayerCurrency string `json:"payer_currency"`
}

type GoodsDetail struct {
	GoodsRemark    string `json:"goods_remark"`
	Quantity       int    `json:"quantity"`
	DiscountAmount int    `json:"discount_amount"`
	GoodsID        string `json:"goods_id"`
	UnitPrice      int    `json:"unit_price"`
}

type PromotionDetail struct {
	Amount              int           `json:"amount"`
	WechatpayContribute int           `json:"wechatpay_contribute"`
	CouponID            string        `json:"coupon_id"`
	Scope               string        `json:"scope"`
	MerchantContribute  int           `json:"merchant_contribute"`
	Name                string        `json:"name"`
	OtherContribute     int           `json:"other_contribute"`
	Currency            string        `json:"currency"`
	StockID             string        `json:"stock_id"`
	GoodsDetail         []GoodsDetail `json:"goods_detail"`
}

type Payer struct {
	OpenID string `json:"openid"`
}

type SceneInfo struct {
	DeviceID string `json:"device_id"`
}

// 根据Resource.Ciphertext字段解析出来的结构
type Transaction struct {
	TransactionID   string            `json:"transaction_id"`
	Amount          Amount            `json:"amount"`
	MchID           string            `json:"mchid"`
	TradeState      string            `json:"trade_state"`
	BankType        string            `json:"bank_type"`
	PromotionDetail []PromotionDetail `json:"promotion_detail"`
	SuccessTime     time.Time         `json:"success_time"`
	Payer           Payer             `json:"payer"`
	OutTradeNo      string            `json:"out_trade_no"`
	AppID           string            `json:"AppID"`
	TradeStateDesc  string            `json:"trade_state_desc"`
	TradeType       string            `json:"trade_type"`
	Attach          string            `json:"attach"`
	SceneInfo       SceneInfo         `json:"scene_info"`
}

type WxPayCallbackReq struct {
	ID           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     Resource  `json:"resource"`
}

type WxPayCallbackRsp struct {
	Code    string `json:"code"` //FAIL, SUCCESS
	Message Amount `json:"message"`
}
