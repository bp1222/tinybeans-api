# PaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**AmountInCents** | Pointer to **int64** |  | [optional] 
**AmountDisplay** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**PaymentId** | Pointer to **string** |  | [optional] 
**ReceiptNumber** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PaymentGatewayType** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewPaymentTransaction

`func NewPaymentTransaction() *PaymentTransaction`

NewPaymentTransaction instantiates a new PaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTransactionWithDefaults

`func NewPaymentTransactionWithDefaults() *PaymentTransaction`

NewPaymentTransactionWithDefaults instantiates a new PaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *PaymentTransaction) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *PaymentTransaction) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *PaymentTransaction) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *PaymentTransaction) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetId

`func (o *PaymentTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTransaction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *PaymentTransaction) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PaymentTransaction) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PaymentTransaction) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PaymentTransaction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PaymentTransaction) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PaymentTransaction) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PaymentTransaction) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PaymentTransaction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetTransactionDate

`func (o *PaymentTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *PaymentTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *PaymentTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *PaymentTransaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetAmountInCents

`func (o *PaymentTransaction) GetAmountInCents() int64`

GetAmountInCents returns the AmountInCents field if non-nil, zero value otherwise.

### GetAmountInCentsOk

`func (o *PaymentTransaction) GetAmountInCentsOk() (*int64, bool)`

GetAmountInCentsOk returns a tuple with the AmountInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInCents

`func (o *PaymentTransaction) SetAmountInCents(v int64)`

SetAmountInCents sets AmountInCents field to given value.

### HasAmountInCents

`func (o *PaymentTransaction) HasAmountInCents() bool`

HasAmountInCents returns a boolean if a field has been set.

### GetAmountDisplay

`func (o *PaymentTransaction) GetAmountDisplay() string`

GetAmountDisplay returns the AmountDisplay field if non-nil, zero value otherwise.

### GetAmountDisplayOk

`func (o *PaymentTransaction) GetAmountDisplayOk() (*string, bool)`

GetAmountDisplayOk returns a tuple with the AmountDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDisplay

`func (o *PaymentTransaction) SetAmountDisplay(v string)`

SetAmountDisplay sets AmountDisplay field to given value.

### HasAmountDisplay

`func (o *PaymentTransaction) HasAmountDisplay() bool`

HasAmountDisplay returns a boolean if a field has been set.

### GetDiscount

`func (o *PaymentTransaction) GetDiscount() int64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PaymentTransaction) GetDiscountOk() (*int64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PaymentTransaction) SetDiscount(v int64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *PaymentTransaction) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetUserId

`func (o *PaymentTransaction) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentTransaction) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentTransaction) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PaymentTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProductId

`func (o *PaymentTransaction) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PaymentTransaction) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PaymentTransaction) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PaymentTransaction) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentTransaction) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentTransaction) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentTransaction) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentTransaction) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetReceiptNumber

`func (o *PaymentTransaction) GetReceiptNumber() string`

GetReceiptNumber returns the ReceiptNumber field if non-nil, zero value otherwise.

### GetReceiptNumberOk

`func (o *PaymentTransaction) GetReceiptNumberOk() (*string, bool)`

GetReceiptNumberOk returns a tuple with the ReceiptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNumber

`func (o *PaymentTransaction) SetReceiptNumber(v string)`

SetReceiptNumber sets ReceiptNumber field to given value.

### HasReceiptNumber

`func (o *PaymentTransaction) HasReceiptNumber() bool`

HasReceiptNumber returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPaymentGatewayType

`func (o *PaymentTransaction) GetPaymentGatewayType() string`

GetPaymentGatewayType returns the PaymentGatewayType field if non-nil, zero value otherwise.

### GetPaymentGatewayTypeOk

`func (o *PaymentTransaction) GetPaymentGatewayTypeOk() (*string, bool)`

GetPaymentGatewayTypeOk returns a tuple with the PaymentGatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGatewayType

`func (o *PaymentTransaction) SetPaymentGatewayType(v string)`

SetPaymentGatewayType sets PaymentGatewayType field to given value.

### HasPaymentGatewayType

`func (o *PaymentTransaction) HasPaymentGatewayType() bool`

HasPaymentGatewayType returns a boolean if a field has been set.

### GetLast4

`func (o *PaymentTransaction) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentTransaction) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentTransaction) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *PaymentTransaction) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetUser

`func (o *PaymentTransaction) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PaymentTransaction) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PaymentTransaction) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *PaymentTransaction) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


