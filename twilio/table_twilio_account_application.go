package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountApplication(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_application",
		Description: "Application that you have created with Twilio.",
		List: &plugin.ListConfig{
			Hydrate: listAccountApplications,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "friendly_name",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountApplication,
			KeyColumns: plugin.SingleColumn("sid"),
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
				Name:        "sms_url",
				Description: "The URL we call when the phone number receives an incoming SMS message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_url",
				Description: "The URL we call when the phone number receives a call.",
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
				Name:        "api_version",
				Description: "The API version used to start a new TwiML session.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "message_status_callback",
				Description: "The URL to send message status information to your application.",
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
				Name:        "sms_status_callback",
				Description: "The URL to SMS send status information to your application.",
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
				Name:        "uri",
				Description: "The URI of the resource, relative to 'https://api.twilio.com'.",
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
				Description: "The URL we call when a TwiML error occurs.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "voice_method",
				Description: "The HTTP method used with the voice_url.",
				Type:        proto.ColumnType_STRING,
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

func listAccountApplications(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_application.listAccountApplications", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListApplicationParams{}

	// Additional filters
	if d.KeyColumnQuals["friendly_name"] != nil {
		req.SetFriendlyName(d.KeyColumnQuals["friendly_name"].GetStringValue())
	}

	// Retrieve the list of applications
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ApiV2010.ListApplication(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, application := range resp {
		d.StreamListItem(ctx, application)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountApplication(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccountApplication")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_application.getAccountApplication", "connection_error", err)
		return nil, err
	}
	applicationSid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if applicationSid == "" {
		return nil, nil
	}

	resp, err := client.ApiV2010.FetchApplication(applicationSid, &openapi.FetchApplicationParams{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
