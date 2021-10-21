package twilio

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type twilioConfig struct {
	AccountSid *string `cty:"account_sid"`
	AuthToken  *string `cty:"auth_token"`
	ApiKey     *string `cty:"api_key"`
	ApiSecret  *string `cty:"api_secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"account_sid": {
		Type: schema.TypeString,
	},
	"auth_token": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
	"api_secret": {
		Type: schema.TypeString,
	},
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
