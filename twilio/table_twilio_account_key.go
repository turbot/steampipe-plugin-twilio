package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_key",
		Description: "API keys in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listTwilioAccountKeys,
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
				Transform:   transform.FromField("DateCreated").Transform(ensureTimestamp),
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the resource was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateUpdated").Transform(ensureTimestamp),
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
				Hydrate:     plugin.HydrateFunc(getAccountSID).WithCache(),
				Transform:   transform.FromValue(),
			},
		},
	}
}

//// LIST FUNCTION

func listTwilioAccountKeys(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_key.listTwilioAccountKeys", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListKeyParams{}

	// Retrieve the list of API keys
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ApiV2010.ListKey(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, key := range resp {
		d.StreamListItem(ctx, key)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
