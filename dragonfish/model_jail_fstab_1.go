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

// JailFstab1 struct for JailFstab1
type JailFstab1 struct {
	Action      *string `json:"action,omitempty"`
	Source      *string `json:"source,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Fstype      *string `json:"fstype,omitempty"`
	Fsoptions   *string `json:"fsoptions,omitempty"`
	Dump        *string `json:"dump,omitempty"`
	Pass        *string `json:"pass,omitempty"`
	Index       *int32  `json:"index,omitempty"`
}

// NewJailFstab1 instantiates a new JailFstab1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJailFstab1() *JailFstab1 {
	this := JailFstab1{}
	return &this
}

// NewJailFstab1WithDefaults instantiates a new JailFstab1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJailFstab1WithDefaults() *JailFstab1 {
	this := JailFstab1{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *JailFstab1) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *JailFstab1) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *JailFstab1) SetAction(v string) {
	o.Action = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JailFstab1) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JailFstab1) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *JailFstab1) SetSource(v string) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *JailFstab1) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *JailFstab1) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *JailFstab1) SetDestination(v string) {
	o.Destination = &v
}

// GetFstype returns the Fstype field value if set, zero value otherwise.
func (o *JailFstab1) GetFstype() string {
	if o == nil || o.Fstype == nil {
		var ret string
		return ret
	}
	return *o.Fstype
}

// GetFstypeOk returns a tuple with the Fstype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetFstypeOk() (*string, bool) {
	if o == nil || o.Fstype == nil {
		return nil, false
	}
	return o.Fstype, true
}

// HasFstype returns a boolean if a field has been set.
func (o *JailFstab1) HasFstype() bool {
	if o != nil && o.Fstype != nil {
		return true
	}

	return false
}

// SetFstype gets a reference to the given string and assigns it to the Fstype field.
func (o *JailFstab1) SetFstype(v string) {
	o.Fstype = &v
}

// GetFsoptions returns the Fsoptions field value if set, zero value otherwise.
func (o *JailFstab1) GetFsoptions() string {
	if o == nil || o.Fsoptions == nil {
		var ret string
		return ret
	}
	return *o.Fsoptions
}

// GetFsoptionsOk returns a tuple with the Fsoptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetFsoptionsOk() (*string, bool) {
	if o == nil || o.Fsoptions == nil {
		return nil, false
	}
	return o.Fsoptions, true
}

// HasFsoptions returns a boolean if a field has been set.
func (o *JailFstab1) HasFsoptions() bool {
	if o != nil && o.Fsoptions != nil {
		return true
	}

	return false
}

// SetFsoptions gets a reference to the given string and assigns it to the Fsoptions field.
func (o *JailFstab1) SetFsoptions(v string) {
	o.Fsoptions = &v
}

// GetDump returns the Dump field value if set, zero value otherwise.
func (o *JailFstab1) GetDump() string {
	if o == nil || o.Dump == nil {
		var ret string
		return ret
	}
	return *o.Dump
}

// GetDumpOk returns a tuple with the Dump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetDumpOk() (*string, bool) {
	if o == nil || o.Dump == nil {
		return nil, false
	}
	return o.Dump, true
}

// HasDump returns a boolean if a field has been set.
func (o *JailFstab1) HasDump() bool {
	if o != nil && o.Dump != nil {
		return true
	}

	return false
}

// SetDump gets a reference to the given string and assigns it to the Dump field.
func (o *JailFstab1) SetDump(v string) {
	o.Dump = &v
}

// GetPass returns the Pass field value if set, zero value otherwise.
func (o *JailFstab1) GetPass() string {
	if o == nil || o.Pass == nil {
		var ret string
		return ret
	}
	return *o.Pass
}

// GetPassOk returns a tuple with the Pass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetPassOk() (*string, bool) {
	if o == nil || o.Pass == nil {
		return nil, false
	}
	return o.Pass, true
}

// HasPass returns a boolean if a field has been set.
func (o *JailFstab1) HasPass() bool {
	if o != nil && o.Pass != nil {
		return true
	}

	return false
}

// SetPass gets a reference to the given string and assigns it to the Pass field.
func (o *JailFstab1) SetPass(v string) {
	o.Pass = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *JailFstab1) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailFstab1) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *JailFstab1) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *JailFstab1) SetIndex(v int32) {
	o.Index = &v
}

func (o JailFstab1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Fstype != nil {
		toSerialize["fstype"] = o.Fstype
	}
	if o.Fsoptions != nil {
		toSerialize["fsoptions"] = o.Fsoptions
	}
	if o.Dump != nil {
		toSerialize["dump"] = o.Dump
	}
	if o.Pass != nil {
		toSerialize["pass"] = o.Pass
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableJailFstab1 struct {
	value *JailFstab1
	isSet bool
}

func (v NullableJailFstab1) Get() *JailFstab1 {
	return v.value
}

func (v *NullableJailFstab1) Set(val *JailFstab1) {
	v.value = val
	v.isSet = true
}

func (v NullableJailFstab1) IsSet() bool {
	return v.isSet
}

func (v *NullableJailFstab1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJailFstab1(val *JailFstab1) *NullableJailFstab1 {
	return &NullableJailFstab1{value: val, isSet: true}
}

func (v NullableJailFstab1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJailFstab1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
