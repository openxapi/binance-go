/*
Binance Umfutures API

Testing UserDataStreamsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package umfutures

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/binance-go/derivatives/umfutures"
)

func Test_umfutures_UserDataStreamsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserDataStreamsAPIService UmfuturesCreateListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.UmfuturesCreateListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserDataStreamsAPIService UmfuturesDeleteListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.UmfuturesDeleteListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserDataStreamsAPIService UmfuturesUpdateListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.UmfuturesUpdateListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
