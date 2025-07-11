# GetMiningStatisticsUserStatusV1RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algo** | Pointer to **string** |  | [optional] 
**DayHashRate** | Pointer to **string** |  | [optional] 
**FifteenMinHashRate** | Pointer to **string** |  | [optional] 
**InvalidNum** | Pointer to **int32** |  | [optional] 
**ProfitToday** | Pointer to [**GetMiningStatisticsUserStatusV1RespDataProfitToday**](GetMiningStatisticsUserStatusV1RespDataProfitToday.md) |  | [optional] 
**ProfitYesterday** | Pointer to [**GetMiningStatisticsUserStatusV1RespDataProfitToday**](GetMiningStatisticsUserStatusV1RespDataProfitToday.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**ValidNum** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMiningStatisticsUserStatusV1RespData

`func NewGetMiningStatisticsUserStatusV1RespData() *GetMiningStatisticsUserStatusV1RespData`

NewGetMiningStatisticsUserStatusV1RespData instantiates a new GetMiningStatisticsUserStatusV1RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMiningStatisticsUserStatusV1RespDataWithDefaults

`func NewGetMiningStatisticsUserStatusV1RespDataWithDefaults() *GetMiningStatisticsUserStatusV1RespData`

NewGetMiningStatisticsUserStatusV1RespDataWithDefaults instantiates a new GetMiningStatisticsUserStatusV1RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgo

`func (o *GetMiningStatisticsUserStatusV1RespData) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *GetMiningStatisticsUserStatusV1RespData) SetAlgo(v string)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *GetMiningStatisticsUserStatusV1RespData) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.

### GetDayHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) GetDayHashRate() string`

GetDayHashRate returns the DayHashRate field if non-nil, zero value otherwise.

### GetDayHashRateOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetDayHashRateOk() (*string, bool)`

GetDayHashRateOk returns a tuple with the DayHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) SetDayHashRate(v string)`

SetDayHashRate sets DayHashRate field to given value.

### HasDayHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) HasDayHashRate() bool`

HasDayHashRate returns a boolean if a field has been set.

### GetFifteenMinHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) GetFifteenMinHashRate() string`

GetFifteenMinHashRate returns the FifteenMinHashRate field if non-nil, zero value otherwise.

### GetFifteenMinHashRateOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetFifteenMinHashRateOk() (*string, bool)`

GetFifteenMinHashRateOk returns a tuple with the FifteenMinHashRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifteenMinHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) SetFifteenMinHashRate(v string)`

SetFifteenMinHashRate sets FifteenMinHashRate field to given value.

### HasFifteenMinHashRate

`func (o *GetMiningStatisticsUserStatusV1RespData) HasFifteenMinHashRate() bool`

HasFifteenMinHashRate returns a boolean if a field has been set.

### GetInvalidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) GetInvalidNum() int32`

GetInvalidNum returns the InvalidNum field if non-nil, zero value otherwise.

### GetInvalidNumOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetInvalidNumOk() (*int32, bool)`

GetInvalidNumOk returns a tuple with the InvalidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) SetInvalidNum(v int32)`

SetInvalidNum sets InvalidNum field to given value.

### HasInvalidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) HasInvalidNum() bool`

HasInvalidNum returns a boolean if a field has been set.

### GetProfitToday

`func (o *GetMiningStatisticsUserStatusV1RespData) GetProfitToday() GetMiningStatisticsUserStatusV1RespDataProfitToday`

GetProfitToday returns the ProfitToday field if non-nil, zero value otherwise.

### GetProfitTodayOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetProfitTodayOk() (*GetMiningStatisticsUserStatusV1RespDataProfitToday, bool)`

GetProfitTodayOk returns a tuple with the ProfitToday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitToday

`func (o *GetMiningStatisticsUserStatusV1RespData) SetProfitToday(v GetMiningStatisticsUserStatusV1RespDataProfitToday)`

SetProfitToday sets ProfitToday field to given value.

### HasProfitToday

`func (o *GetMiningStatisticsUserStatusV1RespData) HasProfitToday() bool`

HasProfitToday returns a boolean if a field has been set.

### GetProfitYesterday

`func (o *GetMiningStatisticsUserStatusV1RespData) GetProfitYesterday() GetMiningStatisticsUserStatusV1RespDataProfitToday`

GetProfitYesterday returns the ProfitYesterday field if non-nil, zero value otherwise.

### GetProfitYesterdayOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetProfitYesterdayOk() (*GetMiningStatisticsUserStatusV1RespDataProfitToday, bool)`

GetProfitYesterdayOk returns a tuple with the ProfitYesterday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitYesterday

`func (o *GetMiningStatisticsUserStatusV1RespData) SetProfitYesterday(v GetMiningStatisticsUserStatusV1RespDataProfitToday)`

SetProfitYesterday sets ProfitYesterday field to given value.

### HasProfitYesterday

`func (o *GetMiningStatisticsUserStatusV1RespData) HasProfitYesterday() bool`

HasProfitYesterday returns a boolean if a field has been set.

### GetUnit

`func (o *GetMiningStatisticsUserStatusV1RespData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetMiningStatisticsUserStatusV1RespData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetMiningStatisticsUserStatusV1RespData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUserName

`func (o *GetMiningStatisticsUserStatusV1RespData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetMiningStatisticsUserStatusV1RespData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetMiningStatisticsUserStatusV1RespData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetValidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) GetValidNum() int32`

GetValidNum returns the ValidNum field if non-nil, zero value otherwise.

### GetValidNumOk

`func (o *GetMiningStatisticsUserStatusV1RespData) GetValidNumOk() (*int32, bool)`

GetValidNumOk returns a tuple with the ValidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) SetValidNum(v int32)`

SetValidNum sets ValidNum field to given value.

### HasValidNum

`func (o *GetMiningStatisticsUserStatusV1RespData) HasValidNum() bool`

HasValidNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


