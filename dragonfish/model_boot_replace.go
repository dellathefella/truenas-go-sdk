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

// BootReplace struct for BootReplace
type BootReplace struct {
	Label                *string `json:"label,omitempty"`
	Dev                  *string `json:"dev,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootReplace BootReplace

// NewBootReplace instantiates a new BootReplace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootReplace() *BootReplace {
	this := BootReplace{}
	return &this
}

// NewBootReplaceWithDefaults instantiates a new BootReplace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootReplaceWithDefaults() *BootReplace {
	this := BootReplace{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BootReplace) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootReplace) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BootReplace) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BootReplace) SetLabel(v string) {
	o.Label = &v
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *BootReplace) GetDev() string {
	if o == nil || o.Dev == nil {
		var ret string
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootReplace) GetDevOk() (*string, bool) {
	if o == nil || o.Dev == nil {
		return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *BootReplace) HasDev() bool {
	if o != nil && o.Dev != nil {
		return true
	}

	return false
}

// SetDev gets a reference to the given string and assigns it to the Dev field.
func (o *BootReplace) SetDev(v string) {
	o.Dev = &v
}

func (o BootReplace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Dev != nil {
		toSerialize["dev"] = o.Dev
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootReplace) UnmarshalJSON(bytes []byte) (err error) {
	varBootReplace := _BootReplace{}

	if err = json.Unmarshal(bytes, &varBootReplace); err == nil {
		*o = BootReplace(varBootReplace)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "dev")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootReplace struct {
	value *BootReplace
	isSet bool
}

func (v NullableBootReplace) Get() *BootReplace {
	return v.value
}

func (v *NullableBootReplace) Set(val *BootReplace) {
	v.value = val
	v.isSet = true
}

func (v NullableBootReplace) IsSet() bool {
	return v.isSet
}

func (v *NullableBootReplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootReplace(val *BootReplace) *NullableBootReplace {
	return &NullableBootReplace{value: val, isSet: true}
}

func (v NullableBootReplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootReplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
