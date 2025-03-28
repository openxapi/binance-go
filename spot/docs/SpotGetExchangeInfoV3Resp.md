# SpotGetExchangeInfoV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeFilters** | Pointer to [**[]SpotGetExchangeInfoV3RespExchangeFiltersInner**](SpotGetExchangeInfoV3RespExchangeFiltersInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]SpotRateLimit**](SpotRateLimit.md) |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**Sors** | Pointer to [**[]SpotGetExchangeInfoV3RespSorsInner**](SpotGetExchangeInfoV3RespSorsInner.md) |  | [optional] 
**Symbols** | Pointer to [**[]SpotGetExchangeInfoV3RespSymbolsInner**](SpotGetExchangeInfoV3RespSymbolsInner.md) |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotGetExchangeInfoV3Resp

`func NewSpotGetExchangeInfoV3Resp() *SpotGetExchangeInfoV3Resp`

NewSpotGetExchangeInfoV3Resp instantiates a new SpotGetExchangeInfoV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetExchangeInfoV3RespWithDefaults

`func NewSpotGetExchangeInfoV3RespWithDefaults() *SpotGetExchangeInfoV3Resp`

NewSpotGetExchangeInfoV3RespWithDefaults instantiates a new SpotGetExchangeInfoV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeFilters

`func (o *SpotGetExchangeInfoV3Resp) GetExchangeFilters() []SpotGetExchangeInfoV3RespExchangeFiltersInner`

GetExchangeFilters returns the ExchangeFilters field if non-nil, zero value otherwise.

### GetExchangeFiltersOk

`func (o *SpotGetExchangeInfoV3Resp) GetExchangeFiltersOk() (*[]SpotGetExchangeInfoV3RespExchangeFiltersInner, bool)`

GetExchangeFiltersOk returns a tuple with the ExchangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeFilters

`func (o *SpotGetExchangeInfoV3Resp) SetExchangeFilters(v []SpotGetExchangeInfoV3RespExchangeFiltersInner)`

SetExchangeFilters sets ExchangeFilters field to given value.

### HasExchangeFilters

`func (o *SpotGetExchangeInfoV3Resp) HasExchangeFilters() bool`

HasExchangeFilters returns a boolean if a field has been set.

### GetRateLimits

`func (o *SpotGetExchangeInfoV3Resp) GetRateLimits() []SpotRateLimit`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *SpotGetExchangeInfoV3Resp) GetRateLimitsOk() (*[]SpotRateLimit, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *SpotGetExchangeInfoV3Resp) SetRateLimits(v []SpotRateLimit)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *SpotGetExchangeInfoV3Resp) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *SpotGetExchangeInfoV3Resp) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *SpotGetExchangeInfoV3Resp) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *SpotGetExchangeInfoV3Resp) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *SpotGetExchangeInfoV3Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetSors

`func (o *SpotGetExchangeInfoV3Resp) GetSors() []SpotGetExchangeInfoV3RespSorsInner`

GetSors returns the Sors field if non-nil, zero value otherwise.

### GetSorsOk

`func (o *SpotGetExchangeInfoV3Resp) GetSorsOk() (*[]SpotGetExchangeInfoV3RespSorsInner, bool)`

GetSorsOk returns a tuple with the Sors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSors

`func (o *SpotGetExchangeInfoV3Resp) SetSors(v []SpotGetExchangeInfoV3RespSorsInner)`

SetSors sets Sors field to given value.

### HasSors

`func (o *SpotGetExchangeInfoV3Resp) HasSors() bool`

HasSors returns a boolean if a field has been set.

### GetSymbols

`func (o *SpotGetExchangeInfoV3Resp) GetSymbols() []SpotGetExchangeInfoV3RespSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *SpotGetExchangeInfoV3Resp) GetSymbolsOk() (*[]SpotGetExchangeInfoV3RespSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *SpotGetExchangeInfoV3Resp) SetSymbols(v []SpotGetExchangeInfoV3RespSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *SpotGetExchangeInfoV3Resp) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetTimezone

`func (o *SpotGetExchangeInfoV3Resp) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SpotGetExchangeInfoV3Resp) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SpotGetExchangeInfoV3Resp) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SpotGetExchangeInfoV3Resp) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


