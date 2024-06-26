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

// PoolDatasetEncryptionSummary1 struct for PoolDatasetEncryptionSummary1
type PoolDatasetEncryptionSummary1 struct {
	KeyFile  *bool                    `json:"key_file,omitempty"`
	Datasets []map[string]interface{} `json:"datasets,omitempty"`
}

// NewPoolDatasetEncryptionSummary1 instantiates a new PoolDatasetEncryptionSummary1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetEncryptionSummary1() *PoolDatasetEncryptionSummary1 {
	this := PoolDatasetEncryptionSummary1{}
	return &this
}

// NewPoolDatasetEncryptionSummary1WithDefaults instantiates a new PoolDatasetEncryptionSummary1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetEncryptionSummary1WithDefaults() *PoolDatasetEncryptionSummary1 {
	this := PoolDatasetEncryptionSummary1{}
	return &this
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *PoolDatasetEncryptionSummary1) GetKeyFile() bool {
	if o == nil || o.KeyFile == nil {
		var ret bool
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetEncryptionSummary1) GetKeyFileOk() (*bool, bool) {
	if o == nil || o.KeyFile == nil {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *PoolDatasetEncryptionSummary1) HasKeyFile() bool {
	if o != nil && o.KeyFile != nil {
		return true
	}

	return false
}

// SetKeyFile gets a reference to the given bool and assigns it to the KeyFile field.
func (o *PoolDatasetEncryptionSummary1) SetKeyFile(v bool) {
	o.KeyFile = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *PoolDatasetEncryptionSummary1) GetDatasets() []map[string]interface{} {
	if o == nil || o.Datasets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetEncryptionSummary1) GetDatasetsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *PoolDatasetEncryptionSummary1) HasDatasets() bool {
	if o != nil && o.Datasets != nil {
		return true
	}

	return false
}

// SetDatasets gets a reference to the given []map[string]interface{} and assigns it to the Datasets field.
func (o *PoolDatasetEncryptionSummary1) SetDatasets(v []map[string]interface{}) {
	o.Datasets = v
}

func (o PoolDatasetEncryptionSummary1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyFile != nil {
		toSerialize["key_file"] = o.KeyFile
	}
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetEncryptionSummary1 struct {
	value *PoolDatasetEncryptionSummary1
	isSet bool
}

func (v NullablePoolDatasetEncryptionSummary1) Get() *PoolDatasetEncryptionSummary1 {
	return v.value
}

func (v *NullablePoolDatasetEncryptionSummary1) Set(val *PoolDatasetEncryptionSummary1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetEncryptionSummary1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetEncryptionSummary1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetEncryptionSummary1(val *PoolDatasetEncryptionSummary1) *NullablePoolDatasetEncryptionSummary1 {
	return &NullablePoolDatasetEncryptionSummary1{value: val, isSet: true}
}

func (v NullablePoolDatasetEncryptionSummary1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetEncryptionSummary1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
