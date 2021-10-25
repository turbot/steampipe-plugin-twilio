package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/chat/v2"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioChatServiceUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_chat_service_user",
		Description: "Users within specified Chat Service in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate:    listChatServiceUsers,
			KeyColumns: plugin.SingleColumn("service_sid"),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getChatServiceUser,
			KeyColumns: plugin.AllColumns([]string{"sid", "service_sid"}),
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
				Name:        "service_sid",
				Description: "The SID of the Service that the resource is associated with.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "attributes",
				Description: "The JSON string that stores application-specific data.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "identity",
				Description: "The string that identifies the resource's User.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_notifiable",
				Description: "Indicates whether the User has a potentially valid Push Notification registration for the Service instance, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "is_online",
				Description: "Indicates whether the User is actively connected to the Service instance and online, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "joined_channels_count",
				Description: "The number of Channels the User is a Member of.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "role_sid",
				Description: "The SID of the Role assigned to the user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the Service resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "links",
				Description: "A list of absolute URLs of the Channel and Binding resources related to the user.",
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

func listChatServiceUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_chat_service_user.listChatServiceUsers", "connection_error", err)
		return nil, err
	}

	// Get Chat service details
	chatServiceID := d.KeyColumnQuals["service_sid"].GetStringValue()

	// No inputs
	if chatServiceID == "" {
		return nil, nil
	}

	req := &openapi.ListUserParams{}

	// Retrieve the list of chat service users
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ChatV2.ListUser(chatServiceID, req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, user := range resp {
		d.StreamListItem(ctx, user)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getChatServiceUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getChatServiceUser")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_chat_service_user.getChatServiceUser", "connection_error", err)
		return nil, err
	}
	sid := d.KeyColumnQuals["sid"].GetStringValue()
	serviceSid := d.KeyColumnQuals["service_sid"].GetStringValue()

	// No inputs
	if sid == "" && serviceSid == "" {
		return nil, nil
	}

	resp, err := client.ChatV2.FetchUser(serviceSid, sid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
