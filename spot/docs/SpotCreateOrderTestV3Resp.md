# SpotCreateOrderTestV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**GetAccountCommissionV3RespDiscount**](GetAccountCommissionV3RespDiscount.md) |  | [optional] 
**StandardCommissionForOrder** | Pointer to [**SpotCreateOrderTestV3RespStandardCommissionForOrder**](SpotCreateOrderTestV3RespStandardCommissionForOrder.md) |  | [optional] 
**TaxCommissionForOrder** | Pointer to [**SpotCreateOrderTestV3RespStandardCommissionForOrder**](SpotCreateOrderTestV3RespStandardCommissionForOrder.md) |  | [optional] 

## Methods

### NewSpotCreateOrderTestV3Resp

`func NewSpotCreateOrderTestV3Resp() *SpotCreateOrderTestV3Resp`

NewSpotCreateOrderTestV3Resp instantiates a new SpotCreateOrderTestV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotCreateOrderTestV3RespWithDefaults

`func NewSpotCreateOrderTestV3RespWithDefaults() *SpotCreateOrderTestV3Resp`

NewSpotCreateOrderTestV3RespWithDefaults instantiates a new SpotCreateOrderTestV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *SpotCreateOrderTestV3Resp) GetDiscount() GetAccountCommissionV3RespDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SpotCreateOrderTestV3Resp) GetDiscountOk() (*GetAccountCommissionV3RespDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SpotCreateOrderTestV3Resp) SetDiscount(v GetAccountCommissionV3RespDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SpotCreateOrderTestV3Resp) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetStandardCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) GetStandardCommissionForOrder() SpotCreateOrderTestV3RespStandardCommissionForOrder`

GetStandardCommissionForOrder returns the StandardCommissionForOrder field if non-nil, zero value otherwise.

### GetStandardCommissionForOrderOk

`func (o *SpotCreateOrderTestV3Resp) GetStandardCommissionForOrderOk() (*SpotCreateOrderTestV3RespStandardCommissionForOrder, bool)`

GetStandardCommissionForOrderOk returns a tuple with the StandardCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) SetStandardCommissionForOrder(v SpotCreateOrderTestV3RespStandardCommissionForOrder)`

SetStandardCommissionForOrder sets StandardCommissionForOrder field to given value.

### HasStandardCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) HasStandardCommissionForOrder() bool`

HasStandardCommissionForOrder returns a boolean if a field has been set.

### GetTaxCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) GetTaxCommissionForOrder() SpotCreateOrderTestV3RespStandardCommissionForOrder`

GetTaxCommissionForOrder returns the TaxCommissionForOrder field if non-nil, zero value otherwise.

### GetTaxCommissionForOrderOk

`func (o *SpotCreateOrderTestV3Resp) GetTaxCommissionForOrderOk() (*SpotCreateOrderTestV3RespStandardCommissionForOrder, bool)`

GetTaxCommissionForOrderOk returns a tuple with the TaxCommissionForOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) SetTaxCommissionForOrder(v SpotCreateOrderTestV3RespStandardCommissionForOrder)`

SetTaxCommissionForOrder sets TaxCommissionForOrder field to given value.

### HasTaxCommissionForOrder

`func (o *SpotCreateOrderTestV3Resp) HasTaxCommissionForOrder() bool`

HasTaxCommissionForOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


