# ServicesMakeRenewOrderReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnsName** | **string** |  | 
**Description** | **string** |  | 
**Duration** | **int32** |  | 
**Fuses** | Pointer to **int32** |  | [optional] 
**TradeProvider** | Pointer to **string** |  | [optional] 
**TradeType** | **string** |  | 
**WrapperExpiry** | **int32** |  | 

## Methods

### NewServicesMakeRenewOrderReq

`func NewServicesMakeRenewOrderReq(cnsName string, description string, duration int32, tradeType string, wrapperExpiry int32, ) *ServicesMakeRenewOrderReq`

NewServicesMakeRenewOrderReq instantiates a new ServicesMakeRenewOrderReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesMakeRenewOrderReqWithDefaults

`func NewServicesMakeRenewOrderReqWithDefaults() *ServicesMakeRenewOrderReq`

NewServicesMakeRenewOrderReqWithDefaults instantiates a new ServicesMakeRenewOrderReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnsName

`func (o *ServicesMakeRenewOrderReq) GetCnsName() string`

GetCnsName returns the CnsName field if non-nil, zero value otherwise.

### GetCnsNameOk

`func (o *ServicesMakeRenewOrderReq) GetCnsNameOk() (*string, bool)`

GetCnsNameOk returns a tuple with the CnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsName

`func (o *ServicesMakeRenewOrderReq) SetCnsName(v string)`

SetCnsName sets CnsName field to given value.


### GetDescription

`func (o *ServicesMakeRenewOrderReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesMakeRenewOrderReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesMakeRenewOrderReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDuration

`func (o *ServicesMakeRenewOrderReq) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ServicesMakeRenewOrderReq) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ServicesMakeRenewOrderReq) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetFuses

`func (o *ServicesMakeRenewOrderReq) GetFuses() int32`

GetFuses returns the Fuses field if non-nil, zero value otherwise.

### GetFusesOk

`func (o *ServicesMakeRenewOrderReq) GetFusesOk() (*int32, bool)`

GetFusesOk returns a tuple with the Fuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuses

`func (o *ServicesMakeRenewOrderReq) SetFuses(v int32)`

SetFuses sets Fuses field to given value.

### HasFuses

`func (o *ServicesMakeRenewOrderReq) HasFuses() bool`

HasFuses returns a boolean if a field has been set.

### GetTradeProvider

`func (o *ServicesMakeRenewOrderReq) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ServicesMakeRenewOrderReq) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ServicesMakeRenewOrderReq) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ServicesMakeRenewOrderReq) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeType

`func (o *ServicesMakeRenewOrderReq) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ServicesMakeRenewOrderReq) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ServicesMakeRenewOrderReq) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.


### GetWrapperExpiry

`func (o *ServicesMakeRenewOrderReq) GetWrapperExpiry() int32`

GetWrapperExpiry returns the WrapperExpiry field if non-nil, zero value otherwise.

### GetWrapperExpiryOk

`func (o *ServicesMakeRenewOrderReq) GetWrapperExpiryOk() (*int32, bool)`

GetWrapperExpiryOk returns a tuple with the WrapperExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperExpiry

`func (o *ServicesMakeRenewOrderReq) SetWrapperExpiry(v int32)`

SetWrapperExpiry sets WrapperExpiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


