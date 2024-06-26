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

// JailImportImage0 struct for JailImportImage0
type JailImportImage0 struct {
	Jail                 *string        `json:"jail,omitempty"`
	Path                 NullableString `json:"path,omitempty"`
	CompressionAlgorithm NullableString `json:"compression_algorithm,omitempty"`
}

// NewJailImportImage0 instantiates a new JailImportImage0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJailImportImage0() *JailImportImage0 {
	this := JailImportImage0{}
	return &this
}

// NewJailImportImage0WithDefaults instantiates a new JailImportImage0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJailImportImage0WithDefaults() *JailImportImage0 {
	this := JailImportImage0{}
	return &this
}

// GetJail returns the Jail field value if set, zero value otherwise.
func (o *JailImportImage0) GetJail() string {
	if o == nil || o.Jail == nil {
		var ret string
		return ret
	}
	return *o.Jail
}

// GetJailOk returns a tuple with the Jail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailImportImage0) GetJailOk() (*string, bool) {
	if o == nil || o.Jail == nil {
		return nil, false
	}
	return o.Jail, true
}

// HasJail returns a boolean if a field has been set.
func (o *JailImportImage0) HasJail() bool {
	if o != nil && o.Jail != nil {
		return true
	}

	return false
}

// SetJail gets a reference to the given string and assigns it to the Jail field.
func (o *JailImportImage0) SetJail(v string) {
	o.Jail = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JailImportImage0) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JailImportImage0) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *JailImportImage0) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *JailImportImage0) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *JailImportImage0) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *JailImportImage0) UnsetPath() {
	o.Path.Unset()
}

// GetCompressionAlgorithm returns the CompressionAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JailImportImage0) GetCompressionAlgorithm() string {
	if o == nil || o.CompressionAlgorithm.Get() == nil {
		var ret string
		return ret
	}
	return *o.CompressionAlgorithm.Get()
}

// GetCompressionAlgorithmOk returns a tuple with the CompressionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JailImportImage0) GetCompressionAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompressionAlgorithm.Get(), o.CompressionAlgorithm.IsSet()
}

// HasCompressionAlgorithm returns a boolean if a field has been set.
func (o *JailImportImage0) HasCompressionAlgorithm() bool {
	if o != nil && o.CompressionAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetCompressionAlgorithm gets a reference to the given NullableString and assigns it to the CompressionAlgorithm field.
func (o *JailImportImage0) SetCompressionAlgorithm(v string) {
	o.CompressionAlgorithm.Set(&v)
}

// SetCompressionAlgorithmNil sets the value for CompressionAlgorithm to be an explicit nil
func (o *JailImportImage0) SetCompressionAlgorithmNil() {
	o.CompressionAlgorithm.Set(nil)
}

// UnsetCompressionAlgorithm ensures that no value is present for CompressionAlgorithm, not even an explicit nil
func (o *JailImportImage0) UnsetCompressionAlgorithm() {
	o.CompressionAlgorithm.Unset()
}

func (o JailImportImage0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jail != nil {
		toSerialize["jail"] = o.Jail
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.CompressionAlgorithm.IsSet() {
		toSerialize["compression_algorithm"] = o.CompressionAlgorithm.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJailImportImage0 struct {
	value *JailImportImage0
	isSet bool
}

func (v NullableJailImportImage0) Get() *JailImportImage0 {
	return v.value
}

func (v *NullableJailImportImage0) Set(val *JailImportImage0) {
	v.value = val
	v.isSet = true
}

func (v NullableJailImportImage0) IsSet() bool {
	return v.isSet
}

func (v *NullableJailImportImage0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJailImportImage0(val *JailImportImage0) *NullableJailImportImage0 {
	return &NullableJailImportImage0{value: val, isSet: true}
}

func (v NullableJailImportImage0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJailImportImage0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
