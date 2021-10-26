connection "twilio" {
  plugin = "twilio"

  # `account_sid` (Required) - The Account SID of your Twilio account/sub-account.
  # If `account_sid` is not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_ACCOUNT_SID` environment variable.
  # account_sid = "YOUR_ACCOUNT_SID"

  # Authenticate using Authorization Token
  # `auth_token` (optional) - The authorization token of your Twilio account/sub-account.
  # If `auth_token` is not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_AUTH_TOKEN` environment variable.
  # auth_token = "YOUR_AUTH_TOKEN"

  # Authenticate using API Key and API Secret
  # `api_key` (optional) - The API key.
  # `api_secret` (optional) - The secret of your API key.
  # If `api_key` and `api_secret` are not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_API_KEY` and `TWILIO_API_SECRET` environment variables.
  # api_key    = "YOUR_API_KEY"
  # api_secret = "YOUR_API_SECRET"
}
