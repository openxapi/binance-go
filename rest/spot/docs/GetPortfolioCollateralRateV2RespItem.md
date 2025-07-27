# GetPortfolioCollateralRateV2RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**CollateralInfo** | Pointer to [**[]GetPortfolioCollateralRateV2RespItemCollateralInfoInner**](GetPortfolioCollateralRateV2RespItemCollateralInfoInner.md) |  | [optional] 

## Methods

### NewGetPortfolioCollateralRateV2RespItem

`func NewGetPortfolioCollateralRateV2RespItem() *GetPortfolioCollateralRateV2RespItem`

NewGetPortfolioCollateralRateV2RespItem instantiates a new GetPortfolioCollateralRateV2RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortfolioCollateralRateV2RespItemWithDefaults

`func NewGetPortfolioCollateralRateV2RespItemWithDefaults() *GetPortfolioCollateralRateV2RespItem`

NewGetPortfolioCollateralRateV2RespItemWithDefaults instantiates a new GetPortfolioCollateralRateV2RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetPortfolioCollateralRateV2RespItem) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetPortfolioCollateralRateV2RespItem) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetPortfolioCollateralRateV2RespItem) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetPortfolioCollateralRateV2RespItem) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetCollateralInfo

`func (o *GetPortfolioCollateralRateV2RespItem) GetCollateralInfo() []GetPortfolioCollateralRateV2RespItemCollateralInfoInner`

GetCollateralInfo returns the CollateralInfo field if non-nil, zero value otherwise.

### GetCollateralInfoOk

`func (o *GetPortfolioCollateralRateV2RespItem) GetCollateralInfoOk() (*[]GetPortfolioCollateralRateV2RespItemCollateralInfoInner, bool)`

GetCollateralInfoOk returns a tuple with the CollateralInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralInfo

`func (o *GetPortfolioCollateralRateV2RespItem) SetCollateralInfo(v []GetPortfolioCollateralRateV2RespItemCollateralInfoInner)`

SetCollateralInfo sets CollateralInfo field to given value.

### HasCollateralInfo

`func (o *GetPortfolioCollateralRateV2RespItem) HasCollateralInfo() bool`

HasCollateralInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


