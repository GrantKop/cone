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

// checks if the C1ApiUserV1UserView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &C1ApiUserV1UserView{}

// C1ApiUserV1UserView The UserView message.
type C1ApiUserV1UserView struct {
	DelegatedUserPath interface{} `json:"delegatedUserPath,omitempty"`
	DirectoriesPath interface{} `json:"directoriesPath,omitempty"`
	ManagersPath interface{} `json:"managersPath,omitempty"`
	RolesPath interface{} `json:"rolesPath,omitempty"`
	User *C1ApiUserV1User `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _C1ApiUserV1UserView C1ApiUserV1UserView

// NewC1ApiUserV1UserView instantiates a new C1ApiUserV1UserView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewC1ApiUserV1UserView() *C1ApiUserV1UserView {
	this := C1ApiUserV1UserView{}
	return &this
}

// NewC1ApiUserV1UserViewWithDefaults instantiates a new C1ApiUserV1UserView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewC1ApiUserV1UserViewWithDefaults() *C1ApiUserV1UserView {
	this := C1ApiUserV1UserView{}
	return &this
}

// GetDelegatedUserPath returns the DelegatedUserPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserView) GetDelegatedUserPath() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DelegatedUserPath
}

// GetDelegatedUserPathOk returns a tuple with the DelegatedUserPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserView) GetDelegatedUserPathOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DelegatedUserPath) {
		return nil, false
	}
	return &o.DelegatedUserPath, true
}

// HasDelegatedUserPath returns a boolean if a field has been set.
func (o *C1ApiUserV1UserView) HasDelegatedUserPath() bool {
	if o != nil && IsNil(o.DelegatedUserPath) {
		return true
	}

	return false
}

// SetDelegatedUserPath gets a reference to the given interface{} and assigns it to the DelegatedUserPath field.
func (o *C1ApiUserV1UserView) SetDelegatedUserPath(v interface{}) {
	o.DelegatedUserPath = v
}

// GetDirectoriesPath returns the DirectoriesPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserView) GetDirectoriesPath() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DirectoriesPath
}

// GetDirectoriesPathOk returns a tuple with the DirectoriesPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserView) GetDirectoriesPathOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DirectoriesPath) {
		return nil, false
	}
	return &o.DirectoriesPath, true
}

// HasDirectoriesPath returns a boolean if a field has been set.
func (o *C1ApiUserV1UserView) HasDirectoriesPath() bool {
	if o != nil && IsNil(o.DirectoriesPath) {
		return true
	}

	return false
}

// SetDirectoriesPath gets a reference to the given interface{} and assigns it to the DirectoriesPath field.
func (o *C1ApiUserV1UserView) SetDirectoriesPath(v interface{}) {
	o.DirectoriesPath = v
}

// GetManagersPath returns the ManagersPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserView) GetManagersPath() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ManagersPath
}

// GetManagersPathOk returns a tuple with the ManagersPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserView) GetManagersPathOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ManagersPath) {
		return nil, false
	}
	return &o.ManagersPath, true
}

// HasManagersPath returns a boolean if a field has been set.
func (o *C1ApiUserV1UserView) HasManagersPath() bool {
	if o != nil && IsNil(o.ManagersPath) {
		return true
	}

	return false
}

// SetManagersPath gets a reference to the given interface{} and assigns it to the ManagersPath field.
func (o *C1ApiUserV1UserView) SetManagersPath(v interface{}) {
	o.ManagersPath = v
}

// GetRolesPath returns the RolesPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *C1ApiUserV1UserView) GetRolesPath() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RolesPath
}

// GetRolesPathOk returns a tuple with the RolesPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *C1ApiUserV1UserView) GetRolesPathOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RolesPath) {
		return nil, false
	}
	return &o.RolesPath, true
}

// HasRolesPath returns a boolean if a field has been set.
func (o *C1ApiUserV1UserView) HasRolesPath() bool {
	if o != nil && IsNil(o.RolesPath) {
		return true
	}

	return false
}

// SetRolesPath gets a reference to the given interface{} and assigns it to the RolesPath field.
func (o *C1ApiUserV1UserView) SetRolesPath(v interface{}) {
	o.RolesPath = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *C1ApiUserV1UserView) GetUser() C1ApiUserV1User {
	if o == nil || IsNil(o.User) {
		var ret C1ApiUserV1User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *C1ApiUserV1UserView) GetUserOk() (*C1ApiUserV1User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *C1ApiUserV1UserView) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given C1ApiUserV1User and assigns it to the User field.
func (o *C1ApiUserV1UserView) SetUser(v C1ApiUserV1User) {
	o.User = &v
}

func (o C1ApiUserV1UserView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o C1ApiUserV1UserView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DelegatedUserPath != nil {
		toSerialize["delegatedUserPath"] = o.DelegatedUserPath
	}
	if o.DirectoriesPath != nil {
		toSerialize["directoriesPath"] = o.DirectoriesPath
	}
	if o.ManagersPath != nil {
		toSerialize["managersPath"] = o.ManagersPath
	}
	if o.RolesPath != nil {
		toSerialize["rolesPath"] = o.RolesPath
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *C1ApiUserV1UserView) UnmarshalJSON(bytes []byte) (err error) {
	varC1ApiUserV1UserView := _C1ApiUserV1UserView{}

	if err = json.Unmarshal(bytes, &varC1ApiUserV1UserView); err == nil {
		*o = C1ApiUserV1UserView(varC1ApiUserV1UserView)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "delegatedUserPath")
		delete(additionalProperties, "directoriesPath")
		delete(additionalProperties, "managersPath")
		delete(additionalProperties, "rolesPath")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableC1ApiUserV1UserView struct {
	value *C1ApiUserV1UserView
	isSet bool
}

func (v NullableC1ApiUserV1UserView) Get() *C1ApiUserV1UserView {
	return v.value
}

func (v *NullableC1ApiUserV1UserView) Set(val *C1ApiUserV1UserView) {
	v.value = val
	v.isSet = true
}

func (v NullableC1ApiUserV1UserView) IsSet() bool {
	return v.isSet
}

func (v *NullableC1ApiUserV1UserView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableC1ApiUserV1UserView(val *C1ApiUserV1UserView) *NullableC1ApiUserV1UserView {
	return &NullableC1ApiUserV1UserView{value: val, isSet: true}
}

func (v NullableC1ApiUserV1UserView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableC1ApiUserV1UserView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


