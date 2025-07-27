# GetAccountCommissionV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**GetAccountCommissionV3RespDiscount**](GetAccountCommissionV3RespDiscount.md) |  | [optional] 
**StandardCommission** | Pointer to [**GetAccountCommissionV3RespStandardCommission**](GetAccountCommissionV3RespStandardCommission.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TaxCommission** | Pointer to [**GetAccountCommissionV3RespStandardCommission**](GetAccountCommissionV3RespStandardCommission.md) |  | [optional] 

## Methods

### NewGetAccountCommissionV3Resp

`func NewGetAccountCommissionV3Resp() *GetAccountCommissionV3Resp`

NewGetAccountCommissionV3Resp instantiates a new GetAccountCommissionV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountCommissionV3RespWithDefaults

`func NewGetAccountCommissionV3RespWithDefaults() *GetAccountCommissionV3Resp`

NewGetAccountCommissionV3RespWithDefaults instantiates a new GetAccountCommissionV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *GetAccountCommissionV3Resp) GetDiscount() GetAccountCommissionV3RespDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *GetAccountCommissionV3Resp) GetDiscountOk() (*GetAccountCommissionV3RespDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *GetAccountCommissionV3Resp) SetDiscount(v GetAccountCommissionV3RespDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *GetAccountCommissionV3Resp) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetStandardCommission

`func (o *GetAccountCommissionV3Resp) GetStandardCommission() GetAccountCommissionV3RespStandardCommission`

GetStandardCommission returns the StandardCommission field if non-nil, zero value otherwise.

### GetStandardCommissionOk

`func (o *GetAccountCommissionV3Resp) GetStandardCommissionOk() (*GetAccountCommissionV3RespStandardCommission, bool)`

GetStandardCommissionOk returns a tuple with the StandardCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommission

`func (o *GetAccountCommissionV3Resp) SetStandardCommission(v GetAccountCommissionV3RespStandardCommission)`

SetStandardCommission sets StandardCommission field to given value.

### HasStandardCommission

`func (o *GetAccountCommissionV3Resp) HasStandardCommission() bool`

HasStandardCommission returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAccountCommissionV3Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAccountCommissionV3Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAccountCommissionV3Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAccountCommissionV3Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTaxCommission

`func (o *GetAccountCommissionV3Resp) GetTaxCommission() GetAccountCommissionV3RespStandardCommission`

GetTaxCommission returns the TaxCommission field if non-nil, zero value otherwise.

### GetTaxCommissionOk

`func (o *GetAccountCommissionV3Resp) GetTaxCommissionOk() (*GetAccountCommissionV3RespStandardCommission, bool)`

GetTaxCommissionOk returns a tuple with the TaxCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommission

`func (o *GetAccountCommissionV3Resp) SetTaxCommission(v GetAccountCommissionV3RespStandardCommission)`

SetTaxCommission sets TaxCommission field to given value.

### HasTaxCommission

`func (o *GetAccountCommissionV3Resp) HasTaxCommission() bool`

HasTaxCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


