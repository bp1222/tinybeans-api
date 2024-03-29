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

// Journal struct for Journal
type Journal struct {
	URL *string `json:"URL,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	DisableComments *bool `json:"disableComments,omitempty"`
	DisableAnonymousComments *bool `json:"disableAnonymousComments,omitempty"`
	DisableDownloads *bool `json:"disableDownloads,omitempty"`
	Classification *string `json:"classification,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	User *User `json:"user,omitempty"`
	Children []Child `json:"children,omitempty"`
	Pets []Pet `json:"pets,omitempty"`
	CoverTheme *CoverTheme `json:"coverTheme,omitempty"`
	Features []Features `json:"features,omitempty"`
	CurrentSubscription *Subscription `json:"currentSubscription,omitempty"`
	LatestSubscription *Subscription `json:"latestSubscription,omitempty"`
	ShouldShowAds *bool `json:"shouldShowAds,omitempty"`
}

// NewJournal instantiates a new Journal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournal() *Journal {
	this := Journal{}
	return &this
}

// NewJournalWithDefaults instantiates a new Journal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalWithDefaults() *Journal {
	this := Journal{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *Journal) GetURL() string {
	if o == nil || isNil(o.URL) {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetURLOk() (*string, bool) {
	if o == nil || isNil(o.URL) {
    return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *Journal) HasURL() bool {
	if o != nil && !isNil(o.URL) {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *Journal) SetURL(v string) {
	o.URL = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Journal) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Journal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Journal) SetId(v int64) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Journal) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Journal) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Journal) SetTitle(v string) {
	o.Title = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Journal) GetTimestamp() int64 {
	if o == nil || isNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetTimestampOk() (*int64, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Journal) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *Journal) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetDisableComments returns the DisableComments field value if set, zero value otherwise.
func (o *Journal) GetDisableComments() bool {
	if o == nil || isNil(o.DisableComments) {
		var ret bool
		return ret
	}
	return *o.DisableComments
}

// GetDisableCommentsOk returns a tuple with the DisableComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetDisableCommentsOk() (*bool, bool) {
	if o == nil || isNil(o.DisableComments) {
    return nil, false
	}
	return o.DisableComments, true
}

// HasDisableComments returns a boolean if a field has been set.
func (o *Journal) HasDisableComments() bool {
	if o != nil && !isNil(o.DisableComments) {
		return true
	}

	return false
}

// SetDisableComments gets a reference to the given bool and assigns it to the DisableComments field.
func (o *Journal) SetDisableComments(v bool) {
	o.DisableComments = &v
}

// GetDisableAnonymousComments returns the DisableAnonymousComments field value if set, zero value otherwise.
func (o *Journal) GetDisableAnonymousComments() bool {
	if o == nil || isNil(o.DisableAnonymousComments) {
		var ret bool
		return ret
	}
	return *o.DisableAnonymousComments
}

// GetDisableAnonymousCommentsOk returns a tuple with the DisableAnonymousComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetDisableAnonymousCommentsOk() (*bool, bool) {
	if o == nil || isNil(o.DisableAnonymousComments) {
    return nil, false
	}
	return o.DisableAnonymousComments, true
}

// HasDisableAnonymousComments returns a boolean if a field has been set.
func (o *Journal) HasDisableAnonymousComments() bool {
	if o != nil && !isNil(o.DisableAnonymousComments) {
		return true
	}

	return false
}

// SetDisableAnonymousComments gets a reference to the given bool and assigns it to the DisableAnonymousComments field.
func (o *Journal) SetDisableAnonymousComments(v bool) {
	o.DisableAnonymousComments = &v
}

// GetDisableDownloads returns the DisableDownloads field value if set, zero value otherwise.
func (o *Journal) GetDisableDownloads() bool {
	if o == nil || isNil(o.DisableDownloads) {
		var ret bool
		return ret
	}
	return *o.DisableDownloads
}

// GetDisableDownloadsOk returns a tuple with the DisableDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetDisableDownloadsOk() (*bool, bool) {
	if o == nil || isNil(o.DisableDownloads) {
    return nil, false
	}
	return o.DisableDownloads, true
}

// HasDisableDownloads returns a boolean if a field has been set.
func (o *Journal) HasDisableDownloads() bool {
	if o != nil && !isNil(o.DisableDownloads) {
		return true
	}

	return false
}

// SetDisableDownloads gets a reference to the given bool and assigns it to the DisableDownloads field.
func (o *Journal) SetDisableDownloads(v bool) {
	o.DisableDownloads = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *Journal) GetClassification() string {
	if o == nil || isNil(o.Classification) {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetClassificationOk() (*string, bool) {
	if o == nil || isNil(o.Classification) {
    return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *Journal) HasClassification() bool {
	if o != nil && !isNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *Journal) SetClassification(v string) {
	o.Classification = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *Journal) GetIsPublic() bool {
	if o == nil || isNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetIsPublicOk() (*bool, bool) {
	if o == nil || isNil(o.IsPublic) {
    return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *Journal) HasIsPublic() bool {
	if o != nil && !isNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *Journal) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Journal) GetDeleted() bool {
	if o == nil || isNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.Deleted) {
    return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Journal) HasDeleted() bool {
	if o != nil && !isNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Journal) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Journal) GetUser() User {
	if o == nil || isNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetUserOk() (*User, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Journal) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Journal) SetUser(v User) {
	o.User = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Journal) GetChildren() []Child {
	if o == nil || isNil(o.Children) {
		var ret []Child
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetChildrenOk() ([]Child, bool) {
	if o == nil || isNil(o.Children) {
    return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Journal) HasChildren() bool {
	if o != nil && !isNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Child and assigns it to the Children field.
func (o *Journal) SetChildren(v []Child) {
	o.Children = v
}

// GetPets returns the Pets field value if set, zero value otherwise.
func (o *Journal) GetPets() []Pet {
	if o == nil || isNil(o.Pets) {
		var ret []Pet
		return ret
	}
	return o.Pets
}

// GetPetsOk returns a tuple with the Pets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetPetsOk() ([]Pet, bool) {
	if o == nil || isNil(o.Pets) {
    return nil, false
	}
	return o.Pets, true
}

// HasPets returns a boolean if a field has been set.
func (o *Journal) HasPets() bool {
	if o != nil && !isNil(o.Pets) {
		return true
	}

	return false
}

// SetPets gets a reference to the given []Pet and assigns it to the Pets field.
func (o *Journal) SetPets(v []Pet) {
	o.Pets = v
}

// GetCoverTheme returns the CoverTheme field value if set, zero value otherwise.
func (o *Journal) GetCoverTheme() CoverTheme {
	if o == nil || isNil(o.CoverTheme) {
		var ret CoverTheme
		return ret
	}
	return *o.CoverTheme
}

// GetCoverThemeOk returns a tuple with the CoverTheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetCoverThemeOk() (*CoverTheme, bool) {
	if o == nil || isNil(o.CoverTheme) {
    return nil, false
	}
	return o.CoverTheme, true
}

// HasCoverTheme returns a boolean if a field has been set.
func (o *Journal) HasCoverTheme() bool {
	if o != nil && !isNil(o.CoverTheme) {
		return true
	}

	return false
}

// SetCoverTheme gets a reference to the given CoverTheme and assigns it to the CoverTheme field.
func (o *Journal) SetCoverTheme(v CoverTheme) {
	o.CoverTheme = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Journal) GetFeatures() []Features {
	if o == nil || isNil(o.Features) {
		var ret []Features
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetFeaturesOk() ([]Features, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Journal) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []Features and assigns it to the Features field.
func (o *Journal) SetFeatures(v []Features) {
	o.Features = v
}

// GetCurrentSubscription returns the CurrentSubscription field value if set, zero value otherwise.
func (o *Journal) GetCurrentSubscription() Subscription {
	if o == nil || isNil(o.CurrentSubscription) {
		var ret Subscription
		return ret
	}
	return *o.CurrentSubscription
}

// GetCurrentSubscriptionOk returns a tuple with the CurrentSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetCurrentSubscriptionOk() (*Subscription, bool) {
	if o == nil || isNil(o.CurrentSubscription) {
    return nil, false
	}
	return o.CurrentSubscription, true
}

// HasCurrentSubscription returns a boolean if a field has been set.
func (o *Journal) HasCurrentSubscription() bool {
	if o != nil && !isNil(o.CurrentSubscription) {
		return true
	}

	return false
}

// SetCurrentSubscription gets a reference to the given Subscription and assigns it to the CurrentSubscription field.
func (o *Journal) SetCurrentSubscription(v Subscription) {
	o.CurrentSubscription = &v
}

// GetLatestSubscription returns the LatestSubscription field value if set, zero value otherwise.
func (o *Journal) GetLatestSubscription() Subscription {
	if o == nil || isNil(o.LatestSubscription) {
		var ret Subscription
		return ret
	}
	return *o.LatestSubscription
}

// GetLatestSubscriptionOk returns a tuple with the LatestSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetLatestSubscriptionOk() (*Subscription, bool) {
	if o == nil || isNil(o.LatestSubscription) {
    return nil, false
	}
	return o.LatestSubscription, true
}

// HasLatestSubscription returns a boolean if a field has been set.
func (o *Journal) HasLatestSubscription() bool {
	if o != nil && !isNil(o.LatestSubscription) {
		return true
	}

	return false
}

// SetLatestSubscription gets a reference to the given Subscription and assigns it to the LatestSubscription field.
func (o *Journal) SetLatestSubscription(v Subscription) {
	o.LatestSubscription = &v
}

// GetShouldShowAds returns the ShouldShowAds field value if set, zero value otherwise.
func (o *Journal) GetShouldShowAds() bool {
	if o == nil || isNil(o.ShouldShowAds) {
		var ret bool
		return ret
	}
	return *o.ShouldShowAds
}

// GetShouldShowAdsOk returns a tuple with the ShouldShowAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Journal) GetShouldShowAdsOk() (*bool, bool) {
	if o == nil || isNil(o.ShouldShowAds) {
    return nil, false
	}
	return o.ShouldShowAds, true
}

// HasShouldShowAds returns a boolean if a field has been set.
func (o *Journal) HasShouldShowAds() bool {
	if o != nil && !isNil(o.ShouldShowAds) {
		return true
	}

	return false
}

// SetShouldShowAds gets a reference to the given bool and assigns it to the ShouldShowAds field.
func (o *Journal) SetShouldShowAds(v bool) {
	o.ShouldShowAds = &v
}

func (o Journal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.URL) {
		toSerialize["URL"] = o.URL
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.DisableComments) {
		toSerialize["disableComments"] = o.DisableComments
	}
	if !isNil(o.DisableAnonymousComments) {
		toSerialize["disableAnonymousComments"] = o.DisableAnonymousComments
	}
	if !isNil(o.DisableDownloads) {
		toSerialize["disableDownloads"] = o.DisableDownloads
	}
	if !isNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !isNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !isNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !isNil(o.Pets) {
		toSerialize["pets"] = o.Pets
	}
	if !isNil(o.CoverTheme) {
		toSerialize["coverTheme"] = o.CoverTheme
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !isNil(o.CurrentSubscription) {
		toSerialize["currentSubscription"] = o.CurrentSubscription
	}
	if !isNil(o.LatestSubscription) {
		toSerialize["latestSubscription"] = o.LatestSubscription
	}
	if !isNil(o.ShouldShowAds) {
		toSerialize["shouldShowAds"] = o.ShouldShowAds
	}
	return json.Marshal(toSerialize)
}

type NullableJournal struct {
	value *Journal
	isSet bool
}

func (v NullableJournal) Get() *Journal {
	return v.value
}

func (v *NullableJournal) Set(val *Journal) {
	v.value = val
	v.isSet = true
}

func (v NullableJournal) IsSet() bool {
	return v.isSet
}

func (v *NullableJournal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournal(val *Journal) *NullableJournal {
	return &NullableJournal{value: val, isSet: true}
}

func (v NullableJournal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


