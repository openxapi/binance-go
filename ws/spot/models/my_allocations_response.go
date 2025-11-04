package models

// MyAllocationsResponse represents global message '#/components/messages/myAllocationsResponse'
type MyAllocationsResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result []struct {
		AllocationId int64 `json:"allocationId,omitempty"` // allocationId property
		AllocationType string `json:"allocationType,omitempty"` // allocationType property
		Commission string `json:"commission,omitempty"` // commission property
		CommissionAsset string `json:"commissionAsset,omitempty"` // commissionAsset property
		IsAllocator bool `json:"isAllocator,omitempty"` // isAllocator property
		IsBuyer bool `json:"isBuyer,omitempty"` // isBuyer property
		IsMaker bool `json:"isMaker,omitempty"` // isMaker property
		OrderId int64 `json:"orderId,omitempty"` // orderId property
		OrderListId int64 `json:"orderListId,omitempty"` // orderListId property
		Price string `json:"price,omitempty"` // price property
		Qty string `json:"qty,omitempty"` // qty property
		QuoteQty string `json:"quoteQty,omitempty"` // quoteQty property
		Symbol string `json:"symbol,omitempty"` // symbol property
		Time int64 `json:"time,omitempty"` // time property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


