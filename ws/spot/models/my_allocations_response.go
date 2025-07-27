package models

import (
	"encoding/json"
)

// MyAllocationsResponseRateLimitsItem represents a nested object structure
type MyAllocationsResponseRateLimitsItem struct {
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

// MyAllocationsResponseResultItem represents a nested object structure
type MyAllocationsResponseResultItem struct {
	// allocationId property (example: 0)
	AllocationId int64 `json:"allocationId,omitempty"`
	// allocationType property (example: "SOR")
	AllocationType string `json:"allocationType,omitempty"`
	// commission property (example: "0.00000000")
	Commission string `json:"commission,omitempty"`
	// commissionAsset property (example: "BTC")
	CommissionAsset string `json:"commissionAsset,omitempty"`
	// isAllocator property (example: false)
	IsAllocator bool `json:"isAllocator,omitempty"`
	// isBuyer property (example: false)
	IsBuyer bool `json:"isBuyer,omitempty"`
	// isMaker property (example: false)
	IsMaker bool `json:"isMaker,omitempty"`
	// orderId property (example: 500)
	OrderId int64 `json:"orderId,omitempty"`
	// orderListId property (example: -1)
	OrderListId int64 `json:"orderListId,omitempty"`
	// price property (example: "1.00000000")
	Price string `json:"price,omitempty"`
	// qty property (example: "0.10000000")
	Qty string `json:"qty,omitempty"`
	// quoteQty property (example: "0.10000000")
	QuoteQty string `json:"quoteQty,omitempty"`
	// symbol property (example: "BTCUSDT")
	Symbol string `json:"symbol,omitempty"`
	// time property (example: 1687319487614)
	Time int64 `json:"time,omitempty"`
}

// MyAllocationsResponse - Receive response from myAllocations
// Message name: Account allocations (USER_DATA) Response
type MyAllocationsResponse struct {
	// id property
	Id string `json:"id,omitempty"`
	// rateLimits property
	RateLimits []MyAllocationsResponseRateLimitsItem `json:"rateLimits,omitempty"`
	// result property
	Result []MyAllocationsResponseResultItem `json:"result,omitempty"`
	// status property
	Status int64 `json:"status,omitempty"`
}

// String returns string representation of MyAllocationsResponse
func (s MyAllocationsResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}


