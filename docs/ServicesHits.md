# ServicesHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetName** | Pointer to **string** | 违规项目所属数据集名称 | [optional] 
**Probability** | Pointer to **float32** |  | [optional] 
**Words** | Pointer to **[]string** | 违规文本关键字 | [optional] 

## Methods

### NewServicesHits

`func NewServicesHits() *ServicesHits`

NewServicesHits instantiates a new ServicesHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesHitsWithDefaults

`func NewServicesHitsWithDefaults() *ServicesHits`

NewServicesHitsWithDefaults instantiates a new ServicesHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetName

`func (o *ServicesHits) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *ServicesHits) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *ServicesHits) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *ServicesHits) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.

### GetProbability

`func (o *ServicesHits) GetProbability() float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *ServicesHits) GetProbabilityOk() (*float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *ServicesHits) SetProbability(v float32)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *ServicesHits) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetWords

`func (o *ServicesHits) GetWords() []string`

GetWords returns the Words field if non-nil, zero value otherwise.

### GetWordsOk

`func (o *ServicesHits) GetWordsOk() (*[]string, bool)`

GetWordsOk returns a tuple with the Words field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWords

`func (o *ServicesHits) SetWords(v []string)`

SetWords sets Words field to given value.

### HasWords

`func (o *ServicesHits) HasWords() bool`

HasWords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


