# ContractCallMempoolTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** | Transaction ID | 
**Nonce** | **int32** | Used for ordering the transactions originating from and paying from an account. The nonce ensures that a transaction is processed at most once. The nonce counts the number of times an account&#39;s owner(s) have authorized a transaction. The first transaction from an account will have a nonce value equal to 0, the second will have a nonce value equal to 1, and so on. | 
**FeeRate** | **string** | Transaction fee as Integer string (64-bit unsigned integer). | 
**SenderAddress** | **string** | Address of the transaction initiator | 
**SponsorNonce** | Pointer to **int32** |  | [optional] 
**Sponsored** | **bool** | Denotes whether the originating account is the same as the paying account | 
**SponsorAddress** | Pointer to **string** |  | [optional] 
**PostConditionMode** | [**TokenTransferTransaction1PostConditionMode**](TokenTransferTransaction1PostConditionMode.md) |  | 
**PostConditions** | [**[]TokenTransferTransaction1PostConditionsInner**](TokenTransferTransaction1PostConditionsInner.md) |  | 
**AnchorMode** | [**TokenTransferTransaction1AnchorMode**](TokenTransferTransaction1AnchorMode.md) |  | 
**TxStatus** | [**TokenTransferMempoolTransactionTxStatus**](TokenTransferMempoolTransactionTxStatus.md) |  | 
**ReceiptTime** | **int32** | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. | 
**ReceiptTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. | 
**TxType** | **string** |  | 
**ContractCall** | [**ContractCallTransactionContractCall**](ContractCallTransactionContractCall.md) |  | 

## Methods

### NewContractCallMempoolTransaction

`func NewContractCallMempoolTransaction(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransaction1PostConditionMode, postConditions []TokenTransferTransaction1PostConditionsInner, anchorMode TokenTransferTransaction1AnchorMode, txStatus TokenTransferMempoolTransactionTxStatus, receiptTime int32, receiptTimeIso string, txType string, contractCall ContractCallTransactionContractCall, ) *ContractCallMempoolTransaction`

NewContractCallMempoolTransaction instantiates a new ContractCallMempoolTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallMempoolTransactionWithDefaults

`func NewContractCallMempoolTransactionWithDefaults() *ContractCallMempoolTransaction`

NewContractCallMempoolTransactionWithDefaults instantiates a new ContractCallMempoolTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *ContractCallMempoolTransaction) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ContractCallMempoolTransaction) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ContractCallMempoolTransaction) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *ContractCallMempoolTransaction) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ContractCallMempoolTransaction) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ContractCallMempoolTransaction) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *ContractCallMempoolTransaction) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *ContractCallMempoolTransaction) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *ContractCallMempoolTransaction) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *ContractCallMempoolTransaction) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *ContractCallMempoolTransaction) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *ContractCallMempoolTransaction) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *ContractCallMempoolTransaction) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *ContractCallMempoolTransaction) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *ContractCallMempoolTransaction) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *ContractCallMempoolTransaction) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *ContractCallMempoolTransaction) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *ContractCallMempoolTransaction) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *ContractCallMempoolTransaction) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *ContractCallMempoolTransaction) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *ContractCallMempoolTransaction) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *ContractCallMempoolTransaction) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *ContractCallMempoolTransaction) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *ContractCallMempoolTransaction) GetPostConditionMode() TokenTransferTransaction1PostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *ContractCallMempoolTransaction) GetPostConditionModeOk() (*TokenTransferTransaction1PostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *ContractCallMempoolTransaction) SetPostConditionMode(v TokenTransferTransaction1PostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *ContractCallMempoolTransaction) GetPostConditions() []TokenTransferTransaction1PostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *ContractCallMempoolTransaction) GetPostConditionsOk() (*[]TokenTransferTransaction1PostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *ContractCallMempoolTransaction) SetPostConditions(v []TokenTransferTransaction1PostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *ContractCallMempoolTransaction) GetAnchorMode() TokenTransferTransaction1AnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *ContractCallMempoolTransaction) GetAnchorModeOk() (*TokenTransferTransaction1AnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *ContractCallMempoolTransaction) SetAnchorMode(v TokenTransferTransaction1AnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetTxStatus

`func (o *ContractCallMempoolTransaction) GetTxStatus() TokenTransferMempoolTransactionTxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *ContractCallMempoolTransaction) GetTxStatusOk() (*TokenTransferMempoolTransactionTxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *ContractCallMempoolTransaction) SetTxStatus(v TokenTransferMempoolTransactionTxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetReceiptTime

`func (o *ContractCallMempoolTransaction) GetReceiptTime() int32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *ContractCallMempoolTransaction) GetReceiptTimeOk() (*int32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *ContractCallMempoolTransaction) SetReceiptTime(v int32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *ContractCallMempoolTransaction) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *ContractCallMempoolTransaction) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *ContractCallMempoolTransaction) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.


### GetTxType

`func (o *ContractCallMempoolTransaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *ContractCallMempoolTransaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *ContractCallMempoolTransaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetContractCall

`func (o *ContractCallMempoolTransaction) GetContractCall() ContractCallTransactionContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *ContractCallMempoolTransaction) GetContractCallOk() (*ContractCallTransactionContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *ContractCallMempoolTransaction) SetContractCall(v ContractCallTransactionContractCall)`

SetContractCall sets ContractCall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


