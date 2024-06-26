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

// AlertclassesUpdate0 struct for AlertclassesUpdate0
type AlertclassesUpdate0 struct {
	Classes map[string]interface{} `json:"classes,omitempty"`
}

// NewAlertclassesUpdate0 instantiates a new AlertclassesUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertclassesUpdate0() *AlertclassesUpdate0 {
	this := AlertclassesUpdate0{}
	return &this
}

// NewAlertclassesUpdate0WithDefaults instantiates a new AlertclassesUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertclassesUpdate0WithDefaults() *AlertclassesUpdate0 {
	this := AlertclassesUpdate0{}
	return &this
}

// GetClasses returns the Classes field value if set, zero value otherwise.
func (o *AlertclassesUpdate0) GetClasses() map[string]interface{} {
	if o == nil || o.Classes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Classes
}

// GetClassesOk returns a tuple with the Classes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertclassesUpdate0) GetClassesOk() (map[string]interface{}, bool) {
	if o == nil || o.Classes == nil {
		return nil, false
	}
	return o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *AlertclassesUpdate0) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given map[string]interface{} and assigns it to the Classes field.
func (o *AlertclassesUpdate0) SetClasses(v map[string]interface{}) {
	o.Classes = v
}

func (o AlertclassesUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Classes != nil {
		toSerialize["classes"] = o.Classes
	}
	return json.Marshal(toSerialize)
}

type NullableAlertclassesUpdate0 struct {
	value *AlertclassesUpdate0
	isSet bool
}

func (v NullableAlertclassesUpdate0) Get() *AlertclassesUpdate0 {
	return v.value
}

func (v *NullableAlertclassesUpdate0) Set(val *AlertclassesUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertclassesUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertclassesUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertclassesUpdate0(val *AlertclassesUpdate0) *NullableAlertclassesUpdate0 {
	return &NullableAlertclassesUpdate0{value: val, isSet: true}
}

func (v NullableAlertclassesUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertclassesUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
