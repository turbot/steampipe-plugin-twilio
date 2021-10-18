package twilio

import (
	"context"
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/twilio/twilio-go"
)

// getSessionConfig :: returns Twilio client to perform API requests
func getSessionConfig(ctx context.Context, d *plugin.QueryData) (*twilio.RestClient, error) {
	// Load clientOptions from cache
	sessionCacheKey := "twilio.clientoption"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*twilio.RestClient), nil
	}

	var username, password, accountSid string

	// Get twilio config
	twilioConfig := GetConfig(d.Connection)

	if twilioConfig.Username != nil {
		username = *twilioConfig.Username
	}

	if twilioConfig.Password != nil {
		password = *twilioConfig.Password
	}

	if twilioConfig.AccountSid != nil {
		accountSid = *twilioConfig.AccountSid
	}

	// Check for API key and Secret
	if username == "" && password == "" {
		username = os.Getenv("TWILIO_API_KEY")
		password = os.Getenv("TWILIO_API_SECRET")
	}

	// Check for Account SID and Auth token
	if username == "" && password == "" {
		username = os.Getenv("TWILIO_ACCOUNT_SID")
		password = os.Getenv("TWILIO_AUTH_TOKEN")
	}

	// No creds
	if username == "" && password == "" {
		return nil, fmt.Errorf("both username and password must be configured")
	}

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username:   username,
		Password:   password,
		AccountSid: accountSid,
	})

	// save clientOptions in cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, client)

	return client, nil
}
