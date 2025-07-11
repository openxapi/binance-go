# GetAlgoFuturesSubOrdersV1RespSubOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **int64** |  | [optional] 
**AvgPrice** | Pointer to **string** |  | [optional] 
**BookTime** | Pointer to **int64** |  | [optional] 
**ExecutedAmt** | Pointer to **string** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**FeeAmt** | Pointer to **string** |  | [optional] 
**FeeAsset** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**SubId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAlgoFuturesSubOrdersV1RespSubOrdersInner

`func NewGetAlgoFuturesSubOrdersV1RespSubOrdersInner() *GetAlgoFuturesSubOrdersV1RespSubOrdersInner`

NewGetAlgoFuturesSubOrdersV1RespSubOrdersInner instantiates a new GetAlgoFuturesSubOrdersV1RespSubOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlgoFuturesSubOrdersV1RespSubOrdersInnerWithDefaults

`func NewGetAlgoFuturesSubOrdersV1RespSubOrdersInnerWithDefaults() *GetAlgoFuturesSubOrdersV1RespSubOrdersInner`

NewGetAlgoFuturesSubOrdersV1RespSubOrdersInnerWithDefaults instantiates a new GetAlgoFuturesSubOrdersV1RespSubOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAvgPrice

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetBookTime

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetBookTime() int64`

GetBookTime returns the BookTime field if non-nil, zero value otherwise.

### GetBookTimeOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetBookTimeOk() (*int64, bool)`

GetBookTimeOk returns a tuple with the BookTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookTime

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetBookTime(v int64)`

SetBookTime sets BookTime field to given value.

### HasBookTime

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasBookTime() bool`

HasBookTime returns a boolean if a field has been set.

### GetExecutedAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetExecutedAmt() string`

GetExecutedAmt returns the ExecutedAmt field if non-nil, zero value otherwise.

### GetExecutedAmtOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetExecutedAmtOk() (*string, bool)`

GetExecutedAmtOk returns a tuple with the ExecutedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetExecutedAmt(v string)`

SetExecutedAmt sets ExecutedAmt field to given value.

### HasExecutedAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasExecutedAmt() bool`

HasExecutedAmt returns a boolean if a field has been set.

### GetExecutedQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFeeAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetFeeAmt() string`

GetFeeAmt returns the FeeAmt field if non-nil, zero value otherwise.

### GetFeeAmtOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetFeeAmtOk() (*string, bool)`

GetFeeAmtOk returns a tuple with the FeeAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetFeeAmt(v string)`

SetFeeAmt sets FeeAmt field to given value.

### HasFeeAmt

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasFeeAmt() bool`

HasFeeAmt returns a boolean if a field has been set.

### GetFeeAsset

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetFeeAsset() string`

GetFeeAsset returns the FeeAsset field if non-nil, zero value otherwise.

### GetFeeAssetOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetFeeAssetOk() (*string, bool)`

GetFeeAssetOk returns a tuple with the FeeAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAsset

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetFeeAsset(v string)`

SetFeeAsset sets FeeAsset field to given value.

### HasFeeAsset

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasFeeAsset() bool`

HasFeeAsset returns a boolean if a field has been set.

### GetOrderId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrigQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetSide

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSubId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSubId() int64`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSubIdOk() (*int64, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetSubId(v int64)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *GetAlgoFuturesSubOrdersV1RespSubOrdersInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


