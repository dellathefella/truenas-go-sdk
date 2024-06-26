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

// IscsiAuthUpdate1 struct for IscsiAuthUpdate1
type IscsiAuthUpdate1 struct {
	Tag        *int32  `json:"tag,omitempty"`
	User       *string `json:"user,omitempty"`
	Secret     *string `json:"secret,omitempty"`
	Peeruser   *string `json:"peeruser,omitempty"`
	Peersecret *string `json:"peersecret,omitempty"`
}

// NewIscsiAuthUpdate1 instantiates a new IscsiAuthUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiAuthUpdate1() *IscsiAuthUpdate1 {
	this := IscsiAuthUpdate1{}
	return &this
}

// NewIscsiAuthUpdate1WithDefaults instantiates a new IscsiAuthUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiAuthUpdate1WithDefaults() *IscsiAuthUpdate1 {
	this := IscsiAuthUpdate1{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *IscsiAuthUpdate1) GetTag() int32 {
	if o == nil || o.Tag == nil {
		var ret int32
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiAuthUpdate1) GetTagOk() (*int32, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *IscsiAuthUpdate1) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given int32 and assigns it to the Tag field.
func (o *IscsiAuthUpdate1) SetTag(v int32) {
	o.Tag = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IscsiAuthUpdate1) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiAuthUpdate1) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IscsiAuthUpdate1) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *IscsiAuthUpdate1) SetUser(v string) {
	o.User = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IscsiAuthUpdate1) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiAuthUpdate1) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IscsiAuthUpdate1) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IscsiAuthUpdate1) SetSecret(v string) {
	o.Secret = &v
}

// GetPeeruser returns the Peeruser field value if set, zero value otherwise.
func (o *IscsiAuthUpdate1) GetPeeruser() string {
	if o == nil || o.Peeruser == nil {
		var ret string
		return ret
	}
	return *o.Peeruser
}

// GetPeeruserOk returns a tuple with the Peeruser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiAuthUpdate1) GetPeeruserOk() (*string, bool) {
	if o == nil || o.Peeruser == nil {
		return nil, false
	}
	return o.Peeruser, true
}

// HasPeeruser returns a boolean if a field has been set.
func (o *IscsiAuthUpdate1) HasPeeruser() bool {
	if o != nil && o.Peeruser != nil {
		return true
	}

	return false
}

// SetPeeruser gets a reference to the given string and assigns it to the Peeruser field.
func (o *IscsiAuthUpdate1) SetPeeruser(v string) {
	o.Peeruser = &v
}

// GetPeersecret returns the Peersecret field value if set, zero value otherwise.
func (o *IscsiAuthUpdate1) GetPeersecret() string {
	if o == nil || o.Peersecret == nil {
		var ret string
		return ret
	}
	return *o.Peersecret
}

// GetPeersecretOk returns a tuple with the Peersecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiAuthUpdate1) GetPeersecretOk() (*string, bool) {
	if o == nil || o.Peersecret == nil {
		return nil, false
	}
	return o.Peersecret, true
}

// HasPeersecret returns a boolean if a field has been set.
func (o *IscsiAuthUpdate1) HasPeersecret() bool {
	if o != nil && o.Peersecret != nil {
		return true
	}

	return false
}

// SetPeersecret gets a reference to the given string and assigns it to the Peersecret field.
func (o *IscsiAuthUpdate1) SetPeersecret(v string) {
	o.Peersecret = &v
}

func (o IscsiAuthUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Peeruser != nil {
		toSerialize["peeruser"] = o.Peeruser
	}
	if o.Peersecret != nil {
		toSerialize["peersecret"] = o.Peersecret
	}
	return json.Marshal(toSerialize)
}

type NullableIscsiAuthUpdate1 struct {
	value *IscsiAuthUpdate1
	isSet bool
}

func (v NullableIscsiAuthUpdate1) Get() *IscsiAuthUpdate1 {
	return v.value
}

func (v *NullableIscsiAuthUpdate1) Set(val *IscsiAuthUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiAuthUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiAuthUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiAuthUpdate1(val *IscsiAuthUpdate1) *NullableIscsiAuthUpdate1 {
	return &NullableIscsiAuthUpdate1{value: val, isSet: true}
}

func (v NullableIscsiAuthUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiAuthUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
