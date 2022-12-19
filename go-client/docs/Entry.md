# Entry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**PrivateMode** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ClientRef** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Year** | Pointer to **int64** |  | [optional] 
**Month** | Pointer to **int64** |  | [optional] 
**Day** | Pointer to **int64** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** |  | [optional] 
**Blobs** | Pointer to [**Blob**](Blob.md) |  | [optional] 
**JournalId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**Emotions** | Pointer to [**[]Emotion**](Emotion.md) |  | [optional] 

## Methods

### NewEntry

`func NewEntry() *Entry`

NewEntry instantiates a new Entry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryWithDefaults

`func NewEntryWithDefaults() *Entry`

NewEntryWithDefaults instantiates a new Entry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Entry) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Entry) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Entry) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Entry) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *Entry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Entry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *Entry) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Entry) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Entry) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Entry) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetPrivateMode

`func (o *Entry) GetPrivateMode() bool`

GetPrivateMode returns the PrivateMode field if non-nil, zero value otherwise.

### GetPrivateModeOk

`func (o *Entry) GetPrivateModeOk() (*bool, bool)`

GetPrivateModeOk returns a tuple with the PrivateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateMode

`func (o *Entry) SetPrivateMode(v bool)`

SetPrivateMode sets PrivateMode field to given value.

### HasPrivateMode

`func (o *Entry) HasPrivateMode() bool`

HasPrivateMode returns a boolean if a field has been set.

### GetUuid

`func (o *Entry) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Entry) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Entry) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Entry) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetClientRef

`func (o *Entry) GetClientRef() string`

GetClientRef returns the ClientRef field if non-nil, zero value otherwise.

### GetClientRefOk

`func (o *Entry) GetClientRefOk() (*string, bool)`

GetClientRefOk returns a tuple with the ClientRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRef

`func (o *Entry) SetClientRef(v string)`

SetClientRef sets ClientRef field to given value.

### HasClientRef

`func (o *Entry) HasClientRef() bool`

HasClientRef returns a boolean if a field has been set.

### GetType

`func (o *Entry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Entry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetYear

`func (o *Entry) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Entry) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Entry) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *Entry) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetMonth

`func (o *Entry) GetMonth() int64`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Entry) GetMonthOk() (*int64, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Entry) SetMonth(v int64)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *Entry) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetDay

`func (o *Entry) GetDay() int64`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Entry) GetDayOk() (*int64, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Entry) SetDay(v int64)`

SetDay sets Day field to given value.

### HasDay

`func (o *Entry) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetCaption

`func (o *Entry) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *Entry) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *Entry) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *Entry) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetTimestamp

`func (o *Entry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Entry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Entry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Entry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *Entry) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *Entry) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *Entry) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *Entry) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetBlobs

`func (o *Entry) GetBlobs() Blob`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *Entry) GetBlobsOk() (*Blob, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *Entry) SetBlobs(v Blob)`

SetBlobs sets Blobs field to given value.

### HasBlobs

`func (o *Entry) HasBlobs() bool`

HasBlobs returns a boolean if a field has been set.

### GetJournalId

`func (o *Entry) GetJournalId() int64`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *Entry) GetJournalIdOk() (*int64, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *Entry) SetJournalId(v int64)`

SetJournalId sets JournalId field to given value.

### HasJournalId

`func (o *Entry) HasJournalId() bool`

HasJournalId returns a boolean if a field has been set.

### GetUserId

`func (o *Entry) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Entry) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Entry) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Entry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmotions

`func (o *Entry) GetEmotions() []Emotion`

GetEmotions returns the Emotions field if non-nil, zero value otherwise.

### GetEmotionsOk

`func (o *Entry) GetEmotionsOk() (*[]Emotion, bool)`

GetEmotionsOk returns a tuple with the Emotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmotions

`func (o *Entry) SetEmotions(v []Emotion)`

SetEmotions sets Emotions field to given value.

### HasEmotions

`func (o *Entry) HasEmotions() bool`

HasEmotions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


