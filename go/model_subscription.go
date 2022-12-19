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

// Subscription struct for Subscription
type Subscription struct {
	URL *string `json:"URL,omitempty"`
	Id *int64 `json:"id,omitempty"`
	StartDate *int64 `json:"startDate,omitempty"`
	ExpiryDate *int64 `json:"expiryDate,omitempty"`
	StartDateDisplay *string `json:"startDateDisplay,omitempty"`
	ExpiryDateDisplay *string `json:"expiryDateDisplay,omitempty"`
	Status *string `json:"status,omitempty"`
	Renewal *string `json:"renewal,omitempty"`
	SubscriptionProductId *int64 `json:"subscriptionProductId,omitempty"`
	SubscriptionProduct *SubscriptionProduct `json:"subscriptionProduct,omitempty"`
	JournalId *int64 `json:"journalId,omitempty"`
	PaymentTransaction *PaymentTransaction `json:"paymentTransaction,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *Subscription) GetURL() string {
	if o == nil || isNil(o.URL) {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetURLOk() (*string, bool) {
	if o == nil || isNil(o.URL) {
    return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *Subscription) HasURL() bool {
	if o != nil && !isNil(o.URL) {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *Subscription) SetURL(v string) {
	o.URL = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() int64 {
	if o == nil || isNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*int64, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Subscription) SetId(v int64) {
	o.Id = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Subscription) GetStartDate() int64 {
	if o == nil || isNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStartDateOk() (*int64, bool) {
	if o == nil || isNil(o.StartDate) {
    return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Subscription) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *Subscription) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Subscription) GetExpiryDate() int64 {
	if o == nil || isNil(o.ExpiryDate) {
		var ret int64
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetExpiryDateOk() (*int64, bool) {
	if o == nil || isNil(o.ExpiryDate) {
    return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Subscription) HasExpiryDate() bool {
	if o != nil && !isNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int64 and assigns it to the ExpiryDate field.
func (o *Subscription) SetExpiryDate(v int64) {
	o.ExpiryDate = &v
}

// GetStartDateDisplay returns the StartDateDisplay field value if set, zero value otherwise.
func (o *Subscription) GetStartDateDisplay() string {
	if o == nil || isNil(o.StartDateDisplay) {
		var ret string
		return ret
	}
	return *o.StartDateDisplay
}

// GetStartDateDisplayOk returns a tuple with the StartDateDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStartDateDisplayOk() (*string, bool) {
	if o == nil || isNil(o.StartDateDisplay) {
    return nil, false
	}
	return o.StartDateDisplay, true
}

// HasStartDateDisplay returns a boolean if a field has been set.
func (o *Subscription) HasStartDateDisplay() bool {
	if o != nil && !isNil(o.StartDateDisplay) {
		return true
	}

	return false
}

// SetStartDateDisplay gets a reference to the given string and assigns it to the StartDateDisplay field.
func (o *Subscription) SetStartDateDisplay(v string) {
	o.StartDateDisplay = &v
}

// GetExpiryDateDisplay returns the ExpiryDateDisplay field value if set, zero value otherwise.
func (o *Subscription) GetExpiryDateDisplay() string {
	if o == nil || isNil(o.ExpiryDateDisplay) {
		var ret string
		return ret
	}
	return *o.ExpiryDateDisplay
}

// GetExpiryDateDisplayOk returns a tuple with the ExpiryDateDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetExpiryDateDisplayOk() (*string, bool) {
	if o == nil || isNil(o.ExpiryDateDisplay) {
    return nil, false
	}
	return o.ExpiryDateDisplay, true
}

// HasExpiryDateDisplay returns a boolean if a field has been set.
func (o *Subscription) HasExpiryDateDisplay() bool {
	if o != nil && !isNil(o.ExpiryDateDisplay) {
		return true
	}

	return false
}

// SetExpiryDateDisplay gets a reference to the given string and assigns it to the ExpiryDateDisplay field.
func (o *Subscription) SetExpiryDateDisplay(v string) {
	o.ExpiryDateDisplay = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Subscription) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Subscription) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Subscription) SetStatus(v string) {
	o.Status = &v
}

// GetRenewal returns the Renewal field value if set, zero value otherwise.
func (o *Subscription) GetRenewal() string {
	if o == nil || isNil(o.Renewal) {
		var ret string
		return ret
	}
	return *o.Renewal
}

// GetRenewalOk returns a tuple with the Renewal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetRenewalOk() (*string, bool) {
	if o == nil || isNil(o.Renewal) {
    return nil, false
	}
	return o.Renewal, true
}

// HasRenewal returns a boolean if a field has been set.
func (o *Subscription) HasRenewal() bool {
	if o != nil && !isNil(o.Renewal) {
		return true
	}

	return false
}

// SetRenewal gets a reference to the given string and assigns it to the Renewal field.
func (o *Subscription) SetRenewal(v string) {
	o.Renewal = &v
}

// GetSubscriptionProductId returns the SubscriptionProductId field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionProductId() int64 {
	if o == nil || isNil(o.SubscriptionProductId) {
		var ret int64
		return ret
	}
	return *o.SubscriptionProductId
}

// GetSubscriptionProductIdOk returns a tuple with the SubscriptionProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionProductIdOk() (*int64, bool) {
	if o == nil || isNil(o.SubscriptionProductId) {
    return nil, false
	}
	return o.SubscriptionProductId, true
}

// HasSubscriptionProductId returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionProductId() bool {
	if o != nil && !isNil(o.SubscriptionProductId) {
		return true
	}

	return false
}

// SetSubscriptionProductId gets a reference to the given int64 and assigns it to the SubscriptionProductId field.
func (o *Subscription) SetSubscriptionProductId(v int64) {
	o.SubscriptionProductId = &v
}

// GetSubscriptionProduct returns the SubscriptionProduct field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionProduct() SubscriptionProduct {
	if o == nil || isNil(o.SubscriptionProduct) {
		var ret SubscriptionProduct
		return ret
	}
	return *o.SubscriptionProduct
}

// GetSubscriptionProductOk returns a tuple with the SubscriptionProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionProductOk() (*SubscriptionProduct, bool) {
	if o == nil || isNil(o.SubscriptionProduct) {
    return nil, false
	}
	return o.SubscriptionProduct, true
}

// HasSubscriptionProduct returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionProduct() bool {
	if o != nil && !isNil(o.SubscriptionProduct) {
		return true
	}

	return false
}

// SetSubscriptionProduct gets a reference to the given SubscriptionProduct and assigns it to the SubscriptionProduct field.
func (o *Subscription) SetSubscriptionProduct(v SubscriptionProduct) {
	o.SubscriptionProduct = &v
}

// GetJournalId returns the JournalId field value if set, zero value otherwise.
func (o *Subscription) GetJournalId() int64 {
	if o == nil || isNil(o.JournalId) {
		var ret int64
		return ret
	}
	return *o.JournalId
}

// GetJournalIdOk returns a tuple with the JournalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetJournalIdOk() (*int64, bool) {
	if o == nil || isNil(o.JournalId) {
    return nil, false
	}
	return o.JournalId, true
}

// HasJournalId returns a boolean if a field has been set.
func (o *Subscription) HasJournalId() bool {
	if o != nil && !isNil(o.JournalId) {
		return true
	}

	return false
}

// SetJournalId gets a reference to the given int64 and assigns it to the JournalId field.
func (o *Subscription) SetJournalId(v int64) {
	o.JournalId = &v
}

// GetPaymentTransaction returns the PaymentTransaction field value if set, zero value otherwise.
func (o *Subscription) GetPaymentTransaction() PaymentTransaction {
	if o == nil || isNil(o.PaymentTransaction) {
		var ret PaymentTransaction
		return ret
	}
	return *o.PaymentTransaction
}

// GetPaymentTransactionOk returns a tuple with the PaymentTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPaymentTransactionOk() (*PaymentTransaction, bool) {
	if o == nil || isNil(o.PaymentTransaction) {
    return nil, false
	}
	return o.PaymentTransaction, true
}

// HasPaymentTransaction returns a boolean if a field has been set.
func (o *Subscription) HasPaymentTransaction() bool {
	if o != nil && !isNil(o.PaymentTransaction) {
		return true
	}

	return false
}

// SetPaymentTransaction gets a reference to the given PaymentTransaction and assigns it to the PaymentTransaction field.
func (o *Subscription) SetPaymentTransaction(v PaymentTransaction) {
	o.PaymentTransaction = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.URL) {
		toSerialize["URL"] = o.URL
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !isNil(o.ExpiryDate) {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if !isNil(o.StartDateDisplay) {
		toSerialize["startDateDisplay"] = o.StartDateDisplay
	}
	if !isNil(o.ExpiryDateDisplay) {
		toSerialize["expiryDateDisplay"] = o.ExpiryDateDisplay
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Renewal) {
		toSerialize["renewal"] = o.Renewal
	}
	if !isNil(o.SubscriptionProductId) {
		toSerialize["subscriptionProductId"] = o.SubscriptionProductId
	}
	if !isNil(o.SubscriptionProduct) {
		toSerialize["subscriptionProduct"] = o.SubscriptionProduct
	}
	if !isNil(o.JournalId) {
		toSerialize["journalId"] = o.JournalId
	}
	if !isNil(o.PaymentTransaction) {
		toSerialize["paymentTransaction"] = o.PaymentTransaction
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


