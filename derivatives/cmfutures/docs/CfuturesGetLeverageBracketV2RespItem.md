# CfuturesGetLeverageBracketV2RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brackets** | Pointer to [**[]CfuturesGetLeverageBracketV1RespItemBracketsInner**](CfuturesGetLeverageBracketV1RespItemBracketsInner.md) |  | [optional] 
**NotionalCoef** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewCfuturesGetLeverageBracketV2RespItem

`func NewCfuturesGetLeverageBracketV2RespItem() *CfuturesGetLeverageBracketV2RespItem`

NewCfuturesGetLeverageBracketV2RespItem instantiates a new CfuturesGetLeverageBracketV2RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfuturesGetLeverageBracketV2RespItemWithDefaults

`func NewCfuturesGetLeverageBracketV2RespItemWithDefaults() *CfuturesGetLeverageBracketV2RespItem`

NewCfuturesGetLeverageBracketV2RespItemWithDefaults instantiates a new CfuturesGetLeverageBracketV2RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrackets

`func (o *CfuturesGetLeverageBracketV2RespItem) GetBrackets() []CfuturesGetLeverageBracketV1RespItemBracketsInner`

GetBrackets returns the Brackets field if non-nil, zero value otherwise.

### GetBracketsOk

`func (o *CfuturesGetLeverageBracketV2RespItem) GetBracketsOk() (*[]CfuturesGetLeverageBracketV1RespItemBracketsInner, bool)`

GetBracketsOk returns a tuple with the Brackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrackets

`func (o *CfuturesGetLeverageBracketV2RespItem) SetBrackets(v []CfuturesGetLeverageBracketV1RespItemBracketsInner)`

SetBrackets sets Brackets field to given value.

### HasBrackets

`func (o *CfuturesGetLeverageBracketV2RespItem) HasBrackets() bool`

HasBrackets returns a boolean if a field has been set.

### GetNotionalCoef

`func (o *CfuturesGetLeverageBracketV2RespItem) GetNotionalCoef() float32`

GetNotionalCoef returns the NotionalCoef field if non-nil, zero value otherwise.

### GetNotionalCoefOk

`func (o *CfuturesGetLeverageBracketV2RespItem) GetNotionalCoefOk() (*float32, bool)`

GetNotionalCoefOk returns a tuple with the NotionalCoef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCoef

`func (o *CfuturesGetLeverageBracketV2RespItem) SetNotionalCoef(v float32)`

SetNotionalCoef sets NotionalCoef field to given value.

### HasNotionalCoef

`func (o *CfuturesGetLeverageBracketV2RespItem) HasNotionalCoef() bool`

HasNotionalCoef returns a boolean if a field has been set.

### GetSymbol

`func (o *CfuturesGetLeverageBracketV2RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CfuturesGetLeverageBracketV2RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CfuturesGetLeverageBracketV2RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CfuturesGetLeverageBracketV2RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


