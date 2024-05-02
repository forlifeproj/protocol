package order_manager

// 生成订单id
type GenerateOrderIdReq struct {
	Uid int64 `json:"uid"`
}

type GenerateOrderIdRsp struct {
	OrderId string `json:"orderId"`
}

// 订单状态轮询接口
type QueryOrderStatusReq struct {
	OrderId string `json:"orderId"`
}

type QueryOrderStatusRsp struct {
	Status       int    `json:"status"`
	MaxPollTimes int    `json:"maxPollTimes"`
	PollInterval int    `json:"pollInterval"`
	NeedPoll     int    `json:"needPoll"`
	BizInfo      string `json:"bizInfo"`
}

type OrderInfo struct {
}

// 拉取订单列表
type GetOrderListReq struct {
	Uid         int64  `json:"uid"`
	PageSize    int    `json:"pageSize"`
	StrPassBack string `json:"strPassBack"`
}

type GetOrderListRsp struct {
	HasMore     int         `json:"hasMore"`
	StrPassBack string      `json:"strPassBack"`
	OrderList   []OrderInfo `json:"orderList"`
}
