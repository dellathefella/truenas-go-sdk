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

// GroupDelete1 struct for GroupDelete1
type GroupDelete1 struct {
	DeleteUsers *bool `json:"delete_users,omitempty"`
}

// NewGroupDelete1 instantiates a new GroupDelete1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupDelete1() *GroupDelete1 {
	this := GroupDelete1{}
	return &this
}

// NewGroupDelete1WithDefaults instantiates a new GroupDelete1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupDelete1WithDefaults() *GroupDelete1 {
	this := GroupDelete1{}
	return &this
}

// GetDeleteUsers returns the DeleteUsers field value if set, zero value otherwise.
func (o *GroupDelete1) GetDeleteUsers() bool {
	if o == nil || o.DeleteUsers == nil {
		var ret bool
		return ret
	}
	return *o.DeleteUsers
}

// GetDeleteUsersOk returns a tuple with the DeleteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupDelete1) GetDeleteUsersOk() (*bool, bool) {
	if o == nil || o.DeleteUsers == nil {
		return nil, false
	}
	return o.DeleteUsers, true
}

// HasDeleteUsers returns a boolean if a field has been set.
func (o *GroupDelete1) HasDeleteUsers() bool {
	if o != nil && o.DeleteUsers != nil {
		return true
	}

	return false
}

// SetDeleteUsers gets a reference to the given bool and assigns it to the DeleteUsers field.
func (o *GroupDelete1) SetDeleteUsers(v bool) {
	o.DeleteUsers = &v
}

func (o GroupDelete1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteUsers != nil {
		toSerialize["delete_users"] = o.DeleteUsers
	}
	return json.Marshal(toSerialize)
}

type NullableGroupDelete1 struct {
	value *GroupDelete1
	isSet bool
}

func (v NullableGroupDelete1) Get() *GroupDelete1 {
	return v.value
}

func (v *NullableGroupDelete1) Set(val *GroupDelete1) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupDelete1) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupDelete1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupDelete1(val *GroupDelete1) *NullableGroupDelete1 {
	return &NullableGroupDelete1{value: val, isSet: true}
}

func (v NullableGroupDelete1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupDelete1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
