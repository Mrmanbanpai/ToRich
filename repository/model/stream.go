package model

type StreamText struct {
	Stream string         `json:"stream"`
	Data   StreamTextData `json:"data"`
}

type StreamTextData struct {
	EventType    string `json:"e"` // 事件类型
	Timestamp    int64  `json:"E"` // 事件时间
	TradingPair  string `json:"s"` // 交易对
	EventId      int    `json:"t"` // 交易ID
	TradingPrice string `json:"p"` // 成交价格
	TradingCount string `json:"q"` // 成交数量
	BuyOrderId   int    `json:"b"` // 买方订单id
	SellOrderId  int    `json:"a"` // 卖房订单id
	TradingTime  int    `json:"T"` // 交易时间
	IsSeller     bool   `json:"m"` // 买方是否是做市方。如true，则此次成交是一个主动卖出单，否则是一个主动买入单。
}
