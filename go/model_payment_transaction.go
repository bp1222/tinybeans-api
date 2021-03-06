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

// PaymentTransaction struct for PaymentTransaction
type PaymentTransaction struct {
	URL *string `json:"URL,omitempty"`
	Id *int64 `json:"id,omitempty"`
	DateCreated *string `json:"dateCreated,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	TransactionDate *string `json:"transactionDate,omitempty"`
	AmountInCents *int64 `json:"amountInCents,omitempty"`
	AmountDisplay *string `json:"amountDisplay,omitempty"`
	Discount *int64 `json:"discount,omitempty"`
	UserId *int64 `json:"userId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	ReceiptNumber *string `json:"receiptNumber,omitempty"`
	Description *string `json:"description,omitempty"`
	PaymentGatewayType *string `json:"paymentGatewayType,omitempty"`
	Last4 *string `json:"last4,omitempty"`
	User *User `json:"user,omitempty"`
}

// NewPaymentTransaction instantiates a new PaymentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTransaction() *PaymentTransaction {
	this := PaymentTransaction{}
	return &this
}

// NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTransactionWithDefaults() *PaymentTransaction {
	this := PaymentTransaction{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *PaymentTransaction) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *PaymentTransaction) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *PaymentTransaction) SetURL(v string) {
	o.URL = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentTransaction) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentTransaction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PaymentTransaction) SetId(v int64) {
	o.Id = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *PaymentTransaction) GetDateCreated() string {
	if o == nil || o.DateCreated == nil {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetDateCreatedOk() (*string, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *PaymentTransaction) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *PaymentTransaction) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PaymentTransaction) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PaymentTransaction) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *PaymentTransaction) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *PaymentTransaction) GetTransactionDate() string {
	if o == nil || o.TransactionDate == nil {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetTransactionDateOk() (*string, bool) {
	if o == nil || o.TransactionDate == nil {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *PaymentTransaction) HasTransactionDate() bool {
	if o != nil && o.TransactionDate != nil {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *PaymentTransaction) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetAmountInCents returns the AmountInCents field value if set, zero value otherwise.
func (o *PaymentTransaction) GetAmountInCents() int64 {
	if o == nil || o.AmountInCents == nil {
		var ret int64
		return ret
	}
	return *o.AmountInCents
}

// GetAmountInCentsOk returns a tuple with the AmountInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetAmountInCentsOk() (*int64, bool) {
	if o == nil || o.AmountInCents == nil {
		return nil, false
	}
	return o.AmountInCents, true
}

// HasAmountInCents returns a boolean if a field has been set.
func (o *PaymentTransaction) HasAmountInCents() bool {
	if o != nil && o.AmountInCents != nil {
		return true
	}

	return false
}

// SetAmountInCents gets a reference to the given int64 and assigns it to the AmountInCents field.
func (o *PaymentTransaction) SetAmountInCents(v int64) {
	o.AmountInCents = &v
}

// GetAmountDisplay returns the AmountDisplay field value if set, zero value otherwise.
func (o *PaymentTransaction) GetAmountDisplay() string {
	if o == nil || o.AmountDisplay == nil {
		var ret string
		return ret
	}
	return *o.AmountDisplay
}

// GetAmountDisplayOk returns a tuple with the AmountDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetAmountDisplayOk() (*string, bool) {
	if o == nil || o.AmountDisplay == nil {
		return nil, false
	}
	return o.AmountDisplay, true
}

// HasAmountDisplay returns a boolean if a field has been set.
func (o *PaymentTransaction) HasAmountDisplay() bool {
	if o != nil && o.AmountDisplay != nil {
		return true
	}

	return false
}

// SetAmountDisplay gets a reference to the given string and assigns it to the AmountDisplay field.
func (o *PaymentTransaction) SetAmountDisplay(v string) {
	o.AmountDisplay = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *PaymentTransaction) GetDiscount() int64 {
	if o == nil || o.Discount == nil {
		var ret int64
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetDiscountOk() (*int64, bool) {
	if o == nil || o.Discount == nil {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *PaymentTransaction) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given int64 and assigns it to the Discount field.
func (o *PaymentTransaction) SetDiscount(v int64) {
	o.Discount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PaymentTransaction) GetUserId() int64 {
	if o == nil || o.UserId == nil {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetUserIdOk() (*int64, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PaymentTransaction) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *PaymentTransaction) SetUserId(v int64) {
	o.UserId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *PaymentTransaction) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *PaymentTransaction) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *PaymentTransaction) SetProductId(v string) {
	o.ProductId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *PaymentTransaction) GetPaymentId() string {
	if o == nil || o.PaymentId == nil {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetPaymentIdOk() (*string, bool) {
	if o == nil || o.PaymentId == nil {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PaymentTransaction) HasPaymentId() bool {
	if o != nil && o.PaymentId != nil {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *PaymentTransaction) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetReceiptNumber returns the ReceiptNumber field value if set, zero value otherwise.
func (o *PaymentTransaction) GetReceiptNumber() string {
	if o == nil || o.ReceiptNumber == nil {
		var ret string
		return ret
	}
	return *o.ReceiptNumber
}

// GetReceiptNumberOk returns a tuple with the ReceiptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetReceiptNumberOk() (*string, bool) {
	if o == nil || o.ReceiptNumber == nil {
		return nil, false
	}
	return o.ReceiptNumber, true
}

// HasReceiptNumber returns a boolean if a field has been set.
func (o *PaymentTransaction) HasReceiptNumber() bool {
	if o != nil && o.ReceiptNumber != nil {
		return true
	}

	return false
}

// SetReceiptNumber gets a reference to the given string and assigns it to the ReceiptNumber field.
func (o *PaymentTransaction) SetReceiptNumber(v string) {
	o.ReceiptNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentTransaction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentTransaction) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentTransaction) SetDescription(v string) {
	o.Description = &v
}

// GetPaymentGatewayType returns the PaymentGatewayType field value if set, zero value otherwise.
func (o *PaymentTransaction) GetPaymentGatewayType() string {
	if o == nil || o.PaymentGatewayType == nil {
		var ret string
		return ret
	}
	return *o.PaymentGatewayType
}

// GetPaymentGatewayTypeOk returns a tuple with the PaymentGatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetPaymentGatewayTypeOk() (*string, bool) {
	if o == nil || o.PaymentGatewayType == nil {
		return nil, false
	}
	return o.PaymentGatewayType, true
}

// HasPaymentGatewayType returns a boolean if a field has been set.
func (o *PaymentTransaction) HasPaymentGatewayType() bool {
	if o != nil && o.PaymentGatewayType != nil {
		return true
	}

	return false
}

// SetPaymentGatewayType gets a reference to the given string and assigns it to the PaymentGatewayType field.
func (o *PaymentTransaction) SetPaymentGatewayType(v string) {
	o.PaymentGatewayType = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *PaymentTransaction) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *PaymentTransaction) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *PaymentTransaction) SetLast4(v string) {
	o.Last4 = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PaymentTransaction) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransaction) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PaymentTransaction) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *PaymentTransaction) SetUser(v User) {
	o.User = &v
}

func (o PaymentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.URL != nil {
		toSerialize["URL"] = o.URL
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.TransactionDate != nil {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if o.AmountInCents != nil {
		toSerialize["amountInCents"] = o.AmountInCents
	}
	if o.AmountDisplay != nil {
		toSerialize["amountDisplay"] = o.AmountDisplay
	}
	if o.Discount != nil {
		toSerialize["discount"] = o.Discount
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.ProductId != nil {
		toSerialize["productId"] = o.ProductId
	}
	if o.PaymentId != nil {
		toSerialize["paymentId"] = o.PaymentId
	}
	if o.ReceiptNumber != nil {
		toSerialize["receiptNumber"] = o.ReceiptNumber
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PaymentGatewayType != nil {
		toSerialize["paymentGatewayType"] = o.PaymentGatewayType
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentTransaction struct {
	value *PaymentTransaction
	isSet bool
}

func (v NullablePaymentTransaction) Get() *PaymentTransaction {
	return v.value
}

func (v *NullablePaymentTransaction) Set(val *PaymentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTransaction(val *PaymentTransaction) *NullablePaymentTransaction {
	return &NullablePaymentTransaction{value: val, isSet: true}
}

func (v NullablePaymentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


