/*
CNS-BACKEND

The responses of the open api in swagger focus on the data field rather than the code and the message fields

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cnsbackend

import (
	"encoding/json"
)

// ServicesCensorTextReq struct for ServicesCensorTextReq
type ServicesCensorTextReq struct {
	Content string `json:"content"`
	AdditionalProperties map[string]interface{}
}

type _ServicesCensorTextReq ServicesCensorTextReq

// NewServicesCensorTextReq instantiates a new ServicesCensorTextReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesCensorTextReq(content string) *ServicesCensorTextReq {
	this := ServicesCensorTextReq{}
	this.Content = content
	return &this
}

// NewServicesCensorTextReqWithDefaults instantiates a new ServicesCensorTextReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesCensorTextReqWithDefaults() *ServicesCensorTextReq {
	this := ServicesCensorTextReq{}
	return &this
}

// GetContent returns the Content field value
func (o *ServicesCensorTextReq) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ServicesCensorTextReq) GetContentOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ServicesCensorTextReq) SetContent(v string) {
	o.Content = v
}

func (o ServicesCensorTextReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesCensorTextReq) UnmarshalJSON(bytes []byte) (err error) {
	varServicesCensorTextReq := _ServicesCensorTextReq{}

	if err = json.Unmarshal(bytes, &varServicesCensorTextReq); err == nil {
		*o = ServicesCensorTextReq(varServicesCensorTextReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesCensorTextReq struct {
	value *ServicesCensorTextReq
	isSet bool
}

func (v NullableServicesCensorTextReq) Get() *ServicesCensorTextReq {
	return v.value
}

func (v *NullableServicesCensorTextReq) Set(val *ServicesCensorTextReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesCensorTextReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesCensorTextReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesCensorTextReq(val *ServicesCensorTextReq) *NullableServicesCensorTextReq {
	return &NullableServicesCensorTextReq{value: val, isSet: true}
}

func (v NullableServicesCensorTextReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesCensorTextReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


