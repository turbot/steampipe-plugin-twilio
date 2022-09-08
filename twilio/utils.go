package twilio

import (
	"context"
	"strings"
	"time"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

// transforms

func ensureTimestamp(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	inputTime := d.Value.(*string)

	// Parse string to time.Time
	parsedTime, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", *inputTime)
	if err != nil {
		return nil, nil
	}

	return parsedTime.Format(time.RFC3339), nil
}

// Paging library returns an error if there is no result
// NOTE: This handling can be removed once it gets handled by twilio-sdk
// https://github.com/twilio/twilio-go/issues/114
func handleListError(err error) bool {
	return strings.Contains(types.ToString(err), "could not retrieve payload from response")
}

// Get current Twilio account SID
func getAccountSID(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cacheKey := "TwilioAccountSID"

	// if found in cache, return the result
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string), nil
	}

	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAccountSID", "connection_error", err)
		return nil, err
	}
	accountSID := client.Client.AccountSid()

	// save to extension cache
	d.ConnectionManager.Cache.Set(cacheKey, accountSID)

	return accountSID, nil
}
