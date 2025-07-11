# GetTickerV1RespItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **string** |  | [optional] 
**CloseTime** | Pointer to **int64** |  | [optional] 
**ExercisePrice** | Pointer to **string** |  | [optional] 
**FirstTradeId** | Pointer to **int64** |  | [optional] 
**High** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**LastQty** | Pointer to **string** |  | [optional] 
**Low** | Pointer to **string** |  | [optional] 
**Open** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **int64** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TradeCount** | Pointer to **int32** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewGetTickerV1RespItem

`func NewGetTickerV1RespItem() *GetTickerV1RespItem`

NewGetTickerV1RespItem instantiates a new GetTickerV1RespItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickerV1RespItemWithDefaults

`func NewGetTickerV1RespItemWithDefaults() *GetTickerV1RespItem`

NewGetTickerV1RespItemWithDefaults instantiates a new GetTickerV1RespItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetTickerV1RespItem) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTickerV1RespItem) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTickerV1RespItem) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetTickerV1RespItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAskPrice

`func (o *GetTickerV1RespItem) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *GetTickerV1RespItem) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *GetTickerV1RespItem) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *GetTickerV1RespItem) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetBidPrice

`func (o *GetTickerV1RespItem) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *GetTickerV1RespItem) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *GetTickerV1RespItem) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *GetTickerV1RespItem) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetCloseTime

`func (o *GetTickerV1RespItem) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *GetTickerV1RespItem) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *GetTickerV1RespItem) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *GetTickerV1RespItem) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetExercisePrice

`func (o *GetTickerV1RespItem) GetExercisePrice() string`

GetExercisePrice returns the ExercisePrice field if non-nil, zero value otherwise.

### GetExercisePriceOk

`func (o *GetTickerV1RespItem) GetExercisePriceOk() (*string, bool)`

GetExercisePriceOk returns a tuple with the ExercisePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercisePrice

`func (o *GetTickerV1RespItem) SetExercisePrice(v string)`

SetExercisePrice sets ExercisePrice field to given value.

### HasExercisePrice

`func (o *GetTickerV1RespItem) HasExercisePrice() bool`

HasExercisePrice returns a boolean if a field has been set.

### GetFirstTradeId

`func (o *GetTickerV1RespItem) GetFirstTradeId() int64`

GetFirstTradeId returns the FirstTradeId field if non-nil, zero value otherwise.

### GetFirstTradeIdOk

`func (o *GetTickerV1RespItem) GetFirstTradeIdOk() (*int64, bool)`

GetFirstTradeIdOk returns a tuple with the FirstTradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTradeId

`func (o *GetTickerV1RespItem) SetFirstTradeId(v int64)`

SetFirstTradeId sets FirstTradeId field to given value.

### HasFirstTradeId

`func (o *GetTickerV1RespItem) HasFirstTradeId() bool`

HasFirstTradeId returns a boolean if a field has been set.

### GetHigh

`func (o *GetTickerV1RespItem) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *GetTickerV1RespItem) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *GetTickerV1RespItem) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *GetTickerV1RespItem) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetTickerV1RespItem) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetTickerV1RespItem) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetTickerV1RespItem) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetTickerV1RespItem) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastQty

`func (o *GetTickerV1RespItem) GetLastQty() string`

GetLastQty returns the LastQty field if non-nil, zero value otherwise.

### GetLastQtyOk

`func (o *GetTickerV1RespItem) GetLastQtyOk() (*string, bool)`

GetLastQtyOk returns a tuple with the LastQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQty

`func (o *GetTickerV1RespItem) SetLastQty(v string)`

SetLastQty sets LastQty field to given value.

### HasLastQty

`func (o *GetTickerV1RespItem) HasLastQty() bool`

HasLastQty returns a boolean if a field has been set.

### GetLow

`func (o *GetTickerV1RespItem) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *GetTickerV1RespItem) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *GetTickerV1RespItem) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *GetTickerV1RespItem) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetOpen

`func (o *GetTickerV1RespItem) GetOpen() string`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *GetTickerV1RespItem) GetOpenOk() (*string, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *GetTickerV1RespItem) SetOpen(v string)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *GetTickerV1RespItem) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetOpenTime

`func (o *GetTickerV1RespItem) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetTickerV1RespItem) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetTickerV1RespItem) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetTickerV1RespItem) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPriceChange

`func (o *GetTickerV1RespItem) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *GetTickerV1RespItem) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *GetTickerV1RespItem) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *GetTickerV1RespItem) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *GetTickerV1RespItem) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *GetTickerV1RespItem) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *GetTickerV1RespItem) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *GetTickerV1RespItem) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetStrikePrice

`func (o *GetTickerV1RespItem) GetStrikePrice() string`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *GetTickerV1RespItem) GetStrikePriceOk() (*string, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *GetTickerV1RespItem) SetStrikePrice(v string)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *GetTickerV1RespItem) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetSymbol

`func (o *GetTickerV1RespItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTickerV1RespItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTickerV1RespItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetTickerV1RespItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradeCount

`func (o *GetTickerV1RespItem) GetTradeCount() int32`

GetTradeCount returns the TradeCount field if non-nil, zero value otherwise.

### GetTradeCountOk

`func (o *GetTickerV1RespItem) GetTradeCountOk() (*int32, bool)`

GetTradeCountOk returns a tuple with the TradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCount

`func (o *GetTickerV1RespItem) SetTradeCount(v int32)`

SetTradeCount sets TradeCount field to given value.

### HasTradeCount

`func (o *GetTickerV1RespItem) HasTradeCount() bool`

HasTradeCount returns a boolean if a field has been set.

### GetVolume

`func (o *GetTickerV1RespItem) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetTickerV1RespItem) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetTickerV1RespItem) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetTickerV1RespItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


