/*
Tinybeans API - Unofficial

Testing JournalsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tinybeans

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_tinybeans_JournalsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test JournalsApiService JournalEntries", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var journal int64

        resp, httpRes, err := apiClient.JournalsApi.JournalEntries(context.Background(), journal).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test JournalsApiService Journals", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.JournalsApi.Journals(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
