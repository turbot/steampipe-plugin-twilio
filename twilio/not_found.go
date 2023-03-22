package twilio

import (
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	twilioclient "github.com/twilio/twilio-go/client"
)

// function which returns an IsNotFoundErrorPredicate for Twilio API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if twilioErr, ok := err.(*twilioclient.TwilioRestError); ok {
			return helpers.StringSliceContains(notFoundErrors, types.ToString(twilioErr.Status))
		}
		return false
	}
}
