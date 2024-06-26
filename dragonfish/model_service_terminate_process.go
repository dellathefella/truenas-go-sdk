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

// ServiceTerminateProcess struct for ServiceTerminateProcess
type ServiceTerminateProcess struct {
	Pid                  *int32 `json:"pid,omitempty"`
	Timeout              *int32 `json:"timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceTerminateProcess ServiceTerminateProcess

// NewServiceTerminateProcess instantiates a new ServiceTerminateProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTerminateProcess() *ServiceTerminateProcess {
	this := ServiceTerminateProcess{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// NewServiceTerminateProcessWithDefaults instantiates a new ServiceTerminateProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTerminateProcessWithDefaults() *ServiceTerminateProcess {
	this := ServiceTerminateProcess{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *ServiceTerminateProcess) GetPid() int32 {
	if o == nil || o.Pid == nil {
		var ret int32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTerminateProcess) GetPidOk() (*int32, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *ServiceTerminateProcess) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given int32 and assigns it to the Pid field.
func (o *ServiceTerminateProcess) SetPid(v int32) {
	o.Pid = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ServiceTerminateProcess) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTerminateProcess) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ServiceTerminateProcess) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *ServiceTerminateProcess) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o ServiceTerminateProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pid != nil {
		toSerialize["pid"] = o.Pid
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceTerminateProcess) UnmarshalJSON(bytes []byte) (err error) {
	varServiceTerminateProcess := _ServiceTerminateProcess{}

	if err = json.Unmarshal(bytes, &varServiceTerminateProcess); err == nil {
		*o = ServiceTerminateProcess(varServiceTerminateProcess)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pid")
		delete(additionalProperties, "timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceTerminateProcess struct {
	value *ServiceTerminateProcess
	isSet bool
}

func (v NullableServiceTerminateProcess) Get() *ServiceTerminateProcess {
	return v.value
}

func (v *NullableServiceTerminateProcess) Set(val *ServiceTerminateProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTerminateProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTerminateProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTerminateProcess(val *ServiceTerminateProcess) *NullableServiceTerminateProcess {
	return &NullableServiceTerminateProcess{value: val, isSet: true}
}

func (v NullableServiceTerminateProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTerminateProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
