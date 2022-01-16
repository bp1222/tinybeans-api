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

// Entries struct for Entries
type Entries struct {
	Status *string `json:"status,omitempty"`
	NumEntriesRemaining *int64 `json:"numEntriesRemaining,omitempty"`
	Entries *[]Entry `json:"entries,omitempty"`
}

// NewEntries instantiates a new Entries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntries() *Entries {
	this := Entries{}
	return &this
}

// NewEntriesWithDefaults instantiates a new Entries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntriesWithDefaults() *Entries {
	this := Entries{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Entries) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entries) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Entries) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Entries) SetStatus(v string) {
	o.Status = &v
}

// GetNumEntriesRemaining returns the NumEntriesRemaining field value if set, zero value otherwise.
func (o *Entries) GetNumEntriesRemaining() int64 {
	if o == nil || o.NumEntriesRemaining == nil {
		var ret int64
		return ret
	}
	return *o.NumEntriesRemaining
}

// GetNumEntriesRemainingOk returns a tuple with the NumEntriesRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entries) GetNumEntriesRemainingOk() (*int64, bool) {
	if o == nil || o.NumEntriesRemaining == nil {
		return nil, false
	}
	return o.NumEntriesRemaining, true
}

// HasNumEntriesRemaining returns a boolean if a field has been set.
func (o *Entries) HasNumEntriesRemaining() bool {
	if o != nil && o.NumEntriesRemaining != nil {
		return true
	}

	return false
}

// SetNumEntriesRemaining gets a reference to the given int64 and assigns it to the NumEntriesRemaining field.
func (o *Entries) SetNumEntriesRemaining(v int64) {
	o.NumEntriesRemaining = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *Entries) GetEntries() []Entry {
	if o == nil || o.Entries == nil {
		var ret []Entry
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entries) GetEntriesOk() (*[]Entry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *Entries) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []Entry and assigns it to the Entries field.
func (o *Entries) SetEntries(v []Entry) {
	o.Entries = &v
}

func (o Entries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.NumEntriesRemaining != nil {
		toSerialize["numEntriesRemaining"] = o.NumEntriesRemaining
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableEntries struct {
	value *Entries
	isSet bool
}

func (v NullableEntries) Get() *Entries {
	return v.value
}

func (v *NullableEntries) Set(val *Entries) {
	v.value = val
	v.isSet = true
}

func (v NullableEntries) IsSet() bool {
	return v.isSet
}

func (v *NullableEntries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntries(val *Entries) *NullableEntries {
	return &NullableEntries{value: val, isSet: true}
}

func (v NullableEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

