# ServicesCensorTextResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | **string** | LogID          int64  &#x60;json:\&quot;log_id\&quot;&#x60; | 
**ConclusionType** | **int32** | 审核结果类型，可取值1.合规，2.不合规，3.疑似，4.审核失败 | 
**Data** | Pointer to [**[]ServicesData**](ServicesData.md) | 不合规/疑似/命中白名单项详细信息。响应成功并且conclusion为疑似或不合规或命中白名单时才返回，响应失败或conclusion为合规且未命中白名单时不返回 | [optional] 

## Methods

### NewServicesCensorTextResp

`func NewServicesCensorTextResp(conclusion string, conclusionType int32, ) *ServicesCensorTextResp`

NewServicesCensorTextResp instantiates a new ServicesCensorTextResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesCensorTextRespWithDefaults

`func NewServicesCensorTextRespWithDefaults() *ServicesCensorTextResp`

NewServicesCensorTextRespWithDefaults instantiates a new ServicesCensorTextResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *ServicesCensorTextResp) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *ServicesCensorTextResp) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *ServicesCensorTextResp) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### GetConclusionType

`func (o *ServicesCensorTextResp) GetConclusionType() int32`

GetConclusionType returns the ConclusionType field if non-nil, zero value otherwise.

### GetConclusionTypeOk

`func (o *ServicesCensorTextResp) GetConclusionTypeOk() (*int32, bool)`

GetConclusionTypeOk returns a tuple with the ConclusionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusionType

`func (o *ServicesCensorTextResp) SetConclusionType(v int32)`

SetConclusionType sets ConclusionType field to given value.


### GetData

`func (o *ServicesCensorTextResp) GetData() []ServicesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesCensorTextResp) GetDataOk() (*[]ServicesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesCensorTextResp) SetData(v []ServicesData)`

SetData sets Data field to given value.

### HasData

`func (o *ServicesCensorTextResp) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


