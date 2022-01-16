/*
Tinybeans API - Unofficial

Reverse Engineered API for Tinybeans.  This is horribly incomplete, and not supported

API version: 0.0.1
Contact: dave@mudsite.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tinybeans

import (
	"encoding/json"
)

// EmailFrequency struct for EmailFrequency
type EmailFrequency struct {
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewEmailFrequency instantiates a new EmailFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailFrequency() *EmailFrequency {
	this := EmailFrequency{}
	return &this
}

// NewEmailFrequencyWithDefaults instantiates a new EmailFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailFrequencyWithDefaults() *EmailFrequency {
	this := EmailFrequency{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EmailFrequency) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailFrequency) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EmailFrequency) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EmailFrequency) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailFrequency) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailFrequency) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailFrequency) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailFrequency) SetName(v string) {
	o.Name = &v
}

func (o EmailFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEmailFrequency struct {
	value *EmailFrequency
	isSet bool
}

func (v NullableEmailFrequency) Get() *EmailFrequency {
	return v.value
}

func (v *NullableEmailFrequency) Set(val *EmailFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailFrequency(val *EmailFrequency) *NullableEmailFrequency {
	return &NullableEmailFrequency{value: val, isSet: true}
}

func (v NullableEmailFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

