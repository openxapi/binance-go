# GetMarginIsolatedAccountV1RespAssetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAsset** | Pointer to [**GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset**](GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IndexPrice** | Pointer to **string** |  | [optional] 
**IsolatedCreated** | Pointer to **bool** |  | [optional] 
**LiquidatePrice** | Pointer to **string** |  | [optional] 
**LiquidateRate** | Pointer to **string** |  | [optional] 
**MarginLevel** | Pointer to **string** |  | [optional] 
**MarginLevelStatus** | Pointer to **string** |  | [optional] 
**MarginRatio** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to [**GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset**](GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset.md) |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TradeEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetMarginIsolatedAccountV1RespAssetsInner

`func NewGetMarginIsolatedAccountV1RespAssetsInner() *GetMarginIsolatedAccountV1RespAssetsInner`

NewGetMarginIsolatedAccountV1RespAssetsInner instantiates a new GetMarginIsolatedAccountV1RespAssetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarginIsolatedAccountV1RespAssetsInnerWithDefaults

`func NewGetMarginIsolatedAccountV1RespAssetsInnerWithDefaults() *GetMarginIsolatedAccountV1RespAssetsInner`

NewGetMarginIsolatedAccountV1RespAssetsInnerWithDefaults instantiates a new GetMarginIsolatedAccountV1RespAssetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetBaseAsset() GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetBaseAssetOk() (*GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetBaseAsset(v GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIndexPrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.

### HasIndexPrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasIndexPrice() bool`

HasIndexPrice returns a boolean if a field has been set.

### GetIsolatedCreated

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetIsolatedCreated() bool`

GetIsolatedCreated returns the IsolatedCreated field if non-nil, zero value otherwise.

### GetIsolatedCreatedOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetIsolatedCreatedOk() (*bool, bool)`

GetIsolatedCreatedOk returns a tuple with the IsolatedCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedCreated

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetIsolatedCreated(v bool)`

SetIsolatedCreated sets IsolatedCreated field to given value.

### HasIsolatedCreated

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasIsolatedCreated() bool`

HasIsolatedCreated returns a boolean if a field has been set.

### GetLiquidatePrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetLiquidatePrice() string`

GetLiquidatePrice returns the LiquidatePrice field if non-nil, zero value otherwise.

### GetLiquidatePriceOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetLiquidatePriceOk() (*string, bool)`

GetLiquidatePriceOk returns a tuple with the LiquidatePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidatePrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetLiquidatePrice(v string)`

SetLiquidatePrice sets LiquidatePrice field to given value.

### HasLiquidatePrice

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasLiquidatePrice() bool`

HasLiquidatePrice returns a boolean if a field has been set.

### GetLiquidateRate

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetLiquidateRate() string`

GetLiquidateRate returns the LiquidateRate field if non-nil, zero value otherwise.

### GetLiquidateRateOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetLiquidateRateOk() (*string, bool)`

GetLiquidateRateOk returns a tuple with the LiquidateRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidateRate

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetLiquidateRate(v string)`

SetLiquidateRate sets LiquidateRate field to given value.

### HasLiquidateRate

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasLiquidateRate() bool`

HasLiquidateRate returns a boolean if a field has been set.

### GetMarginLevel

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginLevel() string`

GetMarginLevel returns the MarginLevel field if non-nil, zero value otherwise.

### GetMarginLevelOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginLevelOk() (*string, bool)`

GetMarginLevelOk returns a tuple with the MarginLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevel

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetMarginLevel(v string)`

SetMarginLevel sets MarginLevel field to given value.

### HasMarginLevel

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasMarginLevel() bool`

HasMarginLevel returns a boolean if a field has been set.

### GetMarginLevelStatus

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginLevelStatus() string`

GetMarginLevelStatus returns the MarginLevelStatus field if non-nil, zero value otherwise.

### GetMarginLevelStatusOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginLevelStatusOk() (*string, bool)`

GetMarginLevelStatusOk returns a tuple with the MarginLevelStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginLevelStatus

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetMarginLevelStatus(v string)`

SetMarginLevelStatus sets MarginLevelStatus field to given value.

### HasMarginLevelStatus

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasMarginLevelStatus() bool`

HasMarginLevelStatus returns a boolean if a field has been set.

### GetMarginRatio

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginRatio() string`

GetMarginRatio returns the MarginRatio field if non-nil, zero value otherwise.

### GetMarginRatioOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetMarginRatioOk() (*string, bool)`

GetMarginRatioOk returns a tuple with the MarginRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginRatio

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetMarginRatio(v string)`

SetMarginRatio sets MarginRatio field to given value.

### HasMarginRatio

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasMarginRatio() bool`

HasMarginRatio returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetQuoteAsset() GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetQuoteAssetOk() (*GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetQuoteAsset(v GetMarginIsolatedAccountV1RespAssetsInnerBaseAsset)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetSymbol

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradeEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetTradeEnabled() bool`

GetTradeEnabled returns the TradeEnabled field if non-nil, zero value otherwise.

### GetTradeEnabledOk

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) GetTradeEnabledOk() (*bool, bool)`

GetTradeEnabledOk returns a tuple with the TradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) SetTradeEnabled(v bool)`

SetTradeEnabled sets TradeEnabled field to given value.

### HasTradeEnabled

`func (o *GetMarginIsolatedAccountV1RespAssetsInner) HasTradeEnabled() bool`

HasTradeEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


