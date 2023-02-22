# ServicesRnewByAdminResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnsName** | **string** |  | 
**Duration** | **int32** |  | 
**Fuses** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**TxError** | Pointer to **string** |  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**TxState** | Pointer to **int32** |  | [optional] 
**WrapperExpiry** | **int32** |  | 

## Methods

### NewServicesRnewByAdminResp

`func NewServicesRnewByAdminResp(cnsName string, duration int32, wrapperExpiry int32, ) *ServicesRnewByAdminResp`

NewServicesRnewByAdminResp instantiates a new ServicesRnewByAdminResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesRnewByAdminRespWithDefaults

`func NewServicesRnewByAdminRespWithDefaults() *ServicesRnewByAdminResp`

NewServicesRnewByAdminRespWithDefaults instantiates a new ServicesRnewByAdminResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnsName

`func (o *ServicesRnewByAdminResp) GetCnsName() string`

GetCnsName returns the CnsName field if non-nil, zero value otherwise.

### GetCnsNameOk

`func (o *ServicesRnewByAdminResp) GetCnsNameOk() (*string, bool)`

GetCnsNameOk returns a tuple with the CnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsName

`func (o *ServicesRnewByAdminResp) SetCnsName(v string)`

SetCnsName sets CnsName field to given value.


### GetDuration

`func (o *ServicesRnewByAdminResp) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ServicesRnewByAdminResp) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ServicesRnewByAdminResp) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetFuses

`func (o *ServicesRnewByAdminResp) GetFuses() int32`

GetFuses returns the Fuses field if non-nil, zero value otherwise.

### GetFusesOk

`func (o *ServicesRnewByAdminResp) GetFusesOk() (*int32, bool)`

GetFusesOk returns a tuple with the Fuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuses

`func (o *ServicesRnewByAdminResp) SetFuses(v int32)`

SetFuses sets Fuses field to given value.

### HasFuses

`func (o *ServicesRnewByAdminResp) HasFuses() bool`

HasFuses returns a boolean if a field has been set.

### GetId

`func (o *ServicesRnewByAdminResp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicesRnewByAdminResp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicesRnewByAdminResp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServicesRnewByAdminResp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTxError

`func (o *ServicesRnewByAdminResp) GetTxError() string`

GetTxError returns the TxError field if non-nil, zero value otherwise.

### GetTxErrorOk

`func (o *ServicesRnewByAdminResp) GetTxErrorOk() (*string, bool)`

GetTxErrorOk returns a tuple with the TxError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxError

`func (o *ServicesRnewByAdminResp) SetTxError(v string)`

SetTxError sets TxError field to given value.

### HasTxError

`func (o *ServicesRnewByAdminResp) HasTxError() bool`

HasTxError returns a boolean if a field has been set.

### GetTxHash

`func (o *ServicesRnewByAdminResp) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ServicesRnewByAdminResp) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ServicesRnewByAdminResp) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ServicesRnewByAdminResp) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetTxState

`func (o *ServicesRnewByAdminResp) GetTxState() int32`

GetTxState returns the TxState field if non-nil, zero value otherwise.

### GetTxStateOk

`func (o *ServicesRnewByAdminResp) GetTxStateOk() (*int32, bool)`

GetTxStateOk returns a tuple with the TxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxState

`func (o *ServicesRnewByAdminResp) SetTxState(v int32)`

SetTxState sets TxState field to given value.

### HasTxState

`func (o *ServicesRnewByAdminResp) HasTxState() bool`

HasTxState returns a boolean if a field has been set.

### GetWrapperExpiry

`func (o *ServicesRnewByAdminResp) GetWrapperExpiry() int32`

GetWrapperExpiry returns the WrapperExpiry field if non-nil, zero value otherwise.

### GetWrapperExpiryOk

`func (o *ServicesRnewByAdminResp) GetWrapperExpiryOk() (*int32, bool)`

GetWrapperExpiryOk returns a tuple with the WrapperExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperExpiry

`func (o *ServicesRnewByAdminResp) SetWrapperExpiry(v int32)`

SetWrapperExpiry sets WrapperExpiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


