# GetAccountV2RespPositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskNotional** | Pointer to **string** |  | [optional] 
**BidNotional** | Pointer to **string** |  | [optional] 
**EntryPrice** | Pointer to **string** |  | [optional] 
**InitialMargin** | Pointer to **string** |  | [optional] 
**Isolated** | Pointer to **bool** |  | [optional] 
**Leverage** | Pointer to **string** |  | [optional] 
**MaintMargin** | Pointer to **string** |  | [optional] 
**MaxNotional** | Pointer to **string** |  | [optional] 
**OpenOrderInitialMargin** | Pointer to **string** |  | [optional] 
**PositionAmt** | Pointer to **string** |  | [optional] 
**PositionInitialMargin** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**UnrealizedProfit** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountV2RespPositionsInner

`func NewGetAccountV2RespPositionsInner() *GetAccountV2RespPositionsInner`

NewGetAccountV2RespPositionsInner instantiates a new GetAccountV2RespPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountV2RespPositionsInnerWithDefaults

`func NewGetAccountV2RespPositionsInnerWithDefaults() *GetAccountV2RespPositionsInner`

NewGetAccountV2RespPositionsInnerWithDefaults instantiates a new GetAccountV2RespPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskNotional

`func (o *GetAccountV2RespPositionsInner) GetAskNotional() string`

GetAskNotional returns the AskNotional field if non-nil, zero value otherwise.

### GetAskNotionalOk

`func (o *GetAccountV2RespPositionsInner) GetAskNotionalOk() (*string, bool)`

GetAskNotionalOk returns a tuple with the AskNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskNotional

`func (o *GetAccountV2RespPositionsInner) SetAskNotional(v string)`

SetAskNotional sets AskNotional field to given value.

### HasAskNotional

`func (o *GetAccountV2RespPositionsInner) HasAskNotional() bool`

HasAskNotional returns a boolean if a field has been set.

### GetBidNotional

`func (o *GetAccountV2RespPositionsInner) GetBidNotional() string`

GetBidNotional returns the BidNotional field if non-nil, zero value otherwise.

### GetBidNotionalOk

`func (o *GetAccountV2RespPositionsInner) GetBidNotionalOk() (*string, bool)`

GetBidNotionalOk returns a tuple with the BidNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidNotional

`func (o *GetAccountV2RespPositionsInner) SetBidNotional(v string)`

SetBidNotional sets BidNotional field to given value.

### HasBidNotional

`func (o *GetAccountV2RespPositionsInner) HasBidNotional() bool`

HasBidNotional returns a boolean if a field has been set.

### GetEntryPrice

`func (o *GetAccountV2RespPositionsInner) GetEntryPrice() string`

GetEntryPrice returns the EntryPrice field if non-nil, zero value otherwise.

### GetEntryPriceOk

`func (o *GetAccountV2RespPositionsInner) GetEntryPriceOk() (*string, bool)`

GetEntryPriceOk returns a tuple with the EntryPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPrice

`func (o *GetAccountV2RespPositionsInner) SetEntryPrice(v string)`

SetEntryPrice sets EntryPrice field to given value.

### HasEntryPrice

`func (o *GetAccountV2RespPositionsInner) HasEntryPrice() bool`

HasEntryPrice returns a boolean if a field has been set.

### GetInitialMargin

`func (o *GetAccountV2RespPositionsInner) GetInitialMargin() string`

GetInitialMargin returns the InitialMargin field if non-nil, zero value otherwise.

### GetInitialMarginOk

`func (o *GetAccountV2RespPositionsInner) GetInitialMarginOk() (*string, bool)`

GetInitialMarginOk returns a tuple with the InitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialMargin

`func (o *GetAccountV2RespPositionsInner) SetInitialMargin(v string)`

SetInitialMargin sets InitialMargin field to given value.

### HasInitialMargin

`func (o *GetAccountV2RespPositionsInner) HasInitialMargin() bool`

HasInitialMargin returns a boolean if a field has been set.

### GetIsolated

`func (o *GetAccountV2RespPositionsInner) GetIsolated() bool`

GetIsolated returns the Isolated field if non-nil, zero value otherwise.

### GetIsolatedOk

`func (o *GetAccountV2RespPositionsInner) GetIsolatedOk() (*bool, bool)`

GetIsolatedOk returns a tuple with the Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolated

`func (o *GetAccountV2RespPositionsInner) SetIsolated(v bool)`

SetIsolated sets Isolated field to given value.

### HasIsolated

`func (o *GetAccountV2RespPositionsInner) HasIsolated() bool`

HasIsolated returns a boolean if a field has been set.

### GetLeverage

`func (o *GetAccountV2RespPositionsInner) GetLeverage() string`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *GetAccountV2RespPositionsInner) GetLeverageOk() (*string, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *GetAccountV2RespPositionsInner) SetLeverage(v string)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *GetAccountV2RespPositionsInner) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetMaintMargin

`func (o *GetAccountV2RespPositionsInner) GetMaintMargin() string`

GetMaintMargin returns the MaintMargin field if non-nil, zero value otherwise.

### GetMaintMarginOk

`func (o *GetAccountV2RespPositionsInner) GetMaintMarginOk() (*string, bool)`

GetMaintMarginOk returns a tuple with the MaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMargin

`func (o *GetAccountV2RespPositionsInner) SetMaintMargin(v string)`

SetMaintMargin sets MaintMargin field to given value.

### HasMaintMargin

`func (o *GetAccountV2RespPositionsInner) HasMaintMargin() bool`

HasMaintMargin returns a boolean if a field has been set.

### GetMaxNotional

`func (o *GetAccountV2RespPositionsInner) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *GetAccountV2RespPositionsInner) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *GetAccountV2RespPositionsInner) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *GetAccountV2RespPositionsInner) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetOpenOrderInitialMargin

`func (o *GetAccountV2RespPositionsInner) GetOpenOrderInitialMargin() string`

GetOpenOrderInitialMargin returns the OpenOrderInitialMargin field if non-nil, zero value otherwise.

### GetOpenOrderInitialMarginOk

`func (o *GetAccountV2RespPositionsInner) GetOpenOrderInitialMarginOk() (*string, bool)`

GetOpenOrderInitialMarginOk returns a tuple with the OpenOrderInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrderInitialMargin

`func (o *GetAccountV2RespPositionsInner) SetOpenOrderInitialMargin(v string)`

SetOpenOrderInitialMargin sets OpenOrderInitialMargin field to given value.

### HasOpenOrderInitialMargin

`func (o *GetAccountV2RespPositionsInner) HasOpenOrderInitialMargin() bool`

HasOpenOrderInitialMargin returns a boolean if a field has been set.

### GetPositionAmt

`func (o *GetAccountV2RespPositionsInner) GetPositionAmt() string`

GetPositionAmt returns the PositionAmt field if non-nil, zero value otherwise.

### GetPositionAmtOk

`func (o *GetAccountV2RespPositionsInner) GetPositionAmtOk() (*string, bool)`

GetPositionAmtOk returns a tuple with the PositionAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAmt

`func (o *GetAccountV2RespPositionsInner) SetPositionAmt(v string)`

SetPositionAmt sets PositionAmt field to given value.

### HasPositionAmt

`func (o *GetAccountV2RespPositionsInner) HasPositionAmt() bool`

HasPositionAmt returns a boolean if a field has been set.

### GetPositionInitialMargin

`func (o *GetAccountV2RespPositionsInner) GetPositionInitialMargin() string`

GetPositionInitialMargin returns the PositionInitialMargin field if non-nil, zero value otherwise.

### GetPositionInitialMarginOk

`func (o *GetAccountV2RespPositionsInner) GetPositionInitialMarginOk() (*string, bool)`

GetPositionInitialMarginOk returns a tuple with the PositionInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionInitialMargin

`func (o *GetAccountV2RespPositionsInner) SetPositionInitialMargin(v string)`

SetPositionInitialMargin sets PositionInitialMargin field to given value.

### HasPositionInitialMargin

`func (o *GetAccountV2RespPositionsInner) HasPositionInitialMargin() bool`

HasPositionInitialMargin returns a boolean if a field has been set.

### GetPositionSide

`func (o *GetAccountV2RespPositionsInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *GetAccountV2RespPositionsInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *GetAccountV2RespPositionsInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *GetAccountV2RespPositionsInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAccountV2RespPositionsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAccountV2RespPositionsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAccountV2RespPositionsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAccountV2RespPositionsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnrealizedProfit

`func (o *GetAccountV2RespPositionsInner) GetUnrealizedProfit() string`

GetUnrealizedProfit returns the UnrealizedProfit field if non-nil, zero value otherwise.

### GetUnrealizedProfitOk

`func (o *GetAccountV2RespPositionsInner) GetUnrealizedProfitOk() (*string, bool)`

GetUnrealizedProfitOk returns a tuple with the UnrealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedProfit

`func (o *GetAccountV2RespPositionsInner) SetUnrealizedProfit(v string)`

SetUnrealizedProfit sets UnrealizedProfit field to given value.

### HasUnrealizedProfit

`func (o *GetAccountV2RespPositionsInner) HasUnrealizedProfit() bool`

HasUnrealizedProfit returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAccountV2RespPositionsInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAccountV2RespPositionsInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAccountV2RespPositionsInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAccountV2RespPositionsInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


