/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dragonfish

import (
	"encoding/json"
	"fmt"
)

// SystemFeatureEnabled0 the model 'SystemFeatureEnabled0'
type SystemFeatureEnabled0 string

// List of system_feature_enabled_0
const (
	DEDUP        SystemFeatureEnabled0 = "DEDUP"
	FIBRECHANNEL SystemFeatureEnabled0 = "FIBRECHANNEL"
	JAILS        SystemFeatureEnabled0 = "JAILS"
	VM           SystemFeatureEnabled0 = "VM"
)

// All allowed values of SystemFeatureEnabled0 enum
var AllowedSystemFeatureEnabled0EnumValues = []SystemFeatureEnabled0{
	"DEDUP",
	"FIBRECHANNEL",
	"JAILS",
	"VM",
}

func (v *SystemFeatureEnabled0) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemFeatureEnabled0(value)
	for _, existing := range AllowedSystemFeatureEnabled0EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemFeatureEnabled0", value)
}

// NewSystemFeatureEnabled0FromValue returns a pointer to a valid SystemFeatureEnabled0
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemFeatureEnabled0FromValue(v string) (*SystemFeatureEnabled0, error) {
	ev := SystemFeatureEnabled0(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemFeatureEnabled0: valid values are %v", v, AllowedSystemFeatureEnabled0EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemFeatureEnabled0) IsValid() bool {
	for _, existing := range AllowedSystemFeatureEnabled0EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to system_feature_enabled_0 value
func (v SystemFeatureEnabled0) Ptr() *SystemFeatureEnabled0 {
	return &v
}

type NullableSystemFeatureEnabled0 struct {
	value *SystemFeatureEnabled0
	isSet bool
}

func (v NullableSystemFeatureEnabled0) Get() *SystemFeatureEnabled0 {
	return v.value
}

func (v *NullableSystemFeatureEnabled0) Set(val *SystemFeatureEnabled0) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemFeatureEnabled0) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemFeatureEnabled0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemFeatureEnabled0(val *SystemFeatureEnabled0) *NullableSystemFeatureEnabled0 {
	return &NullableSystemFeatureEnabled0{value: val, isSet: true}
}

func (v NullableSystemFeatureEnabled0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemFeatureEnabled0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
