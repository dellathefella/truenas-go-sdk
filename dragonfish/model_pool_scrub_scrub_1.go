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

// PoolScrubScrub1 the model 'PoolScrubScrub1'
type PoolScrubScrub1 string

// List of pool_scrub_scrub_1
const (
	START PoolScrubScrub1 = "START"
	STOP  PoolScrubScrub1 = "STOP"
	PAUSE PoolScrubScrub1 = "PAUSE"
)

// All allowed values of PoolScrubScrub1 enum
var AllowedPoolScrubScrub1EnumValues = []PoolScrubScrub1{
	"START",
	"STOP",
	"PAUSE",
}

func (v *PoolScrubScrub1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PoolScrubScrub1(value)
	for _, existing := range AllowedPoolScrubScrub1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PoolScrubScrub1", value)
}

// NewPoolScrubScrub1FromValue returns a pointer to a valid PoolScrubScrub1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPoolScrubScrub1FromValue(v string) (*PoolScrubScrub1, error) {
	ev := PoolScrubScrub1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PoolScrubScrub1: valid values are %v", v, AllowedPoolScrubScrub1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PoolScrubScrub1) IsValid() bool {
	for _, existing := range AllowedPoolScrubScrub1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to pool_scrub_scrub_1 value
func (v PoolScrubScrub1) Ptr() *PoolScrubScrub1 {
	return &v
}

type NullablePoolScrubScrub1 struct {
	value *PoolScrubScrub1
	isSet bool
}

func (v NullablePoolScrubScrub1) Get() *PoolScrubScrub1 {
	return v.value
}

func (v *NullablePoolScrubScrub1) Set(val *PoolScrubScrub1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolScrubScrub1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolScrubScrub1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolScrubScrub1(val *PoolScrubScrub1) *NullablePoolScrubScrub1 {
	return &NullablePoolScrubScrub1{value: val, isSet: true}
}

func (v NullablePoolScrubScrub1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolScrubScrub1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
