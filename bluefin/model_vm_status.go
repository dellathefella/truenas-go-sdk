/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// VMStatus struct for VMStatus
type VMStatus struct {
	State *string `json:"state,omitempty"`
	Pid *int32 `json:"pid,omitempty"`
	DomainState *string `json:"domain_state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VMStatus VMStatus

// NewVMStatus instantiates a new VMStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMStatus() *VMStatus {
	this := VMStatus{}
	return &this
}

// NewVMStatusWithDefaults instantiates a new VMStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMStatusWithDefaults() *VMStatus {
	this := VMStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VMStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VMStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VMStatus) SetState(v string) {
	o.State = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *VMStatus) GetPid() int32 {
	if o == nil || o.Pid == nil {
		var ret int32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMStatus) GetPidOk() (*int32, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *VMStatus) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given int32 and assigns it to the Pid field.
func (o *VMStatus) SetPid(v int32) {
	o.Pid = &v
}

// GetDomainState returns the DomainState field value if set, zero value otherwise.
func (o *VMStatus) GetDomainState() string {
	if o == nil || o.DomainState == nil {
		var ret string
		return ret
	}
	return *o.DomainState
}

// GetDomainStateOk returns a tuple with the DomainState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMStatus) GetDomainStateOk() (*string, bool) {
	if o == nil || o.DomainState == nil {
		return nil, false
	}
	return o.DomainState, true
}

// HasDomainState returns a boolean if a field has been set.
func (o *VMStatus) HasDomainState() bool {
	if o != nil && o.DomainState != nil {
		return true
	}

	return false
}

// SetDomainState gets a reference to the given string and assigns it to the DomainState field.
func (o *VMStatus) SetDomainState(v string) {
	o.DomainState = &v
}

func (o VMStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Pid != nil {
		toSerialize["pid"] = o.Pid
	}
	if o.DomainState != nil {
		toSerialize["domain_state"] = o.DomainState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VMStatus) UnmarshalJSON(bytes []byte) (err error) {
	varVMStatus := _VMStatus{}

	if err = json.Unmarshal(bytes, &varVMStatus); err == nil {
		*o = VMStatus(varVMStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "pid")
		delete(additionalProperties, "domain_state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMStatus struct {
	value *VMStatus
	isSet bool
}

func (v NullableVMStatus) Get() *VMStatus {
	return v.value
}

func (v *NullableVMStatus) Set(val *VMStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVMStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVMStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMStatus(val *VMStatus) *NullableVMStatus {
	return &NullableVMStatus{value: val, isSet: true}
}

func (v NullableVMStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


