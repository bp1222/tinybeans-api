/*
Tinybeans API - Unofficial

Reverse Engineered API for Tinybeans.  This is horribly incomplete, and not supported

API version: 0.0.1
Contact: dave@mudsite.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tinybeans

import (
	"encoding/json"
)

// SubscriptionProduct struct for SubscriptionProduct
type SubscriptionProduct struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DateCreated *string `json:"dateCreated,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	ProductStatus *string `json:"productStatus,omitempty"`
	CodeAppStore *string `json:"codeAppStore,omitempty"`
	CodePlayStore *string `json:"codePlayStore,omitempty"`
	AmountInCents *int64 `json:"amountInCents,omitempty"`
	AmountInDollarsAndCents *float32 `json:"amountInDollarsAndCents,omitempty"`
	Price *Price `json:"price,omitempty"`
	URL *string `json:"URL,omitempty"`
	Tier *string `json:"tier,omitempty"`
	CycleDuration *int64 `json:"cycleDuration,omitempty"`
	Cycle *Cycle `json:"cycle,omitempty"`
	Features []Features `json:"features,omitempty"`
}

// NewSubscriptionProduct instantiates a new SubscriptionProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionProduct() *SubscriptionProduct {
	this := SubscriptionProduct{}
	return &this
}

// NewSubscriptionProductWithDefaults instantiates a new SubscriptionProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionProductWithDefaults() *SubscriptionProduct {
	this := SubscriptionProduct{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SubscriptionProduct) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SubscriptionProduct) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionProduct) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubscriptionProduct) SetDescription(v string) {
	o.Description = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetDateCreated() string {
	if o == nil || isNil(o.DateCreated) {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetDateCreatedOk() (*string, bool) {
	if o == nil || isNil(o.DateCreated) {
    return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *SubscriptionProduct) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetLastUpdated() string {
	if o == nil || isNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetLastUpdatedOk() (*string, bool) {
	if o == nil || isNil(o.LastUpdated) {
    return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *SubscriptionProduct) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetProductStatus returns the ProductStatus field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetProductStatus() string {
	if o == nil || isNil(o.ProductStatus) {
		var ret string
		return ret
	}
	return *o.ProductStatus
}

// GetProductStatusOk returns a tuple with the ProductStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetProductStatusOk() (*string, bool) {
	if o == nil || isNil(o.ProductStatus) {
    return nil, false
	}
	return o.ProductStatus, true
}

// HasProductStatus returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasProductStatus() bool {
	if o != nil && !isNil(o.ProductStatus) {
		return true
	}

	return false
}

// SetProductStatus gets a reference to the given string and assigns it to the ProductStatus field.
func (o *SubscriptionProduct) SetProductStatus(v string) {
	o.ProductStatus = &v
}

// GetCodeAppStore returns the CodeAppStore field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetCodeAppStore() string {
	if o == nil || isNil(o.CodeAppStore) {
		var ret string
		return ret
	}
	return *o.CodeAppStore
}

// GetCodeAppStoreOk returns a tuple with the CodeAppStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetCodeAppStoreOk() (*string, bool) {
	if o == nil || isNil(o.CodeAppStore) {
    return nil, false
	}
	return o.CodeAppStore, true
}

// HasCodeAppStore returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasCodeAppStore() bool {
	if o != nil && !isNil(o.CodeAppStore) {
		return true
	}

	return false
}

// SetCodeAppStore gets a reference to the given string and assigns it to the CodeAppStore field.
func (o *SubscriptionProduct) SetCodeAppStore(v string) {
	o.CodeAppStore = &v
}

// GetCodePlayStore returns the CodePlayStore field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetCodePlayStore() string {
	if o == nil || isNil(o.CodePlayStore) {
		var ret string
		return ret
	}
	return *o.CodePlayStore
}

// GetCodePlayStoreOk returns a tuple with the CodePlayStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetCodePlayStoreOk() (*string, bool) {
	if o == nil || isNil(o.CodePlayStore) {
    return nil, false
	}
	return o.CodePlayStore, true
}

// HasCodePlayStore returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasCodePlayStore() bool {
	if o != nil && !isNil(o.CodePlayStore) {
		return true
	}

	return false
}

// SetCodePlayStore gets a reference to the given string and assigns it to the CodePlayStore field.
func (o *SubscriptionProduct) SetCodePlayStore(v string) {
	o.CodePlayStore = &v
}

// GetAmountInCents returns the AmountInCents field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetAmountInCents() int64 {
	if o == nil || isNil(o.AmountInCents) {
		var ret int64
		return ret
	}
	return *o.AmountInCents
}

// GetAmountInCentsOk returns a tuple with the AmountInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetAmountInCentsOk() (*int64, bool) {
	if o == nil || isNil(o.AmountInCents) {
    return nil, false
	}
	return o.AmountInCents, true
}

// HasAmountInCents returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasAmountInCents() bool {
	if o != nil && !isNil(o.AmountInCents) {
		return true
	}

	return false
}

// SetAmountInCents gets a reference to the given int64 and assigns it to the AmountInCents field.
func (o *SubscriptionProduct) SetAmountInCents(v int64) {
	o.AmountInCents = &v
}

// GetAmountInDollarsAndCents returns the AmountInDollarsAndCents field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetAmountInDollarsAndCents() float32 {
	if o == nil || isNil(o.AmountInDollarsAndCents) {
		var ret float32
		return ret
	}
	return *o.AmountInDollarsAndCents
}

// GetAmountInDollarsAndCentsOk returns a tuple with the AmountInDollarsAndCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetAmountInDollarsAndCentsOk() (*float32, bool) {
	if o == nil || isNil(o.AmountInDollarsAndCents) {
    return nil, false
	}
	return o.AmountInDollarsAndCents, true
}

// HasAmountInDollarsAndCents returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasAmountInDollarsAndCents() bool {
	if o != nil && !isNil(o.AmountInDollarsAndCents) {
		return true
	}

	return false
}

// SetAmountInDollarsAndCents gets a reference to the given float32 and assigns it to the AmountInDollarsAndCents field.
func (o *SubscriptionProduct) SetAmountInDollarsAndCents(v float32) {
	o.AmountInDollarsAndCents = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetPrice() Price {
	if o == nil || isNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetPriceOk() (*Price, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *SubscriptionProduct) SetPrice(v Price) {
	o.Price = &v
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetURL() string {
	if o == nil || isNil(o.URL) {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetURLOk() (*string, bool) {
	if o == nil || isNil(o.URL) {
    return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasURL() bool {
	if o != nil && !isNil(o.URL) {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *SubscriptionProduct) SetURL(v string) {
	o.URL = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetTier() string {
	if o == nil || isNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetTierOk() (*string, bool) {
	if o == nil || isNil(o.Tier) {
    return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasTier() bool {
	if o != nil && !isNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *SubscriptionProduct) SetTier(v string) {
	o.Tier = &v
}

// GetCycleDuration returns the CycleDuration field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetCycleDuration() int64 {
	if o == nil || isNil(o.CycleDuration) {
		var ret int64
		return ret
	}
	return *o.CycleDuration
}

// GetCycleDurationOk returns a tuple with the CycleDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetCycleDurationOk() (*int64, bool) {
	if o == nil || isNil(o.CycleDuration) {
    return nil, false
	}
	return o.CycleDuration, true
}

// HasCycleDuration returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasCycleDuration() bool {
	if o != nil && !isNil(o.CycleDuration) {
		return true
	}

	return false
}

// SetCycleDuration gets a reference to the given int64 and assigns it to the CycleDuration field.
func (o *SubscriptionProduct) SetCycleDuration(v int64) {
	o.CycleDuration = &v
}

// GetCycle returns the Cycle field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetCycle() Cycle {
	if o == nil || isNil(o.Cycle) {
		var ret Cycle
		return ret
	}
	return *o.Cycle
}

// GetCycleOk returns a tuple with the Cycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetCycleOk() (*Cycle, bool) {
	if o == nil || isNil(o.Cycle) {
    return nil, false
	}
	return o.Cycle, true
}

// HasCycle returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasCycle() bool {
	if o != nil && !isNil(o.Cycle) {
		return true
	}

	return false
}

// SetCycle gets a reference to the given Cycle and assigns it to the Cycle field.
func (o *SubscriptionProduct) SetCycle(v Cycle) {
	o.Cycle = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *SubscriptionProduct) GetFeatures() []Features {
	if o == nil || isNil(o.Features) {
		var ret []Features
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionProduct) GetFeaturesOk() ([]Features, bool) {
	if o == nil || isNil(o.Features) {
    return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *SubscriptionProduct) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []Features and assigns it to the Features field.
func (o *SubscriptionProduct) SetFeatures(v []Features) {
	o.Features = v
}

func (o SubscriptionProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.ProductStatus) {
		toSerialize["productStatus"] = o.ProductStatus
	}
	if !isNil(o.CodeAppStore) {
		toSerialize["codeAppStore"] = o.CodeAppStore
	}
	if !isNil(o.CodePlayStore) {
		toSerialize["codePlayStore"] = o.CodePlayStore
	}
	if !isNil(o.AmountInCents) {
		toSerialize["amountInCents"] = o.AmountInCents
	}
	if !isNil(o.AmountInDollarsAndCents) {
		toSerialize["amountInDollarsAndCents"] = o.AmountInDollarsAndCents
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.URL) {
		toSerialize["URL"] = o.URL
	}
	if !isNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !isNil(o.CycleDuration) {
		toSerialize["cycleDuration"] = o.CycleDuration
	}
	if !isNil(o.Cycle) {
		toSerialize["cycle"] = o.Cycle
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionProduct struct {
	value *SubscriptionProduct
	isSet bool
}

func (v NullableSubscriptionProduct) Get() *SubscriptionProduct {
	return v.value
}

func (v *NullableSubscriptionProduct) Set(val *SubscriptionProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionProduct(val *SubscriptionProduct) *NullableSubscriptionProduct {
	return &NullableSubscriptionProduct{value: val, isSet: true}
}

func (v NullableSubscriptionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


