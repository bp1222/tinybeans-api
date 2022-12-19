# Entries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**NumEntriesRemaining** | Pointer to **int64** |  | [optional] 
**Entries** | Pointer to [**[]Entry**](Entry.md) |  | [optional] 

## Methods

### NewEntries

`func NewEntries() *Entries`

NewEntries instantiates a new Entries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesWithDefaults

`func NewEntriesWithDefaults() *Entries`

NewEntriesWithDefaults instantiates a new Entries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Entries) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Entries) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Entries) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Entries) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNumEntriesRemaining

`func (o *Entries) GetNumEntriesRemaining() int64`

GetNumEntriesRemaining returns the NumEntriesRemaining field if non-nil, zero value otherwise.

### GetNumEntriesRemainingOk

`func (o *Entries) GetNumEntriesRemainingOk() (*int64, bool)`

GetNumEntriesRemainingOk returns a tuple with the NumEntriesRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEntriesRemaining

`func (o *Entries) SetNumEntriesRemaining(v int64)`

SetNumEntriesRemaining sets NumEntriesRemaining field to given value.

### HasNumEntriesRemaining

`func (o *Entries) HasNumEntriesRemaining() bool`

HasNumEntriesRemaining returns a boolean if a field has been set.

### GetEntries

`func (o *Entries) GetEntries() []Entry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Entries) GetEntriesOk() (*[]Entry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Entries) SetEntries(v []Entry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *Entries) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


