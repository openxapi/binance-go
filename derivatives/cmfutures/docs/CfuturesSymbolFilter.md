# CfuturesSymbolFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**MaxPrice** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**MinPrice** | Pointer to **string** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**MultiplierDecimal** | Pointer to **int32** |  | [optional] 
**MultiplierDown** | Pointer to **string** |  | [optional] 
**MultiplierUp** | Pointer to **string** |  | [optional] 
**StepSize** | Pointer to **string** |  | [optional] 
**TickSize** | Pointer to **string** |  | [optional] 

## Methods

### NewCfuturesSymbolFilter

`func NewCfuturesSymbolFilter() *CfuturesSymbolFilter`

NewCfuturesSymbolFilter instantiates a new CfuturesSymbolFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfuturesSymbolFilterWithDefaults

`func NewCfuturesSymbolFilterWithDefaults() *CfuturesSymbolFilter`

NewCfuturesSymbolFilterWithDefaults instantiates a new CfuturesSymbolFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *CfuturesSymbolFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *CfuturesSymbolFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *CfuturesSymbolFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *CfuturesSymbolFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetLimit

`func (o *CfuturesSymbolFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CfuturesSymbolFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CfuturesSymbolFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CfuturesSymbolFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxPrice

`func (o *CfuturesSymbolFilter) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *CfuturesSymbolFilter) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *CfuturesSymbolFilter) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *CfuturesSymbolFilter) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetMaxQty

`func (o *CfuturesSymbolFilter) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *CfuturesSymbolFilter) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *CfuturesSymbolFilter) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *CfuturesSymbolFilter) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMinPrice

`func (o *CfuturesSymbolFilter) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *CfuturesSymbolFilter) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *CfuturesSymbolFilter) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *CfuturesSymbolFilter) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMinQty

`func (o *CfuturesSymbolFilter) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *CfuturesSymbolFilter) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *CfuturesSymbolFilter) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *CfuturesSymbolFilter) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMultiplierDecimal

`func (o *CfuturesSymbolFilter) GetMultiplierDecimal() int32`

GetMultiplierDecimal returns the MultiplierDecimal field if non-nil, zero value otherwise.

### GetMultiplierDecimalOk

`func (o *CfuturesSymbolFilter) GetMultiplierDecimalOk() (*int32, bool)`

GetMultiplierDecimalOk returns a tuple with the MultiplierDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDecimal

`func (o *CfuturesSymbolFilter) SetMultiplierDecimal(v int32)`

SetMultiplierDecimal sets MultiplierDecimal field to given value.

### HasMultiplierDecimal

`func (o *CfuturesSymbolFilter) HasMultiplierDecimal() bool`

HasMultiplierDecimal returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *CfuturesSymbolFilter) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *CfuturesSymbolFilter) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *CfuturesSymbolFilter) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *CfuturesSymbolFilter) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *CfuturesSymbolFilter) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *CfuturesSymbolFilter) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *CfuturesSymbolFilter) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *CfuturesSymbolFilter) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetStepSize

`func (o *CfuturesSymbolFilter) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *CfuturesSymbolFilter) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *CfuturesSymbolFilter) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *CfuturesSymbolFilter) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetTickSize

`func (o *CfuturesSymbolFilter) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *CfuturesSymbolFilter) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *CfuturesSymbolFilter) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *CfuturesSymbolFilter) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


