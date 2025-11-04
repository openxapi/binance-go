package models

// TickerResponse represents global message '#/components/messages/tickerResponse'
type TickerResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		CloseTime int64 `json:"closeTime,omitempty"` // closeTime property
		Count int `json:"count,omitempty"` // count property
		FirstId int64 `json:"firstId,omitempty"` // firstId property
		HighPrice string `json:"highPrice,omitempty"` // highPrice property
		LastId int64 `json:"lastId,omitempty"` // lastId property
		LastPrice string `json:"lastPrice,omitempty"` // lastPrice property
		LowPrice string `json:"lowPrice,omitempty"` // lowPrice property
		OpenPrice string `json:"openPrice,omitempty"` // openPrice property
		OpenTime int64 `json:"openTime,omitempty"` // openTime property
		PriceChange string `json:"priceChange,omitempty"` // priceChange property
		PriceChangePercent string `json:"priceChangePercent,omitempty"` // priceChangePercent property
		QuoteVolume string `json:"quoteVolume,omitempty"` // quoteVolume property
		Symbol string `json:"symbol,omitempty"` // symbol property
		Volume string `json:"volume,omitempty"` // volume property
		WeightedAvgPrice string `json:"weightedAvgPrice,omitempty"` // weightedAvgPrice property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


