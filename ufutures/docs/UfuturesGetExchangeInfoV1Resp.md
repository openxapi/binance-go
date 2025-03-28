# UfuturesGetExchangeInfoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UfuturesGetExchangeInfoV1RespAssetsInner**](UfuturesGetExchangeInfoV1RespAssetsInner.md) |  | [optional] 
**ExchangeFilters** | Pointer to **[]string** |  | [optional] 
**RateLimits** | Pointer to [**[]UfuturesGetExchangeInfoV1RespRateLimitsInner**](UfuturesGetExchangeInfoV1RespRateLimitsInner.md) |  | [optional] 
**ServerTime** | Pointer to **int32** |  | [optional] 
**Symbols** | Pointer to [**[]UfuturesGetExchangeInfoV1RespSymbolsInner**](UfuturesGetExchangeInfoV1RespSymbolsInner.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetExchangeInfoV1Resp

`func NewUfuturesGetExchangeInfoV1Resp() *UfuturesGetExchangeInfoV1Resp`

NewUfuturesGetExchangeInfoV1Resp instantiates a new UfuturesGetExchangeInfoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetExchangeInfoV1RespWithDefaults

`func NewUfuturesGetExchangeInfoV1RespWithDefaults() *UfuturesGetExchangeInfoV1Resp`

NewUfuturesGetExchangeInfoV1RespWithDefaults instantiates a new UfuturesGetExchangeInfoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UfuturesGetExchangeInfoV1Resp) GetAssets() []UfuturesGetExchangeInfoV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetAssetsOk() (*[]UfuturesGetExchangeInfoV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UfuturesGetExchangeInfoV1Resp) SetAssets(v []UfuturesGetExchangeInfoV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UfuturesGetExchangeInfoV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetExchangeFilters

`func (o *UfuturesGetExchangeInfoV1Resp) GetExchangeFilters() []string`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetExchangeFiltersOk() (*[]string, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *UfuturesGetExchangeInfoV1Resp) SetExchangeFilters(v []string)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *UfuturesGetExchangeInfoV1Resp) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *UfuturesGetExchangeInfoV1Resp) GetRateLimits() []UfuturesGetExchangeInfoV1RespRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetRateLimitsOk() (*[]UfuturesGetExchangeInfoV1RespRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *UfuturesGetExchangeInfoV1Resp) SetRateLimits(v []UfuturesGetExchangeInfoV1RespRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *UfuturesGetExchangeInfoV1Resp) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *UfuturesGetExchangeInfoV1Resp) GetServerTime() int32`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetServerTimeOk() (*int32, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *UfuturesGetExchangeInfoV1Resp) SetServerTime(v int32)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *UfuturesGetExchangeInfoV1Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetSymbols

`func (o *UfuturesGetExchangeInfoV1Resp) GetSymbols() []UfuturesGetExchangeInfoV1RespSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetSymbolsOk() (*[]UfuturesGetExchangeInfoV1RespSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *UfuturesGetExchangeInfoV1Resp) SetSymbols(v []UfuturesGetExchangeInfoV1RespSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *UfuturesGetExchangeInfoV1Resp) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTimezone

`func (o *UfuturesGetExchangeInfoV1Resp) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UfuturesGetExchangeInfoV1Resp) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UfuturesGetExchangeInfoV1Resp) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UfuturesGetExchangeInfoV1Resp) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


