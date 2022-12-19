# Journal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**DisableComments** | Pointer to **bool** |  | [optional] 
**DisableAnonymousComments** | Pointer to **bool** |  | [optional] 
**DisableDownloads** | Pointer to **bool** |  | [optional] 
**Classification** | Pointer to **string** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Children** | Pointer to [**[]Child**](Child.md) |  | [optional] 
**Pets** | Pointer to [**[]Pet**](Pet.md) |  | [optional] 
**CoverTheme** | Pointer to [**CoverTheme**](CoverTheme.md) |  | [optional] 
**Features** | Pointer to [**[]Features**](Features.md) |  | [optional] 
**CurrentSubscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**LatestSubscription** | Pointer to [**Subscription**](Subscription.md) |  | [optional] 
**ShouldShowAds** | Pointer to **bool** |  | [optional] 

## Methods

### NewJournal

`func NewJournal() *Journal`

NewJournal instantiates a new Journal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalWithDefaults

`func NewJournalWithDefaults() *Journal`

NewJournalWithDefaults instantiates a new Journal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Journal) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Journal) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Journal) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Journal) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *Journal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Journal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Journal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Journal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Journal) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Journal) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Journal) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Journal) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTimestamp

`func (o *Journal) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Journal) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Journal) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Journal) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDisableComments

`func (o *Journal) GetDisableComments() bool`

GetDisableComments returns the DisableComments field if non-nil, zero value otherwise.

### GetDisableCommentsOk

`func (o *Journal) GetDisableCommentsOk() (*bool, bool)`

GetDisableCommentsOk returns a tuple with the DisableComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableComments

`func (o *Journal) SetDisableComments(v bool)`

SetDisableComments sets DisableComments field to given value.

### HasDisableComments

`func (o *Journal) HasDisableComments() bool`

HasDisableComments returns a boolean if a field has been set.

### GetDisableAnonymousComments

`func (o *Journal) GetDisableAnonymousComments() bool`

GetDisableAnonymousComments returns the DisableAnonymousComments field if non-nil, zero value otherwise.

### GetDisableAnonymousCommentsOk

`func (o *Journal) GetDisableAnonymousCommentsOk() (*bool, bool)`

GetDisableAnonymousCommentsOk returns a tuple with the DisableAnonymousComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAnonymousComments

`func (o *Journal) SetDisableAnonymousComments(v bool)`

SetDisableAnonymousComments sets DisableAnonymousComments field to given value.

### HasDisableAnonymousComments

`func (o *Journal) HasDisableAnonymousComments() bool`

HasDisableAnonymousComments returns a boolean if a field has been set.

### GetDisableDownloads

`func (o *Journal) GetDisableDownloads() bool`

GetDisableDownloads returns the DisableDownloads field if non-nil, zero value otherwise.

### GetDisableDownloadsOk

`func (o *Journal) GetDisableDownloadsOk() (*bool, bool)`

GetDisableDownloadsOk returns a tuple with the DisableDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDownloads

`func (o *Journal) SetDisableDownloads(v bool)`

SetDisableDownloads sets DisableDownloads field to given value.

### HasDisableDownloads

`func (o *Journal) HasDisableDownloads() bool`

HasDisableDownloads returns a boolean if a field has been set.

### GetClassification

`func (o *Journal) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Journal) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Journal) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Journal) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetIsPublic

`func (o *Journal) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Journal) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Journal) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Journal) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetDeleted

`func (o *Journal) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Journal) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Journal) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Journal) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUser

`func (o *Journal) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Journal) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Journal) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Journal) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetChildren

`func (o *Journal) GetChildren() []Child`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Journal) GetChildrenOk() (*[]Child, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Journal) SetChildren(v []Child)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Journal) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetPets

`func (o *Journal) GetPets() []Pet`

GetPets returns the Pets field if non-nil, zero value otherwise.

### GetPetsOk

`func (o *Journal) GetPetsOk() (*[]Pet, bool)`

GetPetsOk returns a tuple with the Pets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPets

`func (o *Journal) SetPets(v []Pet)`

SetPets sets Pets field to given value.

### HasPets

`func (o *Journal) HasPets() bool`

HasPets returns a boolean if a field has been set.

### GetCoverTheme

`func (o *Journal) GetCoverTheme() CoverTheme`

GetCoverTheme returns the CoverTheme field if non-nil, zero value otherwise.

### GetCoverThemeOk

`func (o *Journal) GetCoverThemeOk() (*CoverTheme, bool)`

GetCoverThemeOk returns a tuple with the CoverTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverTheme

`func (o *Journal) SetCoverTheme(v CoverTheme)`

SetCoverTheme sets CoverTheme field to given value.

### HasCoverTheme

`func (o *Journal) HasCoverTheme() bool`

HasCoverTheme returns a boolean if a field has been set.

### GetFeatures

`func (o *Journal) GetFeatures() []Features`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Journal) GetFeaturesOk() (*[]Features, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Journal) SetFeatures(v []Features)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Journal) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCurrentSubscription

`func (o *Journal) GetCurrentSubscription() Subscription`

GetCurrentSubscription returns the CurrentSubscription field if non-nil, zero value otherwise.

### GetCurrentSubscriptionOk

`func (o *Journal) GetCurrentSubscriptionOk() (*Subscription, bool)`

GetCurrentSubscriptionOk returns a tuple with the CurrentSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubscription

`func (o *Journal) SetCurrentSubscription(v Subscription)`

SetCurrentSubscription sets CurrentSubscription field to given value.

### HasCurrentSubscription

`func (o *Journal) HasCurrentSubscription() bool`

HasCurrentSubscription returns a boolean if a field has been set.

### GetLatestSubscription

`func (o *Journal) GetLatestSubscription() Subscription`

GetLatestSubscription returns the LatestSubscription field if non-nil, zero value otherwise.

### GetLatestSubscriptionOk

`func (o *Journal) GetLatestSubscriptionOk() (*Subscription, bool)`

GetLatestSubscriptionOk returns a tuple with the LatestSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSubscription

`func (o *Journal) SetLatestSubscription(v Subscription)`

SetLatestSubscription sets LatestSubscription field to given value.

### HasLatestSubscription

`func (o *Journal) HasLatestSubscription() bool`

HasLatestSubscription returns a boolean if a field has been set.

### GetShouldShowAds

`func (o *Journal) GetShouldShowAds() bool`

GetShouldShowAds returns the ShouldShowAds field if non-nil, zero value otherwise.

### GetShouldShowAdsOk

`func (o *Journal) GetShouldShowAdsOk() (*bool, bool)`

GetShouldShowAdsOk returns a tuple with the ShouldShowAds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldShowAds

`func (o *Journal) SetShouldShowAds(v bool)`

SetShouldShowAds sets ShouldShowAds field to given value.

### HasShouldShowAds

`func (o *Journal) HasShouldShowAds() bool`

HasShouldShowAds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


