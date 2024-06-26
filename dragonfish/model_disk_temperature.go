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

// DiskTemperature struct for DiskTemperature
type DiskTemperature struct {
	Name                 *string           `json:"name,omitempty"`
	Powermode            *DiskTemperature1 `json:"powermode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiskTemperature DiskTemperature

// NewDiskTemperature instantiates a new DiskTemperature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskTemperature() *DiskTemperature {
	this := DiskTemperature{}
	var powermode DiskTemperature1 = NEVER
	this.Powermode = &powermode
	return &this
}

// NewDiskTemperatureWithDefaults instantiates a new DiskTemperature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskTemperatureWithDefaults() *DiskTemperature {
	this := DiskTemperature{}
	var powermode DiskTemperature1 = NEVER
	this.Powermode = &powermode
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiskTemperature) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskTemperature) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiskTemperature) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiskTemperature) SetName(v string) {
	o.Name = &v
}

// GetPowermode returns the Powermode field value if set, zero value otherwise.
func (o *DiskTemperature) GetPowermode() DiskTemperature1 {
	if o == nil || o.Powermode == nil {
		var ret DiskTemperature1
		return ret
	}
	return *o.Powermode
}

// GetPowermodeOk returns a tuple with the Powermode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskTemperature) GetPowermodeOk() (*DiskTemperature1, bool) {
	if o == nil || o.Powermode == nil {
		return nil, false
	}
	return o.Powermode, true
}

// HasPowermode returns a boolean if a field has been set.
func (o *DiskTemperature) HasPowermode() bool {
	if o != nil && o.Powermode != nil {
		return true
	}

	return false
}

// SetPowermode gets a reference to the given DiskTemperature1 and assigns it to the Powermode field.
func (o *DiskTemperature) SetPowermode(v DiskTemperature1) {
	o.Powermode = &v
}

func (o DiskTemperature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Powermode != nil {
		toSerialize["powermode"] = o.Powermode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DiskTemperature) UnmarshalJSON(bytes []byte) (err error) {
	varDiskTemperature := _DiskTemperature{}

	if err = json.Unmarshal(bytes, &varDiskTemperature); err == nil {
		*o = DiskTemperature(varDiskTemperature)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "powermode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiskTemperature struct {
	value *DiskTemperature
	isSet bool
}

func (v NullableDiskTemperature) Get() *DiskTemperature {
	return v.value
}

func (v *NullableDiskTemperature) Set(val *DiskTemperature) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskTemperature) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskTemperature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskTemperature(val *DiskTemperature) *NullableDiskTemperature {
	return &NullableDiskTemperature{value: val, isSet: true}
}

func (v NullableDiskTemperature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskTemperature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
