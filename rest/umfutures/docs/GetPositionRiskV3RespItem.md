# GetPositionRiskV3RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adl** | Pointer to **int32** |  | [optional] 
**AskNotional** | Pointer to **string** |  | [optional] 
**BidNotional** | Pointer to **string** |  | [optional] 
**BreakEvenPrice** | Pointer to **string** |  | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**IsolatedMargin** | Pointer to **string** |  | [optional] 
**IsolatedWallet** | Pointer to **string** |  | [optional] 
**LiquidationPrice** | Pointer to **string** |  | [optional] 
**MaintMargin** | Pointer to **string** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**Notional** | Pointer to **string** |  | [optional] 
**OpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**PositionInitialMargin** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**UnRealizedProfit** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetPositionRiskV3RespItem

`func NewGetPositionRiskV3RespItem() *GetPositionRiskV3RespItem`

NewGetPositionRiskV3RespItem instantiates a new GetPositionRiskV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPositionRiskV3RespItemWithDefaults

`func NewGetPositionRiskV3RespItemWithDefaults() *GetPositionRiskV3RespItem`

NewGetPositionRiskV3RespItemWithDefaults instantiates a new GetPositionRiskV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdl

`func (o *GetPositionRiskV3RespItem) GetAdl() int32`

GetAdl returns the Adl field if non-nil, zero value otherwise.

### GetAdlOk

`func (o *GetPositionRiskV3RespItem) GetAdlOk() (*int32, bool)`

GetAdlOk returns a tuple with the Adl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdl

`func (o *GetPositionRiskV3RespItem) SetAdl(v int32)`

SetAdl sets Adl field to given value.

### HasAdl

`func (o *GetPositionRiskV3RespItem) HasAdl() bool`

HasAdl returns a boolean if a field has been set.

### GetAskNotional

`func (o *GetPositionRiskV3RespItem) GetAskNotional() string`

GetAskNotional returns the AskNotional field if non-nil, zero value otherwise.

### GetAskNotionalOk

`func (o *GetPositionRiskV3RespItem) GetAskNotionalOk() (*string, bool)`

GetAskNotionalOk returns a tuple with the AskNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskNotional

`func (o *GetPositionRiskV3RespItem) SetAskNotional(v string)`

SetAskNotional sets AskNotional field to given value.

### HasAskNotional

`func (o *GetPositionRiskV3RespItem) HasAskNotional() bool`

HasAskNotional returns a boolean if a field has been set.

### GetBidNotional

`func (o *GetPositionRiskV3RespItem) GetBidNotional() string`

GetBidNotional returns the BidNotional field if non-nil, zero value otherwise.

### GetBidNotionalOk

`func (o *GetPositionRiskV3RespItem) GetBidNotionalOk() (*string, bool)`

GetBidNotionalOk returns a tuple with the BidNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidNotional

`func (o *GetPositionRiskV3RespItem) SetBidNotional(v string)`

SetBidNotional sets BidNotional field to given value.

### HasBidNotional

`func (o *GetPositionRiskV3RespItem) HasBidNotional() bool`

HasBidNotional returns a boolean if a field has been set.

### GetBreakEvenPrice

`func (o *GetPositionRiskV3RespItem) GetBreakEvenPrice() string`

GetBreakEvenPrice returns the BreakEvenPrice field if non-nil, zero value otherwise.

### GetBreakEvenPriceOk

`func (o *GetPositionRiskV3RespItem) GetBreakEvenPriceOk() (*string, bool)`

GetBreakEvenPriceOk returns a tuple with the BreakEvenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakEvenPrice

`func (o *GetPositionRiskV3RespItem) SetBreakEvenPrice(v string)`

SetBreakEvenPrice sets BreakEvenPrice field to given value.

### HasBreakEvenPrice

`func (o *GetPositionRiskV3RespItem) HasBreakEvenPrice() bool`

HasBreakEvenPrice returns a boolean if a field has been set.

### GetEntryPrice

`func (o *GetPositionRiskV3RespItem) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *GetPositionRiskV3RespItem) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *GetPositionRiskV3RespItem) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *GetPositionRiskV3RespItem) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetInitialMargin

`func (o *GetPositionRiskV3RespItem) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *GetPositionRiskV3RespItem) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *GetPositionRiskV3RespItem) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *GetPositionRiskV3RespItem) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetIsolatedMargin

`func (o *GetPositionRiskV3RespItem) GetIsolatedMargin() string`

GetIsolatedMargin returns the IsolatedMargin field if non-nil, zero value otherwise.

### GetIsolatedMarginOk

`func (o *GetPositionRiskV3RespItem) GetIsolatedMarginOk() (*string, bool)`

GetIsolatedMarginOk returns a tuple with the IsolatedMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedMargin

`func (o *GetPositionRiskV3RespItem) SetIsolatedMargin(v string)`

SetIsolatedMargin sets IsolatedMargin field to given value.

### HasIsolatedMargin

`func (o *GetPositionRiskV3RespItem) HasIsolatedMargin() bool`

HasIsolatedMargin returns a boolean if a field has been set.

### GetIsolatedWallet

`func (o *GetPositionRiskV3RespItem) GetIsolatedWallet() string`

GetIsolatedWallet returns the IsolatedWallet field if non-nil, zero value otherwise.

### GetIsolatedWalletOk

`func (o *GetPositionRiskV3RespItem) GetIsolatedWalletOk() (*string, bool)`

GetIsolatedWalletOk returns a tuple with the IsolatedWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedWallet

`func (o *GetPositionRiskV3RespItem) SetIsolatedWallet(v string)`

SetIsolatedWallet sets IsolatedWallet field to given value.

### HasIsolatedWallet

`func (o *GetPositionRiskV3RespItem) HasIsolatedWallet() bool`

HasIsolatedWallet returns a boolean if a field has been set.

### GetLiquidationPrice

`func (o *GetPositionRiskV3RespItem) GetLiquidationPrice() string`

GetLiquidationPrice returns the LiquidationPrice field if non-nil, zero value otherwise.

### GetLiquidationPriceOk

`func (o *GetPositionRiskV3RespItem) GetLiquidationPriceOk() (*string, bool)`

GetLiquidationPriceOk returns a tuple with the LiquidationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPrice

`func (o *GetPositionRiskV3RespItem) SetLiquidationPrice(v string)`

SetLiquidationPrice sets LiquidationPrice field to given value.

### HasLiquidationPrice

`func (o *GetPositionRiskV3RespItem) HasLiquidationPrice() bool`

HasLiquidationPrice returns a boolean if a field has been set.

### GetMaintMargin

`func (o *GetPositionRiskV3RespItem) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *GetPositionRiskV3RespItem) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *GetPositionRiskV3RespItem) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *GetPositionRiskV3RespItem) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetMarginAsset

`func (o *GetPositionRiskV3RespItem) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *GetPositionRiskV3RespItem) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *GetPositionRiskV3RespItem) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *GetPositionRiskV3RespItem) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetMarkPrice

`func (o *GetPositionRiskV3RespItem) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *GetPositionRiskV3RespItem) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *GetPositionRiskV3RespItem) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *GetPositionRiskV3RespItem) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetNotional

`func (o *GetPositionRiskV3RespItem) GetNotional() string`

GetNotional returns the Notional field if non-nil, zero value otherwise.

### GetNotionalOk

`func (o *GetPositionRiskV3RespItem) GetNotionalOk() (*string, bool)`

GetNotionalOk returns a tuple with the Notional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotional

`func (o *GetPositionRiskV3RespItem) SetNotional(v string)`

SetNotional sets Notional field to given value.

### HasNotional

`func (o *GetPositionRiskV3RespItem) HasNotional() bool`

HasNotional returns a boolean if a field has been set.

### GetOpenOrderInitialMargin

`func (o *GetPositionRiskV3RespItem) GetOpenOrderInitialMargin() string`

GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetOpenOrderInitialMarginOk

`func (o *GetPositionRiskV3RespItem) GetOpenOrderInitialMarginOk() (*string, bool)`

GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderInitialMargin

`func (o *GetPositionRiskV3RespItem) SetOpenOrderInitialMargin(v string)`

SetOpenOrderInitialMargin sets OpenOrderInitialMargin field to given value.

### HasOpenOrderInitialMargin

`func (o *GetPositionRiskV3RespItem) HasOpenOrderInitialMargin() bool`

HasOpenOrderInitialMargin returns a boolean if a field has been set.

### GetPositionAmt

`func (o *GetPositionRiskV3RespItem) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *GetPositionRiskV3RespItem) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *GetPositionRiskV3RespItem) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *GetPositionRiskV3RespItem) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetPositionInitialMargin

`func (o *GetPositionRiskV3RespItem) GetPositionInitialMargin() string`

GetPositionInitialMargin returns the PositionInitialMargin field if non-nil, zero value otherwise.

### GetPositionInitialMarginOk

`func (o *GetPositionRiskV3RespItem) GetPositionInitialMarginOk() (*string, bool)`

GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInitialMargin

`func (o *GetPositionRiskV3RespItem) SetPositionInitialMargin(v string)`

SetPositionInitialMargin sets PositionInitialMargin field to given value.

### HasPositionInitialMargin

`func (o *GetPositionRiskV3RespItem) HasPositionInitialMargin() bool`

HasPositionInitialMargin returns a boolean if a field has been set.

### GetPositionSide

`func (o *GetPositionRiskV3RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *GetPositionRiskV3RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *GetPositionRiskV3RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *GetPositionRiskV3RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetSymbol

`func (o *GetPositionRiskV3RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetPositionRiskV3RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetPositionRiskV3RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetPositionRiskV3RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnRealizedProfit

`func (o *GetPositionRiskV3RespItem) GetUnRealizedProfit() string`

GetUnRealizedProfit returns the UnRealizedProfit field if non-nil, zero value otherwise.

### GetUnRealizedProfitOk

`func (o *GetPositionRiskV3RespItem) GetUnRealizedProfitOk() (*string, bool)`

GetUnRealizedProfitOk returns a tuple with the UnRealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnRealizedProfit

`func (o *GetPositionRiskV3RespItem) SetUnRealizedProfit(v string)`

SetUnRealizedProfit sets UnRealizedProfit field to given value.

### HasUnRealizedProfit

`func (o *GetPositionRiskV3RespItem) HasUnRealizedProfit() bool`

HasUnRealizedProfit returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetPositionRiskV3RespItem) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetPositionRiskV3RespItem) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetPositionRiskV3RespItem) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetPositionRiskV3RespItem) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


