# SpotGetAccountCommissionV3Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**SpotCreateOrderTestV3RespDiscount**](SpotCreateOrderTestV3RespDiscount.md) |  | [optional] 
**StandardCommission** | Pointer to [**SpotGetAccountCommissionV3RespStandardCommission**](SpotGetAccountCommissionV3RespStandardCommission.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TaxCommission** | Pointer to [**SpotGetAccountCommissionV3RespStandardCommission**](SpotGetAccountCommissionV3RespStandardCommission.md) |  | [optional] 

## Methods

### NewSpotGetAccountCommissionV3Resp

`func NewSpotGetAccountCommissionV3Resp() *SpotGetAccountCommissionV3Resp`

NewSpotGetAccountCommissionV3Resp instantiates a new SpotGetAccountCommissionV3Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotGetAccountCommissionV3RespWithDefaults

`func NewSpotGetAccountCommissionV3RespWithDefaults() *SpotGetAccountCommissionV3Resp`

NewSpotGetAccountCommissionV3RespWithDefaults instantiates a new SpotGetAccountCommissionV3Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *SpotGetAccountCommissionV3Resp) GetDiscount() SpotCreateOrderTestV3RespDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *SpotGetAccountCommissionV3Resp) GetDiscountOk() (*SpotCreateOrderTestV3RespDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *SpotGetAccountCommissionV3Resp) SetDiscount(v SpotCreateOrderTestV3RespDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *SpotGetAccountCommissionV3Resp) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetStandardCommission

`func (o *SpotGetAccountCommissionV3Resp) GetStandardCommission() SpotGetAccountCommissionV3RespStandardCommission`

GetStandardCommission returns the StandardCommission field if non-nil, zero value otherwise.

### GetStandardCommissionOk

`func (o *SpotGetAccountCommissionV3Resp) GetStandardCommissionOk() (*SpotGetAccountCommissionV3RespStandardCommission, bool)`

GetStandardCommissionOk returns a tuple with the StandardCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCommission

`func (o *SpotGetAccountCommissionV3Resp) SetStandardCommission(v SpotGetAccountCommissionV3RespStandardCommission)`

SetStandardCommission sets StandardCommission field to given value.

### HasStandardCommission

`func (o *SpotGetAccountCommissionV3Resp) HasStandardCommission() bool`

HasStandardCommission returns a boolean if a field has been set.

### GetSymbol

`func (o *SpotGetAccountCommissionV3Resp) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SpotGetAccountCommissionV3Resp) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SpotGetAccountCommissionV3Resp) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SpotGetAccountCommissionV3Resp) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTaxCommission

`func (o *SpotGetAccountCommissionV3Resp) GetTaxCommission() SpotGetAccountCommissionV3RespStandardCommission`

GetTaxCommission returns the TaxCommission field if non-nil, zero value otherwise.

### GetTaxCommissionOk

`func (o *SpotGetAccountCommissionV3Resp) GetTaxCommissionOk() (*SpotGetAccountCommissionV3RespStandardCommission, bool)`

GetTaxCommissionOk returns a tuple with the TaxCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCommission

`func (o *SpotGetAccountCommissionV3Resp) SetTaxCommission(v SpotGetAccountCommissionV3RespStandardCommission)`

SetTaxCommission sets TaxCommission field to given value.

### HasTaxCommission

`func (o *SpotGetAccountCommissionV3Resp) HasTaxCommission() bool`

HasTaxCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


