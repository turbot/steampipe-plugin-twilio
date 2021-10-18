package twilio

import (
	"context"
	"time"

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
