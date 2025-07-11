# GetLoanVipCollateralDataV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanVipCollateralDataV1RespRowsInner**](GetLoanVipCollateralDataV1RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLoanVipCollateralDataV1Resp

`func NewGetLoanVipCollateralDataV1Resp() *GetLoanVipCollateralDataV1Resp`

NewGetLoanVipCollateralDataV1Resp instantiates a new GetLoanVipCollateralDataV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanVipCollateralDataV1RespWithDefaults

`func NewGetLoanVipCollateralDataV1RespWithDefaults() *GetLoanVipCollateralDataV1Resp`

NewGetLoanVipCollateralDataV1RespWithDefaults instantiates a new GetLoanVipCollateralDataV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanVipCollateralDataV1Resp) GetRows() []GetLoanVipCollateralDataV1RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanVipCollateralDataV1Resp) GetRowsOk() (*[]GetLoanVipCollateralDataV1RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanVipCollateralDataV1Resp) SetRows(v []GetLoanVipCollateralDataV1RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanVipCollateralDataV1Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanVipCollateralDataV1Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanVipCollateralDataV1Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanVipCollateralDataV1Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanVipCollateralDataV1Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


