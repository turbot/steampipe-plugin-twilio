package twilio

import (
	"context"
	"slices"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	twilioclient "github.com/twilio/twilio-go/client"
)

// function which returns an IsNotFoundErrorPredicate for Twilio API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		if twilioErr, ok := err.(*twilioclient.TwilioRestError); ok {
			return slices.Contains(notFoundErrors, types.ToString(twilioErr.Status))
		}
		return false
	}
}
