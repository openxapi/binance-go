# SpotSymbolFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyMaxToMarket** | Pointer to **bool** |  | [optional] 
**ApplyMinToMarket** | Pointer to **bool** |  | [optional] 
**ApplyToMarket** | Pointer to **bool** |  | [optional] 
**AskMultiplierDown** | Pointer to **string** |  | [optional] 
**AskMultiplierUp** | Pointer to **string** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 
**BidMultiplierDown** | Pointer to **string** |  | [optional] 
**BidMultiplierUp** | Pointer to **string** |  | [optional] 
**FilterType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**MaxNotionalValue** | Pointer to **string** |  | [optional] 
**MaxNumAlgoOrders** | Pointer to **int32** |  | [optional] 
**MaxNumIcebergOrders** | Pointer to **int32** |  | [optional] 
**MaxNumOrders** | Pointer to **int32** |  | [optional] 
**MaxPosition** | Pointer to **string** |  | [optional] 
**MaxPrice** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**MaxTrailingAboveDelta** | Pointer to **int32** |  | [optional] 
**MaxTrailingBelowDelta** | Pointer to **int32** |  | [optional] 
**MinNotionalValue** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to **string** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**MinTrailingAboveDelta** | Pointer to **int32** |  | [optional] 
**MinTrailingBelowDelta** | Pointer to **int32** |  | [optional] 
**MultiplierDown** | Pointer to **string** |  | [optional] 
**MultiplierUp** | Pointer to **string** |  | [optional] 
**StepSize** | Pointer to **string** |  | [optional] 
**TickSize** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotSymbolFilter

`func NewSpotSymbolFilter() *SpotSymbolFilter`

NewSpotSymbolFilter instantiates a new SpotSymbolFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotSymbolFilterWithDefaults

`func NewSpotSymbolFilterWithDefaults() *SpotSymbolFilter`

NewSpotSymbolFilterWithDefaults instantiates a new SpotSymbolFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyMaxToMarket

`func (o *SpotSymbolFilter) GetApplyMaxToMarket() bool`

GetApplyMaxToMarket returns the ApplyMaxToMarket field if non-nil, zero value otherwise.

### GetApplyMaxToMarketOk

`func (o *SpotSymbolFilter) GetApplyMaxToMarketOk() (*bool, bool)`

GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMaxToMarket

`func (o *SpotSymbolFilter) SetApplyMaxToMarket(v bool)`

SetApplyMaxToMarket sets ApplyMaxToMarket field to given value.

### HasApplyMaxToMarket

`func (o *SpotSymbolFilter) HasApplyMaxToMarket() bool`

HasApplyMaxToMarket returns a boolean if a field has been set.

### GetApplyMinToMarket

`func (o *SpotSymbolFilter) GetApplyMinToMarket() bool`

GetApplyMinToMarket returns the ApplyMinToMarket field if non-nil, zero value otherwise.

### GetApplyMinToMarketOk

`func (o *SpotSymbolFilter) GetApplyMinToMarketOk() (*bool, bool)`

GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMinToMarket

`func (o *SpotSymbolFilter) SetApplyMinToMarket(v bool)`

SetApplyMinToMarket sets ApplyMinToMarket field to given value.

### HasApplyMinToMarket

`func (o *SpotSymbolFilter) HasApplyMinToMarket() bool`

HasApplyMinToMarket returns a boolean if a field has been set.

### GetApplyToMarket

`func (o *SpotSymbolFilter) GetApplyToMarket() bool`

GetApplyToMarket returns the ApplyToMarket field if non-nil, zero value otherwise.

### GetApplyToMarketOk

`func (o *SpotSymbolFilter) GetApplyToMarketOk() (*bool, bool)`

GetApplyToMarketOk returns a tuple with the ApplyToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToMarket

`func (o *SpotSymbolFilter) SetApplyToMarket(v bool)`

SetApplyToMarket sets ApplyToMarket field to given value.

### HasApplyToMarket

`func (o *SpotSymbolFilter) HasApplyToMarket() bool`

HasApplyToMarket returns a boolean if a field has been set.

### GetAskMultiplierDown

`func (o *SpotSymbolFilter) GetAskMultiplierDown() string`

GetAskMultiplierDown returns the AskMultiplierDown field if non-nil, zero value otherwise.

### GetAskMultiplierDownOk

`func (o *SpotSymbolFilter) GetAskMultiplierDownOk() (*string, bool)`

GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierDown

`func (o *SpotSymbolFilter) SetAskMultiplierDown(v string)`

SetAskMultiplierDown sets AskMultiplierDown field to given value.

### HasAskMultiplierDown

`func (o *SpotSymbolFilter) HasAskMultiplierDown() bool`

HasAskMultiplierDown returns a boolean if a field has been set.

### GetAskMultiplierUp

`func (o *SpotSymbolFilter) GetAskMultiplierUp() string`

GetAskMultiplierUp returns the AskMultiplierUp field if non-nil, zero value otherwise.

### GetAskMultiplierUpOk

`func (o *SpotSymbolFilter) GetAskMultiplierUpOk() (*string, bool)`

GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierUp

`func (o *SpotSymbolFilter) SetAskMultiplierUp(v string)`

SetAskMultiplierUp sets AskMultiplierUp field to given value.

### HasAskMultiplierUp

`func (o *SpotSymbolFilter) HasAskMultiplierUp() bool`

HasAskMultiplierUp returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *SpotSymbolFilter) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *SpotSymbolFilter) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *SpotSymbolFilter) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *SpotSymbolFilter) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.

### GetBidMultiplierDown

`func (o *SpotSymbolFilter) GetBidMultiplierDown() string`

GetBidMultiplierDown returns the BidMultiplierDown field if non-nil, zero value otherwise.

### GetBidMultiplierDownOk

`func (o *SpotSymbolFilter) GetBidMultiplierDownOk() (*string, bool)`

GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierDown

`func (o *SpotSymbolFilter) SetBidMultiplierDown(v string)`

SetBidMultiplierDown sets BidMultiplierDown field to given value.

### HasBidMultiplierDown

`func (o *SpotSymbolFilter) HasBidMultiplierDown() bool`

HasBidMultiplierDown returns a boolean if a field has been set.

### GetBidMultiplierUp

`func (o *SpotSymbolFilter) GetBidMultiplierUp() string`

GetBidMultiplierUp returns the BidMultiplierUp field if non-nil, zero value otherwise.

### GetBidMultiplierUpOk

`func (o *SpotSymbolFilter) GetBidMultiplierUpOk() (*string, bool)`

GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierUp

`func (o *SpotSymbolFilter) SetBidMultiplierUp(v string)`

SetBidMultiplierUp sets BidMultiplierUp field to given value.

### HasBidMultiplierUp

`func (o *SpotSymbolFilter) HasBidMultiplierUp() bool`

HasBidMultiplierUp returns a boolean if a field has been set.

### GetFilterType

`func (o *SpotSymbolFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *SpotSymbolFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *SpotSymbolFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *SpotSymbolFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetLimit

`func (o *SpotSymbolFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SpotSymbolFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SpotSymbolFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SpotSymbolFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxNotionalValue

`func (o *SpotSymbolFilter) GetMaxNotionalValue() string`

GetMaxNotionalValue returns the MaxNotionalValue field if non-nil, zero value otherwise.

### GetMaxNotionalValueOk

`func (o *SpotSymbolFilter) GetMaxNotionalValueOk() (*string, bool)`

GetMaxNotionalValueOk returns a tuple with the MaxNotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotionalValue

`func (o *SpotSymbolFilter) SetMaxNotionalValue(v string)`

SetMaxNotionalValue sets MaxNotionalValue field to given value.

### HasMaxNotionalValue

`func (o *SpotSymbolFilter) HasMaxNotionalValue() bool`

HasMaxNotionalValue returns a boolean if a field has been set.

### GetMaxNumAlgoOrders

`func (o *SpotSymbolFilter) GetMaxNumAlgoOrders() int32`

GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field if non-nil, zero value otherwise.

### GetMaxNumAlgoOrdersOk

`func (o *SpotSymbolFilter) GetMaxNumAlgoOrdersOk() (*int32, bool)`

GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumAlgoOrders

`func (o *SpotSymbolFilter) SetMaxNumAlgoOrders(v int32)`

SetMaxNumAlgoOrders sets MaxNumAlgoOrders field to given value.

### HasMaxNumAlgoOrders

`func (o *SpotSymbolFilter) HasMaxNumAlgoOrders() bool`

HasMaxNumAlgoOrders returns a boolean if a field has been set.

### GetMaxNumIcebergOrders

`func (o *SpotSymbolFilter) GetMaxNumIcebergOrders() int32`

GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field if non-nil, zero value otherwise.

### GetMaxNumIcebergOrdersOk

`func (o *SpotSymbolFilter) GetMaxNumIcebergOrdersOk() (*int32, bool)`

GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumIcebergOrders

`func (o *SpotSymbolFilter) SetMaxNumIcebergOrders(v int32)`

SetMaxNumIcebergOrders sets MaxNumIcebergOrders field to given value.

### HasMaxNumIcebergOrders

`func (o *SpotSymbolFilter) HasMaxNumIcebergOrders() bool`

HasMaxNumIcebergOrders returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *SpotSymbolFilter) GetMaxNumOrders() int32`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *SpotSymbolFilter) GetMaxNumOrdersOk() (*int32, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *SpotSymbolFilter) SetMaxNumOrders(v int32)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *SpotSymbolFilter) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetMaxPosition

`func (o *SpotSymbolFilter) GetMaxPosition() string`

GetMaxPosition returns the MaxPosition field if non-nil, zero value otherwise.

### GetMaxPositionOk

`func (o *SpotSymbolFilter) GetMaxPositionOk() (*string, bool)`

GetMaxPositionOk returns a tuple with the MaxPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPosition

`func (o *SpotSymbolFilter) SetMaxPosition(v string)`

SetMaxPosition sets MaxPosition field to given value.

### HasMaxPosition

`func (o *SpotSymbolFilter) HasMaxPosition() bool`

HasMaxPosition returns a boolean if a field has been set.

### GetMaxPrice

`func (o *SpotSymbolFilter) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *SpotSymbolFilter) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *SpotSymbolFilter) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *SpotSymbolFilter) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMaxQty

`func (o *SpotSymbolFilter) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *SpotSymbolFilter) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *SpotSymbolFilter) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *SpotSymbolFilter) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMaxTrailingAboveDelta

`func (o *SpotSymbolFilter) GetMaxTrailingAboveDelta() int32`

GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMaxTrailingAboveDeltaOk

`func (o *SpotSymbolFilter) GetMaxTrailingAboveDeltaOk() (*int32, bool)`

GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingAboveDelta

`func (o *SpotSymbolFilter) SetMaxTrailingAboveDelta(v int32)`

SetMaxTrailingAboveDelta sets MaxTrailingAboveDelta field to given value.

### HasMaxTrailingAboveDelta

`func (o *SpotSymbolFilter) HasMaxTrailingAboveDelta() bool`

HasMaxTrailingAboveDelta returns a boolean if a field has been set.

### GetMaxTrailingBelowDelta

`func (o *SpotSymbolFilter) GetMaxTrailingBelowDelta() int32`

GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMaxTrailingBelowDeltaOk

`func (o *SpotSymbolFilter) GetMaxTrailingBelowDeltaOk() (*int32, bool)`

GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingBelowDelta

`func (o *SpotSymbolFilter) SetMaxTrailingBelowDelta(v int32)`

SetMaxTrailingBelowDelta sets MaxTrailingBelowDelta field to given value.

### HasMaxTrailingBelowDelta

`func (o *SpotSymbolFilter) HasMaxTrailingBelowDelta() bool`

HasMaxTrailingBelowDelta returns a boolean if a field has been set.

### GetMinNotionalValue

`func (o *SpotSymbolFilter) GetMinNotionalValue() string`

GetMinNotionalValue returns the MinNotionalValue field if non-nil, zero value otherwise.

### GetMinNotionalValueOk

`func (o *SpotSymbolFilter) GetMinNotionalValueOk() (*string, bool)`

GetMinNotionalValueOk returns a tuple with the MinNotionalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotionalValue

`func (o *SpotSymbolFilter) SetMinNotionalValue(v string)`

SetMinNotionalValue sets MinNotionalValue field to given value.

### HasMinNotionalValue

`func (o *SpotSymbolFilter) HasMinNotionalValue() bool`

HasMinNotionalValue returns a boolean if a field has been set.

### GetMinPrice

`func (o *SpotSymbolFilter) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *SpotSymbolFilter) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *SpotSymbolFilter) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *SpotSymbolFilter) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMinQty

`func (o *SpotSymbolFilter) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *SpotSymbolFilter) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *SpotSymbolFilter) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *SpotSymbolFilter) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMinTrailingAboveDelta

`func (o *SpotSymbolFilter) GetMinTrailingAboveDelta() int32`

GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMinTrailingAboveDeltaOk

`func (o *SpotSymbolFilter) GetMinTrailingAboveDeltaOk() (*int32, bool)`

GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingAboveDelta

`func (o *SpotSymbolFilter) SetMinTrailingAboveDelta(v int32)`

SetMinTrailingAboveDelta sets MinTrailingAboveDelta field to given value.

### HasMinTrailingAboveDelta

`func (o *SpotSymbolFilter) HasMinTrailingAboveDelta() bool`

HasMinTrailingAboveDelta returns a boolean if a field has been set.

### GetMinTrailingBelowDelta

`func (o *SpotSymbolFilter) GetMinTrailingBelowDelta() int32`

GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMinTrailingBelowDeltaOk

`func (o *SpotSymbolFilter) GetMinTrailingBelowDeltaOk() (*int32, bool)`

GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingBelowDelta

`func (o *SpotSymbolFilter) SetMinTrailingBelowDelta(v int32)`

SetMinTrailingBelowDelta sets MinTrailingBelowDelta field to given value.

### HasMinTrailingBelowDelta

`func (o *SpotSymbolFilter) HasMinTrailingBelowDelta() bool`

HasMinTrailingBelowDelta returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *SpotSymbolFilter) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *SpotSymbolFilter) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *SpotSymbolFilter) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *SpotSymbolFilter) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *SpotSymbolFilter) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *SpotSymbolFilter) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *SpotSymbolFilter) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *SpotSymbolFilter) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetStepSize

`func (o *SpotSymbolFilter) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *SpotSymbolFilter) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *SpotSymbolFilter) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *SpotSymbolFilter) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetTickSize

`func (o *SpotSymbolFilter) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *SpotSymbolFilter) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *SpotSymbolFilter) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *SpotSymbolFilter) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


