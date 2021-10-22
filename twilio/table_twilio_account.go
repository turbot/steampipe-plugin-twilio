package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account",
		Description: "The set of Accounts belonging to the Twilio Account.",
		List: &plugin.ListConfig{
			Hydrate: listTwilioAccounts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "friendly_name",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
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
				Description: "A human readable description of this account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "type",
				Description: "The type of this account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status",
				Description: "The status of this account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owner_account_sid",
				Description: "The unique string representing the parent of this account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uri",
				Description: "The URI for this resource, relative to 'https://api.twilio.com'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "auth_token",
				Description: "The authorization token for this account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that the account was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateCreated").Transform(ensureTimestamp),
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the account was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateUpdated").Transform(ensureTimestamp),
			},
			{
				Name:        "sub_resource_uris",
				Description: "A list of account instance sub-resources.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("SubresourceUris"),
			},
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("FriendlyName"),
			},
		},
	}
}

//// LIST FUNCTION

func listTwilioAccounts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account.listTwilioAccounts", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListAccountParams{}

	// Additional filters
	if d.KeyColumnQuals["friendly_name"] != nil {
		req.SetFriendlyName(d.KeyColumnQuals["friendly_name"].GetStringValue())
	}

	if d.KeyColumnQuals["status"] != nil {
		req.SetStatus(d.KeyColumnQuals["status"].GetStringValue())
	}

	// Retrieve the list of Accounts
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ApiV2010.ListAccount(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, acc := range resp {
		d.StreamListItem(ctx, acc)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
