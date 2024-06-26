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

// NfsAddPrincipal0 struct for NfsAddPrincipal0
type NfsAddPrincipal0 struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewNfsAddPrincipal0 instantiates a new NfsAddPrincipal0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsAddPrincipal0() *NfsAddPrincipal0 {
	this := NfsAddPrincipal0{}
	return &this
}

// NewNfsAddPrincipal0WithDefaults instantiates a new NfsAddPrincipal0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsAddPrincipal0WithDefaults() *NfsAddPrincipal0 {
	this := NfsAddPrincipal0{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NfsAddPrincipal0) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsAddPrincipal0) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NfsAddPrincipal0) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NfsAddPrincipal0) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NfsAddPrincipal0) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsAddPrincipal0) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NfsAddPrincipal0) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NfsAddPrincipal0) SetPassword(v string) {
	o.Password = &v
}

func (o NfsAddPrincipal0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableNfsAddPrincipal0 struct {
	value *NfsAddPrincipal0
	isSet bool
}

func (v NullableNfsAddPrincipal0) Get() *NfsAddPrincipal0 {
	return v.value
}

func (v *NullableNfsAddPrincipal0) Set(val *NfsAddPrincipal0) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsAddPrincipal0) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsAddPrincipal0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsAddPrincipal0(val *NfsAddPrincipal0) *NullableNfsAddPrincipal0 {
	return &NullableNfsAddPrincipal0{value: val, isSet: true}
}

func (v NullableNfsAddPrincipal0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsAddPrincipal0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
