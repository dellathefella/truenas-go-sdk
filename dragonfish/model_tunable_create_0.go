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

// TunableCreate0 struct for TunableCreate0
type TunableCreate0 struct {
	Var     *string `json:"var,omitempty"`
	Value   *string `json:"value,omitempty"`
	Type    *string `json:"type,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
}

// NewTunableCreate0 instantiates a new TunableCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunableCreate0() *TunableCreate0 {
	this := TunableCreate0{}
	return &this
}

// NewTunableCreate0WithDefaults instantiates a new TunableCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunableCreate0WithDefaults() *TunableCreate0 {
	this := TunableCreate0{}
	return &this
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *TunableCreate0) GetVar() string {
	if o == nil || o.Var == nil {
		var ret string
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetVarOk() (*string, bool) {
	if o == nil || o.Var == nil {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *TunableCreate0) HasVar() bool {
	if o != nil && o.Var != nil {
		return true
	}

	return false
}

// SetVar gets a reference to the given string and assigns it to the Var field.
func (o *TunableCreate0) SetVar(v string) {
	o.Var = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TunableCreate0) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TunableCreate0) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TunableCreate0) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TunableCreate0) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TunableCreate0) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TunableCreate0) SetType(v string) {
	o.Type = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TunableCreate0) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TunableCreate0) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TunableCreate0) SetComment(v string) {
	o.Comment = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TunableCreate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunableCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TunableCreate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TunableCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o TunableCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Var != nil {
		toSerialize["var"] = o.Var
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableTunableCreate0 struct {
	value *TunableCreate0
	isSet bool
}

func (v NullableTunableCreate0) Get() *TunableCreate0 {
	return v.value
}

func (v *NullableTunableCreate0) Set(val *TunableCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableTunableCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableTunableCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunableCreate0(val *TunableCreate0) *NullableTunableCreate0 {
	return &NullableTunableCreate0{value: val, isSet: true}
}

func (v NullableTunableCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunableCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
