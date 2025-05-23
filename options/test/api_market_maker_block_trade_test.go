/*
Binance Options API

Testing MarketMakerBlockTradeAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package options

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/binance-go/options"
)

func Test_options_MarketMakerBlockTradeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MarketMakerBlockTradeAPIService CreateBlockOrderCreateV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MarketMakerBlockTradeAPI.CreateBlockOrderCreateV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MarketMakerBlockTradeAPIService DeleteBlockOrderCreateV1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MarketMakerBlockTradeAPI.DeleteBlockOrderCreateV1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
