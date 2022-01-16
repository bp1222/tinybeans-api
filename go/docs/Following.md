# Following

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ViewEntries** | Pointer to **bool** |  | [optional] 
**AddEntries** | Pointer to **bool** |  | [optional] 
**ViewMilestones** | Pointer to **bool** |  | [optional] 
**EditMilestones** | Pointer to **bool** |  | [optional] 
**CoOwner** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SendFlashback** | Pointer to **bool** |  | [optional] 
**Relationship** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**EmailFrequencyOnNewEvent** | Pointer to [**EmailFrequency**](EmailFrequency.md) |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**JournalId** | Pointer to **int64** |  | [optional] 
**Journal** | Pointer to [**Journal**](Journal.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewFollowing

`func NewFollowing() *Following`

NewFollowing instantiates a new Following object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowingWithDefaults

`func NewFollowingWithDefaults() *Following`

NewFollowingWithDefaults instantiates a new Following object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Following) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Following) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Following) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Following) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *Following) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Following) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Following) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Following) HasId() bool`

HasId returns a boolean if a field has been set.

### GetViewEntries

`func (o *Following) GetViewEntries() bool`

GetViewEntries returns the ViewEntries field if non-nil, zero value otherwise.

### GetViewEntriesOk

`func (o *Following) GetViewEntriesOk() (*bool, bool)`

GetViewEntriesOk returns a tuple with the ViewEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewEntries

`func (o *Following) SetViewEntries(v bool)`

SetViewEntries sets ViewEntries field to given value.

### HasViewEntries

`func (o *Following) HasViewEntries() bool`

HasViewEntries returns a boolean if a field has been set.

### GetAddEntries

`func (o *Following) GetAddEntries() bool`

GetAddEntries returns the AddEntries field if non-nil, zero value otherwise.

### GetAddEntriesOk

`func (o *Following) GetAddEntriesOk() (*bool, bool)`

GetAddEntriesOk returns a tuple with the AddEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEntries

`func (o *Following) SetAddEntries(v bool)`

SetAddEntries sets AddEntries field to given value.

### HasAddEntries

`func (o *Following) HasAddEntries() bool`

HasAddEntries returns a boolean if a field has been set.

### GetViewMilestones

`func (o *Following) GetViewMilestones() bool`

GetViewMilestones returns the ViewMilestones field if non-nil, zero value otherwise.

### GetViewMilestonesOk

`func (o *Following) GetViewMilestonesOk() (*bool, bool)`

GetViewMilestonesOk returns a tuple with the ViewMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMilestones

`func (o *Following) SetViewMilestones(v bool)`

SetViewMilestones sets ViewMilestones field to given value.

### HasViewMilestones

`func (o *Following) HasViewMilestones() bool`

HasViewMilestones returns a boolean if a field has been set.

### GetEditMilestones

`func (o *Following) GetEditMilestones() bool`

GetEditMilestones returns the EditMilestones field if non-nil, zero value otherwise.

### GetEditMilestonesOk

`func (o *Following) GetEditMilestonesOk() (*bool, bool)`

GetEditMilestonesOk returns a tuple with the EditMilestones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditMilestones

`func (o *Following) SetEditMilestones(v bool)`

SetEditMilestones sets EditMilestones field to given value.

### HasEditMilestones

`func (o *Following) HasEditMilestones() bool`

HasEditMilestones returns a boolean if a field has been set.

### GetCoOwner

`func (o *Following) GetCoOwner() bool`

GetCoOwner returns the CoOwner field if non-nil, zero value otherwise.

### GetCoOwnerOk

`func (o *Following) GetCoOwnerOk() (*bool, bool)`

GetCoOwnerOk returns a tuple with the CoOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoOwner

`func (o *Following) SetCoOwner(v bool)`

SetCoOwner sets CoOwner field to given value.

### HasCoOwner

`func (o *Following) HasCoOwner() bool`

HasCoOwner returns a boolean if a field has been set.

### GetSortOrder

`func (o *Following) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *Following) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *Following) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *Following) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSendFlashback

`func (o *Following) GetSendFlashback() bool`

GetSendFlashback returns the SendFlashback field if non-nil, zero value otherwise.

### GetSendFlashbackOk

`func (o *Following) GetSendFlashbackOk() (*bool, bool)`

GetSendFlashbackOk returns a tuple with the SendFlashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFlashback

`func (o *Following) SetSendFlashback(v bool)`

SetSendFlashback sets SendFlashback field to given value.

### HasSendFlashback

`func (o *Following) HasSendFlashback() bool`

HasSendFlashback returns a boolean if a field has been set.

### GetRelationship

`func (o *Following) GetRelationship() Relationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *Following) GetRelationshipOk() (*Relationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *Following) SetRelationship(v Relationship)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *Following) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetEmailFrequencyOnNewEvent

`func (o *Following) GetEmailFrequencyOnNewEvent() EmailFrequency`

GetEmailFrequencyOnNewEvent returns the EmailFrequencyOnNewEvent field if non-nil, zero value otherwise.

### GetEmailFrequencyOnNewEventOk

`func (o *Following) GetEmailFrequencyOnNewEventOk() (*EmailFrequency, bool)`

GetEmailFrequencyOnNewEventOk returns a tuple with the EmailFrequencyOnNewEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrequencyOnNewEvent

`func (o *Following) SetEmailFrequencyOnNewEvent(v EmailFrequency)`

SetEmailFrequencyOnNewEvent sets EmailFrequencyOnNewEvent field to given value.

### HasEmailFrequencyOnNewEvent

`func (o *Following) HasEmailFrequencyOnNewEvent() bool`

HasEmailFrequencyOnNewEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *Following) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Following) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Following) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Following) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetJournalId

`func (o *Following) GetJournalId() int64`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *Following) GetJournalIdOk() (*int64, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *Following) SetJournalId(v int64)`

SetJournalId sets JournalId field to given value.

### HasJournalId

`func (o *Following) HasJournalId() bool`

HasJournalId returns a boolean if a field has been set.

### GetJournal

`func (o *Following) GetJournal() Journal`

GetJournal returns the Journal field if non-nil, zero value otherwise.

### GetJournalOk

`func (o *Following) GetJournalOk() (*Journal, bool)`

GetJournalOk returns a tuple with the Journal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournal

`func (o *Following) SetJournal(v Journal)`

SetJournal sets Journal field to given value.

### HasJournal

`func (o *Following) HasJournal() bool`

HasJournal returns a boolean if a field has been set.

### GetUser

`func (o *Following) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Following) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Following) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Following) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


