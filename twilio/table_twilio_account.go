package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account",
		Description: "Retrieve the current Twilio Account.",
		List: &plugin.ListConfig{
			Hydrate:           listTwilioAccount,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
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
				Name:        "is_sub_account",
				Description: "Indicates whether this account is a sub-account, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.From(checkTwilioSubAccount),
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

			// Steampipe standard columns
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

func listTwilioAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account.listTwilioAccount", "connection_error", err)
		return nil, err
	}

	resp, err := client.ApiV2010.FetchAccount(client.Client.AccountSid())
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, resp)

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func checkTwilioSubAccount(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*openapi.ApiV2010Account)

	return *data.OwnerAccountSid != *data.Sid, nil
}
