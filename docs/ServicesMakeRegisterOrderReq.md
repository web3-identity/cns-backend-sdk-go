# ServicesMakeRegisterOrderReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**TradeProvider** | Pointer to **string** |  | [optional] 
**TradeType** | **string** |  | 

## Methods

### NewServicesMakeRegisterOrderReq

`func NewServicesMakeRegisterOrderReq(description string, tradeType string, ) *ServicesMakeRegisterOrderReq`

NewServicesMakeRegisterOrderReq instantiates a new ServicesMakeRegisterOrderReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesMakeRegisterOrderReqWithDefaults

`func NewServicesMakeRegisterOrderReqWithDefaults() *ServicesMakeRegisterOrderReq`

NewServicesMakeRegisterOrderReqWithDefaults instantiates a new ServicesMakeRegisterOrderReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServicesMakeRegisterOrderReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicesMakeRegisterOrderReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicesMakeRegisterOrderReq) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTradeProvider

`func (o *ServicesMakeRegisterOrderReq) GetTradeProvider() string`

GetTradeProvider returns the TradeProvider field if non-nil, zero value otherwise.

### GetTradeProviderOk

`func (o *ServicesMakeRegisterOrderReq) GetTradeProviderOk() (*string, bool)`

GetTradeProviderOk returns a tuple with the TradeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeProvider

`func (o *ServicesMakeRegisterOrderReq) SetTradeProvider(v string)`

SetTradeProvider sets TradeProvider field to given value.

### HasTradeProvider

`func (o *ServicesMakeRegisterOrderReq) HasTradeProvider() bool`

HasTradeProvider returns a boolean if a field has been set.

### GetTradeType

`func (o *ServicesMakeRegisterOrderReq) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *ServicesMakeRegisterOrderReq) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *ServicesMakeRegisterOrderReq) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


