# SpotCreateOrderCancelReplaceV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**SpotCreateOrderCancelReplaceV3Data**](SpotCreateOrderCancelReplaceV3Data.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**CancelResponse** | Pointer to [**SpotCreateOrderCancelReplaceV3DataCancelResponse**](SpotCreateOrderCancelReplaceV3DataCancelResponse.md) |  | [optional] 
**CancelResult** | Pointer to **string** |  | [optional] 
**NewOrderResponse** | Pointer to [**SpotCreateOrderCancelReplaceV3DataNewOrderResponse**](SpotCreateOrderCancelReplaceV3DataNewOrderResponse.md) |  | [optional] 
**NewOrderResult** | Pointer to **string** |  | [optional] 

## Methods

### NewSpotCreateOrderCancelReplaceV3Resp

`func NewSpotCreateOrderCancelReplaceV3Resp() *SpotCreateOrderCancelReplaceV3Resp`

NewSpotCreateOrderCancelReplaceV3Resp instantiates a new SpotCreateOrderCancelReplaceV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotCreateOrderCancelReplaceV3RespWithDefaults

`func NewSpotCreateOrderCancelReplaceV3RespWithDefaults() *SpotCreateOrderCancelReplaceV3Resp`

NewSpotCreateOrderCancelReplaceV3RespWithDefaults instantiates a new SpotCreateOrderCancelReplaceV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetData() SpotCreateOrderCancelReplaceV3Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetDataOk() (*SpotCreateOrderCancelReplaceV3Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetData(v SpotCreateOrderCancelReplaceV3Data)`

SetData sets Data field to given value.

### HasData

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetCancelResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCancelResponse() SpotCreateOrderCancelReplaceV3DataCancelResponse`

GetCancelResponse returns the CancelResponse field if non-nil, zero value otherwise.

### GetCancelResponseOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCancelResponseOk() (*SpotCreateOrderCancelReplaceV3DataCancelResponse, bool)`

GetCancelResponseOk returns a tuple with the CancelResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetCancelResponse(v SpotCreateOrderCancelReplaceV3DataCancelResponse)`

SetCancelResponse sets CancelResponse field to given value.

### HasCancelResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasCancelResponse() bool`

HasCancelResponse returns a boolean if a field has been set.

### GetCancelResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCancelResult() string`

GetCancelResult returns the CancelResult field if non-nil, zero value otherwise.

### GetCancelResultOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetCancelResultOk() (*string, bool)`

GetCancelResultOk returns a tuple with the CancelResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetCancelResult(v string)`

SetCancelResult sets CancelResult field to given value.

### HasCancelResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasCancelResult() bool`

HasCancelResult returns a boolean if a field has been set.

### GetNewOrderResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetNewOrderResponse() SpotCreateOrderCancelReplaceV3DataNewOrderResponse`

GetNewOrderResponse returns the NewOrderResponse field if non-nil, zero value otherwise.

### GetNewOrderResponseOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetNewOrderResponseOk() (*SpotCreateOrderCancelReplaceV3DataNewOrderResponse, bool)`

GetNewOrderResponseOk returns a tuple with the NewOrderResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetNewOrderResponse(v SpotCreateOrderCancelReplaceV3DataNewOrderResponse)`

SetNewOrderResponse sets NewOrderResponse field to given value.

### HasNewOrderResponse

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasNewOrderResponse() bool`

HasNewOrderResponse returns a boolean if a field has been set.

### GetNewOrderResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetNewOrderResult() string`

GetNewOrderResult returns the NewOrderResult field if non-nil, zero value otherwise.

### GetNewOrderResultOk

`func (o *SpotCreateOrderCancelReplaceV3Resp) GetNewOrderResultOk() (*string, bool)`

GetNewOrderResultOk returns a tuple with the NewOrderResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) SetNewOrderResult(v string)`

SetNewOrderResult sets NewOrderResult field to given value.

### HasNewOrderResult

`func (o *SpotCreateOrderCancelReplaceV3Resp) HasNewOrderResult() bool`

HasNewOrderResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


