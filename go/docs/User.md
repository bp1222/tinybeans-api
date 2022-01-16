# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Avatars** | Pointer to [**Avatars**](Avatars.md) |  | [optional] 
**DateStyle** | Pointer to [**DateStyle**](DateStyle.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**EmailFrequencyOnNewComment** | Pointer to [**EmailFrequency**](EmailFrequency.md) |  | [optional] 
**EmailFrequencyOnNewEmotion** | Pointer to [**EmailFrequency**](EmailFrequency.md) |  | [optional] 
**EmailMarketingOptOut** | Pointer to **bool** |  | [optional] 
**EmailOptOut** | Pointer to **bool** |  | [optional] 
**EmailWeeklySummary** | Pointer to **bool** |  | [optional] 
**ExternalProviderRef** | Pointer to **string** |  | [optional] 
**FacebookId** | Pointer to **string** |  | [optional] 
**FirstLoggedInTimestamp** | Pointer to **int64** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**HasMemoriesAccess** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InstagramId** | Pointer to **string** |  | [optional] 
**InstagramUsername** | Pointer to **string** |  | [optional] 
**LastEmailOpenTimestamp** | Pointer to **int64** |  | [optional] 
**LastLoggedInTimestamp** | Pointer to **int64** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** |  | [optional] 
**MeasurementSystem** | Pointer to [**MeasurementSystem**](MeasurementSystem.md) |  | [optional] 
**PublicUsername** | Pointer to **string** |  | [optional] 
**ReferralCode** | Pointer to **string** |  | [optional] 
**ReferralUrl** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to [**TimeZone**](TimeZone.md) |  | [optional] 
**TimeZoneOffset** | Pointer to **int64** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *User) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *User) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *User) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *User) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetAvatars

`func (o *User) GetAvatars() Avatars`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *User) GetAvatarsOk() (*Avatars, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *User) SetAvatars(v Avatars)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *User) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetDateStyle

`func (o *User) GetDateStyle() DateStyle`

GetDateStyle returns the DateStyle field if non-nil, zero value otherwise.

### GetDateStyleOk

`func (o *User) GetDateStyleOk() (*DateStyle, bool)`

GetDateStyleOk returns a tuple with the DateStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStyle

`func (o *User) SetDateStyle(v DateStyle)`

SetDateStyle sets DateStyle field to given value.

### HasDateStyle

`func (o *User) HasDateStyle() bool`

HasDateStyle returns a boolean if a field has been set.

### GetDeleted

`func (o *User) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *User) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *User) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *User) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *User) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailFrequencyOnNewComment

`func (o *User) GetEmailFrequencyOnNewComment() EmailFrequency`

GetEmailFrequencyOnNewComment returns the EmailFrequencyOnNewComment field if non-nil, zero value otherwise.

### GetEmailFrequencyOnNewCommentOk

`func (o *User) GetEmailFrequencyOnNewCommentOk() (*EmailFrequency, bool)`

GetEmailFrequencyOnNewCommentOk returns a tuple with the EmailFrequencyOnNewComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrequencyOnNewComment

`func (o *User) SetEmailFrequencyOnNewComment(v EmailFrequency)`

SetEmailFrequencyOnNewComment sets EmailFrequencyOnNewComment field to given value.

### HasEmailFrequencyOnNewComment

`func (o *User) HasEmailFrequencyOnNewComment() bool`

HasEmailFrequencyOnNewComment returns a boolean if a field has been set.

### GetEmailFrequencyOnNewEmotion

`func (o *User) GetEmailFrequencyOnNewEmotion() EmailFrequency`

GetEmailFrequencyOnNewEmotion returns the EmailFrequencyOnNewEmotion field if non-nil, zero value otherwise.

### GetEmailFrequencyOnNewEmotionOk

`func (o *User) GetEmailFrequencyOnNewEmotionOk() (*EmailFrequency, bool)`

GetEmailFrequencyOnNewEmotionOk returns a tuple with the EmailFrequencyOnNewEmotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrequencyOnNewEmotion

`func (o *User) SetEmailFrequencyOnNewEmotion(v EmailFrequency)`

SetEmailFrequencyOnNewEmotion sets EmailFrequencyOnNewEmotion field to given value.

### HasEmailFrequencyOnNewEmotion

`func (o *User) HasEmailFrequencyOnNewEmotion() bool`

HasEmailFrequencyOnNewEmotion returns a boolean if a field has been set.

### GetEmailMarketingOptOut

`func (o *User) GetEmailMarketingOptOut() bool`

GetEmailMarketingOptOut returns the EmailMarketingOptOut field if non-nil, zero value otherwise.

### GetEmailMarketingOptOutOk

`func (o *User) GetEmailMarketingOptOutOk() (*bool, bool)`

GetEmailMarketingOptOutOk returns a tuple with the EmailMarketingOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMarketingOptOut

`func (o *User) SetEmailMarketingOptOut(v bool)`

SetEmailMarketingOptOut sets EmailMarketingOptOut field to given value.

### HasEmailMarketingOptOut

`func (o *User) HasEmailMarketingOptOut() bool`

HasEmailMarketingOptOut returns a boolean if a field has been set.

### GetEmailOptOut

`func (o *User) GetEmailOptOut() bool`

GetEmailOptOut returns the EmailOptOut field if non-nil, zero value otherwise.

### GetEmailOptOutOk

`func (o *User) GetEmailOptOutOk() (*bool, bool)`

GetEmailOptOutOk returns a tuple with the EmailOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOptOut

`func (o *User) SetEmailOptOut(v bool)`

SetEmailOptOut sets EmailOptOut field to given value.

### HasEmailOptOut

`func (o *User) HasEmailOptOut() bool`

HasEmailOptOut returns a boolean if a field has been set.

### GetEmailWeeklySummary

`func (o *User) GetEmailWeeklySummary() bool`

GetEmailWeeklySummary returns the EmailWeeklySummary field if non-nil, zero value otherwise.

### GetEmailWeeklySummaryOk

`func (o *User) GetEmailWeeklySummaryOk() (*bool, bool)`

GetEmailWeeklySummaryOk returns a tuple with the EmailWeeklySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailWeeklySummary

`func (o *User) SetEmailWeeklySummary(v bool)`

SetEmailWeeklySummary sets EmailWeeklySummary field to given value.

### HasEmailWeeklySummary

`func (o *User) HasEmailWeeklySummary() bool`

HasEmailWeeklySummary returns a boolean if a field has been set.

### GetExternalProviderRef

`func (o *User) GetExternalProviderRef() string`

GetExternalProviderRef returns the ExternalProviderRef field if non-nil, zero value otherwise.

### GetExternalProviderRefOk

`func (o *User) GetExternalProviderRefOk() (*string, bool)`

GetExternalProviderRefOk returns a tuple with the ExternalProviderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviderRef

`func (o *User) SetExternalProviderRef(v string)`

SetExternalProviderRef sets ExternalProviderRef field to given value.

### HasExternalProviderRef

`func (o *User) HasExternalProviderRef() bool`

HasExternalProviderRef returns a boolean if a field has been set.

### GetFacebookId

`func (o *User) GetFacebookId() string`

GetFacebookId returns the FacebookId field if non-nil, zero value otherwise.

### GetFacebookIdOk

`func (o *User) GetFacebookIdOk() (*string, bool)`

GetFacebookIdOk returns a tuple with the FacebookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookId

`func (o *User) SetFacebookId(v string)`

SetFacebookId sets FacebookId field to given value.

### HasFacebookId

`func (o *User) HasFacebookId() bool`

HasFacebookId returns a boolean if a field has been set.

### GetFirstLoggedInTimestamp

`func (o *User) GetFirstLoggedInTimestamp() int64`

GetFirstLoggedInTimestamp returns the FirstLoggedInTimestamp field if non-nil, zero value otherwise.

### GetFirstLoggedInTimestampOk

`func (o *User) GetFirstLoggedInTimestampOk() (*int64, bool)`

GetFirstLoggedInTimestampOk returns a tuple with the FirstLoggedInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLoggedInTimestamp

`func (o *User) SetFirstLoggedInTimestamp(v int64)`

SetFirstLoggedInTimestamp sets FirstLoggedInTimestamp field to given value.

### HasFirstLoggedInTimestamp

`func (o *User) HasFirstLoggedInTimestamp() bool`

HasFirstLoggedInTimestamp returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *User) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *User) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *User) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *User) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHasMemoriesAccess

`func (o *User) GetHasMemoriesAccess() bool`

GetHasMemoriesAccess returns the HasMemoriesAccess field if non-nil, zero value otherwise.

### GetHasMemoriesAccessOk

`func (o *User) GetHasMemoriesAccessOk() (*bool, bool)`

GetHasMemoriesAccessOk returns a tuple with the HasMemoriesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMemoriesAccess

`func (o *User) SetHasMemoriesAccess(v bool)`

SetHasMemoriesAccess sets HasMemoriesAccess field to given value.

### HasHasMemoriesAccess

`func (o *User) HasHasMemoriesAccess() bool`

HasHasMemoriesAccess returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstagramId

`func (o *User) GetInstagramId() string`

GetInstagramId returns the InstagramId field if non-nil, zero value otherwise.

### GetInstagramIdOk

`func (o *User) GetInstagramIdOk() (*string, bool)`

GetInstagramIdOk returns a tuple with the InstagramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramId

`func (o *User) SetInstagramId(v string)`

SetInstagramId sets InstagramId field to given value.

### HasInstagramId

`func (o *User) HasInstagramId() bool`

HasInstagramId returns a boolean if a field has been set.

### GetInstagramUsername

`func (o *User) GetInstagramUsername() string`

GetInstagramUsername returns the InstagramUsername field if non-nil, zero value otherwise.

### GetInstagramUsernameOk

`func (o *User) GetInstagramUsernameOk() (*string, bool)`

GetInstagramUsernameOk returns a tuple with the InstagramUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUsername

`func (o *User) SetInstagramUsername(v string)`

SetInstagramUsername sets InstagramUsername field to given value.

### HasInstagramUsername

`func (o *User) HasInstagramUsername() bool`

HasInstagramUsername returns a boolean if a field has been set.

### GetLastEmailOpenTimestamp

`func (o *User) GetLastEmailOpenTimestamp() int64`

GetLastEmailOpenTimestamp returns the LastEmailOpenTimestamp field if non-nil, zero value otherwise.

### GetLastEmailOpenTimestampOk

`func (o *User) GetLastEmailOpenTimestampOk() (*int64, bool)`

GetLastEmailOpenTimestampOk returns a tuple with the LastEmailOpenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEmailOpenTimestamp

`func (o *User) SetLastEmailOpenTimestamp(v int64)`

SetLastEmailOpenTimestamp sets LastEmailOpenTimestamp field to given value.

### HasLastEmailOpenTimestamp

`func (o *User) HasLastEmailOpenTimestamp() bool`

HasLastEmailOpenTimestamp returns a boolean if a field has been set.

### GetLastLoggedInTimestamp

`func (o *User) GetLastLoggedInTimestamp() int64`

GetLastLoggedInTimestamp returns the LastLoggedInTimestamp field if non-nil, zero value otherwise.

### GetLastLoggedInTimestampOk

`func (o *User) GetLastLoggedInTimestampOk() (*int64, bool)`

GetLastLoggedInTimestampOk returns a tuple with the LastLoggedInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInTimestamp

`func (o *User) SetLastLoggedInTimestamp(v int64)`

SetLastLoggedInTimestamp sets LastLoggedInTimestamp field to given value.

### HasLastLoggedInTimestamp

`func (o *User) HasLastLoggedInTimestamp() bool`

HasLastLoggedInTimestamp returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *User) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *User) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *User) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *User) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetMeasurementSystem

`func (o *User) GetMeasurementSystem() MeasurementSystem`

GetMeasurementSystem returns the MeasurementSystem field if non-nil, zero value otherwise.

### GetMeasurementSystemOk

`func (o *User) GetMeasurementSystemOk() (*MeasurementSystem, bool)`

GetMeasurementSystemOk returns a tuple with the MeasurementSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementSystem

`func (o *User) SetMeasurementSystem(v MeasurementSystem)`

SetMeasurementSystem sets MeasurementSystem field to given value.

### HasMeasurementSystem

`func (o *User) HasMeasurementSystem() bool`

HasMeasurementSystem returns a boolean if a field has been set.

### GetPublicUsername

`func (o *User) GetPublicUsername() string`

GetPublicUsername returns the PublicUsername field if non-nil, zero value otherwise.

### GetPublicUsernameOk

`func (o *User) GetPublicUsernameOk() (*string, bool)`

GetPublicUsernameOk returns a tuple with the PublicUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUsername

`func (o *User) SetPublicUsername(v string)`

SetPublicUsername sets PublicUsername field to given value.

### HasPublicUsername

`func (o *User) HasPublicUsername() bool`

HasPublicUsername returns a boolean if a field has been set.

### GetReferralCode

`func (o *User) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *User) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *User) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *User) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.

### GetReferralUrl

`func (o *User) GetReferralUrl() string`

GetReferralUrl returns the ReferralUrl field if non-nil, zero value otherwise.

### GetReferralUrlOk

`func (o *User) GetReferralUrlOk() (*string, bool)`

GetReferralUrlOk returns a tuple with the ReferralUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralUrl

`func (o *User) SetReferralUrl(v string)`

SetReferralUrl sets ReferralUrl field to given value.

### HasReferralUrl

`func (o *User) HasReferralUrl() bool`

HasReferralUrl returns a boolean if a field has been set.

### GetTimeZone

`func (o *User) GetTimeZone() TimeZone`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *User) GetTimeZoneOk() (*TimeZone, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *User) SetTimeZone(v TimeZone)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *User) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *User) GetTimeZoneOffset() int64`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *User) GetTimeZoneOffsetOk() (*int64, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *User) SetTimeZoneOffset(v int64)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *User) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.

### GetTimestamp

`func (o *User) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *User) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *User) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *User) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


