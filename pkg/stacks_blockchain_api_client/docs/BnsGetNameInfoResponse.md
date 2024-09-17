# BnsGetNameInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Blockchain** | **string** |  | 
**ExpireBlock** | Pointer to **int32** |  | [optional] 
**GracePeriod** | Pointer to **int32** |  | [optional] 
**LastTxid** | **string** |  | 
**Resolver** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Zonefile** | Pointer to **string** |  | [optional] 
**ZonefileHash** | **string** |  | 

## Methods

### NewBnsGetNameInfoResponse

`func NewBnsGetNameInfoResponse(address string, blockchain string, lastTxid string, status string, zonefileHash string, ) *BnsGetNameInfoResponse`

NewBnsGetNameInfoResponse instantiates a new BnsGetNameInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsGetNameInfoResponseWithDefaults

`func NewBnsGetNameInfoResponseWithDefaults() *BnsGetNameInfoResponse`

NewBnsGetNameInfoResponseWithDefaults instantiates a new BnsGetNameInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BnsGetNameInfoResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BnsGetNameInfoResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BnsGetNameInfoResponse) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBlockchain

`func (o *BnsGetNameInfoResponse) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *BnsGetNameInfoResponse) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *BnsGetNameInfoResponse) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetExpireBlock

`func (o *BnsGetNameInfoResponse) GetExpireBlock() int32`

GetExpireBlock returns the ExpireBlock field if non-nil, zero value otherwise.

### GetExpireBlockOk

`func (o *BnsGetNameInfoResponse) GetExpireBlockOk() (*int32, bool)`

GetExpireBlockOk returns a tuple with the ExpireBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireBlock

`func (o *BnsGetNameInfoResponse) SetExpireBlock(v int32)`

SetExpireBlock sets ExpireBlock field to given value.

### HasExpireBlock

`func (o *BnsGetNameInfoResponse) HasExpireBlock() bool`

HasExpireBlock returns a boolean if a field has been set.

### GetGracePeriod

`func (o *BnsGetNameInfoResponse) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *BnsGetNameInfoResponse) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *BnsGetNameInfoResponse) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *BnsGetNameInfoResponse) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetLastTxid

`func (o *BnsGetNameInfoResponse) GetLastTxid() string`

GetLastTxid returns the LastTxid field if non-nil, zero value otherwise.

### GetLastTxidOk

`func (o *BnsGetNameInfoResponse) GetLastTxidOk() (*string, bool)`

GetLastTxidOk returns a tuple with the LastTxid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTxid

`func (o *BnsGetNameInfoResponse) SetLastTxid(v string)`

SetLastTxid sets LastTxid field to given value.


### GetResolver

`func (o *BnsGetNameInfoResponse) GetResolver() string`

GetResolver returns the Resolver field if non-nil, zero value otherwise.

### GetResolverOk

`func (o *BnsGetNameInfoResponse) GetResolverOk() (*string, bool)`

GetResolverOk returns a tuple with the Resolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolver

`func (o *BnsGetNameInfoResponse) SetResolver(v string)`

SetResolver sets Resolver field to given value.

### HasResolver

`func (o *BnsGetNameInfoResponse) HasResolver() bool`

HasResolver returns a boolean if a field has been set.

### GetStatus

`func (o *BnsGetNameInfoResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsGetNameInfoResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsGetNameInfoResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetZonefile

`func (o *BnsGetNameInfoResponse) GetZonefile() string`

GetZonefile returns the Zonefile field if non-nil, zero value otherwise.

### GetZonefileOk

`func (o *BnsGetNameInfoResponse) GetZonefileOk() (*string, bool)`

GetZonefileOk returns a tuple with the Zonefile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonefile

`func (o *BnsGetNameInfoResponse) SetZonefile(v string)`

SetZonefile sets Zonefile field to given value.

### HasZonefile

`func (o *BnsGetNameInfoResponse) HasZonefile() bool`

HasZonefile returns a boolean if a field has been set.

### GetZonefileHash

`func (o *BnsGetNameInfoResponse) GetZonefileHash() string`

GetZonefileHash returns the ZonefileHash field if non-nil, zero value otherwise.

### GetZonefileHashOk

`func (o *BnsGetNameInfoResponse) GetZonefileHashOk() (*string, bool)`

GetZonefileHashOk returns a tuple with the ZonefileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonefileHash

`func (o *BnsGetNameInfoResponse) SetZonefileHash(v string)`

SetZonefileHash sets ZonefileHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


