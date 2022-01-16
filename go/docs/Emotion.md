# Emotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**EmotionType**](EmotionType.md) |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EntryId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewEmotion

`func NewEmotion() *Emotion`

NewEmotion instantiates a new Emotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmotionWithDefaults

`func NewEmotionWithDefaults() *Emotion`

NewEmotionWithDefaults instantiates a new Emotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Emotion) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Emotion) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Emotion) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Emotion) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *Emotion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Emotion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Emotion) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Emotion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Emotion) GetType() EmotionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Emotion) GetTypeOk() (*EmotionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Emotion) SetType(v EmotionType)`

SetType sets Type field to given value.

### HasType

`func (o *Emotion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimestamp

`func (o *Emotion) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Emotion) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Emotion) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Emotion) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *Emotion) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *Emotion) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *Emotion) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *Emotion) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetDeleted

`func (o *Emotion) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Emotion) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Emotion) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Emotion) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEntryId

`func (o *Emotion) GetEntryId() int64`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *Emotion) GetEntryIdOk() (*int64, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *Emotion) SetEntryId(v int64)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *Emotion) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetUserId

`func (o *Emotion) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Emotion) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Emotion) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Emotion) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


