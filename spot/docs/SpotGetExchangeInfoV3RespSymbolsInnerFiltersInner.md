# SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**MaxPrice** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to **string** |  | [optional] 
**TickSize** | Pointer to **string** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 
**MultiplierDown** | Pointer to **string** |  | [optional] 
**MultiplierUp** | Pointer to **string** |  | [optional] 
**AskMultiplierDown** | Pointer to **string** |  | [optional] 
**AskMultiplierUp** | Pointer to **string** |  | [optional] 
**BidMultiplierDown** | Pointer to **string** |  | [optional] 
**BidMultiplierUp** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**StepSize** | Pointer to **string** |  | [optional] 
**ApplyToMarket** | Pointer to **bool** |  | [optional] 
**MinNotionalValue** | Pointer to **string** |  | [optional] 
**ApplyMaxToMarket** | Pointer to **bool** |  | [optional] 
**ApplyMinToMarket** | Pointer to **bool** |  | [optional] 
**MaxNotionalValue** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**MaxNumOrders** | Pointer to **int32** |  | [optional] 
**MaxNumAlgoOrders** | Pointer to **int32** |  | [optional] 
**MaxNumIcebergOrders** | Pointer to **int32** |  | [optional] 
**MaxPosition** | Pointer to **string** |  | [optional] 
**MaxTrailingAboveDelta** | Pointer to **int32** |  | [optional] 
**MaxTrailingBelowDelta** | Pointer to **int32** |  | [optional] 
**MinTrailingAboveDelta** | Pointer to **int32** |  | [optional] 
**MinTrailingBelowDelta** | Pointer to **int32** |  | [optional] 

## Methods

### NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInner

`func NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInner() *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner`

NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInner instantiates a new SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInnerWithDefaults

`func NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInnerWithDefaults() *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner`

NewSpotGetExchangeInfoV3RespSymbolsInnerFiltersInnerWithDefaults instantiates a new SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetMaxPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMinPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetTickSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetAskMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAskMultiplierDown() string`

GetAskMultiplierDown returns the AskMultiplierDown field if non-nil, zero value otherwise.

### GetAskMultiplierDownOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAskMultiplierDownOk() (*string, bool)`

GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetAskMultiplierDown(v string)`

SetAskMultiplierDown sets AskMultiplierDown field to given value.

### HasAskMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasAskMultiplierDown() bool`

HasAskMultiplierDown returns a boolean if a field has been set.

### GetAskMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAskMultiplierUp() string`

GetAskMultiplierUp returns the AskMultiplierUp field if non-nil, zero value otherwise.

### GetAskMultiplierUpOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetAskMultiplierUpOk() (*string, bool)`

GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetAskMultiplierUp(v string)`

SetAskMultiplierUp sets AskMultiplierUp field to given value.

### HasAskMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasAskMultiplierUp() bool`

HasAskMultiplierUp returns a boolean if a field has been set.

### GetBidMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetBidMultiplierDown() string`

GetBidMultiplierDown returns the BidMultiplierDown field if non-nil, zero value otherwise.

### GetBidMultiplierDownOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetBidMultiplierDownOk() (*string, bool)`

GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetBidMultiplierDown(v string)`

SetBidMultiplierDown sets BidMultiplierDown field to given value.

### HasBidMultiplierDown

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasBidMultiplierDown() bool`

HasBidMultiplierDown returns a boolean if a field has been set.

### GetBidMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetBidMultiplierUp() string`

GetBidMultiplierUp returns the BidMultiplierUp field if non-nil, zero value otherwise.

### GetBidMultiplierUpOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetBidMultiplierUpOk() (*string, bool)`

GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetBidMultiplierUp(v string)`

SetBidMultiplierUp sets BidMultiplierUp field to given value.

### HasBidMultiplierUp

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasBidMultiplierUp() bool`

HasBidMultiplierUp returns a boolean if a field has been set.

### GetMaxQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMinQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetStepSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetApplyToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyToMarket() bool`

GetApplyToMarket returns the ApplyToMarket field if non-nil, zero value otherwise.

### GetApplyToMarketOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyToMarketOk() (*bool, bool)`

GetApplyToMarketOk returns a tuple with the ApplyToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetApplyToMarket(v bool)`

SetApplyToMarket sets ApplyToMarket field to given value.

### HasApplyToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasApplyToMarket() bool`

HasApplyToMarket returns a boolean if a field has been set.

### GetMinNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinNotionalValue() string`

GetMinNotionalValue returns the MinNotionalValue field if non-nil, zero value otherwise.

### GetMinNotionalValueOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinNotionalValueOk() (*string, bool)`

GetMinNotionalValueOk returns a tuple with the MinNotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMinNotionalValue(v string)`

SetMinNotionalValue sets MinNotionalValue field to given value.

### HasMinNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMinNotionalValue() bool`

HasMinNotionalValue returns a boolean if a field has been set.

### GetApplyMaxToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyMaxToMarket() bool`

GetApplyMaxToMarket returns the ApplyMaxToMarket field if non-nil, zero value otherwise.

### GetApplyMaxToMarketOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyMaxToMarketOk() (*bool, bool)`

GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMaxToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetApplyMaxToMarket(v bool)`

SetApplyMaxToMarket sets ApplyMaxToMarket field to given value.

### HasApplyMaxToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasApplyMaxToMarket() bool`

HasApplyMaxToMarket returns a boolean if a field has been set.

### GetApplyMinToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyMinToMarket() bool`

GetApplyMinToMarket returns the ApplyMinToMarket field if non-nil, zero value otherwise.

### GetApplyMinToMarketOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetApplyMinToMarketOk() (*bool, bool)`

GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMinToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetApplyMinToMarket(v bool)`

SetApplyMinToMarket sets ApplyMinToMarket field to given value.

### HasApplyMinToMarket

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasApplyMinToMarket() bool`

HasApplyMinToMarket returns a boolean if a field has been set.

### GetMaxNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNotionalValue() string`

GetMaxNotionalValue returns the MaxNotionalValue field if non-nil, zero value otherwise.

### GetMaxNotionalValueOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNotionalValueOk() (*string, bool)`

GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxNotionalValue(v string)`

SetMaxNotionalValue sets MaxNotionalValue field to given value.

### HasMaxNotionalValue

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxNotionalValue() bool`

HasMaxNotionalValue returns a boolean if a field has been set.

### GetLimit

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumOrders() int32`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumOrdersOk() (*int32, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxNumOrders(v int32)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetMaxNumAlgoOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumAlgoOrders() int32`

GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field if non-nil, zero value otherwise.

### GetMaxNumAlgoOrdersOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumAlgoOrdersOk() (*int32, bool)`

GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumAlgoOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxNumAlgoOrders(v int32)`

SetMaxNumAlgoOrders sets MaxNumAlgoOrders field to given value.

### HasMaxNumAlgoOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxNumAlgoOrders() bool`

HasMaxNumAlgoOrders returns a boolean if a field has been set.

### GetMaxNumIcebergOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumIcebergOrders() int32`

GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field if non-nil, zero value otherwise.

### GetMaxNumIcebergOrdersOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxNumIcebergOrdersOk() (*int32, bool)`

GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumIcebergOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxNumIcebergOrders(v int32)`

SetMaxNumIcebergOrders sets MaxNumIcebergOrders field to given value.

### HasMaxNumIcebergOrders

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxNumIcebergOrders() bool`

HasMaxNumIcebergOrders returns a boolean if a field has been set.

### GetMaxPosition

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxPosition() string`

GetMaxPosition returns the MaxPosition field if non-nil, zero value otherwise.

### GetMaxPositionOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxPositionOk() (*string, bool)`

GetMaxPositionOk returns a tuple with the MaxPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPosition

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxPosition(v string)`

SetMaxPosition sets MaxPosition field to given value.

### HasMaxPosition

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxPosition() bool`

HasMaxPosition returns a boolean if a field has been set.

### GetMaxTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxTrailingAboveDelta() int32`

GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMaxTrailingAboveDeltaOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxTrailingAboveDeltaOk() (*int32, bool)`

GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxTrailingAboveDelta(v int32)`

SetMaxTrailingAboveDelta sets MaxTrailingAboveDelta field to given value.

### HasMaxTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxTrailingAboveDelta() bool`

HasMaxTrailingAboveDelta returns a boolean if a field has been set.

### GetMaxTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxTrailingBelowDelta() int32`

GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMaxTrailingBelowDeltaOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMaxTrailingBelowDeltaOk() (*int32, bool)`

GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMaxTrailingBelowDelta(v int32)`

SetMaxTrailingBelowDelta sets MaxTrailingBelowDelta field to given value.

### HasMaxTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMaxTrailingBelowDelta() bool`

HasMaxTrailingBelowDelta returns a boolean if a field has been set.

### GetMinTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinTrailingAboveDelta() int32`

GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMinTrailingAboveDeltaOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinTrailingAboveDeltaOk() (*int32, bool)`

GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMinTrailingAboveDelta(v int32)`

SetMinTrailingAboveDelta sets MinTrailingAboveDelta field to given value.

### HasMinTrailingAboveDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMinTrailingAboveDelta() bool`

HasMinTrailingAboveDelta returns a boolean if a field has been set.

### GetMinTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinTrailingBelowDelta() int32`

GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMinTrailingBelowDeltaOk

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) GetMinTrailingBelowDeltaOk() (*int32, bool)`

GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) SetMinTrailingBelowDelta(v int32)`

SetMinTrailingBelowDelta sets MinTrailingBelowDelta field to given value.

### HasMinTrailingBelowDelta

`func (o *SpotGetExchangeInfoV3RespSymbolsInnerFiltersInner) HasMinTrailingBelowDelta() bool`

HasMinTrailingBelowDelta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


