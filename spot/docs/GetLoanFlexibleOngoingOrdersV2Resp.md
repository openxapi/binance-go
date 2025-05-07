# GetLoanFlexibleOngoingOrdersV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanFlexibleOngoingOrdersV2RespRowsInner**](GetLoanFlexibleOngoingOrdersV2RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLoanFlexibleOngoingOrdersV2Resp

`func NewGetLoanFlexibleOngoingOrdersV2Resp() *GetLoanFlexibleOngoingOrdersV2Resp`

NewGetLoanFlexibleOngoingOrdersV2Resp instantiates a new GetLoanFlexibleOngoingOrdersV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanFlexibleOngoingOrdersV2RespWithDefaults

`func NewGetLoanFlexibleOngoingOrdersV2RespWithDefaults() *GetLoanFlexibleOngoingOrdersV2Resp`

NewGetLoanFlexibleOngoingOrdersV2RespWithDefaults instantiates a new GetLoanFlexibleOngoingOrdersV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetRows() []GetLoanFlexibleOngoingOrdersV2RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetRowsOk() (*[]GetLoanFlexibleOngoingOrdersV2RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) SetRows(v []GetLoanFlexibleOngoingOrdersV2RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanFlexibleOngoingOrdersV2Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


