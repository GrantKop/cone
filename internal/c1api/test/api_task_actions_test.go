/*
ConductorOne API

Testing TaskActionsAPIService

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

func Test_c1api_TaskActionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskActionsAPIService C1ApiTaskV1TaskActionsServiceApprove", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceApprove(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskActionsAPIService C1ApiTaskV1TaskActionsServiceComment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceComment(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskActionsAPIService C1ApiTaskV1TaskActionsServiceDeny", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskId string

		resp, httpRes, err := apiClient.TaskActionsAPI.C1ApiTaskV1TaskActionsServiceDeny(context.Background(), taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
