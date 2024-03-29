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

// AuthenticateResponse struct for AuthenticateResponse
type AuthenticateResponse struct {
	AccessToken *string `json:"accessToken,omitempty"`
	DidCreateUser *bool `json:"didCreateUser,omitempty"`
	Status *string `json:"status,omitempty"`
	User *User `json:"user,omitempty"`
}

// NewAuthenticateResponse instantiates a new AuthenticateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateResponse() *AuthenticateResponse {
	this := AuthenticateResponse{}
	return &this
}

// NewAuthenticateResponseWithDefaults instantiates a new AuthenticateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateResponseWithDefaults() *AuthenticateResponse {
	this := AuthenticateResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AuthenticateResponse) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
    return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthenticateResponse) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AuthenticateResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetDidCreateUser returns the DidCreateUser field value if set, zero value otherwise.
func (o *AuthenticateResponse) GetDidCreateUser() bool {
	if o == nil || isNil(o.DidCreateUser) {
		var ret bool
		return ret
	}
	return *o.DidCreateUser
}

// GetDidCreateUserOk returns a tuple with the DidCreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateResponse) GetDidCreateUserOk() (*bool, bool) {
	if o == nil || isNil(o.DidCreateUser) {
    return nil, false
	}
	return o.DidCreateUser, true
}

// HasDidCreateUser returns a boolean if a field has been set.
func (o *AuthenticateResponse) HasDidCreateUser() bool {
	if o != nil && !isNil(o.DidCreateUser) {
		return true
	}

	return false
}

// SetDidCreateUser gets a reference to the given bool and assigns it to the DidCreateUser field.
func (o *AuthenticateResponse) SetDidCreateUser(v bool) {
	o.DidCreateUser = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthenticateResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthenticateResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthenticateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthenticateResponse) GetUser() User {
	if o == nil || isNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateResponse) GetUserOk() (*User, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthenticateResponse) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *AuthenticateResponse) SetUser(v User) {
	o.User = &v
}

func (o AuthenticateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.DidCreateUser) {
		toSerialize["didCreateUser"] = o.DidCreateUser
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateResponse struct {
	value *AuthenticateResponse
	isSet bool
}

func (v NullableAuthenticateResponse) Get() *AuthenticateResponse {
	return v.value
}

func (v *NullableAuthenticateResponse) Set(val *AuthenticateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateResponse(val *AuthenticateResponse) *NullableAuthenticateResponse {
	return &NullableAuthenticateResponse{value: val, isSet: true}
}

func (v NullableAuthenticateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


