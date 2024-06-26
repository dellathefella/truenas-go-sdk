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

// BootenvSetAttribute1 struct for BootenvSetAttribute1
type BootenvSetAttribute1 struct {
	Keep *bool `json:"keep,omitempty"`
}

// NewBootenvSetAttribute1 instantiates a new BootenvSetAttribute1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootenvSetAttribute1() *BootenvSetAttribute1 {
	this := BootenvSetAttribute1{}
	return &this
}

// NewBootenvSetAttribute1WithDefaults instantiates a new BootenvSetAttribute1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootenvSetAttribute1WithDefaults() *BootenvSetAttribute1 {
	this := BootenvSetAttribute1{}
	return &this
}

// GetKeep returns the Keep field value if set, zero value otherwise.
func (o *BootenvSetAttribute1) GetKeep() bool {
	if o == nil || o.Keep == nil {
		var ret bool
		return ret
	}
	return *o.Keep
}

// GetKeepOk returns a tuple with the Keep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootenvSetAttribute1) GetKeepOk() (*bool, bool) {
	if o == nil || o.Keep == nil {
		return nil, false
	}
	return o.Keep, true
}

// HasKeep returns a boolean if a field has been set.
func (o *BootenvSetAttribute1) HasKeep() bool {
	if o != nil && o.Keep != nil {
		return true
	}

	return false
}

// SetKeep gets a reference to the given bool and assigns it to the Keep field.
func (o *BootenvSetAttribute1) SetKeep(v bool) {
	o.Keep = &v
}

func (o BootenvSetAttribute1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keep != nil {
		toSerialize["keep"] = o.Keep
	}
	return json.Marshal(toSerialize)
}

type NullableBootenvSetAttribute1 struct {
	value *BootenvSetAttribute1
	isSet bool
}

func (v NullableBootenvSetAttribute1) Get() *BootenvSetAttribute1 {
	return v.value
}

func (v *NullableBootenvSetAttribute1) Set(val *BootenvSetAttribute1) {
	v.value = val
	v.isSet = true
}

func (v NullableBootenvSetAttribute1) IsSet() bool {
	return v.isSet
}

func (v *NullableBootenvSetAttribute1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootenvSetAttribute1(val *BootenvSetAttribute1) *NullableBootenvSetAttribute1 {
	return &NullableBootenvSetAttribute1{value: val, isSet: true}
}

func (v NullableBootenvSetAttribute1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootenvSetAttribute1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
