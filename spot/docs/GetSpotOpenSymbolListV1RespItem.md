# GetSpotOpenSymbolListV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenTime** | Pointer to **int64** |  | [optional] 
**Symbols** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetSpotOpenSymbolListV1RespItem

`func NewGetSpotOpenSymbolListV1RespItem() *GetSpotOpenSymbolListV1RespItem`

NewGetSpotOpenSymbolListV1RespItem instantiates a new GetSpotOpenSymbolListV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpotOpenSymbolListV1RespItemWithDefaults

`func NewGetSpotOpenSymbolListV1RespItemWithDefaults() *GetSpotOpenSymbolListV1RespItem`

NewGetSpotOpenSymbolListV1RespItemWithDefaults instantiates a new GetSpotOpenSymbolListV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenTime

`func (o *GetSpotOpenSymbolListV1RespItem) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetSpotOpenSymbolListV1RespItem) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetSpotOpenSymbolListV1RespItem) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetSpotOpenSymbolListV1RespItem) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetSymbols

`func (o *GetSpotOpenSymbolListV1RespItem) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *GetSpotOpenSymbolListV1RespItem) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *GetSpotOpenSymbolListV1RespItem) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *GetSpotOpenSymbolListV1RespItem) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


