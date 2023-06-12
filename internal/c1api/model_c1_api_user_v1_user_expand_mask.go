/*
ConductorOne API

The ConductorOne API is a HTTP API for managing ConductorOne resources.

API version: 0.1.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package c1api

import (
	"encoding/json"
)

// checks if the C1ApiUserV1UserExpandMask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiUserV1UserExpandMask{}

// C1ApiUserV1UserExpandMask The UserExpandMask message.
type C1ApiUserV1UserExpandMask struct {
	// The paths field.
	Paths []string `json:"paths,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiUserV1UserExpandMask C1ApiUserV1UserExpandMask

// NewC1ApiUserV1UserExpandMask instantiates a new C1ApiUserV1UserExpandMask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiUserV1UserExpandMask() *C1ApiUserV1UserExpandMask {
	this := C1ApiUserV1UserExpandMask{}
	return &this
}

// NewC1ApiUserV1UserExpandMaskWithDefaults instantiates a new C1ApiUserV1UserExpandMask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiUserV1UserExpandMaskWithDefaults() *C1ApiUserV1UserExpandMask {
	this := C1ApiUserV1UserExpandMask{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserExpandMask) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserExpandMask) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *C1ApiUserV1UserExpandMask) HasPaths() bool {
	if o != nil && IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *C1ApiUserV1UserExpandMask) SetPaths(v []string) {
	o.Paths = v
}

func (o C1ApiUserV1UserExpandMask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiUserV1UserExpandMask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiUserV1UserExpandMask) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiUserV1UserExpandMask := _C1ApiUserV1UserExpandMask{}

	if err = json.Unmarshal(bytes, &varC1ApiUserV1UserExpandMask); err == nil {
		*o = C1ApiUserV1UserExpandMask(varC1ApiUserV1UserExpandMask)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "paths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiUserV1UserExpandMask struct {
	value *C1ApiUserV1UserExpandMask
	isSet bool
}

func (v NullableC1ApiUserV1UserExpandMask) Get() *C1ApiUserV1UserExpandMask {
	return v.value
}

func (v *NullableC1ApiUserV1UserExpandMask) Set(val *C1ApiUserV1UserExpandMask) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiUserV1UserExpandMask) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiUserV1UserExpandMask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiUserV1UserExpandMask(val *C1ApiUserV1UserExpandMask) *NullableC1ApiUserV1UserExpandMask {
	return &NullableC1ApiUserV1UserExpandMask{value: val, isSet: true}
}

func (v NullableC1ApiUserV1UserExpandMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiUserV1UserExpandMask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


