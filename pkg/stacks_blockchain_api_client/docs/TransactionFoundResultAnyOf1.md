# TransactionFoundResultAnyOf1

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
**TokenTransfer** | [**TokenTransferTransactionTokenTransfer**](TokenTransferTransactionTokenTransfer.md) |  | 
**SmartContract** | [**SmartContractTransactionSmartContract**](SmartContractTransactionSmartContract.md) |  | 
**ContractCall** | [**ContractCallTransactionContractCall**](ContractCallTransactionContractCall.md) |  | 
**PoisonMicroblock** | [**PoisonMicroblockTransactionPoisonMicroblock**](PoisonMicroblockTransactionPoisonMicroblock.md) |  | 
**CoinbasePayload** | [**CoinbaseTransactionCoinbasePayload**](CoinbaseTransactionCoinbasePayload.md) |  | 
**TenureChangePayload** | [**TenureChangeTransaction1TenureChangePayload**](TenureChangeTransaction1TenureChangePayload.md) |  | 

## Methods

### NewTransactionFoundResultAnyOf1

`func NewTransactionFoundResultAnyOf1(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransaction1PostConditionMode, postConditions []TokenTransferTransaction1PostConditionsInner, anchorMode TokenTransferTransaction1AnchorMode, txStatus TokenTransferMempoolTransactionTxStatus, receiptTime int32, receiptTimeIso string, txType string, tokenTransfer TokenTransferTransactionTokenTransfer, smartContract SmartContractTransactionSmartContract, contractCall ContractCallTransactionContractCall, poisonMicroblock PoisonMicroblockTransactionPoisonMicroblock, coinbasePayload CoinbaseTransactionCoinbasePayload, tenureChangePayload TenureChangeTransaction1TenureChangePayload, ) *TransactionFoundResultAnyOf1`

NewTransactionFoundResultAnyOf1 instantiates a new TransactionFoundResultAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFoundResultAnyOf1WithDefaults

`func NewTransactionFoundResultAnyOf1WithDefaults() *TransactionFoundResultAnyOf1`

NewTransactionFoundResultAnyOf1WithDefaults instantiates a new TransactionFoundResultAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *TransactionFoundResultAnyOf1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionFoundResultAnyOf1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionFoundResultAnyOf1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *TransactionFoundResultAnyOf1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransactionFoundResultAnyOf1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransactionFoundResultAnyOf1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *TransactionFoundResultAnyOf1) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *TransactionFoundResultAnyOf1) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *TransactionFoundResultAnyOf1) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *TransactionFoundResultAnyOf1) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *TransactionFoundResultAnyOf1) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *TransactionFoundResultAnyOf1) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *TransactionFoundResultAnyOf1) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *TransactionFoundResultAnyOf1) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *TransactionFoundResultAnyOf1) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *TransactionFoundResultAnyOf1) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *TransactionFoundResultAnyOf1) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *TransactionFoundResultAnyOf1) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *TransactionFoundResultAnyOf1) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *TransactionFoundResultAnyOf1) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *TransactionFoundResultAnyOf1) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *TransactionFoundResultAnyOf1) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *TransactionFoundResultAnyOf1) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *TransactionFoundResultAnyOf1) GetPostConditionMode() TokenTransferTransaction1PostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *TransactionFoundResultAnyOf1) GetPostConditionModeOk() (*TokenTransferTransaction1PostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *TransactionFoundResultAnyOf1) SetPostConditionMode(v TokenTransferTransaction1PostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *TransactionFoundResultAnyOf1) GetPostConditions() []TokenTransferTransaction1PostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *TransactionFoundResultAnyOf1) GetPostConditionsOk() (*[]TokenTransferTransaction1PostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *TransactionFoundResultAnyOf1) SetPostConditions(v []TokenTransferTransaction1PostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *TransactionFoundResultAnyOf1) GetAnchorMode() TokenTransferTransaction1AnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *TransactionFoundResultAnyOf1) GetAnchorModeOk() (*TokenTransferTransaction1AnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *TransactionFoundResultAnyOf1) SetAnchorMode(v TokenTransferTransaction1AnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetTxStatus

`func (o *TransactionFoundResultAnyOf1) GetTxStatus() TokenTransferMempoolTransactionTxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *TransactionFoundResultAnyOf1) GetTxStatusOk() (*TokenTransferMempoolTransactionTxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *TransactionFoundResultAnyOf1) SetTxStatus(v TokenTransferMempoolTransactionTxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetReceiptTime

`func (o *TransactionFoundResultAnyOf1) GetReceiptTime() int32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *TransactionFoundResultAnyOf1) GetReceiptTimeOk() (*int32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *TransactionFoundResultAnyOf1) SetReceiptTime(v int32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *TransactionFoundResultAnyOf1) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *TransactionFoundResultAnyOf1) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *TransactionFoundResultAnyOf1) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.


### GetTxType

`func (o *TransactionFoundResultAnyOf1) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TransactionFoundResultAnyOf1) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TransactionFoundResultAnyOf1) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *TransactionFoundResultAnyOf1) GetTokenTransfer() TokenTransferTransactionTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *TransactionFoundResultAnyOf1) GetTokenTransferOk() (*TokenTransferTransactionTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *TransactionFoundResultAnyOf1) SetTokenTransfer(v TokenTransferTransactionTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.


### GetSmartContract

`func (o *TransactionFoundResultAnyOf1) GetSmartContract() SmartContractTransactionSmartContract`

GetSmartContract returns the SmartContract field if non-nil, zero value otherwise.

### GetSmartContractOk

`func (o *TransactionFoundResultAnyOf1) GetSmartContractOk() (*SmartContractTransactionSmartContract, bool)`

GetSmartContractOk returns a tuple with the SmartContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContract

`func (o *TransactionFoundResultAnyOf1) SetSmartContract(v SmartContractTransactionSmartContract)`

SetSmartContract sets SmartContract field to given value.


### GetContractCall

`func (o *TransactionFoundResultAnyOf1) GetContractCall() ContractCallTransactionContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *TransactionFoundResultAnyOf1) GetContractCallOk() (*ContractCallTransactionContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *TransactionFoundResultAnyOf1) SetContractCall(v ContractCallTransactionContractCall)`

SetContractCall sets ContractCall field to given value.


### GetPoisonMicroblock

`func (o *TransactionFoundResultAnyOf1) GetPoisonMicroblock() PoisonMicroblockTransactionPoisonMicroblock`

GetPoisonMicroblock returns the PoisonMicroblock field if non-nil, zero value otherwise.

### GetPoisonMicroblockOk

`func (o *TransactionFoundResultAnyOf1) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionPoisonMicroblock, bool)`

GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoisonMicroblock

`func (o *TransactionFoundResultAnyOf1) SetPoisonMicroblock(v PoisonMicroblockTransactionPoisonMicroblock)`

SetPoisonMicroblock sets PoisonMicroblock field to given value.


### GetCoinbasePayload

`func (o *TransactionFoundResultAnyOf1) GetCoinbasePayload() CoinbaseTransactionCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *TransactionFoundResultAnyOf1) GetCoinbasePayloadOk() (*CoinbaseTransactionCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *TransactionFoundResultAnyOf1) SetCoinbasePayload(v CoinbaseTransactionCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.


### GetTenureChangePayload

`func (o *TransactionFoundResultAnyOf1) GetTenureChangePayload() TenureChangeTransaction1TenureChangePayload`

GetTenureChangePayload returns the TenureChangePayload field if non-nil, zero value otherwise.

### GetTenureChangePayloadOk

`func (o *TransactionFoundResultAnyOf1) GetTenureChangePayloadOk() (*TenureChangeTransaction1TenureChangePayload, bool)`

GetTenureChangePayloadOk returns a tuple with the TenureChangePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenureChangePayload

`func (o *TransactionFoundResultAnyOf1) SetTenureChangePayload(v TenureChangeTransaction1TenureChangePayload)`

SetTenureChangePayload sets TenureChangePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


