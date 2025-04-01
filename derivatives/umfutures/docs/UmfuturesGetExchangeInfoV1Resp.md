# UmfuturesGetExchangeInfoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]UmfuturesGetExchangeInfoV1RespAssetsInner**](UmfuturesGetExchangeInfoV1RespAssetsInner.md) |  | [optional] 
**ExchangeFilters** | Pointer to **[]string** |  | [optional] 
**RateLimits** | Pointer to [**[]UmfuturesGetExchangeInfoV1RespRateLimitsInner**](UmfuturesGetExchangeInfoV1RespRateLimitsInner.md) |  | [optional] 
**ServerTime** | Pointer to **int32** |  | [optional] 
**Symbols** | Pointer to [**[]UmfuturesGetExchangeInfoV1RespSymbolsInner**](UmfuturesGetExchangeInfoV1RespSymbolsInner.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewUmfuturesGetExchangeInfoV1Resp

`func NewUmfuturesGetExchangeInfoV1Resp() *UmfuturesGetExchangeInfoV1Resp`

NewUmfuturesGetExchangeInfoV1Resp instantiates a new UmfuturesGetExchangeInfoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetExchangeInfoV1RespWithDefaults

`func NewUmfuturesGetExchangeInfoV1RespWithDefaults() *UmfuturesGetExchangeInfoV1Resp`

NewUmfuturesGetExchangeInfoV1RespWithDefaults instantiates a new UmfuturesGetExchangeInfoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *UmfuturesGetExchangeInfoV1Resp) GetAssets() []UmfuturesGetExchangeInfoV1RespAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetAssetsOk() (*[]UmfuturesGetExchangeInfoV1RespAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *UmfuturesGetExchangeInfoV1Resp) SetAssets(v []UmfuturesGetExchangeInfoV1RespAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *UmfuturesGetExchangeInfoV1Resp) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetExchangeFilters

`func (o *UmfuturesGetExchangeInfoV1Resp) GetExchangeFilters() []string`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetExchangeFiltersOk() (*[]string, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *UmfuturesGetExchangeInfoV1Resp) SetExchangeFilters(v []string)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *UmfuturesGetExchangeInfoV1Resp) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *UmfuturesGetExchangeInfoV1Resp) GetRateLimits() []UmfuturesGetExchangeInfoV1RespRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetRateLimitsOk() (*[]UmfuturesGetExchangeInfoV1RespRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *UmfuturesGetExchangeInfoV1Resp) SetRateLimits(v []UmfuturesGetExchangeInfoV1RespRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *UmfuturesGetExchangeInfoV1Resp) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *UmfuturesGetExchangeInfoV1Resp) GetServerTime() int32`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetServerTimeOk() (*int32, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *UmfuturesGetExchangeInfoV1Resp) SetServerTime(v int32)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *UmfuturesGetExchangeInfoV1Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetSymbols

`func (o *UmfuturesGetExchangeInfoV1Resp) GetSymbols() []UmfuturesGetExchangeInfoV1RespSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetSymbolsOk() (*[]UmfuturesGetExchangeInfoV1RespSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *UmfuturesGetExchangeInfoV1Resp) SetSymbols(v []UmfuturesGetExchangeInfoV1RespSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *UmfuturesGetExchangeInfoV1Resp) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTimezone

`func (o *UmfuturesGetExchangeInfoV1Resp) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UmfuturesGetExchangeInfoV1Resp) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UmfuturesGetExchangeInfoV1Resp) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UmfuturesGetExchangeInfoV1Resp) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


