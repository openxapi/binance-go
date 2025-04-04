/*
Binance Margin Trading API

Testing TradeDataStreamAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package margin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/binance-go/margin"
)

func Test_margin_TradeDataStreamAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TradeDataStreamAPIService MarginCreateUserDataStreamIsolatedV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginCreateUserDataStreamIsolatedV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TradeDataStreamAPIService MarginCreateUserDataStreamV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginCreateUserDataStreamV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TradeDataStreamAPIService MarginDeleteUserDataStreamIsolatedV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginDeleteUserDataStreamIsolatedV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TradeDataStreamAPIService MarginDeleteUserDataStreamV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginDeleteUserDataStreamV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TradeDataStreamAPIService MarginUpdateUserDataStreamIsolatedV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginUpdateUserDataStreamIsolatedV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TradeDataStreamAPIService MarginUpdateUserDataStreamV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TradeDataStreamAPI.MarginUpdateUserDataStreamV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
