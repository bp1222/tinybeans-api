# \JournalsApi

All URIs are relative to *https://tinybeans.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JournalEntries**](JournalsApi.md#JournalEntries) | **Get** /journals/{journal}/entries | Tinybeans Journal Entries
[**Journals**](JournalsApi.md#Journals) | **Get** /journals | Tinybeans Journals



## JournalEntries

> Entries JournalEntries(ctx, journal).FetchSize(fetchSize).IdsOnly(idsOnly).Last(last).Since(since).Execute()

Tinybeans Journal Entries

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    journal := int64(789) // int64 | ID of journal to pull entries from
    fetchSize := int64(789) // int64 | How many to fetch
    idsOnly := int64(789) // int64 | ID's Only? (optional)
    last := int64(789) // int64 | Last (timestamp) you viewed [non inclusive] (optional)
    since := int64(789) // int64 | Since (timestamp) most recent (timestamp) you know about [inclusive] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JournalsApi.JournalEntries(context.Background(), journal).FetchSize(fetchSize).IdsOnly(idsOnly).Last(last).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.JournalEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JournalEntries`: Entries
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.JournalEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journal** | **int64** | ID of journal to pull entries from | 

### Other Parameters

Other parameters are passed through a pointer to a apiJournalEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchSize** | **int64** | How many to fetch | 
 **idsOnly** | **int64** | ID&#39;s Only? | 
 **last** | **int64** | Last (timestamp) you viewed [non inclusive] | 
 **since** | **int64** | Since (timestamp) most recent (timestamp) you know about [inclusive] | 

### Return type

[**Entries**](Entries.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Journals

> Journals Journals(ctx).Execute()

Tinybeans Journals

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JournalsApi.Journals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JournalsApi.Journals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Journals`: Journals
    fmt.Fprintf(os.Stdout, "Response from `JournalsApi.Journals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiJournalsRequest struct via the builder pattern


### Return type

[**Journals**](Journals.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

