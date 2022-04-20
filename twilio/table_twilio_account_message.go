package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_message",
		Description: "Messages in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountMessages,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "date_sent",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:    "sent_from",
					Require: plugin.Optional,
				},
				{
					Name:    "sent_to",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountMessage,
			KeyColumns: plugin.SingleColumn("sid"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "sid",
				Description: "The unique string that identifies the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sent_to",
				Description: "The phone number that received the message.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("To"),
			},
			{
				Name:        "sent_from",
				Description: "The phone number that initiated the message.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("From"),
			},
			{
				Name:        "status",
				Description: "The status of the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_version",
				Description: "The API version used to process the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "body",
				Description: "The message text.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that the resource was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateCreated").Transform(ensureTimestamp),
			},
			{
				Name:        "date_sent",
				Description: "The date and time when the message was sent.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateSent").Transform(ensureTimestamp),
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the resource was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateUpdated").Transform(ensureTimestamp),
			},
			{
				Name:        "direction",
				Description: "The direction of the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "error_code",
				Description: "The error code associated with the message.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "error_message",
				Description: "The description of the error_code.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "messaging_service_sid",
				Description: "The SID of the Messaging Service used with the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "num_media",
				Description: "The number of media files associated with the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "num_segments",
				Description: "The number of messages used to deliver the message body.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "price",
				Description: "The amount billed for the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "price_unit",
				Description: "The currency in which price is measured.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uri",
				Description: "The URI of the resource, relative to 'https://api.twilio.com'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sub_resource_uris",
				Description: "A list of related resources identified by their relative URIs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SubresourceUris"),
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

func listAccountMessages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_message.listAccountMessages", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListMessageParams{}

	// Additional filters
	if d.KeyColumnQuals["sent_from"] != nil {
		req.SetFrom(d.KeyColumnQuals["sent_from"].GetStringValue())
	}

	if d.KeyColumnQuals["sent_to"] != nil {
		req.SetTo(d.KeyColumnQuals["sent_to"].GetStringValue())
	}

	if d.Quals["date_sent"] != nil {
		for _, q := range d.Quals["date_sent"].Quals {
			switch q.Operator {
			case "<", "<=":
				req.SetDateSentBefore(q.Value.GetTimestampValue().AsTime())
			case "=":
				req.SetDateSent(q.Value.GetTimestampValue().AsTime())
			case ">", ">=":
				req.SetDateSentAfter(q.Value.GetTimestampValue().AsTime())
			}
		}
	}

	// Retrieve the list of messages
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ApiV2010.ListMessage(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, message := range resp {
		d.StreamListItem(ctx, message)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountMessage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccountMessage")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_message.getAccountMessage", "connection_error", err)
		return nil, err
	}
	messageSid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if messageSid == "" {
		return nil, nil
	}

	resp, err := client.ApiV2010.FetchMessage(messageSid, &openapi.FetchMessageParams{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
