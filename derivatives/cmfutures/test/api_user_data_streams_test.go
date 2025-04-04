/*
Binance COIN-M Futures API

Testing UserDataStreamsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cmfutures

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/binance-go/derivatives/cmfutures"
)

func Test_cmfutures_UserDataStreamsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserDataStreamsAPIService CmfuturesCreateListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.CmfuturesCreateListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserDataStreamsAPIService CmfuturesDeleteListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.CmfuturesDeleteListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserDataStreamsAPIService CmfuturesUpdateListenKeyV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserDataStreamsAPI.CmfuturesUpdateListenKeyV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
