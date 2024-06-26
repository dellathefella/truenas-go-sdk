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

// DeviceGetInfo0 the model 'DeviceGetInfo0'
type DeviceGetInfo0 string

// List of device_get_info_0
const (
	SERIAL DeviceGetInfo0 = "SERIAL"
	DISK   DeviceGetInfo0 = "DISK"
)

// All allowed values of DeviceGetInfo0 enum
var AllowedDeviceGetInfo0EnumValues = []DeviceGetInfo0{
	"SERIAL",
	"DISK",
}

func (v *DeviceGetInfo0) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceGetInfo0(value)
	for _, existing := range AllowedDeviceGetInfo0EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceGetInfo0", value)
}

// NewDeviceGetInfo0FromValue returns a pointer to a valid DeviceGetInfo0
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceGetInfo0FromValue(v string) (*DeviceGetInfo0, error) {
	ev := DeviceGetInfo0(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceGetInfo0: valid values are %v", v, AllowedDeviceGetInfo0EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceGetInfo0) IsValid() bool {
	for _, existing := range AllowedDeviceGetInfo0EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to device_get_info_0 value
func (v DeviceGetInfo0) Ptr() *DeviceGetInfo0 {
	return &v
}

type NullableDeviceGetInfo0 struct {
	value *DeviceGetInfo0
	isSet bool
}

func (v NullableDeviceGetInfo0) Get() *DeviceGetInfo0 {
	return v.value
}

func (v *NullableDeviceGetInfo0) Set(val *DeviceGetInfo0) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGetInfo0) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGetInfo0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGetInfo0(val *DeviceGetInfo0) *NullableDeviceGetInfo0 {
	return &NullableDeviceGetInfo0{value: val, isSet: true}
}

func (v NullableDeviceGetInfo0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGetInfo0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
