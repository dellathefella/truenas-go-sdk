/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dragonfish

import (
	"encoding/json"
)

// VmStop1 struct for VmStop1
type VmStop1 struct {
	Force             *bool `json:"force,omitempty"`
	ForceAfterTimeout *bool `json:"force_after_timeout,omitempty"`
}

// NewVmStop1 instantiates a new VmStop1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmStop1() *VmStop1 {
	this := VmStop1{}
	return &this
}

// NewVmStop1WithDefaults instantiates a new VmStop1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmStop1WithDefaults() *VmStop1 {
	this := VmStop1{}
	return &this
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *VmStop1) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStop1) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *VmStop1) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *VmStop1) SetForce(v bool) {
	o.Force = &v
}

// GetForceAfterTimeout returns the ForceAfterTimeout field value if set, zero value otherwise.
func (o *VmStop1) GetForceAfterTimeout() bool {
	if o == nil || o.ForceAfterTimeout == nil {
		var ret bool
		return ret
	}
	return *o.ForceAfterTimeout
}

// GetForceAfterTimeoutOk returns a tuple with the ForceAfterTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStop1) GetForceAfterTimeoutOk() (*bool, bool) {
	if o == nil || o.ForceAfterTimeout == nil {
		return nil, false
	}
	return o.ForceAfterTimeout, true
}

// HasForceAfterTimeout returns a boolean if a field has been set.
func (o *VmStop1) HasForceAfterTimeout() bool {
	if o != nil && o.ForceAfterTimeout != nil {
		return true
	}

	return false
}

// SetForceAfterTimeout gets a reference to the given bool and assigns it to the ForceAfterTimeout field.
func (o *VmStop1) SetForceAfterTimeout(v bool) {
	o.ForceAfterTimeout = &v
}

func (o VmStop1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if o.ForceAfterTimeout != nil {
		toSerialize["force_after_timeout"] = o.ForceAfterTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableVmStop1 struct {
	value *VmStop1
	isSet bool
}

func (v NullableVmStop1) Get() *VmStop1 {
	return v.value
}

func (v *NullableVmStop1) Set(val *VmStop1) {
	v.value = val
	v.isSet = true
}

func (v NullableVmStop1) IsSet() bool {
	return v.isSet
}

func (v *NullableVmStop1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmStop1(val *VmStop1) *NullableVmStop1 {
	return &NullableVmStop1{value: val, isSet: true}
}

func (v NullableVmStop1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmStop1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
