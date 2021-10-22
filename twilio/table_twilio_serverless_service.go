package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/serverless/v1"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioServerlessService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_serverless_service",
		Description: "Serverless Services in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listServerlessServices,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServerlessService,
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
				Name:        "account_sid",
				Description: "The SID of the Account that created the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_base",
				Description: "The base domain name for this Service, which is a combination of the unique name and a randomly generated string.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "include_credentials",
				Description: "Indicates whether to inject Account credentials into a function invocation context, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "ui_editable",
				Description: "Indicates whether the Service resource's properties and subresources can be edited via the UI, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "unique_name",
				Description: "An user-defined string that uniquely identifies the Service resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the Service resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "links",
				Description: "A list of URLs of the Service's nested resources.",
				Type:        proto.ColumnType_JSON,
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

func listServerlessServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_serverless_service.listServerlessServices", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListServiceParams{}

	// Retrieve the list of serverless services
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ServerlessV1.ListService(req)
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

func getServerlessService(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getServerlessService")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_serverless_service.getServerlessService", "connection_error", err)
		return nil, err
	}
	sid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if sid == "" {
		return nil, nil
	}

	resp, err := client.ServerlessV1.FetchService(sid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
