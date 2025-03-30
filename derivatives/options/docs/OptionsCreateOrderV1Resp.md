# OptionsCreateOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ExecutedQty** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**Mmp** | Pointer to **bool** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**PostOnly** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PriceScale** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**QuantityScale** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionsCreateOrderV1Resp

`func NewOptionsCreateOrderV1Resp() *OptionsCreateOrderV1Resp`

NewOptionsCreateOrderV1Resp instantiates a new OptionsCreateOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsCreateOrderV1RespWithDefaults

`func NewOptionsCreateOrderV1RespWithDefaults() *OptionsCreateOrderV1Resp`

NewOptionsCreateOrderV1RespWithDefaults instantiates a new OptionsCreateOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *OptionsCreateOrderV1Resp) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *OptionsCreateOrderV1Resp) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *OptionsCreateOrderV1Resp) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *OptionsCreateOrderV1Resp) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OptionsCreateOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OptionsCreateOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OptionsCreateOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OptionsCreateOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCreateDate

`func (o *OptionsCreateOrderV1Resp) GetCreateDate() int32`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *OptionsCreateOrderV1Resp) GetCreateDateOk() (*int32, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *OptionsCreateOrderV1Resp) SetCreateDate(v int32)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *OptionsCreateOrderV1Resp) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCreateTime

`func (o *OptionsCreateOrderV1Resp) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OptionsCreateOrderV1Resp) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OptionsCreateOrderV1Resp) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OptionsCreateOrderV1Resp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OptionsCreateOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OptionsCreateOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OptionsCreateOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OptionsCreateOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFee

`func (o *OptionsCreateOrderV1Resp) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OptionsCreateOrderV1Resp) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OptionsCreateOrderV1Resp) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *OptionsCreateOrderV1Resp) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetMmp

`func (o *OptionsCreateOrderV1Resp) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *OptionsCreateOrderV1Resp) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *OptionsCreateOrderV1Resp) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *OptionsCreateOrderV1Resp) HasMmp() bool`

HasMmp returns a boolean if a field has been set.

### GetOptionSide

`func (o *OptionsCreateOrderV1Resp) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *OptionsCreateOrderV1Resp) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *OptionsCreateOrderV1Resp) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *OptionsCreateOrderV1Resp) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetOrderId

`func (o *OptionsCreateOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OptionsCreateOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OptionsCreateOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OptionsCreateOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPostOnly

`func (o *OptionsCreateOrderV1Resp) GetPostOnly() bool`

GetPostOnly returns the PostOnly field if non-nil, zero value otherwise.

### GetPostOnlyOk

`func (o *OptionsCreateOrderV1Resp) GetPostOnlyOk() (*bool, bool)`

GetPostOnlyOk returns a tuple with the PostOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOnly

`func (o *OptionsCreateOrderV1Resp) SetPostOnly(v bool)`

SetPostOnly sets PostOnly field to given value.

### HasPostOnly

`func (o *OptionsCreateOrderV1Resp) HasPostOnly() bool`

HasPostOnly returns a boolean if a field has been set.

### GetPrice

`func (o *OptionsCreateOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OptionsCreateOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OptionsCreateOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OptionsCreateOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceScale

`func (o *OptionsCreateOrderV1Resp) GetPriceScale() int32`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *OptionsCreateOrderV1Resp) GetPriceScaleOk() (*int32, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *OptionsCreateOrderV1Resp) SetPriceScale(v int32)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *OptionsCreateOrderV1Resp) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionsCreateOrderV1Resp) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionsCreateOrderV1Resp) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionsCreateOrderV1Resp) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionsCreateOrderV1Resp) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityScale

`func (o *OptionsCreateOrderV1Resp) GetQuantityScale() int32`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *OptionsCreateOrderV1Resp) GetQuantityScaleOk() (*int32, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *OptionsCreateOrderV1Resp) SetQuantityScale(v int32)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *OptionsCreateOrderV1Resp) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *OptionsCreateOrderV1Resp) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *OptionsCreateOrderV1Resp) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *OptionsCreateOrderV1Resp) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *OptionsCreateOrderV1Resp) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetReduceOnly

`func (o *OptionsCreateOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *OptionsCreateOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *OptionsCreateOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *OptionsCreateOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *OptionsCreateOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionsCreateOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionsCreateOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionsCreateOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetStatus

`func (o *OptionsCreateOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionsCreateOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionsCreateOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionsCreateOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionsCreateOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionsCreateOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionsCreateOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionsCreateOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OptionsCreateOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OptionsCreateOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OptionsCreateOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OptionsCreateOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OptionsCreateOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionsCreateOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionsCreateOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionsCreateOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OptionsCreateOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OptionsCreateOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OptionsCreateOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OptionsCreateOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


