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

// Callback struct for Callback
type Callback struct {
	[]interface{} *[]interface{}
	bool *bool
	int32 *int32
	map[string]interface{} *map[string]interface{}
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Callback) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into []interface{}
	err = json.Unmarshal(data, &dst.[]interface{});
	if err == nil {
		json[]interface{}, _ := json.Marshal(dst.[]interface{})
		if string(json[]interface{}) == "{}" { // empty struct
			dst.[]interface{} = nil
		} else {
			return nil // data stored in dst.[]interface{}, return on the first match
		}
	} else {
		dst.[]interface{} = nil
	}

	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool);
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.map[string]interface{});
	if err == nil {
		jsonmap[string]interface{}, _ := json.Marshal(dst.map[string]interface{})
		if string(jsonmap[string]interface{}) == "{}" { // empty struct
			dst.map[string]interface{} = nil
		} else {
			return nil // data stored in dst.map[string]interface{}, return on the first match
		}
	} else {
		dst.map[string]interface{} = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(Callback)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Callback) MarshalJSON() ([]byte, error) {
	if src.[]interface{} != nil {
		return json.Marshal(&src.[]interface{})
	}

	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCallback struct {
	value *Callback
	isSet bool
}

func (v NullableCallback) Get() *Callback {
	return v.value
}

func (v *NullableCallback) Set(val *Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallback(val *Callback) *NullableCallback {
	return &NullableCallback{value: val, isSet: true}
}

func (v NullableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


