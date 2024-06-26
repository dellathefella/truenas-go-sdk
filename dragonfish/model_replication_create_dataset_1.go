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

// ReplicationCreateDataset1 the model 'ReplicationCreateDataset1'
type ReplicationCreateDataset1 string

// List of replication_create_dataset_1
const (
	SSH       ReplicationCreateDataset1 = "SSH"
	SSHNETCAT ReplicationCreateDataset1 = "SSH+NETCAT"
	LOCAL     ReplicationCreateDataset1 = "LOCAL"
)

// All allowed values of ReplicationCreateDataset1 enum
var AllowedReplicationCreateDataset1EnumValues = []ReplicationCreateDataset1{
	"SSH",
	"SSH+NETCAT",
	"LOCAL",
}

func (v *ReplicationCreateDataset1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReplicationCreateDataset1(value)
	for _, existing := range AllowedReplicationCreateDataset1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReplicationCreateDataset1", value)
}

// NewReplicationCreateDataset1FromValue returns a pointer to a valid ReplicationCreateDataset1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReplicationCreateDataset1FromValue(v string) (*ReplicationCreateDataset1, error) {
	ev := ReplicationCreateDataset1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReplicationCreateDataset1: valid values are %v", v, AllowedReplicationCreateDataset1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReplicationCreateDataset1) IsValid() bool {
	for _, existing := range AllowedReplicationCreateDataset1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to replication_create_dataset_1 value
func (v ReplicationCreateDataset1) Ptr() *ReplicationCreateDataset1 {
	return &v
}

type NullableReplicationCreateDataset1 struct {
	value *ReplicationCreateDataset1
	isSet bool
}

func (v NullableReplicationCreateDataset1) Get() *ReplicationCreateDataset1 {
	return v.value
}

func (v *NullableReplicationCreateDataset1) Set(val *ReplicationCreateDataset1) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationCreateDataset1) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationCreateDataset1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationCreateDataset1(val *ReplicationCreateDataset1) *NullableReplicationCreateDataset1 {
	return &NullableReplicationCreateDataset1{value: val, isSet: true}
}

func (v NullableReplicationCreateDataset1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationCreateDataset1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
