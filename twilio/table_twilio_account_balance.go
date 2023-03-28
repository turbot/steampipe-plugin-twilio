package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableTwilioAccountBalance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_balance",
		Description: "Retrieve balance of the current Twilio Account",
		List: &plugin.ListConfig{
			Hydrate:           listTwilioAccountBalance,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "balance",
				Description: "Balance of the account.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "currency",
				Description: "Currency of this balance.",
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

func listTwilioAccountBalance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_balance.listTwilioAccountBalance", "connection_error", err)
		return nil, err
	}

	balance, err := client.Api.FetchBalance(&openapi.FetchBalanceParams{})
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_balance.listTwilioAccountBalance", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, balance)

	return nil, nil
}
