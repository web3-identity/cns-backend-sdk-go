# ModelsRenewOrderArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnsName** | **string** |  | 
**Duration** | **int32** |  | 
**Fuses** | Pointer to **int32** |  | [optional] 
**WrapperExpiry** | **int32** |  | 

## Methods

### NewModelsRenewOrderArgs

`func NewModelsRenewOrderArgs(cnsName string, duration int32, wrapperExpiry int32, ) *ModelsRenewOrderArgs`

NewModelsRenewOrderArgs instantiates a new ModelsRenewOrderArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRenewOrderArgsWithDefaults

`func NewModelsRenewOrderArgsWithDefaults() *ModelsRenewOrderArgs`

NewModelsRenewOrderArgsWithDefaults instantiates a new ModelsRenewOrderArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnsName

`func (o *ModelsRenewOrderArgs) GetCnsName() string`

GetCnsName returns the CnsName field if non-nil, zero value otherwise.

### GetCnsNameOk

`func (o *ModelsRenewOrderArgs) GetCnsNameOk() (*string, bool)`

GetCnsNameOk returns a tuple with the CnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsName

`func (o *ModelsRenewOrderArgs) SetCnsName(v string)`

SetCnsName sets CnsName field to given value.


### GetDuration

`func (o *ModelsRenewOrderArgs) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ModelsRenewOrderArgs) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ModelsRenewOrderArgs) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetFuses

`func (o *ModelsRenewOrderArgs) GetFuses() int32`

GetFuses returns the Fuses field if non-nil, zero value otherwise.

### GetFusesOk

`func (o *ModelsRenewOrderArgs) GetFusesOk() (*int32, bool)`

GetFusesOk returns a tuple with the Fuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuses

`func (o *ModelsRenewOrderArgs) SetFuses(v int32)`

SetFuses sets Fuses field to given value.

### HasFuses

`func (o *ModelsRenewOrderArgs) HasFuses() bool`

HasFuses returns a boolean if a field has been set.

### GetWrapperExpiry

`func (o *ModelsRenewOrderArgs) GetWrapperExpiry() int32`

GetWrapperExpiry returns the WrapperExpiry field if non-nil, zero value otherwise.

### GetWrapperExpiryOk

`func (o *ModelsRenewOrderArgs) GetWrapperExpiryOk() (*int32, bool)`

GetWrapperExpiryOk returns a tuple with the WrapperExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperExpiry

`func (o *ModelsRenewOrderArgs) SetWrapperExpiry(v int32)`

SetWrapperExpiry sets WrapperExpiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


