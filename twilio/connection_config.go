package twilio

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type twilioConfig struct {
	Username   *string `cty:"username"`
	Password   *string `cty:"password"`
	AccountSid *string `cty:"account_sid"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"account_sid": {
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
