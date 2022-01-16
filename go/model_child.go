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

// Child struct for Child
type Child struct {
	URL *string `json:"URL,omitempty"`
	Id *int64 `json:"id,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	FullName *string `json:"fullName,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Dob *string `json:"dob,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	Avatars *Avatars `json:"avatars,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	User *User `json:"user,omitempty"`
}

// NewChild instantiates a new Child object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChild() *Child {
	this := Child{}
	return &this
}

// NewChildWithDefaults instantiates a new Child object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildWithDefaults() *Child {
	this := Child{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *Child) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *Child) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *Child) SetURL(v string) {
	o.URL = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Child) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Child) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Child) SetId(v int64) {
	o.Id = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Child) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Child) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Child) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Child) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Child) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Child) SetLastName(v string) {
	o.LastName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *Child) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *Child) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *Child) SetFullName(v string) {
	o.FullName = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Child) GetGender() string {
	if o == nil || o.Gender == nil {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetGenderOk() (*string, bool) {
	if o == nil || o.Gender == nil {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Child) HasGender() bool {
	if o != nil && o.Gender != nil {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *Child) SetGender(v string) {
	o.Gender = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *Child) GetDob() string {
	if o == nil || o.Dob == nil {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetDobOk() (*string, bool) {
	if o == nil || o.Dob == nil {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *Child) HasDob() bool {
	if o != nil && o.Dob != nil {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *Child) SetDob(v string) {
	o.Dob = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Child) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Child) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Child) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *Child) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *Child) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *Child) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *Child) GetAvatars() Avatars {
	if o == nil || o.Avatars == nil {
		var ret Avatars
		return ret
	}
	return *o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetAvatarsOk() (*Avatars, bool) {
	if o == nil || o.Avatars == nil {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *Child) HasAvatars() bool {
	if o != nil && o.Avatars != nil {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given Avatars and assigns it to the Avatars field.
func (o *Child) SetAvatars(v Avatars) {
	o.Avatars = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Child) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Child) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Child) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Child) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Child) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Child) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Child) SetUser(v User) {
	o.User = &v
}

func (o Child) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.URL != nil {
		toSerialize["URL"] = o.URL
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.FullName != nil {
		toSerialize["fullName"] = o.FullName
	}
	if o.Gender != nil {
		toSerialize["gender"] = o.Gender
	}
	if o.Dob != nil {
		toSerialize["dob"] = o.Dob
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["lastUpdatedTimestamp"] = o.LastUpdatedTimestamp
	}
	if o.Avatars != nil {
		toSerialize["avatars"] = o.Avatars
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableChild struct {
	value *Child
	isSet bool
}

func (v NullableChild) Get() *Child {
	return v.value
}

func (v *NullableChild) Set(val *Child) {
	v.value = val
	v.isSet = true
}

func (v NullableChild) IsSet() bool {
	return v.isSet
}

func (v *NullableChild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChild(val *Child) *NullableChild {
	return &NullableChild{value: val, isSet: true}
}

func (v NullableChild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

