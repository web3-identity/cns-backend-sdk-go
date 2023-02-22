# ModelsCommitCore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitHash** | **string** |  | 
**Data** | **[]string** | hex 数组 | 
**Duration** | **int32** |  | 
**Fuses** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**OrderState** | Pointer to **int32** |  | [optional] 
**Owner** | **string** | base32地址或hex地址 | 
**Resolver** | **string** |  | 
**ReverseRecord** | **bool** |  | 
**Secret** | **string** | 32字节 | 
**WrapperExpiry** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsCommitCore

`func NewModelsCommitCore(commitHash string, data []string, duration int32, name string, owner string, resolver string, reverseRecord bool, secret string, ) *ModelsCommitCore`

NewModelsCommitCore instantiates a new ModelsCommitCore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCommitCoreWithDefaults

`func NewModelsCommitCoreWithDefaults() *ModelsCommitCore`

NewModelsCommitCoreWithDefaults instantiates a new ModelsCommitCore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitHash

`func (o *ModelsCommitCore) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ModelsCommitCore) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ModelsCommitCore) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetData

`func (o *ModelsCommitCore) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelsCommitCore) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelsCommitCore) SetData(v []string)`

SetData sets Data field to given value.


### GetDuration

`func (o *ModelsCommitCore) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ModelsCommitCore) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ModelsCommitCore) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetFuses

`func (o *ModelsCommitCore) GetFuses() int32`

GetFuses returns the Fuses field if non-nil, zero value otherwise.

### GetFusesOk

`func (o *ModelsCommitCore) GetFusesOk() (*int32, bool)`

GetFusesOk returns a tuple with the Fuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuses

`func (o *ModelsCommitCore) SetFuses(v int32)`

SetFuses sets Fuses field to given value.

### HasFuses

`func (o *ModelsCommitCore) HasFuses() bool`

HasFuses returns a boolean if a field has been set.

### GetName

`func (o *ModelsCommitCore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCommitCore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCommitCore) SetName(v string)`

SetName sets Name field to given value.


### GetOrderState

`func (o *ModelsCommitCore) GetOrderState() int32`

GetOrderState returns the OrderState field if non-nil, zero value otherwise.

### GetOrderStateOk

`func (o *ModelsCommitCore) GetOrderStateOk() (*int32, bool)`

GetOrderStateOk returns a tuple with the OrderState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderState

`func (o *ModelsCommitCore) SetOrderState(v int32)`

SetOrderState sets OrderState field to given value.

### HasOrderState

`func (o *ModelsCommitCore) HasOrderState() bool`

HasOrderState returns a boolean if a field has been set.

### GetOwner

`func (o *ModelsCommitCore) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelsCommitCore) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelsCommitCore) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetResolver

`func (o *ModelsCommitCore) GetResolver() string`

GetResolver returns the Resolver field if non-nil, zero value otherwise.

### GetResolverOk

`func (o *ModelsCommitCore) GetResolverOk() (*string, bool)`

GetResolverOk returns a tuple with the Resolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolver

`func (o *ModelsCommitCore) SetResolver(v string)`

SetResolver sets Resolver field to given value.


### GetReverseRecord

`func (o *ModelsCommitCore) GetReverseRecord() bool`

GetReverseRecord returns the ReverseRecord field if non-nil, zero value otherwise.

### GetReverseRecordOk

`func (o *ModelsCommitCore) GetReverseRecordOk() (*bool, bool)`

GetReverseRecordOk returns a tuple with the ReverseRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseRecord

`func (o *ModelsCommitCore) SetReverseRecord(v bool)`

SetReverseRecord sets ReverseRecord field to given value.


### GetSecret

`func (o *ModelsCommitCore) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ModelsCommitCore) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ModelsCommitCore) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetWrapperExpiry

`func (o *ModelsCommitCore) GetWrapperExpiry() int32`

GetWrapperExpiry returns the WrapperExpiry field if non-nil, zero value otherwise.

### GetWrapperExpiryOk

`func (o *ModelsCommitCore) GetWrapperExpiryOk() (*int32, bool)`

GetWrapperExpiryOk returns a tuple with the WrapperExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperExpiry

`func (o *ModelsCommitCore) SetWrapperExpiry(v int32)`

SetWrapperExpiry sets WrapperExpiry field to given value.

### HasWrapperExpiry

`func (o *ModelsCommitCore) HasWrapperExpiry() bool`

HasWrapperExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


