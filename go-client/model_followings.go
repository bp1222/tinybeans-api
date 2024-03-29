/*
Tinybeans API - Unofficial

Reverse Engineered API for Tinybeans.  This is horribly incomplete, and not supported

API version: 0.1.1
Contact: dave@mudsite.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tinybeans

import (
	"encoding/json"
)

// Followings struct for Followings
type Followings struct {
	Status *string `json:"status,omitempty"`
	Followings []Following `json:"followings,omitempty"`
}

// NewFollowings instantiates a new Followings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowings() *Followings {
	this := Followings{}
	return &this
}

// NewFollowingsWithDefaults instantiates a new Followings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowingsWithDefaults() *Followings {
	this := Followings{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Followings) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Followings) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Followings) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Followings) SetStatus(v string) {
	o.Status = &v
}

// GetFollowings returns the Followings field value if set, zero value otherwise.
func (o *Followings) GetFollowings() []Following {
	if o == nil || isNil(o.Followings) {
		var ret []Following
		return ret
	}
	return o.Followings
}

// GetFollowingsOk returns a tuple with the Followings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Followings) GetFollowingsOk() ([]Following, bool) {
	if o == nil || isNil(o.Followings) {
    return nil, false
	}
	return o.Followings, true
}

// HasFollowings returns a boolean if a field has been set.
func (o *Followings) HasFollowings() bool {
	if o != nil && !isNil(o.Followings) {
		return true
	}

	return false
}

// SetFollowings gets a reference to the given []Following and assigns it to the Followings field.
func (o *Followings) SetFollowings(v []Following) {
	o.Followings = v
}

func (o Followings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Followings) {
		toSerialize["followings"] = o.Followings
	}
	return json.Marshal(toSerialize)
}

type NullableFollowings struct {
	value *Followings
	isSet bool
}

func (v NullableFollowings) Get() *Followings {
	return v.value
}

func (v *NullableFollowings) Set(val *Followings) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowings) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowings(val *Followings) *NullableFollowings {
	return &NullableFollowings{value: val, isSet: true}
}

func (v NullableFollowings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


