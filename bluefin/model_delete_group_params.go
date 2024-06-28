/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// DeleteGroupParams struct for DeleteGroupParams
type DeleteGroupParams struct {
	DeleteUsers *bool `json:"delete_users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteGroupParams DeleteGroupParams

// NewDeleteGroupParams instantiates a new DeleteGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteGroupParams() *DeleteGroupParams {
	this := DeleteGroupParams{}
	return &this
}

// NewDeleteGroupParamsWithDefaults instantiates a new DeleteGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteGroupParamsWithDefaults() *DeleteGroupParams {
	this := DeleteGroupParams{}
	return &this
}

// GetDeleteUsers returns the DeleteUsers field value if set, zero value otherwise.
func (o *DeleteGroupParams) GetDeleteUsers() bool {
	if o == nil || o.DeleteUsers == nil {
		var ret bool
		return ret
	}
	return *o.DeleteUsers
}

// GetDeleteUsersOk returns a tuple with the DeleteUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteGroupParams) GetDeleteUsersOk() (*bool, bool) {
	if o == nil || o.DeleteUsers == nil {
		return nil, false
	}
	return o.DeleteUsers, true
}

// HasDeleteUsers returns a boolean if a field has been set.
func (o *DeleteGroupParams) HasDeleteUsers() bool {
	if o != nil && o.DeleteUsers != nil {
		return true
	}

	return false
}

// SetDeleteUsers gets a reference to the given bool and assigns it to the DeleteUsers field.
func (o *DeleteGroupParams) SetDeleteUsers(v bool) {
	o.DeleteUsers = &v
}

func (o DeleteGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteUsers != nil {
		toSerialize["delete_users"] = o.DeleteUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteGroupParams) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteGroupParams := _DeleteGroupParams{}

	if err = json.Unmarshal(bytes, &varDeleteGroupParams); err == nil {
		*o = DeleteGroupParams(varDeleteGroupParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delete_users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteGroupParams struct {
	value *DeleteGroupParams
	isSet bool
}

func (v NullableDeleteGroupParams) Get() *DeleteGroupParams {
	return v.value
}

func (v *NullableDeleteGroupParams) Set(val *DeleteGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteGroupParams(val *DeleteGroupParams) *NullableDeleteGroupParams {
	return &NullableDeleteGroupParams{value: val, isSet: true}
}

func (v NullableDeleteGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


