package twilio

import (
	"context"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTwilioAccountCall(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "twilio_account_call",
		Description: "Messages in the Twilio account.",
		List: &plugin.ListConfig{
			Hydrate: listAccountCalls,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "start_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:      "end_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:    "called_from",
					Require: plugin.Optional,
				},
				{
					Name:    "called_to",
					Require: plugin.Optional,
				},
				{
					Name:    "status",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAccountCall,
			KeyColumns: plugin.SingleColumn("sid"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "sid",
				Description: "The unique string that identifies the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "called_to",
				Description: "The phone number, SIP address or Client identifier that received this call.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("To"),
			},
			{
				Name:        "called_from",
				Description: "The phone number, SIP address or Client identifier that made this call.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("From"),
			},
			{
				Name:        "status",
				Description: "The status of this call.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "direction",
				Description: "A string describing the direction of the call. `inbound` for inbound calls, `outbound-api` for calls initiated via the REST API or `outbound-dial` for calls initiated by a `Dial` verb.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "start_time",
				Description: "The start time of the call. Null if the call has not yet been dialed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("StartTime").Transform(ensureTimestamp),
			},
			{
				Name:        "end_time",
				Description: "The end time of the call. Null if the call did not complete successfully.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("EndTime").Transform(ensureTimestamp),
			},
			{
				Name:        "annotation",
				Description: "The annotation provided for the call.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "answered_by",
				Description: "Either `human` or `machine` if this call was initiated with answering machine detection. Empty otherwise.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "api_version",
				Description: "The API version used to process the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "caller_name",
				Description: "The caller's name if this call was an incoming call to a phone number with caller ID Lookup enabled. Otherwise, empty.",
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
				Name:        "duration",
				Description: "The length of the call in seconds.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "forwarded_from",
				Description: "The forwarding phone number if this call was an incoming call forwarded from another number (depends on carrier supporting forwarding). Otherwise, empty.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "from_formatted",
				Description: "The calling phone number, SIP address, or Client identifier formatted for display.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_sid",
				Description: "The Group SID associated with this call. If no Group is associated with the call, the field is empty.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_call_sid",
				Description: "The SID that identifies the call that created this leg.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "phone_number_sid",
				Description: "If the call was inbound, this is the SID of the IncomingPhoneNumber resource that received the call. If the call was outbound, it is the SID of the OutgoingCallerId resource from which the call was placed.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "price",
				Description: "The charge for this call, in the currency associated with the account. Populated after the call is completed. May not be immediately available.",
				Type:        proto.ColumnType_DOUBLE,
			},
			{
				Name:        "price_unit",
				Description: "The currency in which 'price' is measured.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "queue_time",
				Description: "The wait time in milliseconds before the call is placed.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "to_formatted",
				Description: "The phone number, SIP address or Client identifier that received this call. Formatted for display.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "trunk_sid",
				Description: "The unique identifier of the trunk resource that was used for this call.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "uri",
				Description: "The URI of the resource, relative to 'https://api.twilio.com'.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "sub_resource_uris",
				Description: "A list of related subresources identified by their relative URIs.",
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

func listAccountCalls(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_call.listAccountCalls", "connection_error", err)
		return nil, err
	}

	req := &openapi.ListCallParams{}

	// Additional filters
	if d.KeyColumnQuals["called_from"] != nil {
		req.SetFrom(d.KeyColumnQuals["called_from"].GetStringValue())
	}

	if d.KeyColumnQuals["called_to"] != nil {
		req.SetTo(d.KeyColumnQuals["called_to"].GetStringValue())
	}

	if d.KeyColumnQuals["status"] != nil {
		req.SetStatus(d.KeyColumnQuals["status"].GetStringValue())
	}

	if d.Quals["start_time"] != nil {
		for _, q := range d.Quals["start_time"].Quals {
			switch q.Operator {
			case "<", "<=":
				req.SetStartTimeBefore(q.Value.GetTimestampValue().AsTime())
			case "=":
				req.SetStartTime(q.Value.GetTimestampValue().AsTime())
			case ">", ">=":
				req.SetStartTimeAfter(q.Value.GetTimestampValue().AsTime())
			}
		}
	}

	if d.Quals["end_time"] != nil {
		for _, q := range d.Quals["end_time"].Quals {
			switch q.Operator {
			case "<", "<=":
				req.SetEndTimeBefore(q.Value.GetTimestampValue().AsTime())
			case "=":
				req.SetEndTime(q.Value.GetTimestampValue().AsTime())
			case ">", ">=":
				req.SetEndTimeAfter(q.Value.GetTimestampValue().AsTime())
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

	resp, err := client.ApiV2010.ListCall(req)
	if err != nil {
		if handleListError(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, call := range resp {
		d.StreamListItem(ctx, call)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getAccountCall(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAccountCall")

	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("twilio_account_call.getAccountCall", "connection_error", err)
		return nil, err
	}
	callSid := d.KeyColumnQuals["sid"].GetStringValue()

	// No inputs
	if callSid == "" {
		return nil, nil
	}

	resp, err := client.ApiV2010.FetchCall(callSid, &openapi.FetchCallParams{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
