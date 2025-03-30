# PmarginCreateMarginOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CummulativeQuoteQty** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fills** | Pointer to [**[]PmarginCreateMarginOrderV1RespFillsInner**](PmarginCreateMarginOrderV1RespFillsInner.md) |  | [optional] 
**MarginBuyBorrowAmount** | Pointer to **int32** |  | [optional] 
**MarginBuyBorrowAsset** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPmarginCreateMarginOrderV1Resp

`func NewPmarginCreateMarginOrderV1Resp() *PmarginCreateMarginOrderV1Resp`

NewPmarginCreateMarginOrderV1Resp instantiates a new PmarginCreateMarginOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPmarginCreateMarginOrderV1RespWithDefaults

`func NewPmarginCreateMarginOrderV1RespWithDefaults() *PmarginCreateMarginOrderV1Resp`

NewPmarginCreateMarginOrderV1RespWithDefaults instantiates a new PmarginCreateMarginOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *PmarginCreateMarginOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PmarginCreateMarginOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PmarginCreateMarginOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PmarginCreateMarginOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCummulativeQuoteQty

`func (o *PmarginCreateMarginOrderV1Resp) GetCummulativeQuoteQty() string`

GetCummulativeQuoteQty returns the CummulativeQuoteQty field if non-nil, zero value otherwise.

### GetCummulativeQuoteQtyOk

`func (o *PmarginCreateMarginOrderV1Resp) GetCummulativeQuoteQtyOk() (*string, bool)`

GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCummulativeQuoteQty

`func (o *PmarginCreateMarginOrderV1Resp) SetCummulativeQuoteQty(v string)`

SetCummulativeQuoteQty sets CummulativeQuoteQty field to given value.

### HasCummulativeQuoteQty

`func (o *PmarginCreateMarginOrderV1Resp) HasCummulativeQuoteQty() bool`

HasCummulativeQuoteQty returns a boolean if a field has been set.

### GetExecutedQty

`func (o *PmarginCreateMarginOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *PmarginCreateMarginOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *PmarginCreateMarginOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *PmarginCreateMarginOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFills

`func (o *PmarginCreateMarginOrderV1Resp) GetFills() []PmarginCreateMarginOrderV1RespFillsInner`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *PmarginCreateMarginOrderV1Resp) GetFillsOk() (*[]PmarginCreateMarginOrderV1RespFillsInner, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *PmarginCreateMarginOrderV1Resp) SetFills(v []PmarginCreateMarginOrderV1RespFillsInner)`

SetFills sets Fills field to given value.

### HasFills

`func (o *PmarginCreateMarginOrderV1Resp) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAmount() int32`

GetMarginBuyBorrowAmount returns the MarginBuyBorrowAmount field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAmountOk

`func (o *PmarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAmountOk() (*int32, bool)`

GetMarginBuyBorrowAmountOk returns a tuple with the MarginBuyBorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderV1Resp) SetMarginBuyBorrowAmount(v int32)`

SetMarginBuyBorrowAmount sets MarginBuyBorrowAmount field to given value.

### HasMarginBuyBorrowAmount

`func (o *PmarginCreateMarginOrderV1Resp) HasMarginBuyBorrowAmount() bool`

HasMarginBuyBorrowAmount returns a boolean if a field has been set.

### GetMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAsset() string`

GetMarginBuyBorrowAsset returns the MarginBuyBorrowAsset field if non-nil, zero value otherwise.

### GetMarginBuyBorrowAssetOk

`func (o *PmarginCreateMarginOrderV1Resp) GetMarginBuyBorrowAssetOk() (*string, bool)`

GetMarginBuyBorrowAssetOk returns a tuple with the MarginBuyBorrowAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderV1Resp) SetMarginBuyBorrowAsset(v string)`

SetMarginBuyBorrowAsset sets MarginBuyBorrowAsset field to given value.

### HasMarginBuyBorrowAsset

`func (o *PmarginCreateMarginOrderV1Resp) HasMarginBuyBorrowAsset() bool`

HasMarginBuyBorrowAsset returns a boolean if a field has been set.

### GetOrderId

`func (o *PmarginCreateMarginOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PmarginCreateMarginOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PmarginCreateMarginOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PmarginCreateMarginOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *PmarginCreateMarginOrderV1Resp) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *PmarginCreateMarginOrderV1Resp) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *PmarginCreateMarginOrderV1Resp) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *PmarginCreateMarginOrderV1Resp) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetPrice

`func (o *PmarginCreateMarginOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PmarginCreateMarginOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PmarginCreateMarginOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PmarginCreateMarginOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSide

`func (o *PmarginCreateMarginOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PmarginCreateMarginOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PmarginCreateMarginOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *PmarginCreateMarginOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *PmarginCreateMarginOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PmarginCreateMarginOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PmarginCreateMarginOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PmarginCreateMarginOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *PmarginCreateMarginOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PmarginCreateMarginOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PmarginCreateMarginOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PmarginCreateMarginOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PmarginCreateMarginOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PmarginCreateMarginOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PmarginCreateMarginOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PmarginCreateMarginOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTransactTime

`func (o *PmarginCreateMarginOrderV1Resp) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *PmarginCreateMarginOrderV1Resp) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *PmarginCreateMarginOrderV1Resp) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *PmarginCreateMarginOrderV1Resp) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetType

`func (o *PmarginCreateMarginOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PmarginCreateMarginOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PmarginCreateMarginOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PmarginCreateMarginOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


