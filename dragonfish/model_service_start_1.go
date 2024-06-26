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

// ServiceStart1 struct for ServiceStart1
type ServiceStart1 struct {
	HaPropagate *bool `json:"ha_propagate,omitempty"`
}

// NewServiceStart1 instantiates a new ServiceStart1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStart1() *ServiceStart1 {
	this := ServiceStart1{}
	return &this
}

// NewServiceStart1WithDefaults instantiates a new ServiceStart1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStart1WithDefaults() *ServiceStart1 {
	this := ServiceStart1{}
	return &this
}

// GetHaPropagate returns the HaPropagate field value if set, zero value otherwise.
func (o *ServiceStart1) GetHaPropagate() bool {
	if o == nil || o.HaPropagate == nil {
		var ret bool
		return ret
	}
	return *o.HaPropagate
}

// GetHaPropagateOk returns a tuple with the HaPropagate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStart1) GetHaPropagateOk() (*bool, bool) {
	if o == nil || o.HaPropagate == nil {
		return nil, false
	}
	return o.HaPropagate, true
}

// HasHaPropagate returns a boolean if a field has been set.
func (o *ServiceStart1) HasHaPropagate() bool {
	if o != nil && o.HaPropagate != nil {
		return true
	}

	return false
}

// SetHaPropagate gets a reference to the given bool and assigns it to the HaPropagate field.
func (o *ServiceStart1) SetHaPropagate(v bool) {
	o.HaPropagate = &v
}

func (o ServiceStart1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HaPropagate != nil {
		toSerialize["ha_propagate"] = o.HaPropagate
	}
	return json.Marshal(toSerialize)
}

type NullableServiceStart1 struct {
	value *ServiceStart1
	isSet bool
}

func (v NullableServiceStart1) Get() *ServiceStart1 {
	return v.value
}

func (v *NullableServiceStart1) Set(val *ServiceStart1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStart1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStart1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStart1(val *ServiceStart1) *NullableServiceStart1 {
	return &NullableServiceStart1{value: val, isSet: true}
}

func (v NullableServiceStart1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStart1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
