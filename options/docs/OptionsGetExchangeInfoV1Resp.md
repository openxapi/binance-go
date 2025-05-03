# OptionsGetExchangeInfoV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionAssets** | Pointer to [**[]OptionsGetExchangeInfoV1RespOptionAssetsInner**](OptionsGetExchangeInfoV1RespOptionAssetsInner.md) |  | [optional] 
**OptionContracts** | Pointer to [**[]OptionsGetExchangeInfoV1RespOptionContractsInner**](OptionsGetExchangeInfoV1RespOptionContractsInner.md) |  | [optional] 
**OptionSymbols** | Pointer to [**[]OptionsGetExchangeInfoV1RespOptionSymbolsInner**](OptionsGetExchangeInfoV1RespOptionSymbolsInner.md) |  | [optional] 
**RateLimits** | Pointer to [**[]OptionsGetExchangeInfoV1RespRateLimitsInner**](OptionsGetExchangeInfoV1RespRateLimitsInner.md) |  | [optional] 
**ServerTime** | Pointer to **int64** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionsGetExchangeInfoV1Resp

`func NewOptionsGetExchangeInfoV1Resp() *OptionsGetExchangeInfoV1Resp`

NewOptionsGetExchangeInfoV1Resp instantiates a new OptionsGetExchangeInfoV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsGetExchangeInfoV1RespWithDefaults

`func NewOptionsGetExchangeInfoV1RespWithDefaults() *OptionsGetExchangeInfoV1Resp`

NewOptionsGetExchangeInfoV1RespWithDefaults instantiates a new OptionsGetExchangeInfoV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionAssets

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionAssets() []OptionsGetExchangeInfoV1RespOptionAssetsInner`

GetOptionAssets returns the OptionAssets field if non-nil, zero value otherwise.

### GetOptionAssetsOk

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionAssetsOk() (*[]OptionsGetExchangeInfoV1RespOptionAssetsInner, bool)`

GetOptionAssetsOk returns a tuple with the OptionAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionAssets

`func (o *OptionsGetExchangeInfoV1Resp) SetOptionAssets(v []OptionsGetExchangeInfoV1RespOptionAssetsInner)`

SetOptionAssets sets OptionAssets field to given value.

### HasOptionAssets

`func (o *OptionsGetExchangeInfoV1Resp) HasOptionAssets() bool`

HasOptionAssets returns a boolean if a field has been set.

### GetOptionContracts

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionContracts() []OptionsGetExchangeInfoV1RespOptionContractsInner`

GetOptionContracts returns the OptionContracts field if non-nil, zero value otherwise.

### GetOptionContractsOk

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionContractsOk() (*[]OptionsGetExchangeInfoV1RespOptionContractsInner, bool)`

GetOptionContractsOk returns a tuple with the OptionContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionContracts

`func (o *OptionsGetExchangeInfoV1Resp) SetOptionContracts(v []OptionsGetExchangeInfoV1RespOptionContractsInner)`

SetOptionContracts sets OptionContracts field to given value.

### HasOptionContracts

`func (o *OptionsGetExchangeInfoV1Resp) HasOptionContracts() bool`

HasOptionContracts returns a boolean if a field has been set.

### GetOptionSymbols

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionSymbols() []OptionsGetExchangeInfoV1RespOptionSymbolsInner`

GetOptionSymbols returns the OptionSymbols field if non-nil, zero value otherwise.

### GetOptionSymbolsOk

`func (o *OptionsGetExchangeInfoV1Resp) GetOptionSymbolsOk() (*[]OptionsGetExchangeInfoV1RespOptionSymbolsInner, bool)`

GetOptionSymbolsOk returns a tuple with the OptionSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSymbols

`func (o *OptionsGetExchangeInfoV1Resp) SetOptionSymbols(v []OptionsGetExchangeInfoV1RespOptionSymbolsInner)`

SetOptionSymbols sets OptionSymbols field to given value.

### HasOptionSymbols

`func (o *OptionsGetExchangeInfoV1Resp) HasOptionSymbols() bool`

HasOptionSymbols returns a boolean if a field has been set.

### GetRateLimits

`func (o *OptionsGetExchangeInfoV1Resp) GetRateLimits() []OptionsGetExchangeInfoV1RespRateLimitsInner`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *OptionsGetExchangeInfoV1Resp) GetRateLimitsOk() (*[]OptionsGetExchangeInfoV1RespRateLimitsInner, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *OptionsGetExchangeInfoV1Resp) SetRateLimits(v []OptionsGetExchangeInfoV1RespRateLimitsInner)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *OptionsGetExchangeInfoV1Resp) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### GetServerTime

`func (o *OptionsGetExchangeInfoV1Resp) GetServerTime() int64`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *OptionsGetExchangeInfoV1Resp) GetServerTimeOk() (*int64, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *OptionsGetExchangeInfoV1Resp) SetServerTime(v int64)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *OptionsGetExchangeInfoV1Resp) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetTimezone

`func (o *OptionsGetExchangeInfoV1Resp) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *OptionsGetExchangeInfoV1Resp) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *OptionsGetExchangeInfoV1Resp) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *OptionsGetExchangeInfoV1Resp) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


