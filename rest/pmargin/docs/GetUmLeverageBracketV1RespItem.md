# GetUmLeverageBracketV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]GetUmLeverageBracketV1RespItemBracketsInner**](GetUmLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**NotionalCoef** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUmLeverageBracketV1RespItem

`func NewGetUmLeverageBracketV1RespItem() *GetUmLeverageBracketV1RespItem`

NewGetUmLeverageBracketV1RespItem instantiates a new GetUmLeverageBracketV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmLeverageBracketV1RespItemWithDefaults

`func NewGetUmLeverageBracketV1RespItemWithDefaults() *GetUmLeverageBracketV1RespItem`

NewGetUmLeverageBracketV1RespItemWithDefaults instantiates a new GetUmLeverageBracketV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *GetUmLeverageBracketV1RespItem) GetBrackets() []GetUmLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *GetUmLeverageBracketV1RespItem) GetBracketsOk() (*[]GetUmLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *GetUmLeverageBracketV1RespItem) SetBrackets(v []GetUmLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *GetUmLeverageBracketV1RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *GetUmLeverageBracketV1RespItem) GetNotionalCoef() string`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *GetUmLeverageBracketV1RespItem) GetNotionalCoefOk() (*string, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *GetUmLeverageBracketV1RespItem) SetNotionalCoef(v string)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *GetUmLeverageBracketV1RespItem) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetSymbol

`func (o *GetUmLeverageBracketV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetUmLeverageBracketV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetUmLeverageBracketV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetUmLeverageBracketV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


