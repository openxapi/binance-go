# UmfuturesGetLeverageBracketV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]UmfuturesGetLeverageBracketV1RespItemBracketsInner**](UmfuturesGetLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewUmfuturesGetLeverageBracketV1Resp

`func NewUmfuturesGetLeverageBracketV1Resp() *UmfuturesGetLeverageBracketV1Resp`

NewUmfuturesGetLeverageBracketV1Resp instantiates a new UmfuturesGetLeverageBracketV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetLeverageBracketV1RespWithDefaults

`func NewUmfuturesGetLeverageBracketV1RespWithDefaults() *UmfuturesGetLeverageBracketV1Resp`

NewUmfuturesGetLeverageBracketV1RespWithDefaults instantiates a new UmfuturesGetLeverageBracketV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *UmfuturesGetLeverageBracketV1Resp) GetBrackets() []UmfuturesGetLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *UmfuturesGetLeverageBracketV1Resp) GetBracketsOk() (*[]UmfuturesGetLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *UmfuturesGetLeverageBracketV1Resp) SetBrackets(v []UmfuturesGetLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *UmfuturesGetLeverageBracketV1Resp) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *UmfuturesGetLeverageBracketV1Resp) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *UmfuturesGetLeverageBracketV1Resp) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *UmfuturesGetLeverageBracketV1Resp) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *UmfuturesGetLeverageBracketV1Resp) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetSymbol

`func (o *UmfuturesGetLeverageBracketV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesGetLeverageBracketV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesGetLeverageBracketV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UmfuturesGetLeverageBracketV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


