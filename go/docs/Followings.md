# Followings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Followings** | Pointer to [**[]Following**](Following.md) |  | [optional] 

## Methods

### NewFollowings

`func NewFollowings() *Followings`

NewFollowings instantiates a new Followings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowingsWithDefaults

`func NewFollowingsWithDefaults() *Followings`

NewFollowingsWithDefaults instantiates a new Followings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Followings) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Followings) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Followings) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Followings) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFollowings

`func (o *Followings) GetFollowings() []Following`

GetFollowings returns the Followings field if non-nil, zero value otherwise.

### GetFollowingsOk

`func (o *Followings) GetFollowingsOk() (*[]Following, bool)`

GetFollowingsOk returns a tuple with the Followings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowings

`func (o *Followings) SetFollowings(v []Following)`

SetFollowings sets Followings field to given value.

### HasFollowings

`func (o *Followings) HasFollowings() bool`

HasFollowings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


