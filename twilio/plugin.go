package twilio

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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
			"twilio_account_address":               tableTwilioAccountAddress(ctx),
			"twilio_account_application":           tableTwilioAccountApplication(ctx),
			"twilio_account_bundle":                tableTwilioAccountBundle(ctx),
			"twilio_account_call":                  tableTwilioAccountCall(ctx),
			"twilio_account_incoming_phone_number": tableTwilioAccountIncomingPhoneNumber(ctx),
			"twilio_account_key":                   tableTwilioAccountKey(ctx),
			"twilio_account_message":               tableTwilioAccountMessage(ctx),
			"twilio_chat_service":                  tableTwilioChatService(ctx),
			"twilio_chat_service_user":             tableTwilioChatServiceUser(ctx),
			"twilio_messaging_service":             tableTwilioMessagingService(ctx),
			"twilio_serverless_service":            tableTwilioServerlessService(ctx),
			"twilio_serverless_service_function":   tableTwilioServerlessServiceFunction(ctx),
		},
	}

	return p
}
