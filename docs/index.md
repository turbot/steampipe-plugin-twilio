---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/twilio.svg"
brand_color: "#E30810"
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
  phone_number,
  status,
  date_created,
  account_sid
from
  twilio_account_incoming_phone_number;
```

```
+------------------------------------+--------------+--------+----------------------+------------------------------------+
| sid                                | phone_number | status | date_created         | account_sid                        |
+------------------------------------+--------------+--------+----------------------+------------------------------------+
| PN91973970d5c9d01cf98068ad29bc4b72 | +13515239901 | in-use | 2021-10-18T11:09:48Z | ACe0ad2cff256c79c17a75fafd74ac483d |
+------------------------------------+--------------+--------+----------------------+------------------------------------+
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

  # Required
  account_sid = "YOUR_ACCOUNT_SID"

  # Option1 - Authenticate using Authorization Token
  # auth_token = "YOUR_AUTH_TOKEN"

  # Option2 - Authenticate using API Key and API Secret
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

## Configuring Twilio Credentials

### Credentials from Environment Variables

The Twilio plugin will use the standard Twilio environment variables to obtain credentials **only if other arguments (`account_sid`, `api_key`, `api_secret`, `auth_token`) are not specified** in the connection:

```sh
export TWILIO_ACCOUNT_SID=<ACCOUNT_SID>
export TWILIO_API_KEY=<YOUR_API_KEY>
export TWILIO_API_SECRET=<YOUR_API_SECRET>
export TWILIO_AUTH_TOKEN=<ACCOUNT_AUTH_TOKEN>
```

```hcl
connection "twilio" {
  plugin = "twilio"
}
```