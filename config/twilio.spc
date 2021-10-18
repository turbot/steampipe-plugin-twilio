connection "twilio" {
  plugin  = "twilio"

  # Set the static credential with the `username` and `password` arguments.
  # Authenticate using AccountSId and Authorization token
  username = "YOUR_ACCOUNT_SID"
  password = "YOUR_AUTH_TOKEN"

  # Authenticate using API Key and Secret
  username    = "YOUR_API_KEY"
  password    = "YOUR_API_SECRET"
  account_sid = "YOUR_ACCOUNT_SID"
}
