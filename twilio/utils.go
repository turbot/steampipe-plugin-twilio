package twilio

import (
	"context"
	"strings"
	"time"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
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

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getAccountSidMemoized = plugin.HydrateFunc(getAccountSIDUncached).Memoize(memoize.WithCacheKeyFunction(getAccountSidCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getAccountSID(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getAccountSidMemoized(ctx, d, h)
}

// Build a cache key for the call to getAccountSidCacheKey.
func getAccountSidCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getAccountSid"
	return key, nil
}

// Get current Twilio account SID
func getAccountSIDUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAccountSID", "connection_error", err)
		return nil, err
	}
	accountSID := client.Client.AccountSid()

	return accountSID, nil
}
