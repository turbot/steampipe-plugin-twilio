package twilio

import (
	"context"
	"strings"
	"time"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
func handleListError(err error) bool {
	return strings.Contains(types.ToString(err), "could not retrieve payload from response")
}