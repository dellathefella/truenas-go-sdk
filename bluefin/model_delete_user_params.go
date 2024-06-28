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

// DeleteUserParams struct for DeleteUserParams
type DeleteUserParams struct {
	DeleteGroup *bool `json:"delete_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteUserParams DeleteUserParams

// NewDeleteUserParams instantiates a new DeleteUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserParams() *DeleteUserParams {
	this := DeleteUserParams{}
	return &this
}

// NewDeleteUserParamsWithDefaults instantiates a new DeleteUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserParamsWithDefaults() *DeleteUserParams {
	this := DeleteUserParams{}
	return &this
}

// GetDeleteGroup returns the DeleteGroup field value if set, zero value otherwise.
func (o *DeleteUserParams) GetDeleteGroup() bool {
	if o == nil || o.DeleteGroup == nil {
		var ret bool
		return ret
	}
	return *o.DeleteGroup
}

// GetDeleteGroupOk returns a tuple with the DeleteGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteUserParams) GetDeleteGroupOk() (*bool, bool) {
	if o == nil || o.DeleteGroup == nil {
		return nil, false
	}
	return o.DeleteGroup, true
}

// HasDeleteGroup returns a boolean if a field has been set.
func (o *DeleteUserParams) HasDeleteGroup() bool {
	if o != nil && o.DeleteGroup != nil {
		return true
	}

	return false
}

// SetDeleteGroup gets a reference to the given bool and assigns it to the DeleteGroup field.
func (o *DeleteUserParams) SetDeleteGroup(v bool) {
	o.DeleteGroup = &v
}

func (o DeleteUserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteGroup != nil {
		toSerialize["delete_group"] = o.DeleteGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteUserParams) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteUserParams := _DeleteUserParams{}

	if err = json.Unmarshal(bytes, &varDeleteUserParams); err == nil {
		*o = DeleteUserParams(varDeleteUserParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delete_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteUserParams struct {
	value *DeleteUserParams
	isSet bool
}

func (v NullableDeleteUserParams) Get() *DeleteUserParams {
	return v.value
}

func (v *NullableDeleteUserParams) Set(val *DeleteUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserParams(val *DeleteUserParams) *NullableDeleteUserParams {
	return &NullableDeleteUserParams{value: val, isSet: true}
}

func (v NullableDeleteUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


