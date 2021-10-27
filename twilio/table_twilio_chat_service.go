package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioChatService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_chat_service",
		Description: "Chat Services in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listChatServices,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getChatService,
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
				Name:        "default_channel_creator_role_sid",
				Description: "The channel role assigned to a channel creator when they join a new channel.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default_channel_role_sid",
				Description: "The channel role assigned to users when they are added to a channel.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "default_service_role_sid",
				Description: "The service role assigned to users when they are added to the service.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "post_webhook_retry_count",
				Description: "The number of times calls to the 'post_webhook_url' will be retried.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "post_webhook_url",
				Description: "The URL for post-event webhooks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "pre_webhook_retry_count",
				Description: "Count of times webhook will be retried in case of timeout or 429/503/504 HTTP responses.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "pre_webhook_url",
				Description: "The webhook URL for pre-event webhooks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "reachability_enabled",
				Description: "Indicates whether the Reachability Indicator feature is enabled for this Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "read_status_enabled",
				Description: "Indicates whether the Message Consumption Horizon feature is enabled, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "typing_indicator_timeout",
				Description: "How long in seconds to wait before assuming the user is no longer typing.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the Service resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "webhook_method",
				Description: "The HTTP method  to use for both PRE and POST webhooks.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "limits",
				Description: "An object that describes the limits of the service instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "links",
				Description: "A list of absolute URLs of the Service's Channels, Roles, and Users.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "media",
				Description: "The properties of the media that the service supports.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "notifications",
				Description: "The notification configuration for the Service instance.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "webhook_filters",
				Description: "The list of webhook events that are enabled for this Service instance.",
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

func listChatServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_chat_service.listChatServices", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListServiceParams{}

	// Retrieve the list of chat services
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ChatV2.ListService(req)
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

func getChatService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getChatService")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_chat_service.getChatService", "connection_error", err)
		return nil, err
	}
	sid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if sid == "" {
		return nil, nil
	}

	resp, err := client.ChatV2.FetchService(sid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
