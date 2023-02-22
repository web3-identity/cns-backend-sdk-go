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

// ServicesMakeRegisterOrderReq struct for ServicesMakeRegisterOrderReq
type ServicesMakeRegisterOrderReq struct {
	Description string `json:"description"`
	TradeProvider *string `json:"trade_provider,omitempty"`
	TradeType string `json:"trade_type"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMakeRegisterOrderReq ServicesMakeRegisterOrderReq

// NewServicesMakeRegisterOrderReq instantiates a new ServicesMakeRegisterOrderReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMakeRegisterOrderReq(description string, tradeType string) *ServicesMakeRegisterOrderReq {
	this := ServicesMakeRegisterOrderReq{}
	this.Description = description
	this.TradeType = tradeType
	return &this
}

// NewServicesMakeRegisterOrderReqWithDefaults instantiates a new ServicesMakeRegisterOrderReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMakeRegisterOrderReqWithDefaults() *ServicesMakeRegisterOrderReq {
	this := ServicesMakeRegisterOrderReq{}
	return &this
}

// GetDescription returns the Description field value
func (o *ServicesMakeRegisterOrderReq) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRegisterOrderReq) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServicesMakeRegisterOrderReq) SetDescription(v string) {
	o.Description = v
}

// GetTradeProvider returns the TradeProvider field value if set, zero value otherwise.
func (o *ServicesMakeRegisterOrderReq) GetTradeProvider() string {
	if o == nil || isNil(o.TradeProvider) {
		var ret string
		return ret
	}
	return *o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeRegisterOrderReq) GetTradeProviderOk() (*string, bool) {
	if o == nil || isNil(o.TradeProvider) {
    return nil, false
	}
	return o.TradeProvider, true
}

// HasTradeProvider returns a boolean if a field has been set.
func (o *ServicesMakeRegisterOrderReq) HasTradeProvider() bool {
	if o != nil && !isNil(o.TradeProvider) {
		return true
	}

	return false
}

// SetTradeProvider gets a reference to the given string and assigns it to the TradeProvider field.
func (o *ServicesMakeRegisterOrderReq) SetTradeProvider(v string) {
	o.TradeProvider = &v
}

// GetTradeType returns the TradeType field value
func (o *ServicesMakeRegisterOrderReq) GetTradeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRegisterOrderReq) GetTradeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TradeType, true
}

// SetTradeType sets field value
func (o *ServicesMakeRegisterOrderReq) SetTradeType(v string) {
	o.TradeType = v
}

func (o ServicesMakeRegisterOrderReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.TradeProvider) {
		toSerialize["trade_provider"] = o.TradeProvider
	}
	if true {
		toSerialize["trade_type"] = o.TradeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesMakeRegisterOrderReq) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMakeRegisterOrderReq := _ServicesMakeRegisterOrderReq{}

	if err = json.Unmarshal(bytes, &varServicesMakeRegisterOrderReq); err == nil {
		*o = ServicesMakeRegisterOrderReq(varServicesMakeRegisterOrderReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMakeRegisterOrderReq struct {
	value *ServicesMakeRegisterOrderReq
	isSet bool
}

func (v NullableServicesMakeRegisterOrderReq) Get() *ServicesMakeRegisterOrderReq {
	return v.value
}

func (v *NullableServicesMakeRegisterOrderReq) Set(val *ServicesMakeRegisterOrderReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMakeRegisterOrderReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMakeRegisterOrderReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMakeRegisterOrderReq(val *ServicesMakeRegisterOrderReq) *NullableServicesMakeRegisterOrderReq {
	return &NullableServicesMakeRegisterOrderReq{value: val, isSet: true}
}

func (v NullableServicesMakeRegisterOrderReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMakeRegisterOrderReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


