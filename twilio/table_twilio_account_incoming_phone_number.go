package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountIncomingPhoneNumber(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_incoming_phone_number",
		Description: "Incoming phone numbers in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountIncomingPhoneNumbers,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "beta",
					Require:   plugin.Optional,
					Operators: []string{"<>", "="},
				},
				{
					Name:    "friendly_name",
					Require: plugin.Optional,
				},
				{
					Name:    "origin",
					Require: plugin.Optional,
				},
				{
					Name:    "phone_number",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "sid",
				Description: "The unique string that identifies the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "friendly_name",
				Description: "The string that you assigned to describe the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "phone_number",
				Description: "The phone number in E.164 format.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of this resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "origin",
				Description: "The phone number's origin. Can be twilio or hosted.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "address_requirements",
				Description: "Indicates whether the phone number requires an address registered with Twilio, or not.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "address_sid",
				Description: "The SID of the Address resource associated with the phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_version",
				Description: "The API version used to start a new TwiML session.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "beta",
				Description: "Indicates whether the phone number is new to the Twilio platform, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "bundle_sid",
				Description: "The SID of the Bundle resource associated with number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that the resource was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateCreated").Transform(ensureTimestamp),
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the resource was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateUpdated").Transform(ensureTimestamp),
			},
			{
				Name:        "emergency_address_sid",
				Description: "The emergency address configuration to use for emergency calling.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "emergency_address_status",
				Description: "State of the emergency address configuration for the phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "emergency_status",
				Description: "Displays if emergency calling is enabled for this number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "identity_sid",
				Description: "The SID of the Identity resource associated with number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sms_application_sid",
				Description: "The SID of the application that handles SMS messages sent to the phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sms_fallback_method",
				Description: "The HTTP method used with sms_fallback_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sms_fallback_url",
				Description: "The URL that we call when an error occurs while retrieving or executing the TwiML.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sms_method",
				Description: "The HTTP method to use with sms_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sms_url",
				Description: "The URL we call when the phone number receives an incoming SMS message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_callback",
				Description: "The URL to send status information to your application.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_callback_method",
				Description: "The HTTP method we use to call status_callback.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "trunk_sid",
				Description: "The SID of the Trunk that handles calls to the phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uri",
				Description: "The URI of the resource, relative to 'https://api.twilio.com'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_application_sid",
				Description: "The SID of the application that handles calls to the phone number.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_caller_id_lookup",
				Description: "Indicates whether to lookup the caller's name, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "voice_fallback_method",
				Description: "The HTTP method used with voice_fallback_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_fallback_url",
				Description: "The URL we call when an error occurs in TwiML.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_method",
				Description: "The HTTP method used with the voice_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_receive_mode",
				Description: "The mode of the voice receive.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_url",
				Description: "The URL we call when the phone number receives a call.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "capabilities",
				Description: "The set of Boolean properties that indicate whether a phone number can receive calls or messages.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FriendlyName"),
			},

			// Twilio standard columns
			{
				Name:        "account_sid",
				Description: "The SID of the Account that created the resource.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listAccountIncomingPhoneNumbers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_incoming_phone_number.listAccountIncomingPhoneNumbers", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListIncomingPhoneNumberParams{}

	// Additional filters
	if d.EqualsQuals["friendly_name"] != nil {
		req.SetFriendlyName(d.EqualsQuals["friendly_name"].GetStringValue())
	}

	if d.EqualsQuals["origin"] != nil {
		req.SetOrigin(d.EqualsQuals["origin"].GetStringValue())
	}

	if d.EqualsQuals["phone_number"] != nil {
		req.SetPhoneNumber(d.EqualsQuals["phone_number"].GetStringValue())
	}

	if d.EqualsQuals["beta"] != nil {
		req.SetBeta(d.EqualsQuals["beta"].GetBoolValue())
	}

	// Non-Equals Qual Map handling
	if d.Quals["beta"] != nil {
		for _, q := range d.Quals["beta"].Quals {
			value := q.Value.GetBoolValue()
			if q.Operator == "<>" {
				req.SetBeta(false)
				if !value {
					req.SetBeta(true)
				}
			}
		}
	}

	// Retrieve the list of incoming phone numbers
	// Twilio SDK defaults to 1000 as an efficient page size:
	// https://github.com/twilio/twilio-go/blob/bf58569e99f043b8d1453a7d3812b5952bdda329/client/page_util.go#L17-L18
	pageSize := 1000

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < pageSize {
			pageSize = int(*limit)
		}
	}
	req.SetPageSize(pageSize)

	// Twilio SDK handles paging internally
	resp, err := client.Api.ListIncomingPhoneNumber(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, ph := range resp {
		d.StreamListItem(ctx, ph)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
