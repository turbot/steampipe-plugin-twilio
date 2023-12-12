package twilio

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type twilioConfig struct {
	AccountSid *string `hcl:"account_sid"`
	AuthToken  *string `hcl:"auth_token"`
	ApiKey     *string `hcl:"api_key"`
	ApiSecret  *string `hcl:"api_secret"`
}

func ConfigInstance() interface{} {
	return &twilioConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) twilioConfig {
	if connection == nil || connection.Config == nil {
		return twilioConfig{}
	}
	config, _ := connection.Config.(twilioConfig)
	return config
}
