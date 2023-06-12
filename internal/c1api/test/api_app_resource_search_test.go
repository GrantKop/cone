/*
ConductorOne API

Testing AppResourceSearchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package c1api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/conductorone/cone/internal/c1api"
)

func Test_c1api_AppResourceSearchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppResourceSearchAPIService C1ApiAppV1AppResourceSearchSearchAppResourceTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppResourceSearchAPI.C1ApiAppV1AppResourceSearchSearchAppResourceTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
