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

// ServicesRnewByAdminResp struct for ServicesRnewByAdminResp
type ServicesRnewByAdminResp struct {
	CnsName string `json:"cns_name"`
	Duration int32 `json:"duration"`
	Fuses *int32 `json:"fuses,omitempty"`
	Id *int32 `json:"id,omitempty"`
	TxError *string `json:"tx_error,omitempty"`
	TxHash *string `json:"tx_hash,omitempty"`
	TxState *int32 `json:"tx_state,omitempty"`
	WrapperExpiry int32 `json:"wrapper_expiry"`
	AdditionalProperties map[string]interface{}
}

type _ServicesRnewByAdminResp ServicesRnewByAdminResp

// NewServicesRnewByAdminResp instantiates a new ServicesRnewByAdminResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicesRnewByAdminResp(cnsName string, duration int32, wrapperExpiry int32) *ServicesRnewByAdminResp {
	this := ServicesRnewByAdminResp{}
	this.CnsName = cnsName
	this.Duration = duration
	this.WrapperExpiry = wrapperExpiry
	return &this
}

// NewServicesRnewByAdminRespWithDefaults instantiates a new ServicesRnewByAdminResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesRnewByAdminRespWithDefaults() *ServicesRnewByAdminResp {
	this := ServicesRnewByAdminResp{}
	return &this
}

// GetCnsName returns the CnsName field value
func (o *ServicesRnewByAdminResp) GetCnsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CnsName
}

// GetCnsNameOk returns a tuple with the CnsName field value
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetCnsNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CnsName, true
}

// SetCnsName sets field value
func (o *ServicesRnewByAdminResp) SetCnsName(v string) {
	o.CnsName = v
}

// GetDuration returns the Duration field value
func (o *ServicesRnewByAdminResp) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetDurationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ServicesRnewByAdminResp) SetDuration(v int32) {
	o.Duration = v
}

// GetFuses returns the Fuses field value if set, zero value otherwise.
func (o *ServicesRnewByAdminResp) GetFuses() int32 {
	if o == nil || isNil(o.Fuses) {
		var ret int32
		return ret
	}
	return *o.Fuses
}

// GetFusesOk returns a tuple with the Fuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetFusesOk() (*int32, bool) {
	if o == nil || isNil(o.Fuses) {
    return nil, false
	}
	return o.Fuses, true
}

// HasFuses returns a boolean if a field has been set.
func (o *ServicesRnewByAdminResp) HasFuses() bool {
	if o != nil && !isNil(o.Fuses) {
		return true
	}

	return false
}

// SetFuses gets a reference to the given int32 and assigns it to the Fuses field.
func (o *ServicesRnewByAdminResp) SetFuses(v int32) {
	o.Fuses = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServicesRnewByAdminResp) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServicesRnewByAdminResp) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ServicesRnewByAdminResp) SetId(v int32) {
	o.Id = &v
}

// GetTxError returns the TxError field value if set, zero value otherwise.
func (o *ServicesRnewByAdminResp) GetTxError() string {
	if o == nil || isNil(o.TxError) {
		var ret string
		return ret
	}
	return *o.TxError
}

// GetTxErrorOk returns a tuple with the TxError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetTxErrorOk() (*string, bool) {
	if o == nil || isNil(o.TxError) {
    return nil, false
	}
	return o.TxError, true
}

// HasTxError returns a boolean if a field has been set.
func (o *ServicesRnewByAdminResp) HasTxError() bool {
	if o != nil && !isNil(o.TxError) {
		return true
	}

	return false
}

// SetTxError gets a reference to the given string and assigns it to the TxError field.
func (o *ServicesRnewByAdminResp) SetTxError(v string) {
	o.TxError = &v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *ServicesRnewByAdminResp) GetTxHash() string {
	if o == nil || isNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetTxHashOk() (*string, bool) {
	if o == nil || isNil(o.TxHash) {
    return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *ServicesRnewByAdminResp) HasTxHash() bool {
	if o != nil && !isNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *ServicesRnewByAdminResp) SetTxHash(v string) {
	o.TxHash = &v
}

// GetTxState returns the TxState field value if set, zero value otherwise.
func (o *ServicesRnewByAdminResp) GetTxState() int32 {
	if o == nil || isNil(o.TxState) {
		var ret int32
		return ret
	}
	return *o.TxState
}

// GetTxStateOk returns a tuple with the TxState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetTxStateOk() (*int32, bool) {
	if o == nil || isNil(o.TxState) {
    return nil, false
	}
	return o.TxState, true
}

// HasTxState returns a boolean if a field has been set.
func (o *ServicesRnewByAdminResp) HasTxState() bool {
	if o != nil && !isNil(o.TxState) {
		return true
	}

	return false
}

// SetTxState gets a reference to the given int32 and assigns it to the TxState field.
func (o *ServicesRnewByAdminResp) SetTxState(v int32) {
	o.TxState = &v
}

// GetWrapperExpiry returns the WrapperExpiry field value
func (o *ServicesRnewByAdminResp) GetWrapperExpiry() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WrapperExpiry
}

// GetWrapperExpiryOk returns a tuple with the WrapperExpiry field value
// and a boolean to check if the value has been set.
func (o *ServicesRnewByAdminResp) GetWrapperExpiryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WrapperExpiry, true
}

// SetWrapperExpiry sets field value
func (o *ServicesRnewByAdminResp) SetWrapperExpiry(v int32) {
	o.WrapperExpiry = v
}

func (o ServicesRnewByAdminResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cns_name"] = o.CnsName
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Fuses) {
		toSerialize["fuses"] = o.Fuses
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.TxError) {
		toSerialize["tx_error"] = o.TxError
	}
	if !isNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	if !isNil(o.TxState) {
		toSerialize["tx_state"] = o.TxState
	}
	if true {
		toSerialize["wrapper_expiry"] = o.WrapperExpiry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicesRnewByAdminResp) UnmarshalJSON(bytes []byte) (err error) {
	varServicesRnewByAdminResp := _ServicesRnewByAdminResp{}

	if err = json.Unmarshal(bytes, &varServicesRnewByAdminResp); err == nil {
		*o = ServicesRnewByAdminResp(varServicesRnewByAdminResp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cns_name")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "fuses")
		delete(additionalProperties, "id")
		delete(additionalProperties, "tx_error")
		delete(additionalProperties, "tx_hash")
		delete(additionalProperties, "tx_state")
		delete(additionalProperties, "wrapper_expiry")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicesRnewByAdminResp struct {
	value *ServicesRnewByAdminResp
	isSet bool
}

func (v NullableServicesRnewByAdminResp) Get() *ServicesRnewByAdminResp {
	return v.value
}

func (v *NullableServicesRnewByAdminResp) Set(val *ServicesRnewByAdminResp) {
	v.value = val
	v.isSet = true
}

func (v NullableServicesRnewByAdminResp) IsSet() bool {
	return v.isSet
}

func (v *NullableServicesRnewByAdminResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicesRnewByAdminResp(val *ServicesRnewByAdminResp) *NullableServicesRnewByAdminResp {
	return &NullableServicesRnewByAdminResp{value: val, isSet: true}
}

func (v NullableServicesRnewByAdminResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicesRnewByAdminResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

