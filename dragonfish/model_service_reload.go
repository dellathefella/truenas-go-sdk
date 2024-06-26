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

// ServiceReload struct for ServiceReload
type ServiceReload struct {
	Service *string `json:"service,omitempty"`
	ServiceControl *ServiceReload1 `json:"service-control,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceReload ServiceReload

// NewServiceReload instantiates a new ServiceReload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceReload() *ServiceReload {
	this := ServiceReload{}
	var serviceControl ServiceReload1 = {}
	this.ServiceControl = &serviceControl
	return &this
}

// NewServiceReloadWithDefaults instantiates a new ServiceReload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceReloadWithDefaults() *ServiceReload {
	this := ServiceReload{}
	var serviceControl ServiceReload1 = {}
	this.ServiceControl = &serviceControl
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceReload) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceReload) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceReload) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceReload) SetService(v string) {
	o.Service = &v
}

// GetServiceControl returns the ServiceControl field value if set, zero value otherwise.
func (o *ServiceReload) GetServiceControl() ServiceReload1 {
	if o == nil || o.ServiceControl == nil {
		var ret ServiceReload1
		return ret
	}
	return *o.ServiceControl
}

// GetServiceControlOk returns a tuple with the ServiceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceReload) GetServiceControlOk() (*ServiceReload1, bool) {
	if o == nil || o.ServiceControl == nil {
		return nil, false
	}
	return o.ServiceControl, true
}

// HasServiceControl returns a boolean if a field has been set.
func (o *ServiceReload) HasServiceControl() bool {
	if o != nil && o.ServiceControl != nil {
		return true
	}

	return false
}

// SetServiceControl gets a reference to the given ServiceReload1 and assigns it to the ServiceControl field.
func (o *ServiceReload) SetServiceControl(v ServiceReload1) {
	o.ServiceControl = &v
}

func (o ServiceReload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceControl != nil {
		toSerialize["service-control"] = o.ServiceControl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceReload) UnmarshalJSON(bytes []byte) (err error) {
	varServiceReload := _ServiceReload{}

	if err = json.Unmarshal(bytes, &varServiceReload); err == nil {
		*o = ServiceReload(varServiceReload)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		delete(additionalProperties, "service-control")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceReload struct {
	value *ServiceReload
	isSet bool
}

func (v NullableServiceReload) Get() *ServiceReload {
	return v.value
}

func (v *NullableServiceReload) Set(val *ServiceReload) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceReload) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceReload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceReload(val *ServiceReload) *NullableServiceReload {
	return &NullableServiceReload{value: val, isSet: true}
}

func (v NullableServiceReload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceReload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


