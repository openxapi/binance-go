# GetUserTradesV1RespItem

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

### NewGetUserTradesV1RespItem

`func NewGetUserTradesV1RespItem() *GetUserTradesV1RespItem`

NewGetUserTradesV1RespItem instantiates a new GetUserTradesV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTradesV1RespItemWithDefaults

`func NewGetUserTradesV1RespItemWithDefaults() *GetUserTradesV1RespItem`

NewGetUserTradesV1RespItemWithDefaults instantiates a new GetUserTradesV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyer

`func (o *GetUserTradesV1RespItem) GetBuyer() bool`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *GetUserTradesV1RespItem) GetBuyerOk() (*bool, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *GetUserTradesV1RespItem) SetBuyer(v bool)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *GetUserTradesV1RespItem) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCommission

`func (o *GetUserTradesV1RespItem) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *GetUserTradesV1RespItem) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *GetUserTradesV1RespItem) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *GetUserTradesV1RespItem) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *GetUserTradesV1RespItem) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *GetUserTradesV1RespItem) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *GetUserTradesV1RespItem) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *GetUserTradesV1RespItem) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetId

`func (o *GetUserTradesV1RespItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserTradesV1RespItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserTradesV1RespItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserTradesV1RespItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaker

`func (o *GetUserTradesV1RespItem) GetMaker() bool`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *GetUserTradesV1RespItem) GetMakerOk() (*bool, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *GetUserTradesV1RespItem) SetMaker(v bool)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *GetUserTradesV1RespItem) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetOrderId

`func (o *GetUserTradesV1RespItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetUserTradesV1RespItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetUserTradesV1RespItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetUserTradesV1RespItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPositionSide

`func (o *GetUserTradesV1RespItem) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *GetUserTradesV1RespItem) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *GetUserTradesV1RespItem) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *GetUserTradesV1RespItem) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetPrice

`func (o *GetUserTradesV1RespItem) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetUserTradesV1RespItem) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetUserTradesV1RespItem) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetUserTradesV1RespItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *GetUserTradesV1RespItem) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *GetUserTradesV1RespItem) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *GetUserTradesV1RespItem) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *GetUserTradesV1RespItem) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *GetUserTradesV1RespItem) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *GetUserTradesV1RespItem) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *GetUserTradesV1RespItem) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *GetUserTradesV1RespItem) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *GetUserTradesV1RespItem) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *GetUserTradesV1RespItem) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *GetUserTradesV1RespItem) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *GetUserTradesV1RespItem) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSide

`func (o *GetUserTradesV1RespItem) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetUserTradesV1RespItem) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetUserTradesV1RespItem) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetUserTradesV1RespItem) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSymbol

`func (o *GetUserTradesV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetUserTradesV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetUserTradesV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetUserTradesV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTime

`func (o *GetUserTradesV1RespItem) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetUserTradesV1RespItem) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetUserTradesV1RespItem) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetUserTradesV1RespItem) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


