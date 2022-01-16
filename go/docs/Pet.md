# Pet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Species** | Pointer to **string** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Sex** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**DevelopmentOptOut** | Pointer to **bool** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**AdoptionDate** | Pointer to **string** |  | [optional] 
**Avatars** | Pointer to [**Avatars**](Avatars.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewPet

`func NewPet() *Pet`

NewPet instantiates a new Pet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPetWithDefaults

`func NewPetWithDefaults() *Pet`

NewPetWithDefaults instantiates a new Pet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Pet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Pet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpecies

`func (o *Pet) GetSpecies() string`

GetSpecies returns the Species field if non-nil, zero value otherwise.

### GetSpeciesOk

`func (o *Pet) GetSpeciesOk() (*string, bool)`

GetSpeciesOk returns a tuple with the Species field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecies

`func (o *Pet) SetSpecies(v string)`

SetSpecies sets Species field to given value.

### HasSpecies

`func (o *Pet) HasSpecies() bool`

HasSpecies returns a boolean if a field has been set.

### GetBreed

`func (o *Pet) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *Pet) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *Pet) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *Pet) HasBreed() bool`

HasBreed returns a boolean if a field has been set.

### GetSize

`func (o *Pet) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Pet) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Pet) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Pet) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSex

`func (o *Pet) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *Pet) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *Pet) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *Pet) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetDeleted

`func (o *Pet) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Pet) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Pet) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Pet) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDevelopmentOptOut

`func (o *Pet) GetDevelopmentOptOut() bool`

GetDevelopmentOptOut returns the DevelopmentOptOut field if non-nil, zero value otherwise.

### GetDevelopmentOptOutOk

`func (o *Pet) GetDevelopmentOptOutOk() (*bool, bool)`

GetDevelopmentOptOutOk returns a tuple with the DevelopmentOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopmentOptOut

`func (o *Pet) SetDevelopmentOptOut(v bool)`

SetDevelopmentOptOut sets DevelopmentOptOut field to given value.

### HasDevelopmentOptOut

`func (o *Pet) HasDevelopmentOptOut() bool`

HasDevelopmentOptOut returns a boolean if a field has been set.

### GetBirthday

`func (o *Pet) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *Pet) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *Pet) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *Pet) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetAdoptionDate

`func (o *Pet) GetAdoptionDate() string`

GetAdoptionDate returns the AdoptionDate field if non-nil, zero value otherwise.

### GetAdoptionDateOk

`func (o *Pet) GetAdoptionDateOk() (*string, bool)`

GetAdoptionDateOk returns a tuple with the AdoptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptionDate

`func (o *Pet) SetAdoptionDate(v string)`

SetAdoptionDate sets AdoptionDate field to given value.

### HasAdoptionDate

`func (o *Pet) HasAdoptionDate() bool`

HasAdoptionDate returns a boolean if a field has been set.

### GetAvatars

`func (o *Pet) GetAvatars() Avatars`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *Pet) GetAvatarsOk() (*Avatars, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *Pet) SetAvatars(v Avatars)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *Pet) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetUser

`func (o *Pet) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Pet) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Pet) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Pet) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


