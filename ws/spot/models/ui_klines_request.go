package models

// UiKlinesRequest represents global message '#/components/messages/uiKlinesRequest'
type UiKlinesRequest struct {
	Id MessageID `json:"id,omitempty"` // id property
	Method string `json:"method,omitempty"` // method property
	Params struct {
		EndTime int64 `json:"endTime,omitempty"`
		Interval string `json:"interval"` // See klines
		Limit int `json:"limit,omitempty"` // Default: 500; Maximum: 1000
		StartTime int64 `json:"startTime,omitempty"`
		Symbol string `json:"symbol"`
		TimeZone string `json:"timeZone,omitempty"` // Default: 0 (UTC)
	} `json:"params,omitempty"` // params property
}


