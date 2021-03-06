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

// Pet struct for Pet
type Pet struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Species *string `json:"species,omitempty"`
	Breed *string `json:"breed,omitempty"`
	Size *string `json:"size,omitempty"`
	Sex *string `json:"sex,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DevelopmentOptOut *bool `json:"developmentOptOut,omitempty"`
	Birthday *string `json:"birthday,omitempty"`
	AdoptionDate *string `json:"adoptionDate,omitempty"`
	Avatars *Avatars `json:"avatars,omitempty"`
	User *User `json:"user,omitempty"`
}

// NewPet instantiates a new Pet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPet() *Pet {
	this := Pet{}
	return &this
}

// NewPetWithDefaults instantiates a new Pet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPetWithDefaults() *Pet {
	this := Pet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Pet) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Pet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Pet) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Pet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Pet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Pet) SetName(v string) {
	o.Name = &v
}

// GetSpecies returns the Species field value if set, zero value otherwise.
func (o *Pet) GetSpecies() string {
	if o == nil || o.Species == nil {
		var ret string
		return ret
	}
	return *o.Species
}

// GetSpeciesOk returns a tuple with the Species field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetSpeciesOk() (*string, bool) {
	if o == nil || o.Species == nil {
		return nil, false
	}
	return o.Species, true
}

// HasSpecies returns a boolean if a field has been set.
func (o *Pet) HasSpecies() bool {
	if o != nil && o.Species != nil {
		return true
	}

	return false
}

// SetSpecies gets a reference to the given string and assigns it to the Species field.
func (o *Pet) SetSpecies(v string) {
	o.Species = &v
}

// GetBreed returns the Breed field value if set, zero value otherwise.
func (o *Pet) GetBreed() string {
	if o == nil || o.Breed == nil {
		var ret string
		return ret
	}
	return *o.Breed
}

// GetBreedOk returns a tuple with the Breed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetBreedOk() (*string, bool) {
	if o == nil || o.Breed == nil {
		return nil, false
	}
	return o.Breed, true
}

// HasBreed returns a boolean if a field has been set.
func (o *Pet) HasBreed() bool {
	if o != nil && o.Breed != nil {
		return true
	}

	return false
}

// SetBreed gets a reference to the given string and assigns it to the Breed field.
func (o *Pet) SetBreed(v string) {
	o.Breed = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Pet) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Pet) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Pet) SetSize(v string) {
	o.Size = &v
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *Pet) GetSex() string {
	if o == nil || o.Sex == nil {
		var ret string
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetSexOk() (*string, bool) {
	if o == nil || o.Sex == nil {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *Pet) HasSex() bool {
	if o != nil && o.Sex != nil {
		return true
	}

	return false
}

// SetSex gets a reference to the given string and assigns it to the Sex field.
func (o *Pet) SetSex(v string) {
	o.Sex = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Pet) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Pet) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Pet) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDevelopmentOptOut returns the DevelopmentOptOut field value if set, zero value otherwise.
func (o *Pet) GetDevelopmentOptOut() bool {
	if o == nil || o.DevelopmentOptOut == nil {
		var ret bool
		return ret
	}
	return *o.DevelopmentOptOut
}

// GetDevelopmentOptOutOk returns a tuple with the DevelopmentOptOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetDevelopmentOptOutOk() (*bool, bool) {
	if o == nil || o.DevelopmentOptOut == nil {
		return nil, false
	}
	return o.DevelopmentOptOut, true
}

// HasDevelopmentOptOut returns a boolean if a field has been set.
func (o *Pet) HasDevelopmentOptOut() bool {
	if o != nil && o.DevelopmentOptOut != nil {
		return true
	}

	return false
}

// SetDevelopmentOptOut gets a reference to the given bool and assigns it to the DevelopmentOptOut field.
func (o *Pet) SetDevelopmentOptOut(v bool) {
	o.DevelopmentOptOut = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *Pet) GetBirthday() string {
	if o == nil || o.Birthday == nil {
		var ret string
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetBirthdayOk() (*string, bool) {
	if o == nil || o.Birthday == nil {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *Pet) HasBirthday() bool {
	if o != nil && o.Birthday != nil {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given string and assigns it to the Birthday field.
func (o *Pet) SetBirthday(v string) {
	o.Birthday = &v
}

// GetAdoptionDate returns the AdoptionDate field value if set, zero value otherwise.
func (o *Pet) GetAdoptionDate() string {
	if o == nil || o.AdoptionDate == nil {
		var ret string
		return ret
	}
	return *o.AdoptionDate
}

// GetAdoptionDateOk returns a tuple with the AdoptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetAdoptionDateOk() (*string, bool) {
	if o == nil || o.AdoptionDate == nil {
		return nil, false
	}
	return o.AdoptionDate, true
}

// HasAdoptionDate returns a boolean if a field has been set.
func (o *Pet) HasAdoptionDate() bool {
	if o != nil && o.AdoptionDate != nil {
		return true
	}

	return false
}

// SetAdoptionDate gets a reference to the given string and assigns it to the AdoptionDate field.
func (o *Pet) SetAdoptionDate(v string) {
	o.AdoptionDate = &v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *Pet) GetAvatars() Avatars {
	if o == nil || o.Avatars == nil {
		var ret Avatars
		return ret
	}
	return *o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetAvatarsOk() (*Avatars, bool) {
	if o == nil || o.Avatars == nil {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *Pet) HasAvatars() bool {
	if o != nil && o.Avatars != nil {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given Avatars and assigns it to the Avatars field.
func (o *Pet) SetAvatars(v Avatars) {
	o.Avatars = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Pet) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Pet) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Pet) SetUser(v User) {
	o.User = &v
}

func (o Pet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Species != nil {
		toSerialize["species"] = o.Species
	}
	if o.Breed != nil {
		toSerialize["breed"] = o.Breed
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Sex != nil {
		toSerialize["sex"] = o.Sex
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.DevelopmentOptOut != nil {
		toSerialize["developmentOptOut"] = o.DevelopmentOptOut
	}
	if o.Birthday != nil {
		toSerialize["birthday"] = o.Birthday
	}
	if o.AdoptionDate != nil {
		toSerialize["adoptionDate"] = o.AdoptionDate
	}
	if o.Avatars != nil {
		toSerialize["avatars"] = o.Avatars
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullablePet struct {
	value *Pet
	isSet bool
}

func (v NullablePet) Get() *Pet {
	return v.value
}

func (v *NullablePet) Set(val *Pet) {
	v.value = val
	v.isSet = true
}

func (v NullablePet) IsSet() bool {
	return v.isSet
}

func (v *NullablePet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePet(val *Pet) *NullablePet {
	return &NullablePet{value: val, isSet: true}
}

func (v NullablePet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


