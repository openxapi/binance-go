/*
Binance Sub Account API

Testing ApiManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package subaccount

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/binance-go/subaccount"
)

func Test_subaccount_ApiManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApiManagementAPIService SubaccountCreateSubAccountSubAccountApiIpRestrictionV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApiManagementAPI.SubaccountCreateSubAccountSubAccountApiIpRestrictionV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApiManagementAPIService SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApiManagementAPI.SubaccountDeleteSubAccountSubAccountApiIpRestrictionIpListV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApiManagementAPIService SubaccountGetSubAccountSubAccountApiIpRestrictionV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApiManagementAPI.SubaccountGetSubAccountSubAccountApiIpRestrictionV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
