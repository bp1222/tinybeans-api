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

// Emotion struct for Emotion
type Emotion struct {
	URL *string `json:"URL,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Type *EmotionType `json:"type,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	EntryId *int64 `json:"entryId,omitempty"`
	UserId *int64 `json:"userId,omitempty"`
}

// NewEmotion instantiates a new Emotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmotion() *Emotion {
	this := Emotion{}
	return &this
}

// NewEmotionWithDefaults instantiates a new Emotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmotionWithDefaults() *Emotion {
	this := Emotion{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *Emotion) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *Emotion) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *Emotion) SetURL(v string) {
	o.URL = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Emotion) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Emotion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Emotion) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Emotion) GetType() EmotionType {
	if o == nil || o.Type == nil {
		var ret EmotionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetTypeOk() (*EmotionType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Emotion) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given EmotionType and assigns it to the Type field.
func (o *Emotion) SetType(v EmotionType) {
	o.Type = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Emotion) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Emotion) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Emotion) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *Emotion) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *Emotion) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *Emotion) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Emotion) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Emotion) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Emotion) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *Emotion) GetEntryId() int64 {
	if o == nil || o.EntryId == nil {
		var ret int64
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetEntryIdOk() (*int64, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *Emotion) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given int64 and assigns it to the EntryId field.
func (o *Emotion) SetEntryId(v int64) {
	o.EntryId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Emotion) GetUserId() int64 {
	if o == nil || o.UserId == nil {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Emotion) GetUserIdOk() (*int64, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Emotion) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *Emotion) SetUserId(v int64) {
	o.UserId = &v
}

func (o Emotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.URL != nil {
		toSerialize["URL"] = o.URL
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["lastUpdatedTimestamp"] = o.LastUpdatedTimestamp
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableEmotion struct {
	value *Emotion
	isSet bool
}

func (v NullableEmotion) Get() *Emotion {
	return v.value
}

func (v *NullableEmotion) Set(val *Emotion) {
	v.value = val
	v.isSet = true
}

func (v NullableEmotion) IsSet() bool {
	return v.isSet
}

func (v *NullableEmotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmotion(val *Emotion) *NullableEmotion {
	return &NullableEmotion{value: val, isSet: true}
}

func (v NullableEmotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


