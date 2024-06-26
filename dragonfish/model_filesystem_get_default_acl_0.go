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

// FilesystemGetDefaultAcl0 the model 'FilesystemGetDefaultAcl0'
type FilesystemGetDefaultAcl0 string

// List of filesystem_get_default_acl_0
const (
	OPEN        FilesystemGetDefaultAcl0 = "OPEN"
	RESTRICTED  FilesystemGetDefaultAcl0 = "RESTRICTED"
	HOME        FilesystemGetDefaultAcl0 = "HOME"
	DOMAIN_HOME FilesystemGetDefaultAcl0 = "DOMAIN_HOME"
)

// All allowed values of FilesystemGetDefaultAcl0 enum
var AllowedFilesystemGetDefaultAcl0EnumValues = []FilesystemGetDefaultAcl0{
	"OPEN",
	"RESTRICTED",
	"HOME",
	"DOMAIN_HOME",
}

func (v *FilesystemGetDefaultAcl0) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilesystemGetDefaultAcl0(value)
	for _, existing := range AllowedFilesystemGetDefaultAcl0EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilesystemGetDefaultAcl0", value)
}

// NewFilesystemGetDefaultAcl0FromValue returns a pointer to a valid FilesystemGetDefaultAcl0
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilesystemGetDefaultAcl0FromValue(v string) (*FilesystemGetDefaultAcl0, error) {
	ev := FilesystemGetDefaultAcl0(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilesystemGetDefaultAcl0: valid values are %v", v, AllowedFilesystemGetDefaultAcl0EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilesystemGetDefaultAcl0) IsValid() bool {
	for _, existing := range AllowedFilesystemGetDefaultAcl0EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to filesystem_get_default_acl_0 value
func (v FilesystemGetDefaultAcl0) Ptr() *FilesystemGetDefaultAcl0 {
	return &v
}

type NullableFilesystemGetDefaultAcl0 struct {
	value *FilesystemGetDefaultAcl0
	isSet bool
}

func (v NullableFilesystemGetDefaultAcl0) Get() *FilesystemGetDefaultAcl0 {
	return v.value
}

func (v *NullableFilesystemGetDefaultAcl0) Set(val *FilesystemGetDefaultAcl0) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemGetDefaultAcl0) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemGetDefaultAcl0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemGetDefaultAcl0(val *FilesystemGetDefaultAcl0) *NullableFilesystemGetDefaultAcl0 {
	return &NullableFilesystemGetDefaultAcl0{value: val, isSet: true}
}

func (v NullableFilesystemGetDefaultAcl0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemGetDefaultAcl0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
