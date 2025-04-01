# CmfuturesGetExchangeInfoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeFilters** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RateLimits** | Pointer to [**[]CmfuturesGetExchangeInfoV1RespRateLimitsInner**](CmfuturesGetExchangeInfoV1RespRateLimitsInner.md) |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**Symbols** | Pointer to [**[]CmfuturesGetExchangeInfoV1RespSymbolsInner**](CmfuturesGetExchangeInfoV1RespSymbolsInner.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewCmfuturesGetExchangeInfoV1Resp

`func NewCmfuturesGetExchangeInfoV1Resp() *CmfuturesGetExchangeInfoV1Resp`

NewCmfuturesGetExchangeInfoV1Resp instantiates a new CmfuturesGetExchangeInfoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmfuturesGetExchangeInfoV1RespWithDefaults

`func NewCmfuturesGetExchangeInfoV1RespWithDefaults() *CmfuturesGetExchangeInfoV1Resp`

NewCmfuturesGetExchangeInfoV1RespWithDefaults instantiates a new CmfuturesGetExchangeInfoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeFilters

`func (o *CmfuturesGetExchangeInfoV1Resp) GetExchangeFilters() []map[string]interface{}`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *CmfuturesGetExchangeInfoV1Resp) GetExchangeFiltersOk() (*[]map[string]interface{}, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *CmfuturesGetExchangeInfoV1Resp) SetExchangeFilters(v []map[string]interface{})`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *CmfuturesGetExchangeInfoV1Resp) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *CmfuturesGetExchangeInfoV1Resp) GetRateLimits() []CmfuturesGetExchangeInfoV1RespRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *CmfuturesGetExchangeInfoV1Resp) GetRateLimitsOk() (*[]CmfuturesGetExchangeInfoV1RespRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *CmfuturesGetExchangeInfoV1Resp) SetRateLimits(v []CmfuturesGetExchangeInfoV1RespRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *CmfuturesGetExchangeInfoV1Resp) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *CmfuturesGetExchangeInfoV1Resp) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *CmfuturesGetExchangeInfoV1Resp) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *CmfuturesGetExchangeInfoV1Resp) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *CmfuturesGetExchangeInfoV1Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetSymbols

`func (o *CmfuturesGetExchangeInfoV1Resp) GetSymbols() []CmfuturesGetExchangeInfoV1RespSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *CmfuturesGetExchangeInfoV1Resp) GetSymbolsOk() (*[]CmfuturesGetExchangeInfoV1RespSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *CmfuturesGetExchangeInfoV1Resp) SetSymbols(v []CmfuturesGetExchangeInfoV1RespSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *CmfuturesGetExchangeInfoV1Resp) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTimezone

`func (o *CmfuturesGetExchangeInfoV1Resp) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CmfuturesGetExchangeInfoV1Resp) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CmfuturesGetExchangeInfoV1Resp) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CmfuturesGetExchangeInfoV1Resp) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


