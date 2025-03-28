# UfuturesGetLeverageBracketV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]UfuturesGetLeverageBracketV1RespItemBracketsInner**](UfuturesGetLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetLeverageBracketV1Resp

`func NewUfuturesGetLeverageBracketV1Resp() *UfuturesGetLeverageBracketV1Resp`

NewUfuturesGetLeverageBracketV1Resp instantiates a new UfuturesGetLeverageBracketV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetLeverageBracketV1RespWithDefaults

`func NewUfuturesGetLeverageBracketV1RespWithDefaults() *UfuturesGetLeverageBracketV1Resp`

NewUfuturesGetLeverageBracketV1RespWithDefaults instantiates a new UfuturesGetLeverageBracketV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *UfuturesGetLeverageBracketV1Resp) GetBrackets() []UfuturesGetLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *UfuturesGetLeverageBracketV1Resp) GetBracketsOk() (*[]UfuturesGetLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *UfuturesGetLeverageBracketV1Resp) SetBrackets(v []UfuturesGetLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *UfuturesGetLeverageBracketV1Resp) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *UfuturesGetLeverageBracketV1Resp) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *UfuturesGetLeverageBracketV1Resp) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *UfuturesGetLeverageBracketV1Resp) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *UfuturesGetLeverageBracketV1Resp) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesGetLeverageBracketV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesGetLeverageBracketV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesGetLeverageBracketV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesGetLeverageBracketV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


