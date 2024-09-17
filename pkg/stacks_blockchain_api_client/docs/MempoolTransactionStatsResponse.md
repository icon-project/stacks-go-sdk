# MempoolTransactionStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxTypeCounts** | **map[string]int32** | Number of tranasction in the mempool, broken down by transaction type. | 
**TxSimpleFeeAverages** | [**map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue**](MempoolTransactionStatsResponseTxSimpleFeeAveragesValue.md) | The simple mean (average) transaction fee, broken down by transaction type. Note that this does not factor in actual execution costs. The average fee is not a reliable metric for calculating a fee for a new transaction. | 
**TxAges** | [**map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue**](MempoolTransactionStatsResponseTxSimpleFeeAveragesValue.md) | The average time (in blocks) that transactions have lived in the mempool. The start block height is simply the current chain-tip of when the attached Stacks node receives the transaction. This timing can be different across Stacks nodes / API instances due to propagation timing differences in the p2p network. | 
**TxByteSizes** | [**map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue**](MempoolTransactionStatsResponseTxSimpleFeeAveragesValue.md) | The average byte size of transactions in the mempool, broken down by transaction type. | 

## Methods

### NewMempoolTransactionStatsResponse

`func NewMempoolTransactionStatsResponse(txTypeCounts map[string]int32, txSimpleFeeAverages map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, txAges map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, txByteSizes map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, ) *MempoolTransactionStatsResponse`

NewMempoolTransactionStatsResponse instantiates a new MempoolTransactionStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTransactionStatsResponseWithDefaults

`func NewMempoolTransactionStatsResponseWithDefaults() *MempoolTransactionStatsResponse`

NewMempoolTransactionStatsResponseWithDefaults instantiates a new MempoolTransactionStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxTypeCounts

`func (o *MempoolTransactionStatsResponse) GetTxTypeCounts() map[string]int32`

GetTxTypeCounts returns the TxTypeCounts field if non-nil, zero value otherwise.

### GetTxTypeCountsOk

`func (o *MempoolTransactionStatsResponse) GetTxTypeCountsOk() (*map[string]int32, bool)`

GetTxTypeCountsOk returns a tuple with the TxTypeCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTypeCounts

`func (o *MempoolTransactionStatsResponse) SetTxTypeCounts(v map[string]int32)`

SetTxTypeCounts sets TxTypeCounts field to given value.


### GetTxSimpleFeeAverages

`func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAverages() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue`

GetTxSimpleFeeAverages returns the TxSimpleFeeAverages field if non-nil, zero value otherwise.

### GetTxSimpleFeeAveragesOk

`func (o *MempoolTransactionStatsResponse) GetTxSimpleFeeAveragesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool)`

GetTxSimpleFeeAveragesOk returns a tuple with the TxSimpleFeeAverages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSimpleFeeAverages

`func (o *MempoolTransactionStatsResponse) SetTxSimpleFeeAverages(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue)`

SetTxSimpleFeeAverages sets TxSimpleFeeAverages field to given value.


### GetTxAges

`func (o *MempoolTransactionStatsResponse) GetTxAges() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue`

GetTxAges returns the TxAges field if non-nil, zero value otherwise.

### GetTxAgesOk

`func (o *MempoolTransactionStatsResponse) GetTxAgesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool)`

GetTxAgesOk returns a tuple with the TxAges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAges

`func (o *MempoolTransactionStatsResponse) SetTxAges(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue)`

SetTxAges sets TxAges field to given value.


### GetTxByteSizes

`func (o *MempoolTransactionStatsResponse) GetTxByteSizes() map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue`

GetTxByteSizes returns the TxByteSizes field if non-nil, zero value otherwise.

### GetTxByteSizesOk

`func (o *MempoolTransactionStatsResponse) GetTxByteSizesOk() (*map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue, bool)`

GetTxByteSizesOk returns a tuple with the TxByteSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteSizes

`func (o *MempoolTransactionStatsResponse) SetTxByteSizes(v map[string]MempoolTransactionStatsResponseTxSimpleFeeAveragesValue)`

SetTxByteSizes sets TxByteSizes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


