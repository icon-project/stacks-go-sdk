# AddressBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stx** | [**StxBalance**](StxBalance.md) |  | 
**FungibleTokens** | [**map[string]FtBalance**](FtBalance.md) |  | 
**NonFungibleTokens** | [**map[string]NftBalance**](NftBalance.md) |  | 
**TokenOfferingLocked** | Pointer to [**AddressTokenOfferingLocked**](AddressTokenOfferingLocked.md) |  | [optional] 

## Methods

### NewAddressBalanceResponse

`func NewAddressBalanceResponse(stx StxBalance, fungibleTokens map[string]FtBalance, nonFungibleTokens map[string]NftBalance, ) *AddressBalanceResponse`

NewAddressBalanceResponse instantiates a new AddressBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceResponseWithDefaults

`func NewAddressBalanceResponseWithDefaults() *AddressBalanceResponse`

NewAddressBalanceResponseWithDefaults instantiates a new AddressBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStx

`func (o *AddressBalanceResponse) GetStx() StxBalance`

GetStx returns the Stx field if non-nil, zero value otherwise.

### GetStxOk

`func (o *AddressBalanceResponse) GetStxOk() (*StxBalance, bool)`

GetStxOk returns a tuple with the Stx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStx

`func (o *AddressBalanceResponse) SetStx(v StxBalance)`

SetStx sets Stx field to given value.


### GetFungibleTokens

`func (o *AddressBalanceResponse) GetFungibleTokens() map[string]FtBalance`

GetFungibleTokens returns the FungibleTokens field if non-nil, zero value otherwise.

### GetFungibleTokensOk

`func (o *AddressBalanceResponse) GetFungibleTokensOk() (*map[string]FtBalance, bool)`

GetFungibleTokensOk returns a tuple with the FungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleTokens

`func (o *AddressBalanceResponse) SetFungibleTokens(v map[string]FtBalance)`

SetFungibleTokens sets FungibleTokens field to given value.


### GetNonFungibleTokens

`func (o *AddressBalanceResponse) GetNonFungibleTokens() map[string]NftBalance`

GetNonFungibleTokens returns the NonFungibleTokens field if non-nil, zero value otherwise.

### GetNonFungibleTokensOk

`func (o *AddressBalanceResponse) GetNonFungibleTokensOk() (*map[string]NftBalance, bool)`

GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFungibleTokens

`func (o *AddressBalanceResponse) SetNonFungibleTokens(v map[string]NftBalance)`

SetNonFungibleTokens sets NonFungibleTokens field to given value.


### GetTokenOfferingLocked

`func (o *AddressBalanceResponse) GetTokenOfferingLocked() AddressTokenOfferingLocked`

GetTokenOfferingLocked returns the TokenOfferingLocked field if non-nil, zero value otherwise.

### GetTokenOfferingLockedOk

`func (o *AddressBalanceResponse) GetTokenOfferingLockedOk() (*AddressTokenOfferingLocked, bool)`

GetTokenOfferingLockedOk returns a tuple with the TokenOfferingLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOfferingLocked

`func (o *AddressBalanceResponse) SetTokenOfferingLocked(v AddressTokenOfferingLocked)`

SetTokenOfferingLocked sets TokenOfferingLocked field to given value.

### HasTokenOfferingLocked

`func (o *AddressBalanceResponse) HasTokenOfferingLocked() bool`

HasTokenOfferingLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


