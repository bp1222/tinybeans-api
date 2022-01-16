# Journals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Journals** | Pointer to [**[]Journal**](Journal.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewJournals

`func NewJournals() *Journals`

NewJournals instantiates a new Journals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalsWithDefaults

`func NewJournalsWithDefaults() *Journals`

NewJournalsWithDefaults instantiates a new Journals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJournals

`func (o *Journals) GetJournals() []Journal`

GetJournals returns the Journals field if non-nil, zero value otherwise.

### GetJournalsOk

`func (o *Journals) GetJournalsOk() (*[]Journal, bool)`

GetJournalsOk returns a tuple with the Journals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournals

`func (o *Journals) SetJournals(v []Journal)`

SetJournals sets Journals field to given value.

### HasJournals

`func (o *Journals) HasJournals() bool`

HasJournals returns a boolean if a field has been set.

### GetStatus

`func (o *Journals) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Journals) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Journals) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Journals) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


