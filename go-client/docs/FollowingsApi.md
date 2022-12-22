# \FollowingsApi

All URIs are relative to *https://tinybeans.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Followings**](FollowingsApi.md#Followings) | **Get** /followings | Tinybeans Followings



## Followings

> Followings Followings(ctx).Execute()

Tinybeans Followings

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
    resp, r, err := apiClient.FollowingsApi.Followings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowingsApi.Followings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Followings`: Followings
    fmt.Fprintf(os.Stdout, "Response from `FollowingsApi.Followings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFollowingsRequest struct via the builder pattern


### Return type

[**Followings**](Followings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

