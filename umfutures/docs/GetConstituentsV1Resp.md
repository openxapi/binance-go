# GetConstituentsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constituents** | Pointer to [**[]GetConstituentsV1RespConstituentsInner**](GetConstituentsV1RespConstituentsInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetConstituentsV1Resp

`func NewGetConstituentsV1Resp() *GetConstituentsV1Resp`

NewGetConstituentsV1Resp instantiates a new GetConstituentsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConstituentsV1RespWithDefaults

`func NewGetConstituentsV1RespWithDefaults() *GetConstituentsV1Resp`

NewGetConstituentsV1RespWithDefaults instantiates a new GetConstituentsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstituents

`func (o *GetConstituentsV1Resp) GetConstituents() []GetConstituentsV1RespConstituentsInner`

GetConstituents returns the Constituents field if non-nil, zero value otherwise.

### GetConstituentsOk

`func (o *GetConstituentsV1Resp) GetConstituentsOk() (*[]GetConstituentsV1RespConstituentsInner, bool)`

GetConstituentsOk returns a tuple with the Constituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituents

`func (o *GetConstituentsV1Resp) SetConstituents(v []GetConstituentsV1RespConstituentsInner)`

SetConstituents sets Constituents field to given value.

### HasConstituents

`func (o *GetConstituentsV1Resp) HasConstituents() bool`

HasConstituents returns a boolean if a field has been set.

### GetSymbol

`func (o *GetConstituentsV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetConstituentsV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetConstituentsV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetConstituentsV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *GetConstituentsV1Resp) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetConstituentsV1Resp) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetConstituentsV1Resp) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetConstituentsV1Resp) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


