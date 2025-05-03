# GetMarginAccountV1Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCollateralValueInUSDT** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**BorrowEnabled** | Pointer to **bool** |  | [optional] 
**CollateralMarginLevel** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **bool** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**TotalAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalLiabilityOfBtc** | Pointer to **string** |  | [optional] 
**TotalNetAssetOfBtc** | Pointer to **string** |  | [optional] 
**TotalOpenOrderLossInUSDT** | Pointer to **string** |  | [optional] 
**TradeEnabled** | Pointer to **bool** |  | [optional] 
**TransferInEnabled** | Pointer to **bool** |  | [optional] 
**TransferOutEnabled** | Pointer to **bool** |  | [optional] 
**UserAssets** | Pointer to [**[]GetManagedSubaccountMarginAssetV1RespUserAssetsInner**](GetManagedSubaccountMarginAssetV1RespUserAssetsInner.md) |  | [optional] 

## Methods

### NewGetMarginAccountV1Resp

`func NewGetMarginAccountV1Resp() *GetMarginAccountV1Resp`

NewGetMarginAccountV1Resp instantiates a new GetMarginAccountV1Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginAccountV1RespWithDefaults

`func NewGetMarginAccountV1RespWithDefaults() *GetMarginAccountV1Resp`

NewGetMarginAccountV1RespWithDefaults instantiates a new GetMarginAccountV1Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCollateralValueInUSDT

`func (o *GetMarginAccountV1Resp) GetTotalCollateralValueInUSDT() string`

GetTotalCollateralValueInUSDT returns the TotalCollateralValueInUSDT field if non-nil, zero value otherwise.

### GetTotalCollateralValueInUSDTOk

`func (o *GetMarginAccountV1Resp) GetTotalCollateralValueInUSDTOk() (*string, bool)`

GetTotalCollateralValueInUSDTOk returns a tuple with the TotalCollateralValueInUSDT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollateralValueInUSDT

`func (o *GetMarginAccountV1Resp) SetTotalCollateralValueInUSDT(v string)`

SetTotalCollateralValueInUSDT sets TotalCollateralValueInUSDT field to given value.

### HasTotalCollateralValueInUSDT

`func (o *GetMarginAccountV1Resp) HasTotalCollateralValueInUSDT() bool`

HasTotalCollateralValueInUSDT returns a boolean if a field has been set.

### GetAccountType

`func (o *GetMarginAccountV1Resp) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetMarginAccountV1Resp) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetMarginAccountV1Resp) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetMarginAccountV1Resp) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBorrowEnabled

`func (o *GetMarginAccountV1Resp) GetBorrowEnabled() bool`

GetBorrowEnabled returns the BorrowEnabled field if non-nil, zero value otherwise.

### GetBorrowEnabledOk

`func (o *GetMarginAccountV1Resp) GetBorrowEnabledOk() (*bool, bool)`

GetBorrowEnabledOk returns a tuple with the BorrowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowEnabled

`func (o *GetMarginAccountV1Resp) SetBorrowEnabled(v bool)`

SetBorrowEnabled sets BorrowEnabled field to given value.

### HasBorrowEnabled

`func (o *GetMarginAccountV1Resp) HasBorrowEnabled() bool`

HasBorrowEnabled returns a boolean if a field has been set.

### GetCollateralMarginLevel

`func (o *GetMarginAccountV1Resp) GetCollateralMarginLevel() string`

GetCollateralMarginLevel returns the CollateralMarginLevel field if non-nil, zero value otherwise.

### GetCollateralMarginLevelOk

`func (o *GetMarginAccountV1Resp) GetCollateralMarginLevelOk() (*string, bool)`

GetCollateralMarginLevelOk returns a tuple with the CollateralMarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralMarginLevel

`func (o *GetMarginAccountV1Resp) SetCollateralMarginLevel(v string)`

SetCollateralMarginLevel sets CollateralMarginLevel field to given value.

### HasCollateralMarginLevel

`func (o *GetMarginAccountV1Resp) HasCollateralMarginLevel() bool`

HasCollateralMarginLevel returns a boolean if a field has been set.

### GetCreated

`func (o *GetMarginAccountV1Resp) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetMarginAccountV1Resp) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetMarginAccountV1Resp) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetMarginAccountV1Resp) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetMarginLevel

`func (o *GetMarginAccountV1Resp) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *GetMarginAccountV1Resp) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *GetMarginAccountV1Resp) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *GetMarginAccountV1Resp) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetTotalAssetOfBtc

`func (o *GetMarginAccountV1Resp) GetTotalAssetOfBtc() string`

GetTotalAssetOfBtc returns the TotalAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalAssetOfBtcOk

`func (o *GetMarginAccountV1Resp) GetTotalAssetOfBtcOk() (*string, bool)`

GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssetOfBtc

`func (o *GetMarginAccountV1Resp) SetTotalAssetOfBtc(v string)`

SetTotalAssetOfBtc sets TotalAssetOfBtc field to given value.

### HasTotalAssetOfBtc

`func (o *GetMarginAccountV1Resp) HasTotalAssetOfBtc() bool`

HasTotalAssetOfBtc returns a boolean if a field has been set.

### GetTotalLiabilityOfBtc

`func (o *GetMarginAccountV1Resp) GetTotalLiabilityOfBtc() string`

GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field if non-nil, zero value otherwise.

### GetTotalLiabilityOfBtcOk

`func (o *GetMarginAccountV1Resp) GetTotalLiabilityOfBtcOk() (*string, bool)`

GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabilityOfBtc

`func (o *GetMarginAccountV1Resp) SetTotalLiabilityOfBtc(v string)`

SetTotalLiabilityOfBtc sets TotalLiabilityOfBtc field to given value.

### HasTotalLiabilityOfBtc

`func (o *GetMarginAccountV1Resp) HasTotalLiabilityOfBtc() bool`

HasTotalLiabilityOfBtc returns a boolean if a field has been set.

### GetTotalNetAssetOfBtc

`func (o *GetMarginAccountV1Resp) GetTotalNetAssetOfBtc() string`

GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field if non-nil, zero value otherwise.

### GetTotalNetAssetOfBtcOk

`func (o *GetMarginAccountV1Resp) GetTotalNetAssetOfBtcOk() (*string, bool)`

GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetAssetOfBtc

`func (o *GetMarginAccountV1Resp) SetTotalNetAssetOfBtc(v string)`

SetTotalNetAssetOfBtc sets TotalNetAssetOfBtc field to given value.

### HasTotalNetAssetOfBtc

`func (o *GetMarginAccountV1Resp) HasTotalNetAssetOfBtc() bool`

HasTotalNetAssetOfBtc returns a boolean if a field has been set.

### GetTotalOpenOrderLossInUSDT

`func (o *GetMarginAccountV1Resp) GetTotalOpenOrderLossInUSDT() string`

GetTotalOpenOrderLossInUSDT returns the TotalOpenOrderLossInUSDT field if non-nil, zero value otherwise.

### GetTotalOpenOrderLossInUSDTOk

`func (o *GetMarginAccountV1Resp) GetTotalOpenOrderLossInUSDTOk() (*string, bool)`

GetTotalOpenOrderLossInUSDTOk returns a tuple with the TotalOpenOrderLossInUSDT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOpenOrderLossInUSDT

`func (o *GetMarginAccountV1Resp) SetTotalOpenOrderLossInUSDT(v string)`

SetTotalOpenOrderLossInUSDT sets TotalOpenOrderLossInUSDT field to given value.

### HasTotalOpenOrderLossInUSDT

`func (o *GetMarginAccountV1Resp) HasTotalOpenOrderLossInUSDT() bool`

HasTotalOpenOrderLossInUSDT returns a boolean if a field has been set.

### GetTradeEnabled

`func (o *GetMarginAccountV1Resp) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *GetMarginAccountV1Resp) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *GetMarginAccountV1Resp) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.

### HasTradeEnabled

`func (o *GetMarginAccountV1Resp) HasTradeEnabled() bool`

HasTradeEnabled returns a boolean if a field has been set.

### GetTransferInEnabled

`func (o *GetMarginAccountV1Resp) GetTransferInEnabled() bool`

GetTransferInEnabled returns the TransferInEnabled field if non-nil, zero value otherwise.

### GetTransferInEnabledOk

`func (o *GetMarginAccountV1Resp) GetTransferInEnabledOk() (*bool, bool)`

GetTransferInEnabledOk returns a tuple with the TransferInEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferInEnabled

`func (o *GetMarginAccountV1Resp) SetTransferInEnabled(v bool)`

SetTransferInEnabled sets TransferInEnabled field to given value.

### HasTransferInEnabled

`func (o *GetMarginAccountV1Resp) HasTransferInEnabled() bool`

HasTransferInEnabled returns a boolean if a field has been set.

### GetTransferOutEnabled

`func (o *GetMarginAccountV1Resp) GetTransferOutEnabled() bool`

GetTransferOutEnabled returns the TransferOutEnabled field if non-nil, zero value otherwise.

### GetTransferOutEnabledOk

`func (o *GetMarginAccountV1Resp) GetTransferOutEnabledOk() (*bool, bool)`

GetTransferOutEnabledOk returns a tuple with the TransferOutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferOutEnabled

`func (o *GetMarginAccountV1Resp) SetTransferOutEnabled(v bool)`

SetTransferOutEnabled sets TransferOutEnabled field to given value.

### HasTransferOutEnabled

`func (o *GetMarginAccountV1Resp) HasTransferOutEnabled() bool`

HasTransferOutEnabled returns a boolean if a field has been set.

### GetUserAssets

`func (o *GetMarginAccountV1Resp) GetUserAssets() []GetManagedSubaccountMarginAssetV1RespUserAssetsInner`

GetUserAssets returns the UserAssets field if non-nil, zero value otherwise.

### GetUserAssetsOk

`func (o *GetMarginAccountV1Resp) GetUserAssetsOk() (*[]GetManagedSubaccountMarginAssetV1RespUserAssetsInner, bool)`

GetUserAssetsOk returns a tuple with the UserAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssets

`func (o *GetMarginAccountV1Resp) SetUserAssets(v []GetManagedSubaccountMarginAssetV1RespUserAssetsInner)`

SetUserAssets sets UserAssets field to given value.

### HasUserAssets

`func (o *GetMarginAccountV1Resp) HasUserAssets() bool`

HasUserAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


