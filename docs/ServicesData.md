# ServicesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conclusion** | Pointer to **string** | 审核结果，可取值：合规、不合规、疑似、审核失败 | [optional] 
**ConclusionType** | Pointer to **int32** | 审核结果类型，可取值1.合规，2.不合规，3.疑似，4.审核失败 | [optional] 
**Hits** | Pointer to [**[]ServicesHits**](ServicesHits.md) | 命中关键词信息 | [optional] 
**Msg** | Pointer to **string** | 不合规项描述信息 | [optional] 
**SubType** | Pointer to **int32** | 审核子类型，此字段需参照type主类型字段决定其含义： 当type&#x3D;11时subType取值含义： 0:百度官方默认违禁词库 当type&#x3D;12时subType取值含义： 0:低质灌水、1:暴恐违禁、2:文本色情、3:政治敏感、4:恶意推广、5:低俗辱骂 当type&#x3D;13时subType取值含义： 0:自定义文本黑名单 当type&#x3D;14时subType取值含义： 0:自定义文本白名单 | [optional] 
**Type** | Pointer to **int32** | 审核主类型，11：官方违禁词库、12：文本反作弊、13:自定义文本黑名单、14:自定义文本白名单 | [optional] 

## Methods

### NewServicesData

`func NewServicesData() *ServicesData`

NewServicesData instantiates a new ServicesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesDataWithDefaults

`func NewServicesDataWithDefaults() *ServicesData`

NewServicesDataWithDefaults instantiates a new ServicesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConclusion

`func (o *ServicesData) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *ServicesData) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *ServicesData) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *ServicesData) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### GetConclusionType

`func (o *ServicesData) GetConclusionType() int32`

GetConclusionType returns the ConclusionType field if non-nil, zero value otherwise.

### GetConclusionTypeOk

`func (o *ServicesData) GetConclusionTypeOk() (*int32, bool)`

GetConclusionTypeOk returns a tuple with the ConclusionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusionType

`func (o *ServicesData) SetConclusionType(v int32)`

SetConclusionType sets ConclusionType field to given value.

### HasConclusionType

`func (o *ServicesData) HasConclusionType() bool`

HasConclusionType returns a boolean if a field has been set.

### GetHits

`func (o *ServicesData) GetHits() []ServicesHits`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *ServicesData) GetHitsOk() (*[]ServicesHits, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *ServicesData) SetHits(v []ServicesHits)`

SetHits sets Hits field to given value.

### HasHits

`func (o *ServicesData) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetMsg

`func (o *ServicesData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ServicesData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ServicesData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ServicesData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSubType

`func (o *ServicesData) GetSubType() int32`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ServicesData) GetSubTypeOk() (*int32, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ServicesData) SetSubType(v int32)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ServicesData) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *ServicesData) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicesData) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicesData) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ServicesData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


