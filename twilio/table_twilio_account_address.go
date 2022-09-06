package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountAddress(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_address",
		Description: "Address represents your or your customer’s physical location within a country.",
		List: &plugin.ListConfig{
			Hydrate: listAccountAddresses,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "friendly_name",
					Require: plugin.Optional,
				},
				{
					Name:    "customer_name",
					Require: plugin.Optional,
				},
				{
					Name:    "iso_country",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountAddress,
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
				Name:        "customer_name",
				Description: "The name associated with the address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "verified",
				Description: "Indicates whether the address has been verified to comply with regulation.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "validated",
				Description: "Indicates whether the address has been validated to comply with local regulation.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "city",
				Description: "The city in which the address is located.",
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
			{
				Name:        "emergency_enabled",
				Description: "Indicates whether emergency calling has been enabled on this number.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "iso_country",
				Description: "The ISO country code of the address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "postal_code",
				Description: "The postal code of the address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "region",
				Description: "The state or region of the address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "street",
				Description: "The number and street address of the address.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uri",
				Description: "The URI of the resource, relative to 'https://api.twilio.com'.",
				Type:        proto.ColumnType_STRING,
			},

			// Twilio standard columns
			{
				Name:        "account_sid",
				Description: "The SID of the Account that is responsible for the resource.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listAccountAddresses(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_address.listAccountAddresses", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListAddressParams{}

	// Additional filters
	if d.KeyColumnQuals["friendly_name"] != nil {
		req.SetFriendlyName(d.KeyColumnQuals["friendly_name"].GetStringValue())
	}

	if d.KeyColumnQuals["customer_name"] != nil {
		req.SetCustomerName(d.KeyColumnQuals["customer_name"].GetStringValue())
	}

	if d.KeyColumnQuals["iso_country"] != nil {
		req.SetIsoCountry(d.KeyColumnQuals["iso_country"].GetStringValue())
	}

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

	resp, err := client.ApiV2010.ListAddress(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, address := range resp {
		d.StreamListItem(ctx, address)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountAddress(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccountAddress")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_address.getAccountAddress", "connection_error", err)
		return nil, err
	}
	addressSid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if addressSid == "" {
		return nil, nil
	}

	resp, err := client.ApiV2010.FetchAddress(addressSid, &openapi.FetchAddressParams{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
