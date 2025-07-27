package models

import (
	"encoding/json"
)

// MyTradesResponseRateLimitsItem represents a nested object structure
type MyTradesResponseRateLimitsItem struct {
	// count property (example: 20)
	Count int64 `json:"count,omitempty"`
	// interval property (example: "MINUTE")
	Interval string `json:"interval,omitempty"`
	// intervalNum property (example: 1)
	IntervalNum int64 `json:"intervalNum,omitempty"`
	// limit property (example: 6000)
	Limit int64 `json:"limit,omitempty"`
	// rateLimitType property (example: "REQUEST_WEIGHT")
	RateLimitType string `json:"rateLimitType,omitempty"`
}

// MyTradesResponseResultItem represents a nested object structure
type MyTradesResponseResultItem struct {
	// commission property (example: "0.00000000")
	Commission string `json:"commission,omitempty"`
	// commissionAsset property (example: "BNB")
	CommissionAsset string `json:"commissionAsset,omitempty"`
	// id property (example: 1650422481)
	Id int64 `json:"id,omitempty"`
	// isBestMatch property (example: true)
	IsBestMatch bool `json:"isBestMatch,omitempty"`
	// isBuyer property (example: false)
	IsBuyer bool `json:"isBuyer,omitempty"`
	// isMaker property (example: true)
	IsMaker bool `json:"isMaker,omitempty"`
	// orderId property (example: 12569099453)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: -1)
	OrderListId int64 `json:"orderListId,omitempty"`
	// price property (example: "23416.10000000")
	Price string `json:"price,omitempty"`
	// qty property (example: "0.00635000")
	Qty string `json:"qty,omitempty"`
	// quoteQty property (example: "148.69223500")
	QuoteQty string `json:"quoteQty,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// time property (example: 1660801715793)
	Time int64 `json:"time,omitempty"`
}

// MyTradesResponse - Receive response from myTrades
// Message name: Account trade history (USER_DATA) Response
type MyTradesResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []MyTradesResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []MyTradesResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of MyTradesResponse
func (s MyTradesResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


