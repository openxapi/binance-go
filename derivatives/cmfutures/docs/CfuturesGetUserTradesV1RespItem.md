# CfuturesGetUserTradesV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseQty** | Pointer to **string** |  | [optional] 
**Buyer** | Pointer to **bool** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Maker** | Pointer to **bool** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**RealizedPnl** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewCfuturesGetUserTradesV1RespItem

`func NewCfuturesGetUserTradesV1RespItem() *CfuturesGetUserTradesV1RespItem`

NewCfuturesGetUserTradesV1RespItem instantiates a new CfuturesGetUserTradesV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCfuturesGetUserTradesV1RespItemWithDefaults

`func NewCfuturesGetUserTradesV1RespItemWithDefaults() *CfuturesGetUserTradesV1RespItem`

NewCfuturesGetUserTradesV1RespItemWithDefaults instantiates a new CfuturesGetUserTradesV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseQty

`func (o *CfuturesGetUserTradesV1RespItem) GetBaseQty() string`

GetBaseQty returns the BaseQty field if non-nil, zero value otherwise.

### GetBaseQtyOk

`func (o *CfuturesGetUserTradesV1RespItem) GetBaseQtyOk() (*string, bool)`

GetBaseQtyOk returns a tuple with the BaseQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseQty

`func (o *CfuturesGetUserTradesV1RespItem) SetBaseQty(v string)`

SetBaseQty sets BaseQty field to given value.

### HasBaseQty

`func (o *CfuturesGetUserTradesV1RespItem) HasBaseQty() bool`

HasBaseQty returns a boolean if a field has been set.

### GetBuyer

`func (o *CfuturesGetUserTradesV1RespItem) GetBuyer() bool`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *CfuturesGetUserTradesV1RespItem) GetBuyerOk() (*bool, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *CfuturesGetUserTradesV1RespItem) SetBuyer(v bool)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *CfuturesGetUserTradesV1RespItem) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCommission

`func (o *CfuturesGetUserTradesV1RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *CfuturesGetUserTradesV1RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *CfuturesGetUserTradesV1RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *CfuturesGetUserTradesV1RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *CfuturesGetUserTradesV1RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *CfuturesGetUserTradesV1RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *CfuturesGetUserTradesV1RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *CfuturesGetUserTradesV1RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *CfuturesGetUserTradesV1RespItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CfuturesGetUserTradesV1RespItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CfuturesGetUserTradesV1RespItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CfuturesGetUserTradesV1RespItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaker

`func (o *CfuturesGetUserTradesV1RespItem) GetMaker() bool`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *CfuturesGetUserTradesV1RespItem) GetMakerOk() (*bool, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *CfuturesGetUserTradesV1RespItem) SetMaker(v bool)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *CfuturesGetUserTradesV1RespItem) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetMarginAsset

`func (o *CfuturesGetUserTradesV1RespItem) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *CfuturesGetUserTradesV1RespItem) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *CfuturesGetUserTradesV1RespItem) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *CfuturesGetUserTradesV1RespItem) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetOrderId

`func (o *CfuturesGetUserTradesV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CfuturesGetUserTradesV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CfuturesGetUserTradesV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CfuturesGetUserTradesV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPair

`func (o *CfuturesGetUserTradesV1RespItem) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *CfuturesGetUserTradesV1RespItem) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *CfuturesGetUserTradesV1RespItem) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *CfuturesGetUserTradesV1RespItem) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetPositionSide

`func (o *CfuturesGetUserTradesV1RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *CfuturesGetUserTradesV1RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *CfuturesGetUserTradesV1RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *CfuturesGetUserTradesV1RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *CfuturesGetUserTradesV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CfuturesGetUserTradesV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CfuturesGetUserTradesV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CfuturesGetUserTradesV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *CfuturesGetUserTradesV1RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *CfuturesGetUserTradesV1RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *CfuturesGetUserTradesV1RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *CfuturesGetUserTradesV1RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *CfuturesGetUserTradesV1RespItem) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *CfuturesGetUserTradesV1RespItem) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *CfuturesGetUserTradesV1RespItem) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *CfuturesGetUserTradesV1RespItem) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSide

`func (o *CfuturesGetUserTradesV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CfuturesGetUserTradesV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CfuturesGetUserTradesV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CfuturesGetUserTradesV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSymbol

`func (o *CfuturesGetUserTradesV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CfuturesGetUserTradesV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CfuturesGetUserTradesV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CfuturesGetUserTradesV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *CfuturesGetUserTradesV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CfuturesGetUserTradesV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CfuturesGetUserTradesV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *CfuturesGetUserTradesV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


