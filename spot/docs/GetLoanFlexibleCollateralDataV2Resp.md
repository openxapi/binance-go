# GetLoanFlexibleCollateralDataV2Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanFlexibleCollateralDataV2RespRowsInner**](GetLoanFlexibleCollateralDataV2RespRowsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetLoanFlexibleCollateralDataV2Resp

`func NewGetLoanFlexibleCollateralDataV2Resp() *GetLoanFlexibleCollateralDataV2Resp`

NewGetLoanFlexibleCollateralDataV2Resp instantiates a new GetLoanFlexibleCollateralDataV2Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanFlexibleCollateralDataV2RespWithDefaults

`func NewGetLoanFlexibleCollateralDataV2RespWithDefaults() *GetLoanFlexibleCollateralDataV2Resp`

NewGetLoanFlexibleCollateralDataV2RespWithDefaults instantiates a new GetLoanFlexibleCollateralDataV2Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanFlexibleCollateralDataV2Resp) GetRows() []GetLoanFlexibleCollateralDataV2RespRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanFlexibleCollateralDataV2Resp) GetRowsOk() (*[]GetLoanFlexibleCollateralDataV2RespRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanFlexibleCollateralDataV2Resp) SetRows(v []GetLoanFlexibleCollateralDataV2RespRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanFlexibleCollateralDataV2Resp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanFlexibleCollateralDataV2Resp) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanFlexibleCollateralDataV2Resp) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanFlexibleCollateralDataV2Resp) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanFlexibleCollateralDataV2Resp) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


