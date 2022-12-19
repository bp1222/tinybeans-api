# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **int64** |  | [optional] 
**ExpiryDate** | Pointer to **int64** |  | [optional] 
**StartDateDisplay** | Pointer to **string** |  | [optional] 
**ExpiryDateDisplay** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Renewal** | Pointer to **string** |  | [optional] 
**SubscriptionProductId** | Pointer to **int64** |  | [optional] 
**SubscriptionProduct** | Pointer to [**SubscriptionProduct**](SubscriptionProduct.md) |  | [optional] 
**JournalId** | Pointer to **int64** |  | [optional] 
**PaymentTransaction** | Pointer to [**PaymentTransaction**](PaymentTransaction.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *Subscription) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Subscription) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Subscription) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Subscription) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *Subscription) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Subscription) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Subscription) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Subscription) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Subscription) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Subscription) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Subscription) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Subscription) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetStartDateDisplay

`func (o *Subscription) GetStartDateDisplay() string`

GetStartDateDisplay returns the StartDateDisplay field if non-nil, zero value otherwise.

### GetStartDateDisplayOk

`func (o *Subscription) GetStartDateDisplayOk() (*string, bool)`

GetStartDateDisplayOk returns a tuple with the StartDateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateDisplay

`func (o *Subscription) SetStartDateDisplay(v string)`

SetStartDateDisplay sets StartDateDisplay field to given value.

### HasStartDateDisplay

`func (o *Subscription) HasStartDateDisplay() bool`

HasStartDateDisplay returns a boolean if a field has been set.

### GetExpiryDateDisplay

`func (o *Subscription) GetExpiryDateDisplay() string`

GetExpiryDateDisplay returns the ExpiryDateDisplay field if non-nil, zero value otherwise.

### GetExpiryDateDisplayOk

`func (o *Subscription) GetExpiryDateDisplayOk() (*string, bool)`

GetExpiryDateDisplayOk returns a tuple with the ExpiryDateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateDisplay

`func (o *Subscription) SetExpiryDateDisplay(v string)`

SetExpiryDateDisplay sets ExpiryDateDisplay field to given value.

### HasExpiryDateDisplay

`func (o *Subscription) HasExpiryDateDisplay() bool`

HasExpiryDateDisplay returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRenewal

`func (o *Subscription) GetRenewal() string`

GetRenewal returns the Renewal field if non-nil, zero value otherwise.

### GetRenewalOk

`func (o *Subscription) GetRenewalOk() (*string, bool)`

GetRenewalOk returns a tuple with the Renewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewal

`func (o *Subscription) SetRenewal(v string)`

SetRenewal sets Renewal field to given value.

### HasRenewal

`func (o *Subscription) HasRenewal() bool`

HasRenewal returns a boolean if a field has been set.

### GetSubscriptionProductId

`func (o *Subscription) GetSubscriptionProductId() int64`

GetSubscriptionProductId returns the SubscriptionProductId field if non-nil, zero value otherwise.

### GetSubscriptionProductIdOk

`func (o *Subscription) GetSubscriptionProductIdOk() (*int64, bool)`

GetSubscriptionProductIdOk returns a tuple with the SubscriptionProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionProductId

`func (o *Subscription) SetSubscriptionProductId(v int64)`

SetSubscriptionProductId sets SubscriptionProductId field to given value.

### HasSubscriptionProductId

`func (o *Subscription) HasSubscriptionProductId() bool`

HasSubscriptionProductId returns a boolean if a field has been set.

### GetSubscriptionProduct

`func (o *Subscription) GetSubscriptionProduct() SubscriptionProduct`

GetSubscriptionProduct returns the SubscriptionProduct field if non-nil, zero value otherwise.

### GetSubscriptionProductOk

`func (o *Subscription) GetSubscriptionProductOk() (*SubscriptionProduct, bool)`

GetSubscriptionProductOk returns a tuple with the SubscriptionProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionProduct

`func (o *Subscription) SetSubscriptionProduct(v SubscriptionProduct)`

SetSubscriptionProduct sets SubscriptionProduct field to given value.

### HasSubscriptionProduct

`func (o *Subscription) HasSubscriptionProduct() bool`

HasSubscriptionProduct returns a boolean if a field has been set.

### GetJournalId

`func (o *Subscription) GetJournalId() int64`

GetJournalId returns the JournalId field if non-nil, zero value otherwise.

### GetJournalIdOk

`func (o *Subscription) GetJournalIdOk() (*int64, bool)`

GetJournalIdOk returns a tuple with the JournalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalId

`func (o *Subscription) SetJournalId(v int64)`

SetJournalId sets JournalId field to given value.

### HasJournalId

`func (o *Subscription) HasJournalId() bool`

HasJournalId returns a boolean if a field has been set.

### GetPaymentTransaction

`func (o *Subscription) GetPaymentTransaction() PaymentTransaction`

GetPaymentTransaction returns the PaymentTransaction field if non-nil, zero value otherwise.

### GetPaymentTransactionOk

`func (o *Subscription) GetPaymentTransactionOk() (*PaymentTransaction, bool)`

GetPaymentTransactionOk returns a tuple with the PaymentTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTransaction

`func (o *Subscription) SetPaymentTransaction(v PaymentTransaction)`

SetPaymentTransaction sets PaymentTransaction field to given value.

### HasPaymentTransaction

`func (o *Subscription) HasPaymentTransaction() bool`

HasPaymentTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


