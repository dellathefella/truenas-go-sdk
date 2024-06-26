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

// SmbStatus2 struct for SmbStatus2
type SmbStatus2 struct {
	Relationships   *bool                  `json:"relationships,omitempty"`
	Extend          NullableString         `json:"extend,omitempty"`
	ExtendContext   NullableString         `json:"extend_context,omitempty"`
	Prefix          NullableString         `json:"prefix,omitempty"`
	Extra           map[string]interface{} `json:"extra,omitempty"`
	OrderBy         []interface{}          `json:"order_by,omitempty"`
	Select          []interface{}          `json:"select,omitempty"`
	Count           *bool                  `json:"count,omitempty"`
	Get             *bool                  `json:"get,omitempty"`
	Offset          *int32                 `json:"offset,omitempty"`
	Limit           *int32                 `json:"limit,omitempty"`
	ForceSqlFilters *bool                  `json:"force_sql_filters,omitempty"`
}

// NewSmbStatus2 instantiates a new SmbStatus2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStatus2() *SmbStatus2 {
	this := SmbStatus2{}
	return &this
}

// NewSmbStatus2WithDefaults instantiates a new SmbStatus2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStatus2WithDefaults() *SmbStatus2 {
	this := SmbStatus2{}
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SmbStatus2) GetRelationships() bool {
	if o == nil || o.Relationships == nil {
		var ret bool
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetRelationshipsOk() (*bool, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SmbStatus2) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given bool and assigns it to the Relationships field.
func (o *SmbStatus2) SetRelationships(v bool) {
	o.Relationships = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbStatus2) GetExtend() string {
	if o == nil || o.Extend.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extend.Get()
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbStatus2) GetExtendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extend.Get(), o.Extend.IsSet()
}

// HasExtend returns a boolean if a field has been set.
func (o *SmbStatus2) HasExtend() bool {
	if o != nil && o.Extend.IsSet() {
		return true
	}

	return false
}

// SetExtend gets a reference to the given NullableString and assigns it to the Extend field.
func (o *SmbStatus2) SetExtend(v string) {
	o.Extend.Set(&v)
}

// SetExtendNil sets the value for Extend to be an explicit nil
func (o *SmbStatus2) SetExtendNil() {
	o.Extend.Set(nil)
}

// UnsetExtend ensures that no value is present for Extend, not even an explicit nil
func (o *SmbStatus2) UnsetExtend() {
	o.Extend.Unset()
}

// GetExtendContext returns the ExtendContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbStatus2) GetExtendContext() string {
	if o == nil || o.ExtendContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtendContext.Get()
}

// GetExtendContextOk returns a tuple with the ExtendContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbStatus2) GetExtendContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtendContext.Get(), o.ExtendContext.IsSet()
}

// HasExtendContext returns a boolean if a field has been set.
func (o *SmbStatus2) HasExtendContext() bool {
	if o != nil && o.ExtendContext.IsSet() {
		return true
	}

	return false
}

// SetExtendContext gets a reference to the given NullableString and assigns it to the ExtendContext field.
func (o *SmbStatus2) SetExtendContext(v string) {
	o.ExtendContext.Set(&v)
}

// SetExtendContextNil sets the value for ExtendContext to be an explicit nil
func (o *SmbStatus2) SetExtendContextNil() {
	o.ExtendContext.Set(nil)
}

// UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
func (o *SmbStatus2) UnsetExtendContext() {
	o.ExtendContext.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmbStatus2) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmbStatus2) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *SmbStatus2) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *SmbStatus2) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *SmbStatus2) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *SmbStatus2) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *SmbStatus2) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetExtraOk() (map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *SmbStatus2) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *SmbStatus2) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *SmbStatus2) GetOrderBy() []interface{} {
	if o == nil || o.OrderBy == nil {
		var ret []interface{}
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetOrderByOk() ([]interface{}, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *SmbStatus2) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []interface{} and assigns it to the OrderBy field.
func (o *SmbStatus2) SetOrderBy(v []interface{}) {
	o.OrderBy = v
}

// GetSelect returns the Select field value if set, zero value otherwise.
func (o *SmbStatus2) GetSelect() []interface{} {
	if o == nil || o.Select == nil {
		var ret []interface{}
		return ret
	}
	return o.Select
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetSelectOk() ([]interface{}, bool) {
	if o == nil || o.Select == nil {
		return nil, false
	}
	return o.Select, true
}

// HasSelect returns a boolean if a field has been set.
func (o *SmbStatus2) HasSelect() bool {
	if o != nil && o.Select != nil {
		return true
	}

	return false
}

// SetSelect gets a reference to the given []interface{} and assigns it to the Select field.
func (o *SmbStatus2) SetSelect(v []interface{}) {
	o.Select = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SmbStatus2) GetCount() bool {
	if o == nil || o.Count == nil {
		var ret bool
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetCountOk() (*bool, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SmbStatus2) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given bool and assigns it to the Count field.
func (o *SmbStatus2) SetCount(v bool) {
	o.Count = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *SmbStatus2) GetGet() bool {
	if o == nil || o.Get == nil {
		var ret bool
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetGetOk() (*bool, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *SmbStatus2) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given bool and assigns it to the Get field.
func (o *SmbStatus2) SetGet(v bool) {
	o.Get = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SmbStatus2) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SmbStatus2) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SmbStatus2) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SmbStatus2) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SmbStatus2) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SmbStatus2) SetLimit(v int32) {
	o.Limit = &v
}

// GetForceSqlFilters returns the ForceSqlFilters field value if set, zero value otherwise.
func (o *SmbStatus2) GetForceSqlFilters() bool {
	if o == nil || o.ForceSqlFilters == nil {
		var ret bool
		return ret
	}
	return *o.ForceSqlFilters
}

// GetForceSqlFiltersOk returns a tuple with the ForceSqlFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus2) GetForceSqlFiltersOk() (*bool, bool) {
	if o == nil || o.ForceSqlFilters == nil {
		return nil, false
	}
	return o.ForceSqlFilters, true
}

// HasForceSqlFilters returns a boolean if a field has been set.
func (o *SmbStatus2) HasForceSqlFilters() bool {
	if o != nil && o.ForceSqlFilters != nil {
		return true
	}

	return false
}

// SetForceSqlFilters gets a reference to the given bool and assigns it to the ForceSqlFilters field.
func (o *SmbStatus2) SetForceSqlFilters(v bool) {
	o.ForceSqlFilters = &v
}

func (o SmbStatus2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Extend.IsSet() {
		toSerialize["extend"] = o.Extend.Get()
	}
	if o.ExtendContext.IsSet() {
		toSerialize["extend_context"] = o.ExtendContext.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.Select != nil {
		toSerialize["select"] = o.Select
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ForceSqlFilters != nil {
		toSerialize["force_sql_filters"] = o.ForceSqlFilters
	}
	return json.Marshal(toSerialize)
}

type NullableSmbStatus2 struct {
	value *SmbStatus2
	isSet bool
}

func (v NullableSmbStatus2) Get() *SmbStatus2 {
	return v.value
}

func (v *NullableSmbStatus2) Set(val *SmbStatus2) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStatus2) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStatus2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStatus2(val *SmbStatus2) *NullableSmbStatus2 {
	return &NullableSmbStatus2{value: val, isSet: true}
}

func (v NullableSmbStatus2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStatus2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
