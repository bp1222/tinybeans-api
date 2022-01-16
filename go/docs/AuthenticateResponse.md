# AuthenticateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**DidCreateUser** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewAuthenticateResponse

`func NewAuthenticateResponse() *AuthenticateResponse`

NewAuthenticateResponse instantiates a new AuthenticateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateResponseWithDefaults

`func NewAuthenticateResponseWithDefaults() *AuthenticateResponse`

NewAuthenticateResponseWithDefaults instantiates a new AuthenticateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AuthenticateResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthenticateResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthenticateResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthenticateResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetDidCreateUser

`func (o *AuthenticateResponse) GetDidCreateUser() bool`

GetDidCreateUser returns the DidCreateUser field if non-nil, zero value otherwise.

### GetDidCreateUserOk

`func (o *AuthenticateResponse) GetDidCreateUserOk() (*bool, bool)`

GetDidCreateUserOk returns a tuple with the DidCreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidCreateUser

`func (o *AuthenticateResponse) SetDidCreateUser(v bool)`

SetDidCreateUser sets DidCreateUser field to given value.

### HasDidCreateUser

`func (o *AuthenticateResponse) HasDidCreateUser() bool`

HasDidCreateUser returns a boolean if a field has been set.

### GetStatus

`func (o *AuthenticateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *AuthenticateResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticateResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticateResponse) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticateResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


