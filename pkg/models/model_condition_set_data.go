/*
Permit.io API

 Authorization as a service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// checks if the ConditionSetData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionSetData{}

// ConditionSetData struct for ConditionSetData
type ConditionSetData struct {
	Type ConditionSetType `json:"type"`
	Key  string           `json:"key"`
}

// NewConditionSetData instantiates a new ConditionSetData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionSetData(type_ ConditionSetType, key string) *ConditionSetData {
	this := ConditionSetData{}
	this.Type = type_
	this.Key = key
	return &this
}

// NewConditionSetDataWithDefaults instantiates a new ConditionSetData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionSetDataWithDefaults() *ConditionSetData {
	this := ConditionSetData{}
	return &this
}

// GetType returns the Type field value
func (o *ConditionSetData) GetType() ConditionSetType {
	if o == nil {
		var ret ConditionSetType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConditionSetData) GetTypeOk() (*ConditionSetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConditionSetData) SetType(v ConditionSetType) {
	o.Type = v
}

// GetKey returns the Key field value
func (o *ConditionSetData) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConditionSetData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ConditionSetData) SetKey(v string) {
	o.Key = v
}

func (o ConditionSetData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionSetData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableConditionSetData struct {
	value *ConditionSetData
	isSet bool
}

func (v NullableConditionSetData) Get() *ConditionSetData {
	return v.value
}

func (v *NullableConditionSetData) Set(val *ConditionSetData) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionSetData) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionSetData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionSetData(val *ConditionSetData) *NullableConditionSetData {
	return &NullableConditionSetData{value: val, isSet: true}
}

func (v NullableConditionSetData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionSetData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
