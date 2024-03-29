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

// ServicesMakeRenewOrderReq struct for ServicesMakeRenewOrderReq
type ServicesMakeRenewOrderReq struct {
	CnsName string `json:"cns_name"`
	Description string `json:"description"`
	Duration int32 `json:"duration"`
	Fuses *int32 `json:"fuses,omitempty"`
	TradeProvider *string `json:"trade_provider,omitempty"`
	TradeType string `json:"trade_type"`
	WrapperExpiry int32 `json:"wrapper_expiry"`
	AdditionalProperties map[string]interface{}
}

type _ServicesMakeRenewOrderReq ServicesMakeRenewOrderReq

// NewServicesMakeRenewOrderReq instantiates a new ServicesMakeRenewOrderReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesMakeRenewOrderReq(cnsName string, description string, duration int32, tradeType string, wrapperExpiry int32) *ServicesMakeRenewOrderReq {
	this := ServicesMakeRenewOrderReq{}
	this.CnsName = cnsName
	this.Description = description
	this.Duration = duration
	this.TradeType = tradeType
	this.WrapperExpiry = wrapperExpiry
	return &this
}

// NewServicesMakeRenewOrderReqWithDefaults instantiates a new ServicesMakeRenewOrderReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesMakeRenewOrderReqWithDefaults() *ServicesMakeRenewOrderReq {
	this := ServicesMakeRenewOrderReq{}
	return &this
}

// GetCnsName returns the CnsName field value
func (o *ServicesMakeRenewOrderReq) GetCnsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CnsName
}

// GetCnsNameOk returns a tuple with the CnsName field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetCnsNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CnsName, true
}

// SetCnsName sets field value
func (o *ServicesMakeRenewOrderReq) SetCnsName(v string) {
	o.CnsName = v
}

// GetDescription returns the Description field value
func (o *ServicesMakeRenewOrderReq) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServicesMakeRenewOrderReq) SetDescription(v string) {
	o.Description = v
}

// GetDuration returns the Duration field value
func (o *ServicesMakeRenewOrderReq) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetDurationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ServicesMakeRenewOrderReq) SetDuration(v int32) {
	o.Duration = v
}

// GetFuses returns the Fuses field value if set, zero value otherwise.
func (o *ServicesMakeRenewOrderReq) GetFuses() int32 {
	if o == nil || isNil(o.Fuses) {
		var ret int32
		return ret
	}
	return *o.Fuses
}

// GetFusesOk returns a tuple with the Fuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetFusesOk() (*int32, bool) {
	if o == nil || isNil(o.Fuses) {
    return nil, false
	}
	return o.Fuses, true
}

// HasFuses returns a boolean if a field has been set.
func (o *ServicesMakeRenewOrderReq) HasFuses() bool {
	if o != nil && !isNil(o.Fuses) {
		return true
	}

	return false
}

// SetFuses gets a reference to the given int32 and assigns it to the Fuses field.
func (o *ServicesMakeRenewOrderReq) SetFuses(v int32) {
	o.Fuses = &v
}

// GetTradeProvider returns the TradeProvider field value if set, zero value otherwise.
func (o *ServicesMakeRenewOrderReq) GetTradeProvider() string {
	if o == nil || isNil(o.TradeProvider) {
		var ret string
		return ret
	}
	return *o.TradeProvider
}

// GetTradeProviderOk returns a tuple with the TradeProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetTradeProviderOk() (*string, bool) {
	if o == nil || isNil(o.TradeProvider) {
    return nil, false
	}
	return o.TradeProvider, true
}

// HasTradeProvider returns a boolean if a field has been set.
func (o *ServicesMakeRenewOrderReq) HasTradeProvider() bool {
	if o != nil && !isNil(o.TradeProvider) {
		return true
	}

	return false
}

// SetTradeProvider gets a reference to the given string and assigns it to the TradeProvider field.
func (o *ServicesMakeRenewOrderReq) SetTradeProvider(v string) {
	o.TradeProvider = &v
}

// GetTradeType returns the TradeType field value
func (o *ServicesMakeRenewOrderReq) GetTradeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetTradeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TradeType, true
}

// SetTradeType sets field value
func (o *ServicesMakeRenewOrderReq) SetTradeType(v string) {
	o.TradeType = v
}

// GetWrapperExpiry returns the WrapperExpiry field value
func (o *ServicesMakeRenewOrderReq) GetWrapperExpiry() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WrapperExpiry
}

// GetWrapperExpiryOk returns a tuple with the WrapperExpiry field value
// and a boolean to check if the value has been set.
func (o *ServicesMakeRenewOrderReq) GetWrapperExpiryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WrapperExpiry, true
}

// SetWrapperExpiry sets field value
func (o *ServicesMakeRenewOrderReq) SetWrapperExpiry(v int32) {
	o.WrapperExpiry = v
}

func (o ServicesMakeRenewOrderReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cns_name"] = o.CnsName
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Fuses) {
		toSerialize["fuses"] = o.Fuses
	}
	if !isNil(o.TradeProvider) {
		toSerialize["trade_provider"] = o.TradeProvider
	}
	if true {
		toSerialize["trade_type"] = o.TradeType
	}
	if true {
		toSerialize["wrapper_expiry"] = o.WrapperExpiry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesMakeRenewOrderReq) UnmarshalJSON(bytes []byte) (err error) {
	varServicesMakeRenewOrderReq := _ServicesMakeRenewOrderReq{}

	if err = json.Unmarshal(bytes, &varServicesMakeRenewOrderReq); err == nil {
		*o = ServicesMakeRenewOrderReq(varServicesMakeRenewOrderReq)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cns_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "fuses")
		delete(additionalProperties, "trade_provider")
		delete(additionalProperties, "trade_type")
		delete(additionalProperties, "wrapper_expiry")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesMakeRenewOrderReq struct {
	value *ServicesMakeRenewOrderReq
	isSet bool
}

func (v NullableServicesMakeRenewOrderReq) Get() *ServicesMakeRenewOrderReq {
	return v.value
}

func (v *NullableServicesMakeRenewOrderReq) Set(val *ServicesMakeRenewOrderReq) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesMakeRenewOrderReq) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesMakeRenewOrderReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesMakeRenewOrderReq(val *ServicesMakeRenewOrderReq) *NullableServicesMakeRenewOrderReq {
	return &NullableServicesMakeRenewOrderReq{value: val, isSet: true}
}

func (v NullableServicesMakeRenewOrderReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesMakeRenewOrderReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


