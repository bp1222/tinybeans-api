# SubscriptionProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**ProductStatus** | Pointer to **string** |  | [optional] 
**CodeAppStore** | Pointer to **string** |  | [optional] 
**CodePlayStore** | Pointer to **string** |  | [optional] 
**AmountInCents** | Pointer to **int64** |  | [optional] 
**AmountInDollarsAndCents** | Pointer to **float32** |  | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 
**CycleDuration** | Pointer to **int64** |  | [optional] 
**Cycle** | Pointer to [**Cycle**](Cycle.md) |  | [optional] 
**Features** | Pointer to [**[]Features**](Features.md) |  | [optional] 

## Methods

### NewSubscriptionProduct

`func NewSubscriptionProduct() *SubscriptionProduct`

NewSubscriptionProduct instantiates a new SubscriptionProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionProductWithDefaults

`func NewSubscriptionProductWithDefaults() *SubscriptionProduct`

NewSubscriptionProductWithDefaults instantiates a new SubscriptionProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionProduct) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionProduct) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionProduct) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *SubscriptionProduct) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionProduct) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionProduct) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubscriptionProduct) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SubscriptionProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *SubscriptionProduct) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SubscriptionProduct) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SubscriptionProduct) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SubscriptionProduct) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SubscriptionProduct) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SubscriptionProduct) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SubscriptionProduct) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SubscriptionProduct) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProductStatus

`func (o *SubscriptionProduct) GetProductStatus() string`

GetProductStatus returns the ProductStatus field if non-nil, zero value otherwise.

### GetProductStatusOk

`func (o *SubscriptionProduct) GetProductStatusOk() (*string, bool)`

GetProductStatusOk returns a tuple with the ProductStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductStatus

`func (o *SubscriptionProduct) SetProductStatus(v string)`

SetProductStatus sets ProductStatus field to given value.

### HasProductStatus

`func (o *SubscriptionProduct) HasProductStatus() bool`

HasProductStatus returns a boolean if a field has been set.

### GetCodeAppStore

`func (o *SubscriptionProduct) GetCodeAppStore() string`

GetCodeAppStore returns the CodeAppStore field if non-nil, zero value otherwise.

### GetCodeAppStoreOk

`func (o *SubscriptionProduct) GetCodeAppStoreOk() (*string, bool)`

GetCodeAppStoreOk returns a tuple with the CodeAppStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeAppStore

`func (o *SubscriptionProduct) SetCodeAppStore(v string)`

SetCodeAppStore sets CodeAppStore field to given value.

### HasCodeAppStore

`func (o *SubscriptionProduct) HasCodeAppStore() bool`

HasCodeAppStore returns a boolean if a field has been set.

### GetCodePlayStore

`func (o *SubscriptionProduct) GetCodePlayStore() string`

GetCodePlayStore returns the CodePlayStore field if non-nil, zero value otherwise.

### GetCodePlayStoreOk

`func (o *SubscriptionProduct) GetCodePlayStoreOk() (*string, bool)`

GetCodePlayStoreOk returns a tuple with the CodePlayStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePlayStore

`func (o *SubscriptionProduct) SetCodePlayStore(v string)`

SetCodePlayStore sets CodePlayStore field to given value.

### HasCodePlayStore

`func (o *SubscriptionProduct) HasCodePlayStore() bool`

HasCodePlayStore returns a boolean if a field has been set.

### GetAmountInCents

`func (o *SubscriptionProduct) GetAmountInCents() int64`

GetAmountInCents returns the AmountInCents field if non-nil, zero value otherwise.

### GetAmountInCentsOk

`func (o *SubscriptionProduct) GetAmountInCentsOk() (*int64, bool)`

GetAmountInCentsOk returns a tuple with the AmountInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInCents

`func (o *SubscriptionProduct) SetAmountInCents(v int64)`

SetAmountInCents sets AmountInCents field to given value.

### HasAmountInCents

`func (o *SubscriptionProduct) HasAmountInCents() bool`

HasAmountInCents returns a boolean if a field has been set.

### GetAmountInDollarsAndCents

`func (o *SubscriptionProduct) GetAmountInDollarsAndCents() float32`

GetAmountInDollarsAndCents returns the AmountInDollarsAndCents field if non-nil, zero value otherwise.

### GetAmountInDollarsAndCentsOk

`func (o *SubscriptionProduct) GetAmountInDollarsAndCentsOk() (*float32, bool)`

GetAmountInDollarsAndCentsOk returns a tuple with the AmountInDollarsAndCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInDollarsAndCents

`func (o *SubscriptionProduct) SetAmountInDollarsAndCents(v float32)`

SetAmountInDollarsAndCents sets AmountInDollarsAndCents field to given value.

### HasAmountInDollarsAndCents

`func (o *SubscriptionProduct) HasAmountInDollarsAndCents() bool`

HasAmountInDollarsAndCents returns a boolean if a field has been set.

### GetPrice

`func (o *SubscriptionProduct) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SubscriptionProduct) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SubscriptionProduct) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SubscriptionProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetURL

`func (o *SubscriptionProduct) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *SubscriptionProduct) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *SubscriptionProduct) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *SubscriptionProduct) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetTier

`func (o *SubscriptionProduct) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *SubscriptionProduct) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *SubscriptionProduct) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *SubscriptionProduct) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetCycleDuration

`func (o *SubscriptionProduct) GetCycleDuration() int64`

GetCycleDuration returns the CycleDuration field if non-nil, zero value otherwise.

### GetCycleDurationOk

`func (o *SubscriptionProduct) GetCycleDurationOk() (*int64, bool)`

GetCycleDurationOk returns a tuple with the CycleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleDuration

`func (o *SubscriptionProduct) SetCycleDuration(v int64)`

SetCycleDuration sets CycleDuration field to given value.

### HasCycleDuration

`func (o *SubscriptionProduct) HasCycleDuration() bool`

HasCycleDuration returns a boolean if a field has been set.

### GetCycle

`func (o *SubscriptionProduct) GetCycle() Cycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *SubscriptionProduct) GetCycleOk() (*Cycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *SubscriptionProduct) SetCycle(v Cycle)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *SubscriptionProduct) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetFeatures

`func (o *SubscriptionProduct) GetFeatures() []Features`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *SubscriptionProduct) GetFeaturesOk() (*[]Features, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *SubscriptionProduct) SetFeatures(v []Features)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *SubscriptionProduct) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


