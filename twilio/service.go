package twilio

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/twilio/twilio-go"
)

// getSessionConfig :: returns Twilio client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*twilio.RestClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "twilio.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*twilio.RestClient), nil
	}

	var accountSid, authToken, apiKey, apiSecret string

	// Get twilio config
	twilioConfig := GetConfig(d.Connection)

	if twilioConfig.AccountSid != nil {
		accountSid = *twilioConfig.AccountSid
	}

	if accountSid == "" {
		accountSid = os.Getenv("TWILIO_ACCOUNT_SID")
	}

	if accountSid == "" {
		return nil, fmt.Errorf("account_sid must be configured")
	}

	if twilioConfig.AuthToken != nil {
		authToken = *twilioConfig.AuthToken
	}

	if twilioConfig.ApiKey != nil {
		apiKey = *twilioConfig.ApiKey
	}

	if twilioConfig.ApiSecret != nil {
		apiSecret = *twilioConfig.ApiSecret
	}

	// Check for environment variables
	if apiKey == "" {
		apiKey = os.Getenv("TWILIO_API_KEY")
	}

	if apiSecret == "" {
		apiSecret = os.Getenv("TWILIO_API_SECRET")
	}

	if authToken == "" {
		authToken = os.Getenv("TWILIO_AUTH_TOKEN")
	}

	
	clientReq := twilio.RestClientParams{}
	if authToken != "" {
		clientReq.Username = accountSid
		clientReq.Password = authToken
	}

	if apiKey != "" && apiSecret != "" {
		clientReq.Username = apiKey
		clientReq.Password = apiSecret
		clientReq.AccountSid = accountSid
	}

	// No creds
	if clientReq.Username == "" && clientReq.Password == "" {
		return nil, fmt.Errorf("either api_key and api_secret or auth_token must be configured")
	}

	client := twilio.NewRestClientWithParams(clientReq)

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
