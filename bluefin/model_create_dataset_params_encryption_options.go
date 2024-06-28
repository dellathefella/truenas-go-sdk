/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// CreateDatasetParamsEncryptionOptions struct for CreateDatasetParamsEncryptionOptions
type CreateDatasetParamsEncryptionOptions struct {
	Algorithm *string `json:"algorithm,omitempty"`
	GenerateKey *bool `json:"generate_key,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
	Key *string `json:"key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDatasetParamsEncryptionOptions CreateDatasetParamsEncryptionOptions

// NewCreateDatasetParamsEncryptionOptions instantiates a new CreateDatasetParamsEncryptionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatasetParamsEncryptionOptions() *CreateDatasetParamsEncryptionOptions {
	this := CreateDatasetParamsEncryptionOptions{}
	return &this
}

// NewCreateDatasetParamsEncryptionOptionsWithDefaults instantiates a new CreateDatasetParamsEncryptionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatasetParamsEncryptionOptionsWithDefaults() *CreateDatasetParamsEncryptionOptions {
	this := CreateDatasetParamsEncryptionOptions{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *CreateDatasetParamsEncryptionOptions) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParamsEncryptionOptions) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *CreateDatasetParamsEncryptionOptions) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *CreateDatasetParamsEncryptionOptions) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetGenerateKey returns the GenerateKey field value if set, zero value otherwise.
func (o *CreateDatasetParamsEncryptionOptions) GetGenerateKey() bool {
	if o == nil || o.GenerateKey == nil {
		var ret bool
		return ret
	}
	return *o.GenerateKey
}

// GetGenerateKeyOk returns a tuple with the GenerateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParamsEncryptionOptions) GetGenerateKeyOk() (*bool, bool) {
	if o == nil || o.GenerateKey == nil {
		return nil, false
	}
	return o.GenerateKey, true
}

// HasGenerateKey returns a boolean if a field has been set.
func (o *CreateDatasetParamsEncryptionOptions) HasGenerateKey() bool {
	if o != nil && o.GenerateKey != nil {
		return true
	}

	return false
}

// SetGenerateKey gets a reference to the given bool and assigns it to the GenerateKey field.
func (o *CreateDatasetParamsEncryptionOptions) SetGenerateKey(v bool) {
	o.GenerateKey = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *CreateDatasetParamsEncryptionOptions) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParamsEncryptionOptions) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *CreateDatasetParamsEncryptionOptions) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *CreateDatasetParamsEncryptionOptions) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateDatasetParamsEncryptionOptions) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParamsEncryptionOptions) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateDatasetParamsEncryptionOptions) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateDatasetParamsEncryptionOptions) SetKey(v string) {
	o.Key = &v
}

func (o CreateDatasetParamsEncryptionOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.GenerateKey != nil {
		toSerialize["generate_key"] = o.GenerateKey
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateDatasetParamsEncryptionOptions) UnmarshalJSON(bytes []byte) (err error) {
	varCreateDatasetParamsEncryptionOptions := _CreateDatasetParamsEncryptionOptions{}

	if err = json.Unmarshal(bytes, &varCreateDatasetParamsEncryptionOptions); err == nil {
		*o = CreateDatasetParamsEncryptionOptions(varCreateDatasetParamsEncryptionOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "generate_key")
		delete(additionalProperties, "passphrase")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDatasetParamsEncryptionOptions struct {
	value *CreateDatasetParamsEncryptionOptions
	isSet bool
}

func (v NullableCreateDatasetParamsEncryptionOptions) Get() *CreateDatasetParamsEncryptionOptions {
	return v.value
}

func (v *NullableCreateDatasetParamsEncryptionOptions) Set(val *CreateDatasetParamsEncryptionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatasetParamsEncryptionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatasetParamsEncryptionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatasetParamsEncryptionOptions(val *CreateDatasetParamsEncryptionOptions) *NullableCreateDatasetParamsEncryptionOptions {
	return &NullableCreateDatasetParamsEncryptionOptions{value: val, isSet: true}
}

func (v NullableCreateDatasetParamsEncryptionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatasetParamsEncryptionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


