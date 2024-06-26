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

// StaticrouteCreate0 struct for StaticrouteCreate0
type StaticrouteCreate0 struct {
	Destination *string `json:"destination,omitempty"`
	Gateway     *string `json:"gateway,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewStaticrouteCreate0 instantiates a new StaticrouteCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticrouteCreate0() *StaticrouteCreate0 {
	this := StaticrouteCreate0{}
	return &this
}

// NewStaticrouteCreate0WithDefaults instantiates a new StaticrouteCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticrouteCreate0WithDefaults() *StaticrouteCreate0 {
	this := StaticrouteCreate0{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *StaticrouteCreate0) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticrouteCreate0) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *StaticrouteCreate0) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *StaticrouteCreate0) SetDestination(v string) {
	o.Destination = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *StaticrouteCreate0) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticrouteCreate0) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *StaticrouteCreate0) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *StaticrouteCreate0) SetGateway(v string) {
	o.Gateway = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StaticrouteCreate0) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticrouteCreate0) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StaticrouteCreate0) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StaticrouteCreate0) SetDescription(v string) {
	o.Description = &v
}

func (o StaticrouteCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableStaticrouteCreate0 struct {
	value *StaticrouteCreate0
	isSet bool
}

func (v NullableStaticrouteCreate0) Get() *StaticrouteCreate0 {
	return v.value
}

func (v *NullableStaticrouteCreate0) Set(val *StaticrouteCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticrouteCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticrouteCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticrouteCreate0(val *StaticrouteCreate0) *NullableStaticrouteCreate0 {
	return &NullableStaticrouteCreate0{value: val, isSet: true}
}

func (v NullableStaticrouteCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticrouteCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
