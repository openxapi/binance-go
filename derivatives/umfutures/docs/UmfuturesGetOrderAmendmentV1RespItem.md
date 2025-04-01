# UmfuturesGetOrderAmendmentV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amendment** | Pointer to [**UmfuturesGetOrderAmendmentV1RespItemAmendment**](UmfuturesGetOrderAmendmentV1RespItemAmendment.md) |  | [optional] 
**AmendmentId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewUmfuturesGetOrderAmendmentV1RespItem

`func NewUmfuturesGetOrderAmendmentV1RespItem() *UmfuturesGetOrderAmendmentV1RespItem`

NewUmfuturesGetOrderAmendmentV1RespItem instantiates a new UmfuturesGetOrderAmendmentV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetOrderAmendmentV1RespItemWithDefaults

`func NewUmfuturesGetOrderAmendmentV1RespItemWithDefaults() *UmfuturesGetOrderAmendmentV1RespItem`

NewUmfuturesGetOrderAmendmentV1RespItemWithDefaults instantiates a new UmfuturesGetOrderAmendmentV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendment

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetAmendment() UmfuturesGetOrderAmendmentV1RespItemAmendment`

GetAmendment returns the Amendment field if non-nil, zero value otherwise.

### GetAmendmentOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetAmendmentOk() (*UmfuturesGetOrderAmendmentV1RespItemAmendment, bool)`

GetAmendmentOk returns a tuple with the Amendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendment

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetAmendment(v UmfuturesGetOrderAmendmentV1RespItemAmendment)`

SetAmendment sets Amendment field to given value.

### HasAmendment

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasAmendment() bool`

HasAmendment returns a boolean if a field has been set.

### GetAmendmentId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetAmendmentId() int64`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetAmendmentIdOk() (*int64, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetAmendmentId(v int64)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPair

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetSymbol

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UmfuturesGetOrderAmendmentV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UmfuturesGetOrderAmendmentV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UmfuturesGetOrderAmendmentV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


