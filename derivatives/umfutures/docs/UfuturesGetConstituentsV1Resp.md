# UfuturesGetConstituentsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constituents** | Pointer to [**[]UfuturesGetConstituentsV1RespConstituentsInner**](UfuturesGetConstituentsV1RespConstituentsInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewUfuturesGetConstituentsV1Resp

`func NewUfuturesGetConstituentsV1Resp() *UfuturesGetConstituentsV1Resp`

NewUfuturesGetConstituentsV1Resp instantiates a new UfuturesGetConstituentsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetConstituentsV1RespWithDefaults

`func NewUfuturesGetConstituentsV1RespWithDefaults() *UfuturesGetConstituentsV1Resp`

NewUfuturesGetConstituentsV1RespWithDefaults instantiates a new UfuturesGetConstituentsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstituents

`func (o *UfuturesGetConstituentsV1Resp) GetConstituents() []UfuturesGetConstituentsV1RespConstituentsInner`

GetConstituents returns the Constituents field if non-nil, zero value otherwise.

### GetConstituentsOk

`func (o *UfuturesGetConstituentsV1Resp) GetConstituentsOk() (*[]UfuturesGetConstituentsV1RespConstituentsInner, bool)`

GetConstituentsOk returns a tuple with the Constituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituents

`func (o *UfuturesGetConstituentsV1Resp) SetConstituents(v []UfuturesGetConstituentsV1RespConstituentsInner)`

SetConstituents sets Constituents field to given value.

### HasConstituents

`func (o *UfuturesGetConstituentsV1Resp) HasConstituents() bool`

HasConstituents returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesGetConstituentsV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesGetConstituentsV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesGetConstituentsV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesGetConstituentsV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *UfuturesGetConstituentsV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UfuturesGetConstituentsV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UfuturesGetConstituentsV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UfuturesGetConstituentsV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


