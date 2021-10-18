package twilio

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-twilio"

// Plugin creates this (twilio) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"twilio_account":                       tableTwilioAccount(ctx),
			"twilio_account_incoming_phone_number": tableTwilioAccountIncomingPhoneNumber(ctx),
			"twilio_account_key":                   tableTwilioAccountKey(ctx),
			"twilio_account_message":               tableTwilioAccountMessage(ctx),
		},
	}

	return p
}
