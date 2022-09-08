package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/serverless/v1"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioServerlessServiceFunction(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_serverless_service_function",
		Description: "Serverless Functions in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate:       listServerlessServiceFunctions,
			ParentHydrate: listServerlessServices,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getServerlessServiceFunction,
			KeyColumns: plugin.AllColumns([]string{"sid", "service_sid"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "sid",
				Description: "The unique string that identifies the function resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "friendly_name",
				Description: "The string that you assigned to describe the function resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "service_sid",
				Description: "The SID of the Service that the Function resource is associated with.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "date_created",
				Description: "The date and time that the function resource was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "date_updated",
				Description: "The date and time that the function resource was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "url",
				Description: "The absolute URL of the function resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "links",
				Description: "A list of URLs of nested resources of the function resource.",
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
				Description: "The SID of the Account that created the function resource.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listServerlessServiceFunctions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_serverless_service_function.listServerlessServiceFunctions", "connection_error", err)
		return nil, err
	}

	// Get Serverless service details
	serviceID := h.Item.(openapi.ServerlessV1Service).Sid

	req := &openapi.ListFunctionParams{}

	// Retrieve the list of serverless functions
	maxResult := 50

	// Reduce the basic request limit down if the user has only requested a small number of rows
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if int(*limit) < maxResult {
			maxResult = int(*limit)
		}
	}
	req.SetLimit(maxResult)

	resp, err := client.ServerlessV1.ListFunction(*serviceID, req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, function := range resp {
		d.StreamListItem(ctx, function)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getServerlessServiceFunction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getServerlessServiceFunction")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_serverless_service_function.getServerlessServiceFunction", "connection_error", err)
		return nil, err
	}
	sid := d.KeyColumnQuals["sid"].GetStringValue()
	serviceSid := d.KeyColumnQuals["service_sid"].GetStringValue()

	// No inputs
	if sid == "" && serviceSid == "" {
		return nil, nil
	}

	resp, err := client.ServerlessV1.FetchFunction(serviceSid, sid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
