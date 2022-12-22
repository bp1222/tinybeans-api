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

// MeasurementSystem struct for MeasurementSystem
type MeasurementSystem struct {
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewMeasurementSystem instantiates a new MeasurementSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeasurementSystem() *MeasurementSystem {
	this := MeasurementSystem{}
	return &this
}

// NewMeasurementSystemWithDefaults instantiates a new MeasurementSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeasurementSystemWithDefaults() *MeasurementSystem {
	this := MeasurementSystem{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *MeasurementSystem) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementSystem) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *MeasurementSystem) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *MeasurementSystem) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MeasurementSystem) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeasurementSystem) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MeasurementSystem) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MeasurementSystem) SetName(v string) {
	o.Name = &v
}

func (o MeasurementSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableMeasurementSystem struct {
	value *MeasurementSystem
	isSet bool
}

func (v NullableMeasurementSystem) Get() *MeasurementSystem {
	return v.value
}

func (v *NullableMeasurementSystem) Set(val *MeasurementSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementSystem(val *MeasurementSystem) *NullableMeasurementSystem {
	return &NullableMeasurementSystem{value: val, isSet: true}
}

func (v NullableMeasurementSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


