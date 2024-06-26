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

// SystemShutdown0 struct for SystemShutdown0
type SystemShutdown0 struct {
	Delay *int32 `json:"delay,omitempty"`
}

// NewSystemShutdown0 instantiates a new SystemShutdown0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemShutdown0() *SystemShutdown0 {
	this := SystemShutdown0{}
	return &this
}

// NewSystemShutdown0WithDefaults instantiates a new SystemShutdown0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemShutdown0WithDefaults() *SystemShutdown0 {
	this := SystemShutdown0{}
	return &this
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *SystemShutdown0) GetDelay() int32 {
	if o == nil || o.Delay == nil {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemShutdown0) GetDelayOk() (*int32, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *SystemShutdown0) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *SystemShutdown0) SetDelay(v int32) {
	o.Delay = &v
}

func (o SystemShutdown0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	return json.Marshal(toSerialize)
}

type NullableSystemShutdown0 struct {
	value *SystemShutdown0
	isSet bool
}

func (v NullableSystemShutdown0) Get() *SystemShutdown0 {
	return v.value
}

func (v *NullableSystemShutdown0) Set(val *SystemShutdown0) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemShutdown0) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemShutdown0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemShutdown0(val *SystemShutdown0) *NullableSystemShutdown0 {
	return &NullableSystemShutdown0{value: val, isSet: true}
}

func (v NullableSystemShutdown0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemShutdown0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
