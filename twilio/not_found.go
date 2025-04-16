package twilio

import (
	"slices"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	twilioclient "github.com/twilio/twilio-go/client"
)

// function which returns an IsNotFoundErrorPredicate for Twilio API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if twilioErr, ok := err.(*twilioclient.TwilioRestError); ok {
			return slices.Contains(notFoundErrors, types.ToString(twilioErr.Status))
		}
		return false
	}
}
