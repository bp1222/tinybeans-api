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

// Blob struct for Blob
type Blob struct {
	O *string `json:"o,omitempty"`
	O2 *string `json:"o2,omitempty"`
	T *string `json:"t,omitempty"`
	S *string `json:"s,omitempty"`
	S2 *string `json:"s2,omitempty"`
	M *string `json:"m,omitempty"`
	L *string `json:"l,omitempty"`
	P *string `json:"p,omitempty"`
}

// NewBlob instantiates a new Blob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlob() *Blob {
	this := Blob{}
	return &this
}

// NewBlobWithDefaults instantiates a new Blob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobWithDefaults() *Blob {
	this := Blob{}
	return &this
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Blob) GetO() string {
	if o == nil || o.O == nil {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetOOk() (*string, bool) {
	if o == nil || o.O == nil {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *Blob) HasO() bool {
	if o != nil && o.O != nil {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *Blob) SetO(v string) {
	o.O = &v
}

// GetO2 returns the O2 field value if set, zero value otherwise.
func (o *Blob) GetO2() string {
	if o == nil || o.O2 == nil {
		var ret string
		return ret
	}
	return *o.O2
}

// GetO2Ok returns a tuple with the O2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetO2Ok() (*string, bool) {
	if o == nil || o.O2 == nil {
		return nil, false
	}
	return o.O2, true
}

// HasO2 returns a boolean if a field has been set.
func (o *Blob) HasO2() bool {
	if o != nil && o.O2 != nil {
		return true
	}

	return false
}

// SetO2 gets a reference to the given string and assigns it to the O2 field.
func (o *Blob) SetO2(v string) {
	o.O2 = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Blob) GetT() string {
	if o == nil || o.T == nil {
		var ret string
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetTOk() (*string, bool) {
	if o == nil || o.T == nil {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *Blob) HasT() bool {
	if o != nil && o.T != nil {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *Blob) SetT(v string) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Blob) GetS() string {
	if o == nil || o.S == nil {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetSOk() (*string, bool) {
	if o == nil || o.S == nil {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *Blob) HasS() bool {
	if o != nil && o.S != nil {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Blob) SetS(v string) {
	o.S = &v
}

// GetS2 returns the S2 field value if set, zero value otherwise.
func (o *Blob) GetS2() string {
	if o == nil || o.S2 == nil {
		var ret string
		return ret
	}
	return *o.S2
}

// GetS2Ok returns a tuple with the S2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetS2Ok() (*string, bool) {
	if o == nil || o.S2 == nil {
		return nil, false
	}
	return o.S2, true
}

// HasS2 returns a boolean if a field has been set.
func (o *Blob) HasS2() bool {
	if o != nil && o.S2 != nil {
		return true
	}

	return false
}

// SetS2 gets a reference to the given string and assigns it to the S2 field.
func (o *Blob) SetS2(v string) {
	o.S2 = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *Blob) GetM() string {
	if o == nil || o.M == nil {
		var ret string
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetMOk() (*string, bool) {
	if o == nil || o.M == nil {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *Blob) HasM() bool {
	if o != nil && o.M != nil {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *Blob) SetM(v string) {
	o.M = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Blob) GetL() string {
	if o == nil || o.L == nil {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetLOk() (*string, bool) {
	if o == nil || o.L == nil {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *Blob) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Blob) SetL(v string) {
	o.L = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Blob) GetP() string {
	if o == nil || o.P == nil {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Blob) GetPOk() (*string, bool) {
	if o == nil || o.P == nil {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *Blob) HasP() bool {
	if o != nil && o.P != nil {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Blob) SetP(v string) {
	o.P = &v
}

func (o Blob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.O != nil {
		toSerialize["o"] = o.O
	}
	if o.O2 != nil {
		toSerialize["o2"] = o.O2
	}
	if o.T != nil {
		toSerialize["t"] = o.T
	}
	if o.S != nil {
		toSerialize["s"] = o.S
	}
	if o.S2 != nil {
		toSerialize["s2"] = o.S2
	}
	if o.M != nil {
		toSerialize["m"] = o.M
	}
	if o.L != nil {
		toSerialize["l"] = o.L
	}
	if o.P != nil {
		toSerialize["p"] = o.P
	}
	return json.Marshal(toSerialize)
}

type NullableBlob struct {
	value *Blob
	isSet bool
}

func (v NullableBlob) Get() *Blob {
	return v.value
}

func (v *NullableBlob) Set(val *Blob) {
	v.value = val
	v.isSet = true
}

func (v NullableBlob) IsSet() bool {
	return v.isSet
}

func (v *NullableBlob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlob(val *Blob) *NullableBlob {
	return &NullableBlob{value: val, isSet: true}
}

func (v NullableBlob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


