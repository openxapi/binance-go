package models

// V2AccountStatusResponse represents global message '#/components/messages/v2AccountStatusResponse'
type V2AccountStatusResponse struct {
	Id MessageID `json:"id,omitempty"` // id property
	RateLimits []struct {
		Count int `json:"count,omitempty"` // count property
		Interval string `json:"interval,omitempty"` // interval property
		IntervalNum int `json:"intervalNum,omitempty"` // intervalNum property
		Limit int `json:"limit,omitempty"` // limit property
		RateLimitType string `json:"rateLimitType,omitempty"` // rateLimitType property
	} `json:"rateLimits,omitempty"` // rateLimits property
	Result struct {
		Assets []struct {
			Asset string `json:"asset,omitempty"` // asset property
			AvailableBalance string `json:"availableBalance,omitempty"` // availableBalance property
			CrossUnPnl string `json:"crossUnPnl,omitempty"` // crossUnPnl property
			CrossWalletBalance string `json:"crossWalletBalance,omitempty"` // crossWalletBalance property
			InitialMargin string `json:"initialMargin,omitempty"` // initialMargin property
			MaintMargin string `json:"maintMargin,omitempty"` // maintMargin property
			MarginBalance string `json:"marginBalance,omitempty"` // marginBalance property
			MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"` // maxWithdrawAmount property
			OpenOrderInitialMargin string `json:"openOrderInitialMargin,omitempty"` // openOrderInitialMargin property
			PositionInitialMargin string `json:"positionInitialMargin,omitempty"` // positionInitialMargin property
			UnrealizedProfit string `json:"unrealizedProfit,omitempty"` // unrealizedProfit property
			UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
			WalletBalance string `json:"walletBalance,omitempty"` // walletBalance property
		} `json:"assets,omitempty"` // assets property
		AvailableBalance string `json:"availableBalance,omitempty"` // availableBalance property
		MaxWithdrawAmount string `json:"maxWithdrawAmount,omitempty"` // maxWithdrawAmount property
		Positions []struct {
			InitialMargin string `json:"initialMargin,omitempty"` // initialMargin property
			IsolatedMargin string `json:"isolatedMargin,omitempty"` // isolatedMargin property
			IsolatedWallet string `json:"isolatedWallet,omitempty"` // isolatedWallet property
			MaintMargin string `json:"maintMargin,omitempty"` // maintMargin property
			Notional string `json:"notional,omitempty"` // notional property
			PositionAmt string `json:"positionAmt,omitempty"` // positionAmt property
			PositionSide string `json:"positionSide,omitempty"` // positionSide property
			Symbol string `json:"symbol,omitempty"` // symbol property
			UnrealizedProfit string `json:"unrealizedProfit,omitempty"` // unrealizedProfit property
			UpdateTime int64 `json:"updateTime,omitempty"` // updateTime property
		} `json:"positions,omitempty"` // positions property
		TotalCrossUnPnl string `json:"totalCrossUnPnl,omitempty"` // totalCrossUnPnl property
		TotalCrossWalletBalance string `json:"totalCrossWalletBalance,omitempty"` // totalCrossWalletBalance property
		TotalInitialMargin string `json:"totalInitialMargin,omitempty"` // totalInitialMargin property
		TotalMaintMargin string `json:"totalMaintMargin,omitempty"` // totalMaintMargin property
		TotalMarginBalance string `json:"totalMarginBalance,omitempty"` // totalMarginBalance property
		TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin,omitempty"` // totalOpenOrderInitialMargin property
		TotalPositionInitialMargin string `json:"totalPositionInitialMargin,omitempty"` // totalPositionInitialMargin property
		TotalUnrealizedProfit string `json:"totalUnrealizedProfit,omitempty"` // totalUnrealizedProfit property
		TotalWalletBalance string `json:"totalWalletBalance,omitempty"` // totalWalletBalance property
	} `json:"result,omitempty"` // result property
	Status int `json:"status,omitempty"` // status property
}


