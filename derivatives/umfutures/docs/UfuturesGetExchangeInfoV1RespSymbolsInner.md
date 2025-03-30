# UfuturesGetExchangeInfoV1RespSymbolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderType** | Pointer to **[]string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int32** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**DeliveryDate** | Pointer to **int32** |  | [optional] 
**Filters** | Pointer to [**[]UfuturesGetExchangeInfoV1RespSymbolsInnerFiltersInner**](UfuturesGetExchangeInfoV1RespSymbolsInnerFiltersInner.md) |  | [optional] 
**LiquidationFee** | Pointer to **string** |  | [optional] 
**MaintMarginPercent** | Pointer to **string** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**MarketTakeBound** | Pointer to **string** |  | [optional] 
**OnboardDate** | Pointer to **int32** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**PricePrecision** | Pointer to **int32** |  | [optional] 
**QuantityPrecision** | Pointer to **int32** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**QuotePrecision** | Pointer to **int32** |  | [optional] 
**RequiredMarginPercent** | Pointer to **string** |  | [optional] 
**SettlePlan** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **[]string** |  | [optional] 
**TriggerProtect** | Pointer to **string** |  | [optional] 
**UnderlyingSubType** | Pointer to **[]string** |  | [optional] 
**UnderlyingType** | Pointer to **string** |  | [optional] 

## Methods

### NewUfuturesGetExchangeInfoV1RespSymbolsInner

`func NewUfuturesGetExchangeInfoV1RespSymbolsInner() *UfuturesGetExchangeInfoV1RespSymbolsInner`

NewUfuturesGetExchangeInfoV1RespSymbolsInner instantiates a new UfuturesGetExchangeInfoV1RespSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults

`func NewUfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults() *UfuturesGetExchangeInfoV1RespSymbolsInner`

NewUfuturesGetExchangeInfoV1RespSymbolsInnerWithDefaults instantiates a new UfuturesGetExchangeInfoV1RespSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetOrderType() []string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetOrderTypeOk() (*[]string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetOrderType(v []string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetBaseAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetPrecision() int32`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetBaseAssetPrecisionOk() (*int32, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetBaseAssetPrecision(v int32)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetContractType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetDeliveryDate() int32`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetDeliveryDateOk() (*int32, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetDeliveryDate(v int32)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetFilters

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetFilters() []UfuturesGetExchangeInfoV1RespSymbolsInnerFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetFiltersOk() (*[]UfuturesGetExchangeInfoV1RespSymbolsInnerFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetFilters(v []UfuturesGetExchangeInfoV1RespSymbolsInnerFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLiquidationFee

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetLiquidationFee() string`

GetLiquidationFee returns the LiquidationFee field if non-nil, zero value otherwise.

### GetLiquidationFeeOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetLiquidationFeeOk() (*string, bool)`

GetLiquidationFeeOk returns a tuple with the LiquidationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationFee

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetLiquidationFee(v string)`

SetLiquidationFee sets LiquidationFee field to given value.

### HasLiquidationFee

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasLiquidationFee() bool`

HasLiquidationFee returns a boolean if a field has been set.

### GetMaintMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMaintMarginPercent() string`

GetMaintMarginPercent returns the MaintMarginPercent field if non-nil, zero value otherwise.

### GetMaintMarginPercentOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMaintMarginPercentOk() (*string, bool)`

GetMaintMarginPercentOk returns a tuple with the MaintMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetMaintMarginPercent(v string)`

SetMaintMarginPercent sets MaintMarginPercent field to given value.

### HasMaintMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasMaintMarginPercent() bool`

HasMaintMarginPercent returns a boolean if a field has been set.

### GetMarginAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetMarketTakeBound

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMarketTakeBound() string`

GetMarketTakeBound returns the MarketTakeBound field if non-nil, zero value otherwise.

### GetMarketTakeBoundOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetMarketTakeBoundOk() (*string, bool)`

GetMarketTakeBoundOk returns a tuple with the MarketTakeBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTakeBound

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetMarketTakeBound(v string)`

SetMarketTakeBound sets MarketTakeBound field to given value.

### HasMarketTakeBound

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasMarketTakeBound() bool`

HasMarketTakeBound returns a boolean if a field has been set.

### GetOnboardDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetOnboardDate() int32`

GetOnboardDate returns the OnboardDate field if non-nil, zero value otherwise.

### GetOnboardDateOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetOnboardDateOk() (*int32, bool)`

GetOnboardDateOk returns a tuple with the OnboardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetOnboardDate(v int32)`

SetOnboardDate sets OnboardDate field to given value.

### HasOnboardDate

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasOnboardDate() bool`

HasOnboardDate returns a boolean if a field has been set.

### GetPair

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetPricePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetPricePrecision() int32`

GetPricePrecision returns the PricePrecision field if non-nil, zero value otherwise.

### GetPricePrecisionOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetPricePrecisionOk() (*int32, bool)`

GetPricePrecisionOk returns a tuple with the PricePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetPricePrecision(v int32)`

SetPricePrecision sets PricePrecision field to given value.

### HasPricePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasPricePrecision() bool`

HasPricePrecision returns a boolean if a field has been set.

### GetQuantityPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuantityPrecision() int32`

GetQuantityPrecision returns the QuantityPrecision field if non-nil, zero value otherwise.

### GetQuantityPrecisionOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuantityPrecisionOk() (*int32, bool)`

GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetQuantityPrecision(v int32)`

SetQuantityPrecision sets QuantityPrecision field to given value.

### HasQuantityPrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasQuantityPrecision() bool`

HasQuantityPrecision returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuotePrecision() int32`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetQuotePrecisionOk() (*int32, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetQuotePrecision(v int32)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetRequiredMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetRequiredMarginPercent() string`

GetRequiredMarginPercent returns the RequiredMarginPercent field if non-nil, zero value otherwise.

### GetRequiredMarginPercentOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetRequiredMarginPercentOk() (*string, bool)`

GetRequiredMarginPercentOk returns a tuple with the RequiredMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetRequiredMarginPercent(v string)`

SetRequiredMarginPercent sets RequiredMarginPercent field to given value.

### HasRequiredMarginPercent

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasRequiredMarginPercent() bool`

HasRequiredMarginPercent returns a boolean if a field has been set.

### GetSettlePlan

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetSettlePlan() int32`

GetSettlePlan returns the SettlePlan field if non-nil, zero value otherwise.

### GetSettlePlanOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetSettlePlanOk() (*int32, bool)`

GetSettlePlanOk returns a tuple with the SettlePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlePlan

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetSettlePlan(v int32)`

SetSettlePlan sets SettlePlan field to given value.

### HasSettlePlan

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasSettlePlan() bool`

HasSettlePlan returns a boolean if a field has been set.

### GetStatus

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetTimeInForce() []string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetTimeInForceOk() (*[]string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetTimeInForce(v []string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTriggerProtect

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetTriggerProtect() string`

GetTriggerProtect returns the TriggerProtect field if non-nil, zero value otherwise.

### GetTriggerProtectOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetTriggerProtectOk() (*string, bool)`

GetTriggerProtectOk returns a tuple with the TriggerProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerProtect

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetTriggerProtect(v string)`

SetTriggerProtect sets TriggerProtect field to given value.

### HasTriggerProtect

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasTriggerProtect() bool`

HasTriggerProtect returns a boolean if a field has been set.

### GetUnderlyingSubType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingSubType() []string`

GetUnderlyingSubType returns the UnderlyingSubType field if non-nil, zero value otherwise.

### GetUnderlyingSubTypeOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingSubTypeOk() (*[]string, bool)`

GetUnderlyingSubTypeOk returns a tuple with the UnderlyingSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSubType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetUnderlyingSubType(v []string)`

SetUnderlyingSubType sets UnderlyingSubType field to given value.

### HasUnderlyingSubType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasUnderlyingSubType() bool`

HasUnderlyingSubType returns a boolean if a field has been set.

### GetUnderlyingType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingType() string`

GetUnderlyingType returns the UnderlyingType field if non-nil, zero value otherwise.

### GetUnderlyingTypeOk

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) GetUnderlyingTypeOk() (*string, bool)`

GetUnderlyingTypeOk returns a tuple with the UnderlyingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) SetUnderlyingType(v string)`

SetUnderlyingType sets UnderlyingType field to given value.

### HasUnderlyingType

`func (o *UfuturesGetExchangeInfoV1RespSymbolsInner) HasUnderlyingType() bool`

HasUnderlyingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


