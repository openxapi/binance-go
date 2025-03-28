# SpotGetMyTradesV3RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**IsBestMatch** | Pointer to **bool** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**QuoteQty** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpotGetMyTradesV3RespItem

`func NewSpotGetMyTradesV3RespItem() *SpotGetMyTradesV3RespItem`

NewSpotGetMyTradesV3RespItem instantiates a new SpotGetMyTradesV3RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetMyTradesV3RespItemWithDefaults

`func NewSpotGetMyTradesV3RespItemWithDefaults() *SpotGetMyTradesV3RespItem`

NewSpotGetMyTradesV3RespItemWithDefaults instantiates a new SpotGetMyTradesV3RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *SpotGetMyTradesV3RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *SpotGetMyTradesV3RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *SpotGetMyTradesV3RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *SpotGetMyTradesV3RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *SpotGetMyTradesV3RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *SpotGetMyTradesV3RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *SpotGetMyTradesV3RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *SpotGetMyTradesV3RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *SpotGetMyTradesV3RespItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpotGetMyTradesV3RespItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpotGetMyTradesV3RespItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SpotGetMyTradesV3RespItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBestMatch

`func (o *SpotGetMyTradesV3RespItem) GetIsBestMatch() bool`

GetIsBestMatch returns the IsBestMatch field if non-nil, zero value otherwise.

### GetIsBestMatchOk

`func (o *SpotGetMyTradesV3RespItem) GetIsBestMatchOk() (*bool, bool)`

GetIsBestMatchOk returns a tuple with the IsBestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBestMatch

`func (o *SpotGetMyTradesV3RespItem) SetIsBestMatch(v bool)`

SetIsBestMatch sets IsBestMatch field to given value.

### HasIsBestMatch

`func (o *SpotGetMyTradesV3RespItem) HasIsBestMatch() bool`

HasIsBestMatch returns a boolean if a field has been set.

### GetIsBuyer

`func (o *SpotGetMyTradesV3RespItem) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *SpotGetMyTradesV3RespItem) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *SpotGetMyTradesV3RespItem) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *SpotGetMyTradesV3RespItem) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *SpotGetMyTradesV3RespItem) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *SpotGetMyTradesV3RespItem) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *SpotGetMyTradesV3RespItem) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *SpotGetMyTradesV3RespItem) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *SpotGetMyTradesV3RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *SpotGetMyTradesV3RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *SpotGetMyTradesV3RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *SpotGetMyTradesV3RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *SpotGetMyTradesV3RespItem) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *SpotGetMyTradesV3RespItem) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *SpotGetMyTradesV3RespItem) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *SpotGetMyTradesV3RespItem) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetPrice

`func (o *SpotGetMyTradesV3RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SpotGetMyTradesV3RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SpotGetMyTradesV3RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SpotGetMyTradesV3RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *SpotGetMyTradesV3RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *SpotGetMyTradesV3RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *SpotGetMyTradesV3RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *SpotGetMyTradesV3RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *SpotGetMyTradesV3RespItem) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *SpotGetMyTradesV3RespItem) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *SpotGetMyTradesV3RespItem) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *SpotGetMyTradesV3RespItem) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetMyTradesV3RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetMyTradesV3RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetMyTradesV3RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetMyTradesV3RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *SpotGetMyTradesV3RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SpotGetMyTradesV3RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SpotGetMyTradesV3RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SpotGetMyTradesV3RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


