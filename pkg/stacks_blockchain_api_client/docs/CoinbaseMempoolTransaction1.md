# CoinbaseMempoolTransaction1

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
**PostConditionMode** | [**TokenTransferTransactionPostConditionMode**](TokenTransferTransactionPostConditionMode.md) |  | 
**PostConditions** | [**[]TokenTransferTransactionPostConditionsInner**](TokenTransferTransactionPostConditionsInner.md) |  | 
**AnchorMode** | [**TokenTransferTransactionAnchorMode**](TokenTransferTransactionAnchorMode.md) |  | 
**TxStatus** | [**TokenTransferMempoolTransaction1TxStatus**](TokenTransferMempoolTransaction1TxStatus.md) |  | 
**ReceiptTime** | **int32** | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. | 
**ReceiptTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. | 
**TxType** | **string** |  | 
**CoinbasePayload** | [**CoinbaseTransactionCoinbasePayload**](CoinbaseTransactionCoinbasePayload.md) |  | 

## Methods

### NewCoinbaseMempoolTransaction1

`func NewCoinbaseMempoolTransaction1(txId string, nonce int32, feeRate string, senderAddress string, sponsored bool, postConditionMode TokenTransferTransactionPostConditionMode, postConditions []TokenTransferTransactionPostConditionsInner, anchorMode TokenTransferTransactionAnchorMode, txStatus TokenTransferMempoolTransaction1TxStatus, receiptTime int32, receiptTimeIso string, txType string, coinbasePayload CoinbaseTransactionCoinbasePayload, ) *CoinbaseMempoolTransaction1`

NewCoinbaseMempoolTransaction1 instantiates a new CoinbaseMempoolTransaction1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinbaseMempoolTransaction1WithDefaults

`func NewCoinbaseMempoolTransaction1WithDefaults() *CoinbaseMempoolTransaction1`

NewCoinbaseMempoolTransaction1WithDefaults instantiates a new CoinbaseMempoolTransaction1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *CoinbaseMempoolTransaction1) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *CoinbaseMempoolTransaction1) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *CoinbaseMempoolTransaction1) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetNonce

`func (o *CoinbaseMempoolTransaction1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CoinbaseMempoolTransaction1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CoinbaseMempoolTransaction1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetFeeRate

`func (o *CoinbaseMempoolTransaction1) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *CoinbaseMempoolTransaction1) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *CoinbaseMempoolTransaction1) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.


### GetSenderAddress

`func (o *CoinbaseMempoolTransaction1) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *CoinbaseMempoolTransaction1) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *CoinbaseMempoolTransaction1) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.


### GetSponsorNonce

`func (o *CoinbaseMempoolTransaction1) GetSponsorNonce() int32`

GetSponsorNonce returns the SponsorNonce field if non-nil, zero value otherwise.

### GetSponsorNonceOk

`func (o *CoinbaseMempoolTransaction1) GetSponsorNonceOk() (*int32, bool)`

GetSponsorNonceOk returns a tuple with the SponsorNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNonce

`func (o *CoinbaseMempoolTransaction1) SetSponsorNonce(v int32)`

SetSponsorNonce sets SponsorNonce field to given value.

### HasSponsorNonce

`func (o *CoinbaseMempoolTransaction1) HasSponsorNonce() bool`

HasSponsorNonce returns a boolean if a field has been set.

### GetSponsored

`func (o *CoinbaseMempoolTransaction1) GetSponsored() bool`

GetSponsored returns the Sponsored field if non-nil, zero value otherwise.

### GetSponsoredOk

`func (o *CoinbaseMempoolTransaction1) GetSponsoredOk() (*bool, bool)`

GetSponsoredOk returns a tuple with the Sponsored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsored

`func (o *CoinbaseMempoolTransaction1) SetSponsored(v bool)`

SetSponsored sets Sponsored field to given value.


### GetSponsorAddress

`func (o *CoinbaseMempoolTransaction1) GetSponsorAddress() string`

GetSponsorAddress returns the SponsorAddress field if non-nil, zero value otherwise.

### GetSponsorAddressOk

`func (o *CoinbaseMempoolTransaction1) GetSponsorAddressOk() (*string, bool)`

GetSponsorAddressOk returns a tuple with the SponsorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorAddress

`func (o *CoinbaseMempoolTransaction1) SetSponsorAddress(v string)`

SetSponsorAddress sets SponsorAddress field to given value.

### HasSponsorAddress

`func (o *CoinbaseMempoolTransaction1) HasSponsorAddress() bool`

HasSponsorAddress returns a boolean if a field has been set.

### GetPostConditionMode

`func (o *CoinbaseMempoolTransaction1) GetPostConditionMode() TokenTransferTransactionPostConditionMode`

GetPostConditionMode returns the PostConditionMode field if non-nil, zero value otherwise.

### GetPostConditionModeOk

`func (o *CoinbaseMempoolTransaction1) GetPostConditionModeOk() (*TokenTransferTransactionPostConditionMode, bool)`

GetPostConditionModeOk returns a tuple with the PostConditionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditionMode

`func (o *CoinbaseMempoolTransaction1) SetPostConditionMode(v TokenTransferTransactionPostConditionMode)`

SetPostConditionMode sets PostConditionMode field to given value.


### GetPostConditions

`func (o *CoinbaseMempoolTransaction1) GetPostConditions() []TokenTransferTransactionPostConditionsInner`

GetPostConditions returns the PostConditions field if non-nil, zero value otherwise.

### GetPostConditionsOk

`func (o *CoinbaseMempoolTransaction1) GetPostConditionsOk() (*[]TokenTransferTransactionPostConditionsInner, bool)`

GetPostConditionsOk returns a tuple with the PostConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostConditions

`func (o *CoinbaseMempoolTransaction1) SetPostConditions(v []TokenTransferTransactionPostConditionsInner)`

SetPostConditions sets PostConditions field to given value.


### GetAnchorMode

`func (o *CoinbaseMempoolTransaction1) GetAnchorMode() TokenTransferTransactionAnchorMode`

GetAnchorMode returns the AnchorMode field if non-nil, zero value otherwise.

### GetAnchorModeOk

`func (o *CoinbaseMempoolTransaction1) GetAnchorModeOk() (*TokenTransferTransactionAnchorMode, bool)`

GetAnchorModeOk returns a tuple with the AnchorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorMode

`func (o *CoinbaseMempoolTransaction1) SetAnchorMode(v TokenTransferTransactionAnchorMode)`

SetAnchorMode sets AnchorMode field to given value.


### GetTxStatus

`func (o *CoinbaseMempoolTransaction1) GetTxStatus() TokenTransferMempoolTransaction1TxStatus`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *CoinbaseMempoolTransaction1) GetTxStatusOk() (*TokenTransferMempoolTransaction1TxStatus, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *CoinbaseMempoolTransaction1) SetTxStatus(v TokenTransferMempoolTransaction1TxStatus)`

SetTxStatus sets TxStatus field to given value.


### GetReceiptTime

`func (o *CoinbaseMempoolTransaction1) GetReceiptTime() int32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *CoinbaseMempoolTransaction1) GetReceiptTimeOk() (*int32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *CoinbaseMempoolTransaction1) SetReceiptTime(v int32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *CoinbaseMempoolTransaction1) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *CoinbaseMempoolTransaction1) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *CoinbaseMempoolTransaction1) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.


### GetTxType

`func (o *CoinbaseMempoolTransaction1) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *CoinbaseMempoolTransaction1) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *CoinbaseMempoolTransaction1) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetCoinbasePayload

`func (o *CoinbaseMempoolTransaction1) GetCoinbasePayload() CoinbaseTransactionCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *CoinbaseMempoolTransaction1) GetCoinbasePayloadOk() (*CoinbaseTransactionCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *CoinbaseMempoolTransaction1) SetCoinbasePayload(v CoinbaseTransactionCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


