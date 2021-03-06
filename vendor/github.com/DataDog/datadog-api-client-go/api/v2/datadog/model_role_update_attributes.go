/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleUpdateAttributes Attributes of the edited role.
type RoleUpdateAttributes struct {
	// The name of the role.
	Name *string `json:"name,omitempty"`
}

// NewRoleUpdateAttributes instantiates a new RoleUpdateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdateAttributes() *RoleUpdateAttributes {
	this := RoleUpdateAttributes{}
	return &this
}

// NewRoleUpdateAttributesWithDefaults instantiates a new RoleUpdateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateAttributesWithDefaults() *RoleUpdateAttributes {
	this := RoleUpdateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleUpdateAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleUpdateAttributes) SetName(v string) {
	o.Name = &v
}

func (o RoleUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRoleUpdateAttributes struct {
	value *RoleUpdateAttributes
	isSet bool
}

func (v NullableRoleUpdateAttributes) Get() *RoleUpdateAttributes {
	return v.value
}

func (v *NullableRoleUpdateAttributes) Set(val *RoleUpdateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdateAttributes(val *RoleUpdateAttributes) *NullableRoleUpdateAttributes {
	return &NullableRoleUpdateAttributes{value: val, isSet: true}
}

func (v NullableRoleUpdateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
