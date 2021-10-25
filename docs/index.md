---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/twilio.svg"
brand_color: "#F22F46"
display_name: "Twilio"
short_name: "twilio"
description: "Steampipe plugin to query calls, messages and other communication functions from your Twilio project."
og_description: "Query Twilio with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/twilio-social-graphic.png"
---

# Twilio + Steampipe

[Twilio](https://www.twilio.com) is a cloud communications platform, offering developers to programmatically make and receive phone calls, send and receive text messages, and perform other communication functions using its web service APIs.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List all active phone numbers in your Twilio account:

```sql
select
  sid,
  friendly_name,
  extract(day from now() - date_created) as age,
  account_sid
from
  twilio_account_key;
```

```
+------------------------------------+---------------+-----+------------------------------------+
| sid                                | friendly_name | age | account_sid                        |
+------------------------------------+---------------+-----+------------------------------------+
| SK1dd43df2cc722a368ab925c0642d7896 | dev-test      | 7   | ACe0ad3djf256b88c17e75fafd74ac483d |
| SK31fccd42e0071567e86fee58ed433600 | test-main     | 7   | ACe0ad3djf256b88c17e75fafd74ac483d |
+------------------------------------+---------------+-----+------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/twilio/tables)**

## Get started

### Install

Download and install the latest Twilio plugin:

```bash
steampipe plugin install twilio
```

### Credentials

| Item | Description |
| - | - |
| Credentials | 1. Get your `Account SID` and `Auth Token` from [Twilio Console](https://www.twilio.com/console).<br />2. If you want to use `API keys` to authenticate instead of your Twilio account SID and auth token, generate your [API Keys](https://www.twilio.com/console/runtime/api-keys). |
| Radius | Each connection represents a single Twilio account/sub-account. |
| Resolution | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/twilio.spc`).<br />2. Credentials specified in environment variables e.g. `TWILIO_ACCOUNT_SID`, `TWILIO_API_KEY`, `TWILIO_API_SECRET`, and `TWILIO_AUTH_TOKEN`. |

### Configuration

Installing the latest twilio plugin will create a config file (`~/.steampipe/config/twilio.spc`) with a single connection named `twilio`:

```hcl
connection "twilio" {
  plugin = "twilio"

  # `account_sid` (Required) - The Account SID of your Twilio account/sub-account.
  # If `account_sid` is not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_ACCOUNT_SID` environment variable.
  # account_sid = "YOUR_ACCOUNT_SID"

  # Option1 - Authenticate using Authorization Token
  # `auth_token` (optional) - The authorization token of your Twilio account/sub-account.
  # If `auth_token` is not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_AUTH_TOKEN` environment variable.
  # auth_token = "YOUR_AUTH_TOKEN"

  # Option2 - Authenticate using API Key and API Secret
  # `api_key` (optional) - The API key.
  # `api_secret` (optional) - The secret of your API key.
  # If `api_key` and `api_secret` are not specified in a connection, it will be loaded from:
  # The path specified in the `TWILIO_API_KEY` and `TWILIO_API_SECRET` environment variables.
  # api_key    = "YOUR_API_KEY"
  # api_secret = "YOUR_API_SECRET"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-twilio
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)

## Multi-Account Connections

You may create multiple twilio connections:

```hcl
connection "twilio_01" {
  plugin  = "twilio"
}

connection "twilio_02" {
  plugin  = "twilio"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from twilio_02.twilio_account_call;
```

Alternatively, can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/using-steampipe/managing-connections#setting-the-search-path):

```sql
select * from twilio_account_call;
```

You can multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators).Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection:

```hcl
connection "twilio_all" {
  plugin      = "twilio"
  type        = "aggregator"
  connections = ["twilio_01", "twilio_02"]
}
```

Querying tables from this connection will return results from the `twilio_01` and `twilio_02` connections:

```sql
select * from twilio_all.twilio_account_call;
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the Twilio plugin connections whose names begin with `twilio_`:

```hcl
connection "twilio_all" {
  type        = "aggregator"
  plugin      = "twilio"
  connections = ["twilio_*"]
}
```
