# GetLeverageBracketV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]GetLeverageBracketV1RespItemBracketsInner**](GetLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 

## Methods

### NewGetLeverageBracketV1RespItem

`func NewGetLeverageBracketV1RespItem() *GetLeverageBracketV1RespItem`

NewGetLeverageBracketV1RespItem instantiates a new GetLeverageBracketV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLeverageBracketV1RespItemWithDefaults

`func NewGetLeverageBracketV1RespItemWithDefaults() *GetLeverageBracketV1RespItem`

NewGetLeverageBracketV1RespItemWithDefaults instantiates a new GetLeverageBracketV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *GetLeverageBracketV1RespItem) GetBrackets() []GetLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *GetLeverageBracketV1RespItem) GetBracketsOk() (*[]GetLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *GetLeverageBracketV1RespItem) SetBrackets(v []GetLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *GetLeverageBracketV1RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetPair

`func (o *GetLeverageBracketV1RespItem) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *GetLeverageBracketV1RespItem) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *GetLeverageBracketV1RespItem) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *GetLeverageBracketV1RespItem) HasPair() bool`

HasPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


