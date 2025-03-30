# PmarginGetCmOrderAmendmentV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amendment** | Pointer to [**PmarginGetCmOrderAmendmentV1RespItemAmendment**](PmarginGetCmOrderAmendmentV1RespItemAmendment.md) |  | [optional] 
**AmendmentId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewPmarginGetCmOrderAmendmentV1RespItem

`func NewPmarginGetCmOrderAmendmentV1RespItem() *PmarginGetCmOrderAmendmentV1RespItem`

NewPmarginGetCmOrderAmendmentV1RespItem instantiates a new PmarginGetCmOrderAmendmentV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginGetCmOrderAmendmentV1RespItemWithDefaults

`func NewPmarginGetCmOrderAmendmentV1RespItemWithDefaults() *PmarginGetCmOrderAmendmentV1RespItem`

NewPmarginGetCmOrderAmendmentV1RespItemWithDefaults instantiates a new PmarginGetCmOrderAmendmentV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendment

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetAmendment() PmarginGetCmOrderAmendmentV1RespItemAmendment`

GetAmendment returns the Amendment field if non-nil, zero value otherwise.

### GetAmendmentOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetAmendmentOk() (*PmarginGetCmOrderAmendmentV1RespItemAmendment, bool)`

GetAmendmentOk returns a tuple with the Amendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendment

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetAmendment(v PmarginGetCmOrderAmendmentV1RespItemAmendment)`

SetAmendment sets Amendment field to given value.

### HasAmendment

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasAmendment() bool`

HasAmendment returns a boolean if a field has been set.

### GetAmendmentId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetAmendmentId() int64`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetAmendmentIdOk() (*int64, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetAmendmentId(v int64)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPair

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PmarginGetCmOrderAmendmentV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PmarginGetCmOrderAmendmentV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *PmarginGetCmOrderAmendmentV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


