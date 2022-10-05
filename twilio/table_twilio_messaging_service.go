package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/messaging/v1"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioMessagingService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_messaging_service",
		Description: "Messageing Services in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listMessagingServices,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getMessagingService,
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
				Name:        "date_created",
				Description: "The date and time that the resource was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the resource was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "area_code_geo_match",
				Description: "Indicates whether to enable Area Code Geomatch on the Service Instance, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("AreaCodeGeomatch"),
			},
			{
				Name:        "fallback_method",
				Description: "The HTTP method we use to call fallback_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "fallback_to_long_code",
				Description: "Indicates whether to enable Fallback to Long Code for messages sent through the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "fallback_url",
				Description: "The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "inbound_method",
				Description: "The HTTP method we use to call inbound_request_url.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "inbound_request_url",
				Description: "The URL we call using inbound_method when a message is received by any phone number or short code in the Service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mms_converter",
				Description: "Indicates whether to enable the MMS Converter for messages sent through the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "scan_message_content",
				Description: "The type of the scan message content.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "smart_encoding",
				Description: "Indicates whether to enable Encoding for messages sent through the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "status_callback",
				Description: "The URL we call to pass status updates about message delivery.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sticky_sender",
				Description: "Indicates whether to enable Sticky Sender on the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "synchronous_validation",
				Description: "Indicates whether to enable Synchronous Validation on the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the Service resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "use_inbound_webhook_on_number",
				Description: "If enabled, the webhook url configured on the phone number will be used and will override the `inbound_request_url`/`fallback_url` url called when an inbound message is received.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "validity_period",
				Description: "How long, in seconds, messages sent from the Service are valid.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "links",
				Description: "A list of absolute URLs of related resources.",
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

func listMessagingServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_messaging_service.listMessagingServices", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListServiceParams{}

	// Retrieve the list of messaging services
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
	resp, err := client.MessagingV1.ListService(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, service := range resp {
		d.StreamListItem(ctx, service)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getMessagingService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getMessagingService")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_messaging_service.getMessagingService", "connection_error", err)
		return nil, err
	}
	sid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if sid == "" {
		return nil, nil
	}

	resp, err := client.MessagingV1.FetchService(sid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
