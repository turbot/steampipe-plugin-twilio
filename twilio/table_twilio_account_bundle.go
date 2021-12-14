package twilio

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
	openapi "github.com/twilio/twilio-go/rest/numbers/v2"
)

func tableTwilioAccountBundle(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_bundle",
		Description: "Bundle represents your or your customer’s Regulatory bundle",
		List: &plugin.ListConfig{
			Hydrate: listAccountBundles,
		},
		Columns: []*plugin.Column{
			{
				Name:        "account_sid",
				Description: "The SID of the Account that created the Bundle resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sid",
				Description: "The unique string that we created to identify the Bundle resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "friendly_name",
				Description: "The string that you assigned to describe the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that the bundle was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateCreated"),
			},
			{
				Name:        "status",
				Description: "The verification status of the Bundle resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the bundle was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("DateUpdated"),
			},
			{
				Name:        "email",
				Description: "The email address that will receive updates when the Bundle resource changes status.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "regulation_sid",
				Description: "The unique string of a regulation that is associated to the Bundle resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "status_callback",
				Description: "The URL we call to inform your application of status changes.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the Bundle resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "valid_until",
				Description: "The date and time when the resource will be valid until.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

func listAccountBundles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getSessionConfig(ctx, d)

	if err != nil {
		plugin.Logger(ctx).Error("regulatory_bundles.listBundles", "connection_error", err)
		return nil, err
	}
	req := &openapi.ListBundleParams{}

	// Retrieve the list of addresses
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.NumbersV2.StreamBundle(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for bundles := range resp {
		d.StreamListItem(ctx, bundles)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
