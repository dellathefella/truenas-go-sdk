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

// FailoverSyncToPeer0 struct for FailoverSyncToPeer0
type FailoverSyncToPeer0 struct {
	Reboot *bool `json:"reboot,omitempty"`
}

// NewFailoverSyncToPeer0 instantiates a new FailoverSyncToPeer0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverSyncToPeer0() *FailoverSyncToPeer0 {
	this := FailoverSyncToPeer0{}
	return &this
}

// NewFailoverSyncToPeer0WithDefaults instantiates a new FailoverSyncToPeer0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverSyncToPeer0WithDefaults() *FailoverSyncToPeer0 {
	this := FailoverSyncToPeer0{}
	return &this
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *FailoverSyncToPeer0) GetReboot() bool {
	if o == nil || o.Reboot == nil {
		var ret bool
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverSyncToPeer0) GetRebootOk() (*bool, bool) {
	if o == nil || o.Reboot == nil {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *FailoverSyncToPeer0) HasReboot() bool {
	if o != nil && o.Reboot != nil {
		return true
	}

	return false
}

// SetReboot gets a reference to the given bool and assigns it to the Reboot field.
func (o *FailoverSyncToPeer0) SetReboot(v bool) {
	o.Reboot = &v
}

func (o FailoverSyncToPeer0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reboot != nil {
		toSerialize["reboot"] = o.Reboot
	}
	return json.Marshal(toSerialize)
}

type NullableFailoverSyncToPeer0 struct {
	value *FailoverSyncToPeer0
	isSet bool
}

func (v NullableFailoverSyncToPeer0) Get() *FailoverSyncToPeer0 {
	return v.value
}

func (v *NullableFailoverSyncToPeer0) Set(val *FailoverSyncToPeer0) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverSyncToPeer0) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverSyncToPeer0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverSyncToPeer0(val *FailoverSyncToPeer0) *NullableFailoverSyncToPeer0 {
	return &NullableFailoverSyncToPeer0{value: val, isSet: true}
}

func (v NullableFailoverSyncToPeer0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverSyncToPeer0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
