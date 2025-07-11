# GetMiningPaymentListV1RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProfits** | Pointer to [**[]GetMiningPaymentListV1RespDataAccountProfitsInner**](GetMiningPaymentListV1RespDataAccountProfitsInner.md) |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalNum** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMiningPaymentListV1RespData

`func NewGetMiningPaymentListV1RespData() *GetMiningPaymentListV1RespData`

NewGetMiningPaymentListV1RespData instantiates a new GetMiningPaymentListV1RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiningPaymentListV1RespDataWithDefaults

`func NewGetMiningPaymentListV1RespDataWithDefaults() *GetMiningPaymentListV1RespData`

NewGetMiningPaymentListV1RespDataWithDefaults instantiates a new GetMiningPaymentListV1RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProfits

`func (o *GetMiningPaymentListV1RespData) GetAccountProfits() []GetMiningPaymentListV1RespDataAccountProfitsInner`

GetAccountProfits returns the AccountProfits field if non-nil, zero value otherwise.

### GetAccountProfitsOk

`func (o *GetMiningPaymentListV1RespData) GetAccountProfitsOk() (*[]GetMiningPaymentListV1RespDataAccountProfitsInner, bool)`

GetAccountProfitsOk returns a tuple with the AccountProfits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProfits

`func (o *GetMiningPaymentListV1RespData) SetAccountProfits(v []GetMiningPaymentListV1RespDataAccountProfitsInner)`

SetAccountProfits sets AccountProfits field to given value.

### HasAccountProfits

`func (o *GetMiningPaymentListV1RespData) HasAccountProfits() bool`

HasAccountProfits returns a boolean if a field has been set.

### GetPageSize

`func (o *GetMiningPaymentListV1RespData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetMiningPaymentListV1RespData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetMiningPaymentListV1RespData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetMiningPaymentListV1RespData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalNum

`func (o *GetMiningPaymentListV1RespData) GetTotalNum() int32`

GetTotalNum returns the TotalNum field if non-nil, zero value otherwise.

### GetTotalNumOk

`func (o *GetMiningPaymentListV1RespData) GetTotalNumOk() (*int32, bool)`

GetTotalNumOk returns a tuple with the TotalNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNum

`func (o *GetMiningPaymentListV1RespData) SetTotalNum(v int32)`

SetTotalNum sets TotalNum field to given value.

### HasTotalNum

`func (o *GetMiningPaymentListV1RespData) HasTotalNum() bool`

HasTotalNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


