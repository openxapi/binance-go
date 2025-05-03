# GetIndexInfoV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAssetList** | Pointer to [**[]GetIndexInfoV1RespItemBaseAssetListInner**](GetIndexInfoV1RespItemBaseAssetListInner.md) |  | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetIndexInfoV1RespItem

`func NewGetIndexInfoV1RespItem() *GetIndexInfoV1RespItem`

NewGetIndexInfoV1RespItem instantiates a new GetIndexInfoV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIndexInfoV1RespItemWithDefaults

`func NewGetIndexInfoV1RespItemWithDefaults() *GetIndexInfoV1RespItem`

NewGetIndexInfoV1RespItemWithDefaults instantiates a new GetIndexInfoV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAssetList

`func (o *GetIndexInfoV1RespItem) GetBaseAssetList() []GetIndexInfoV1RespItemBaseAssetListInner`

GetBaseAssetList returns the BaseAssetList field if non-nil, zero value otherwise.

### GetBaseAssetListOk

`func (o *GetIndexInfoV1RespItem) GetBaseAssetListOk() (*[]GetIndexInfoV1RespItemBaseAssetListInner, bool)`

GetBaseAssetListOk returns a tuple with the BaseAssetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetList

`func (o *GetIndexInfoV1RespItem) SetBaseAssetList(v []GetIndexInfoV1RespItemBaseAssetListInner)`

SetBaseAssetList sets BaseAssetList field to given value.

### HasBaseAssetList

`func (o *GetIndexInfoV1RespItem) HasBaseAssetList() bool`

HasBaseAssetList returns a boolean if a field has been set.

### GetComponent

`func (o *GetIndexInfoV1RespItem) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GetIndexInfoV1RespItem) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GetIndexInfoV1RespItem) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *GetIndexInfoV1RespItem) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetSymbol

`func (o *GetIndexInfoV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetIndexInfoV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetIndexInfoV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetIndexInfoV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *GetIndexInfoV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetIndexInfoV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetIndexInfoV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetIndexInfoV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


