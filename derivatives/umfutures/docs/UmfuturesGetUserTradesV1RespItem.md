# UmfuturesGetUserTradesV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyer** | Pointer to **bool** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Maker** | Pointer to **bool** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**QuoteQty** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewUmfuturesGetUserTradesV1RespItem

`func NewUmfuturesGetUserTradesV1RespItem() *UmfuturesGetUserTradesV1RespItem`

NewUmfuturesGetUserTradesV1RespItem instantiates a new UmfuturesGetUserTradesV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUmfuturesGetUserTradesV1RespItemWithDefaults

`func NewUmfuturesGetUserTradesV1RespItemWithDefaults() *UmfuturesGetUserTradesV1RespItem`

NewUmfuturesGetUserTradesV1RespItemWithDefaults instantiates a new UmfuturesGetUserTradesV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyer

`func (o *UmfuturesGetUserTradesV1RespItem) GetBuyer() bool`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetBuyerOk() (*bool, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *UmfuturesGetUserTradesV1RespItem) SetBuyer(v bool)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *UmfuturesGetUserTradesV1RespItem) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCommission

`func (o *UmfuturesGetUserTradesV1RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *UmfuturesGetUserTradesV1RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *UmfuturesGetUserTradesV1RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *UmfuturesGetUserTradesV1RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *UmfuturesGetUserTradesV1RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *UmfuturesGetUserTradesV1RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *UmfuturesGetUserTradesV1RespItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UmfuturesGetUserTradesV1RespItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UmfuturesGetUserTradesV1RespItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaker

`func (o *UmfuturesGetUserTradesV1RespItem) GetMaker() bool`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetMakerOk() (*bool, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *UmfuturesGetUserTradesV1RespItem) SetMaker(v bool)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *UmfuturesGetUserTradesV1RespItem) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *UmfuturesGetUserTradesV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UmfuturesGetUserTradesV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *UmfuturesGetUserTradesV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPositionSide

`func (o *UmfuturesGetUserTradesV1RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *UmfuturesGetUserTradesV1RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *UmfuturesGetUserTradesV1RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *UmfuturesGetUserTradesV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UmfuturesGetUserTradesV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UmfuturesGetUserTradesV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *UmfuturesGetUserTradesV1RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *UmfuturesGetUserTradesV1RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *UmfuturesGetUserTradesV1RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *UmfuturesGetUserTradesV1RespItem) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *UmfuturesGetUserTradesV1RespItem) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *UmfuturesGetUserTradesV1RespItem) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *UmfuturesGetUserTradesV1RespItem) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *UmfuturesGetUserTradesV1RespItem) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *UmfuturesGetUserTradesV1RespItem) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSide

`func (o *UmfuturesGetUserTradesV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *UmfuturesGetUserTradesV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *UmfuturesGetUserTradesV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSymbol

`func (o *UmfuturesGetUserTradesV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UmfuturesGetUserTradesV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UmfuturesGetUserTradesV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *UmfuturesGetUserTradesV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UmfuturesGetUserTradesV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UmfuturesGetUserTradesV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UmfuturesGetUserTradesV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


