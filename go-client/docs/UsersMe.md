# UsersMe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewUsersMe

`func NewUsersMe() *UsersMe`

NewUsersMe instantiates a new UsersMe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersMeWithDefaults

`func NewUsersMeWithDefaults() *UsersMe`

NewUsersMeWithDefaults instantiates a new UsersMe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UsersMe) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsersMe) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsersMe) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UsersMe) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *UsersMe) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsersMe) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsersMe) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UsersMe) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


