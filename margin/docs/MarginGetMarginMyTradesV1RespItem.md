# MarginGetMarginMyTradesV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsIsolated** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarginGetMarginMyTradesV1RespItem

`func NewMarginGetMarginMyTradesV1RespItem() *MarginGetMarginMyTradesV1RespItem`

NewMarginGetMarginMyTradesV1RespItem instantiates a new MarginGetMarginMyTradesV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginGetMarginMyTradesV1RespItemWithDefaults

`func NewMarginGetMarginMyTradesV1RespItemWithDefaults() *MarginGetMarginMyTradesV1RespItem`

NewMarginGetMarginMyTradesV1RespItemWithDefaults instantiates a new MarginGetMarginMyTradesV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *MarginGetMarginMyTradesV1RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MarginGetMarginMyTradesV1RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MarginGetMarginMyTradesV1RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *MarginGetMarginMyTradesV1RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MarginGetMarginMyTradesV1RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *MarginGetMarginMyTradesV1RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *MarginGetMarginMyTradesV1RespItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarginGetMarginMyTradesV1RespItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MarginGetMarginMyTradesV1RespItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *MarginGetMarginMyTradesV1RespItem) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *MarginGetMarginMyTradesV1RespItem) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.

### GetIsBuyer

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MarginGetMarginMyTradesV1RespItem) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *MarginGetMarginMyTradesV1RespItem) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsIsolated

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsIsolated() bool`

GetIsIsolated returns the IsIsolated field if non-nil, zero value otherwise.

### GetIsIsolatedOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsIsolatedOk() (*bool, bool)`

GetIsIsolatedOk returns a tuple with the IsIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsolated

`func (o *MarginGetMarginMyTradesV1RespItem) SetIsIsolated(v bool)`

SetIsIsolated sets IsIsolated field to given value.

### HasIsIsolated

`func (o *MarginGetMarginMyTradesV1RespItem) HasIsIsolated() bool`

HasIsIsolated returns a boolean if a field has been set.

### GetIsMaker

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MarginGetMarginMyTradesV1RespItem) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *MarginGetMarginMyTradesV1RespItem) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *MarginGetMarginMyTradesV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MarginGetMarginMyTradesV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MarginGetMarginMyTradesV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPrice

`func (o *MarginGetMarginMyTradesV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MarginGetMarginMyTradesV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MarginGetMarginMyTradesV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *MarginGetMarginMyTradesV1RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MarginGetMarginMyTradesV1RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *MarginGetMarginMyTradesV1RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetSymbol

`func (o *MarginGetMarginMyTradesV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarginGetMarginMyTradesV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarginGetMarginMyTradesV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *MarginGetMarginMyTradesV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarginGetMarginMyTradesV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarginGetMarginMyTradesV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarginGetMarginMyTradesV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


