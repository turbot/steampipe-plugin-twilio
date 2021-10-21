connection "twilio" {
  plugin = "twilio"

  # Required
  account_sid = "YOUR_ACCOUNT_SID"

  # Option1 - Authenticate using Authorization Token
  # auth_token = "YOUR_AUTH_TOKEN"

  # Option2 - Authenticate using API Key and API Secret
  # api_key    = "YOUR_API_KEY"
  # api_secret = "YOUR_API_SECRET"
}
