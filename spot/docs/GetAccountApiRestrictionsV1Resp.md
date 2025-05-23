# GetAccountApiRestrictionsV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**EnableFixApiTrade** | Pointer to **bool** |  | [optional] 
**EnableFixReadOnly** | Pointer to **bool** |  | [optional] 
**EnableFutures** | Pointer to **bool** |  | [optional] 
**EnableInternalTransfer** | Pointer to **bool** |  | [optional] 
**EnableMargin** | Pointer to **bool** |  | [optional] 
**EnablePortfolioMarginTrading** | Pointer to **bool** |  | [optional] 
**EnableReading** | Pointer to **bool** |  | [optional] 
**EnableSpotAndMarginTrading** | Pointer to **bool** |  | [optional] 
**EnableVanillaOptions** | Pointer to **bool** |  | [optional] 
**EnableWithdrawals** | Pointer to **bool** |  | [optional] 
**IpRestrict** | Pointer to **bool** |  | [optional] 
**PermitsUniversalTransfer** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAccountApiRestrictionsV1Resp

`func NewGetAccountApiRestrictionsV1Resp() *GetAccountApiRestrictionsV1Resp`

NewGetAccountApiRestrictionsV1Resp instantiates a new GetAccountApiRestrictionsV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountApiRestrictionsV1RespWithDefaults

`func NewGetAccountApiRestrictionsV1RespWithDefaults() *GetAccountApiRestrictionsV1Resp`

NewGetAccountApiRestrictionsV1RespWithDefaults instantiates a new GetAccountApiRestrictionsV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *GetAccountApiRestrictionsV1Resp) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetAccountApiRestrictionsV1Resp) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetAccountApiRestrictionsV1Resp) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetAccountApiRestrictionsV1Resp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEnableFixApiTrade

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFixApiTrade() bool`

GetEnableFixApiTrade returns the EnableFixApiTrade field if non-nil, zero value otherwise.

### GetEnableFixApiTradeOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFixApiTradeOk() (*bool, bool)`

GetEnableFixApiTradeOk returns a tuple with the EnableFixApiTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixApiTrade

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableFixApiTrade(v bool)`

SetEnableFixApiTrade sets EnableFixApiTrade field to given value.

### HasEnableFixApiTrade

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableFixApiTrade() bool`

HasEnableFixApiTrade returns a boolean if a field has been set.

### GetEnableFixReadOnly

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFixReadOnly() bool`

GetEnableFixReadOnly returns the EnableFixReadOnly field if non-nil, zero value otherwise.

### GetEnableFixReadOnlyOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFixReadOnlyOk() (*bool, bool)`

GetEnableFixReadOnlyOk returns a tuple with the EnableFixReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFixReadOnly

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableFixReadOnly(v bool)`

SetEnableFixReadOnly sets EnableFixReadOnly field to given value.

### HasEnableFixReadOnly

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableFixReadOnly() bool`

HasEnableFixReadOnly returns a boolean if a field has been set.

### GetEnableFutures

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFutures() bool`

GetEnableFutures returns the EnableFutures field if non-nil, zero value otherwise.

### GetEnableFuturesOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableFuturesOk() (*bool, bool)`

GetEnableFuturesOk returns a tuple with the EnableFutures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFutures

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableFutures(v bool)`

SetEnableFutures sets EnableFutures field to given value.

### HasEnableFutures

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableFutures() bool`

HasEnableFutures returns a boolean if a field has been set.

### GetEnableInternalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableInternalTransfer() bool`

GetEnableInternalTransfer returns the EnableInternalTransfer field if non-nil, zero value otherwise.

### GetEnableInternalTransferOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableInternalTransferOk() (*bool, bool)`

GetEnableInternalTransferOk returns a tuple with the EnableInternalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableInternalTransfer(v bool)`

SetEnableInternalTransfer sets EnableInternalTransfer field to given value.

### HasEnableInternalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableInternalTransfer() bool`

HasEnableInternalTransfer returns a boolean if a field has been set.

### GetEnableMargin

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableMargin() bool`

GetEnableMargin returns the EnableMargin field if non-nil, zero value otherwise.

### GetEnableMarginOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableMarginOk() (*bool, bool)`

GetEnableMarginOk returns a tuple with the EnableMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMargin

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableMargin(v bool)`

SetEnableMargin sets EnableMargin field to given value.

### HasEnableMargin

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableMargin() bool`

HasEnableMargin returns a boolean if a field has been set.

### GetEnablePortfolioMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) GetEnablePortfolioMarginTrading() bool`

GetEnablePortfolioMarginTrading returns the EnablePortfolioMarginTrading field if non-nil, zero value otherwise.

### GetEnablePortfolioMarginTradingOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnablePortfolioMarginTradingOk() (*bool, bool)`

GetEnablePortfolioMarginTradingOk returns a tuple with the EnablePortfolioMarginTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePortfolioMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) SetEnablePortfolioMarginTrading(v bool)`

SetEnablePortfolioMarginTrading sets EnablePortfolioMarginTrading field to given value.

### HasEnablePortfolioMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) HasEnablePortfolioMarginTrading() bool`

HasEnablePortfolioMarginTrading returns a boolean if a field has been set.

### GetEnableReading

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableReading() bool`

GetEnableReading returns the EnableReading field if non-nil, zero value otherwise.

### GetEnableReadingOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableReadingOk() (*bool, bool)`

GetEnableReadingOk returns a tuple with the EnableReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReading

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableReading(v bool)`

SetEnableReading sets EnableReading field to given value.

### HasEnableReading

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableReading() bool`

HasEnableReading returns a boolean if a field has been set.

### GetEnableSpotAndMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableSpotAndMarginTrading() bool`

GetEnableSpotAndMarginTrading returns the EnableSpotAndMarginTrading field if non-nil, zero value otherwise.

### GetEnableSpotAndMarginTradingOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableSpotAndMarginTradingOk() (*bool, bool)`

GetEnableSpotAndMarginTradingOk returns a tuple with the EnableSpotAndMarginTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotAndMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableSpotAndMarginTrading(v bool)`

SetEnableSpotAndMarginTrading sets EnableSpotAndMarginTrading field to given value.

### HasEnableSpotAndMarginTrading

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableSpotAndMarginTrading() bool`

HasEnableSpotAndMarginTrading returns a boolean if a field has been set.

### GetEnableVanillaOptions

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableVanillaOptions() bool`

GetEnableVanillaOptions returns the EnableVanillaOptions field if non-nil, zero value otherwise.

### GetEnableVanillaOptionsOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableVanillaOptionsOk() (*bool, bool)`

GetEnableVanillaOptionsOk returns a tuple with the EnableVanillaOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVanillaOptions

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableVanillaOptions(v bool)`

SetEnableVanillaOptions sets EnableVanillaOptions field to given value.

### HasEnableVanillaOptions

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableVanillaOptions() bool`

HasEnableVanillaOptions returns a boolean if a field has been set.

### GetEnableWithdrawals

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableWithdrawals() bool`

GetEnableWithdrawals returns the EnableWithdrawals field if non-nil, zero value otherwise.

### GetEnableWithdrawalsOk

`func (o *GetAccountApiRestrictionsV1Resp) GetEnableWithdrawalsOk() (*bool, bool)`

GetEnableWithdrawalsOk returns a tuple with the EnableWithdrawals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWithdrawals

`func (o *GetAccountApiRestrictionsV1Resp) SetEnableWithdrawals(v bool)`

SetEnableWithdrawals sets EnableWithdrawals field to given value.

### HasEnableWithdrawals

`func (o *GetAccountApiRestrictionsV1Resp) HasEnableWithdrawals() bool`

HasEnableWithdrawals returns a boolean if a field has been set.

### GetIpRestrict

`func (o *GetAccountApiRestrictionsV1Resp) GetIpRestrict() bool`

GetIpRestrict returns the IpRestrict field if non-nil, zero value otherwise.

### GetIpRestrictOk

`func (o *GetAccountApiRestrictionsV1Resp) GetIpRestrictOk() (*bool, bool)`

GetIpRestrictOk returns a tuple with the IpRestrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestrict

`func (o *GetAccountApiRestrictionsV1Resp) SetIpRestrict(v bool)`

SetIpRestrict sets IpRestrict field to given value.

### HasIpRestrict

`func (o *GetAccountApiRestrictionsV1Resp) HasIpRestrict() bool`

HasIpRestrict returns a boolean if a field has been set.

### GetPermitsUniversalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) GetPermitsUniversalTransfer() bool`

GetPermitsUniversalTransfer returns the PermitsUniversalTransfer field if non-nil, zero value otherwise.

### GetPermitsUniversalTransferOk

`func (o *GetAccountApiRestrictionsV1Resp) GetPermitsUniversalTransferOk() (*bool, bool)`

GetPermitsUniversalTransferOk returns a tuple with the PermitsUniversalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitsUniversalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) SetPermitsUniversalTransfer(v bool)`

SetPermitsUniversalTransfer sets PermitsUniversalTransfer field to given value.

### HasPermitsUniversalTransfer

`func (o *GetAccountApiRestrictionsV1Resp) HasPermitsUniversalTransfer() bool`

HasPermitsUniversalTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


