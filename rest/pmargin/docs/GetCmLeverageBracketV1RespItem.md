# GetCmLeverageBracketV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]GetCmLeverageBracketV1RespItemBracketsInner**](GetCmLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCmLeverageBracketV1RespItem

`func NewGetCmLeverageBracketV1RespItem() *GetCmLeverageBracketV1RespItem`

NewGetCmLeverageBracketV1RespItem instantiates a new GetCmLeverageBracketV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCmLeverageBracketV1RespItemWithDefaults

`func NewGetCmLeverageBracketV1RespItemWithDefaults() *GetCmLeverageBracketV1RespItem`

NewGetCmLeverageBracketV1RespItemWithDefaults instantiates a new GetCmLeverageBracketV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *GetCmLeverageBracketV1RespItem) GetBrackets() []GetCmLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *GetCmLeverageBracketV1RespItem) GetBracketsOk() (*[]GetCmLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *GetCmLeverageBracketV1RespItem) SetBrackets(v []GetCmLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *GetCmLeverageBracketV1RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetSymbol

`func (o *GetCmLeverageBracketV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetCmLeverageBracketV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetCmLeverageBracketV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetCmLeverageBracketV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


