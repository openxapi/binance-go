# CmfuturesGetConstituentsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constituents** | Pointer to [**[]CmfuturesGetConstituentsV1RespConstituentsInner**](CmfuturesGetConstituentsV1RespConstituentsInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewCmfuturesGetConstituentsV1Resp

`func NewCmfuturesGetConstituentsV1Resp() *CmfuturesGetConstituentsV1Resp`

NewCmfuturesGetConstituentsV1Resp instantiates a new CmfuturesGetConstituentsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmfuturesGetConstituentsV1RespWithDefaults

`func NewCmfuturesGetConstituentsV1RespWithDefaults() *CmfuturesGetConstituentsV1Resp`

NewCmfuturesGetConstituentsV1RespWithDefaults instantiates a new CmfuturesGetConstituentsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstituents

`func (o *CmfuturesGetConstituentsV1Resp) GetConstituents() []CmfuturesGetConstituentsV1RespConstituentsInner`

GetConstituents returns the Constituents field if non-nil, zero value otherwise.

### GetConstituentsOk

`func (o *CmfuturesGetConstituentsV1Resp) GetConstituentsOk() (*[]CmfuturesGetConstituentsV1RespConstituentsInner, bool)`

GetConstituentsOk returns a tuple with the Constituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituents

`func (o *CmfuturesGetConstituentsV1Resp) SetConstituents(v []CmfuturesGetConstituentsV1RespConstituentsInner)`

SetConstituents sets Constituents field to given value.

### HasConstituents

`func (o *CmfuturesGetConstituentsV1Resp) HasConstituents() bool`

HasConstituents returns a boolean if a field has been set.

### GetSymbol

`func (o *CmfuturesGetConstituentsV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CmfuturesGetConstituentsV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CmfuturesGetConstituentsV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CmfuturesGetConstituentsV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *CmfuturesGetConstituentsV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CmfuturesGetConstituentsV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CmfuturesGetConstituentsV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *CmfuturesGetConstituentsV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


