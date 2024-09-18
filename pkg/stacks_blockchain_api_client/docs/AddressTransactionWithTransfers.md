# AddressTransactionWithTransfers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**GetTransactionList200ResponseResultsInner**](GetTransactionList200ResponseResultsInner.md) |  | 
**StxSent** | **string** | Total sent from the given address, including the tx fee, in micro-STX as an integer string. | 
**StxReceived** | **string** | Total received by the given address in micro-STX as an integer string. | 
**StxTransfers** | [**[]AddressTransactionWithTransfersStxTransfersInner**](AddressTransactionWithTransfersStxTransfersInner.md) |  | 
**FtTransfers** | Pointer to [**[]AddressTransactionWithTransfersFtTransfersInner**](AddressTransactionWithTransfersFtTransfersInner.md) |  | [optional] 
**NftTransfers** | Pointer to [**[]AddressTransactionWithTransfersNftTransfersInner**](AddressTransactionWithTransfersNftTransfersInner.md) |  | [optional] 

## Methods

### NewAddressTransactionWithTransfers

`func NewAddressTransactionWithTransfers(tx GetTransactionList200ResponseResultsInner, stxSent string, stxReceived string, stxTransfers []AddressTransactionWithTransfersStxTransfersInner, ) *AddressTransactionWithTransfers`

NewAddressTransactionWithTransfers instantiates a new AddressTransactionWithTransfers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersWithDefaults

`func NewAddressTransactionWithTransfersWithDefaults() *AddressTransactionWithTransfers`

NewAddressTransactionWithTransfersWithDefaults instantiates a new AddressTransactionWithTransfers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *AddressTransactionWithTransfers) GetTx() GetTransactionList200ResponseResultsInner`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *AddressTransactionWithTransfers) GetTxOk() (*GetTransactionList200ResponseResultsInner, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *AddressTransactionWithTransfers) SetTx(v GetTransactionList200ResponseResultsInner)`

SetTx sets Tx field to given value.


### GetStxSent

`func (o *AddressTransactionWithTransfers) GetStxSent() string`

GetStxSent returns the StxSent field if non-nil, zero value otherwise.

### GetStxSentOk

`func (o *AddressTransactionWithTransfers) GetStxSentOk() (*string, bool)`

GetStxSentOk returns a tuple with the StxSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxSent

`func (o *AddressTransactionWithTransfers) SetStxSent(v string)`

SetStxSent sets StxSent field to given value.


### GetStxReceived

`func (o *AddressTransactionWithTransfers) GetStxReceived() string`

GetStxReceived returns the StxReceived field if non-nil, zero value otherwise.

### GetStxReceivedOk

`func (o *AddressTransactionWithTransfers) GetStxReceivedOk() (*string, bool)`

GetStxReceivedOk returns a tuple with the StxReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxReceived

`func (o *AddressTransactionWithTransfers) SetStxReceived(v string)`

SetStxReceived sets StxReceived field to given value.


### GetStxTransfers

`func (o *AddressTransactionWithTransfers) GetStxTransfers() []AddressTransactionWithTransfersStxTransfersInner`

GetStxTransfers returns the StxTransfers field if non-nil, zero value otherwise.

### GetStxTransfersOk

`func (o *AddressTransactionWithTransfers) GetStxTransfersOk() (*[]AddressTransactionWithTransfersStxTransfersInner, bool)`

GetStxTransfersOk returns a tuple with the StxTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxTransfers

`func (o *AddressTransactionWithTransfers) SetStxTransfers(v []AddressTransactionWithTransfersStxTransfersInner)`

SetStxTransfers sets StxTransfers field to given value.


### GetFtTransfers

`func (o *AddressTransactionWithTransfers) GetFtTransfers() []AddressTransactionWithTransfersFtTransfersInner`

GetFtTransfers returns the FtTransfers field if non-nil, zero value otherwise.

### GetFtTransfersOk

`func (o *AddressTransactionWithTransfers) GetFtTransfersOk() (*[]AddressTransactionWithTransfersFtTransfersInner, bool)`

GetFtTransfersOk returns a tuple with the FtTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtTransfers

`func (o *AddressTransactionWithTransfers) SetFtTransfers(v []AddressTransactionWithTransfersFtTransfersInner)`

SetFtTransfers sets FtTransfers field to given value.

### HasFtTransfers

`func (o *AddressTransactionWithTransfers) HasFtTransfers() bool`

HasFtTransfers returns a boolean if a field has been set.

### GetNftTransfers

`func (o *AddressTransactionWithTransfers) GetNftTransfers() []AddressTransactionWithTransfersNftTransfersInner`

GetNftTransfers returns the NftTransfers field if non-nil, zero value otherwise.

### GetNftTransfersOk

`func (o *AddressTransactionWithTransfers) GetNftTransfersOk() (*[]AddressTransactionWithTransfersNftTransfersInner, bool)`

GetNftTransfersOk returns a tuple with the NftTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftTransfers

`func (o *AddressTransactionWithTransfers) SetNftTransfers(v []AddressTransactionWithTransfersNftTransfersInner)`

SetNftTransfers sets NftTransfers field to given value.

### HasNftTransfers

`func (o *AddressTransactionWithTransfers) HasNftTransfers() bool`

HasNftTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


