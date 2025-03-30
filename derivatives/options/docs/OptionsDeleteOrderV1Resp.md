# OptionsDeleteOrderV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPrice** | Pointer to **string** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **int32** |  | [optional] 
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
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewOptionsDeleteOrderV1Resp

`func NewOptionsDeleteOrderV1Resp() *OptionsDeleteOrderV1Resp`

NewOptionsDeleteOrderV1Resp instantiates a new OptionsDeleteOrderV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionsDeleteOrderV1RespWithDefaults

`func NewOptionsDeleteOrderV1RespWithDefaults() *OptionsDeleteOrderV1Resp`

NewOptionsDeleteOrderV1RespWithDefaults instantiates a new OptionsDeleteOrderV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPrice

`func (o *OptionsDeleteOrderV1Resp) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *OptionsDeleteOrderV1Resp) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *OptionsDeleteOrderV1Resp) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *OptionsDeleteOrderV1Resp) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetClientOrderId

`func (o *OptionsDeleteOrderV1Resp) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OptionsDeleteOrderV1Resp) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OptionsDeleteOrderV1Resp) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OptionsDeleteOrderV1Resp) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetCreateDate

`func (o *OptionsDeleteOrderV1Resp) GetCreateDate() int32`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *OptionsDeleteOrderV1Resp) GetCreateDateOk() (*int32, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *OptionsDeleteOrderV1Resp) SetCreateDate(v int32)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *OptionsDeleteOrderV1Resp) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetExecutedQty

`func (o *OptionsDeleteOrderV1Resp) GetExecutedQty() string`

GetExecutedQty returns the ExecutedQty field if non-nil, zero value otherwise.

### GetExecutedQtyOk

`func (o *OptionsDeleteOrderV1Resp) GetExecutedQtyOk() (*string, bool)`

GetExecutedQtyOk returns a tuple with the ExecutedQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQty

`func (o *OptionsDeleteOrderV1Resp) SetExecutedQty(v string)`

SetExecutedQty sets ExecutedQty field to given value.

### HasExecutedQty

`func (o *OptionsDeleteOrderV1Resp) HasExecutedQty() bool`

HasExecutedQty returns a boolean if a field has been set.

### GetFee

`func (o *OptionsDeleteOrderV1Resp) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *OptionsDeleteOrderV1Resp) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *OptionsDeleteOrderV1Resp) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *OptionsDeleteOrderV1Resp) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetMmp

`func (o *OptionsDeleteOrderV1Resp) GetMmp() bool`

GetMmp returns the Mmp field if non-nil, zero value otherwise.

### GetMmpOk

`func (o *OptionsDeleteOrderV1Resp) GetMmpOk() (*bool, bool)`

GetMmpOk returns a tuple with the Mmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmp

`func (o *OptionsDeleteOrderV1Resp) SetMmp(v bool)`

SetMmp sets Mmp field to given value.

### HasMmp

`func (o *OptionsDeleteOrderV1Resp) HasMmp() bool`

HasMmp returns a boolean if a field has been set.

### GetOptionSide

`func (o *OptionsDeleteOrderV1Resp) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *OptionsDeleteOrderV1Resp) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *OptionsDeleteOrderV1Resp) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *OptionsDeleteOrderV1Resp) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.

### GetOrderId

`func (o *OptionsDeleteOrderV1Resp) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OptionsDeleteOrderV1Resp) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OptionsDeleteOrderV1Resp) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OptionsDeleteOrderV1Resp) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetPostOnly

`func (o *OptionsDeleteOrderV1Resp) GetPostOnly() bool`

GetPostOnly returns the PostOnly field if non-nil, zero value otherwise.

### GetPostOnlyOk

`func (o *OptionsDeleteOrderV1Resp) GetPostOnlyOk() (*bool, bool)`

GetPostOnlyOk returns a tuple with the PostOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOnly

`func (o *OptionsDeleteOrderV1Resp) SetPostOnly(v bool)`

SetPostOnly sets PostOnly field to given value.

### HasPostOnly

`func (o *OptionsDeleteOrderV1Resp) HasPostOnly() bool`

HasPostOnly returns a boolean if a field has been set.

### GetPrice

`func (o *OptionsDeleteOrderV1Resp) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OptionsDeleteOrderV1Resp) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OptionsDeleteOrderV1Resp) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OptionsDeleteOrderV1Resp) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceScale

`func (o *OptionsDeleteOrderV1Resp) GetPriceScale() int32`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *OptionsDeleteOrderV1Resp) GetPriceScaleOk() (*int32, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *OptionsDeleteOrderV1Resp) SetPriceScale(v int32)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *OptionsDeleteOrderV1Resp) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantity

`func (o *OptionsDeleteOrderV1Resp) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OptionsDeleteOrderV1Resp) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OptionsDeleteOrderV1Resp) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OptionsDeleteOrderV1Resp) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityScale

`func (o *OptionsDeleteOrderV1Resp) GetQuantityScale() int32`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *OptionsDeleteOrderV1Resp) GetQuantityScaleOk() (*int32, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *OptionsDeleteOrderV1Resp) SetQuantityScale(v int32)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *OptionsDeleteOrderV1Resp) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *OptionsDeleteOrderV1Resp) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *OptionsDeleteOrderV1Resp) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *OptionsDeleteOrderV1Resp) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *OptionsDeleteOrderV1Resp) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetReduceOnly

`func (o *OptionsDeleteOrderV1Resp) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *OptionsDeleteOrderV1Resp) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *OptionsDeleteOrderV1Resp) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *OptionsDeleteOrderV1Resp) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *OptionsDeleteOrderV1Resp) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OptionsDeleteOrderV1Resp) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OptionsDeleteOrderV1Resp) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OptionsDeleteOrderV1Resp) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSource

`func (o *OptionsDeleteOrderV1Resp) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OptionsDeleteOrderV1Resp) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OptionsDeleteOrderV1Resp) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OptionsDeleteOrderV1Resp) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *OptionsDeleteOrderV1Resp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptionsDeleteOrderV1Resp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptionsDeleteOrderV1Resp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OptionsDeleteOrderV1Resp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionsDeleteOrderV1Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionsDeleteOrderV1Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionsDeleteOrderV1Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionsDeleteOrderV1Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *OptionsDeleteOrderV1Resp) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *OptionsDeleteOrderV1Resp) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *OptionsDeleteOrderV1Resp) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *OptionsDeleteOrderV1Resp) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetType

`func (o *OptionsDeleteOrderV1Resp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionsDeleteOrderV1Resp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionsDeleteOrderV1Resp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionsDeleteOrderV1Resp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OptionsDeleteOrderV1Resp) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OptionsDeleteOrderV1Resp) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OptionsDeleteOrderV1Resp) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OptionsDeleteOrderV1Resp) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


