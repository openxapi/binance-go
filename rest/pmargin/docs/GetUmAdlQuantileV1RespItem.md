# GetUmAdlQuantileV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdlQuantile** | Pointer to [**GetCmAdlQuantileV1RespItemAdlQuantile**](GetCmAdlQuantileV1RespItemAdlQuantile.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUmAdlQuantileV1RespItem

`func NewGetUmAdlQuantileV1RespItem() *GetUmAdlQuantileV1RespItem`

NewGetUmAdlQuantileV1RespItem instantiates a new GetUmAdlQuantileV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUmAdlQuantileV1RespItemWithDefaults

`func NewGetUmAdlQuantileV1RespItemWithDefaults() *GetUmAdlQuantileV1RespItem`

NewGetUmAdlQuantileV1RespItemWithDefaults instantiates a new GetUmAdlQuantileV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdlQuantile

`func (o *GetUmAdlQuantileV1RespItem) GetAdlQuantile() GetCmAdlQuantileV1RespItemAdlQuantile`

GetAdlQuantile returns the AdlQuantile field if non-nil, zero value otherwise.

### GetAdlQuantileOk

`func (o *GetUmAdlQuantileV1RespItem) GetAdlQuantileOk() (*GetCmAdlQuantileV1RespItemAdlQuantile, bool)`

GetAdlQuantileOk returns a tuple with the AdlQuantile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdlQuantile

`func (o *GetUmAdlQuantileV1RespItem) SetAdlQuantile(v GetCmAdlQuantileV1RespItemAdlQuantile)`

SetAdlQuantile sets AdlQuantile field to given value.

### HasAdlQuantile

`func (o *GetUmAdlQuantileV1RespItem) HasAdlQuantile() bool`

HasAdlQuantile returns a boolean if a field has been set.

### GetSymbol

`func (o *GetUmAdlQuantileV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetUmAdlQuantileV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetUmAdlQuantileV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetUmAdlQuantileV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


