# UfuturesGetLeverageBracketV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]UfuturesGetLeverageBracketV1RespItemBracketsInner**](UfuturesGetLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetLeverageBracketV1RespItem

`func NewUfuturesGetLeverageBracketV1RespItem() *UfuturesGetLeverageBracketV1RespItem`

NewUfuturesGetLeverageBracketV1RespItem instantiates a new UfuturesGetLeverageBracketV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetLeverageBracketV1RespItemWithDefaults

`func NewUfuturesGetLeverageBracketV1RespItemWithDefaults() *UfuturesGetLeverageBracketV1RespItem`

NewUfuturesGetLeverageBracketV1RespItemWithDefaults instantiates a new UfuturesGetLeverageBracketV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *UfuturesGetLeverageBracketV1RespItem) GetBrackets() []UfuturesGetLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *UfuturesGetLeverageBracketV1RespItem) GetBracketsOk() (*[]UfuturesGetLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *UfuturesGetLeverageBracketV1RespItem) SetBrackets(v []UfuturesGetLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *UfuturesGetLeverageBracketV1RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *UfuturesGetLeverageBracketV1RespItem) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *UfuturesGetLeverageBracketV1RespItem) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *UfuturesGetLeverageBracketV1RespItem) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *UfuturesGetLeverageBracketV1RespItem) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesGetLeverageBracketV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesGetLeverageBracketV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesGetLeverageBracketV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesGetLeverageBracketV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


