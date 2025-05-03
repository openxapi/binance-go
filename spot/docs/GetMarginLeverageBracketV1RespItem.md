# GetMarginLeverageBracketV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetNames** | Pointer to **[]string** |  | [optional] 
**Brackets** | Pointer to [**[]GetMarginLeverageBracketV1RespItemBracketsInner**](GetMarginLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMarginLeverageBracketV1RespItem

`func NewGetMarginLeverageBracketV1RespItem() *GetMarginLeverageBracketV1RespItem`

NewGetMarginLeverageBracketV1RespItem instantiates a new GetMarginLeverageBracketV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginLeverageBracketV1RespItemWithDefaults

`func NewGetMarginLeverageBracketV1RespItemWithDefaults() *GetMarginLeverageBracketV1RespItem`

NewGetMarginLeverageBracketV1RespItemWithDefaults instantiates a new GetMarginLeverageBracketV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetNames

`func (o *GetMarginLeverageBracketV1RespItem) GetAssetNames() []string`

GetAssetNames returns the AssetNames field if non-nil, zero value otherwise.

### GetAssetNamesOk

`func (o *GetMarginLeverageBracketV1RespItem) GetAssetNamesOk() (*[]string, bool)`

GetAssetNamesOk returns a tuple with the AssetNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetNames

`func (o *GetMarginLeverageBracketV1RespItem) SetAssetNames(v []string)`

SetAssetNames sets AssetNames field to given value.

### HasAssetNames

`func (o *GetMarginLeverageBracketV1RespItem) HasAssetNames() bool`

HasAssetNames returns a boolean if a field has been set.

### GetBrackets

`func (o *GetMarginLeverageBracketV1RespItem) GetBrackets() []GetMarginLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *GetMarginLeverageBracketV1RespItem) GetBracketsOk() (*[]GetMarginLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *GetMarginLeverageBracketV1RespItem) SetBrackets(v []GetMarginLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *GetMarginLeverageBracketV1RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetRank

`func (o *GetMarginLeverageBracketV1RespItem) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *GetMarginLeverageBracketV1RespItem) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *GetMarginLeverageBracketV1RespItem) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *GetMarginLeverageBracketV1RespItem) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


