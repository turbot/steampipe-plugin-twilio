---
title: "Steampipe Table: twilio_account_key - Query Twilio Account Keys using SQL"
description: "Allows users to query Twilio Account Keys, providing vital information about each key, including the date of creation, date of update, and its current status."
---

# Table: twilio_account_key - Query Twilio Account Keys using SQL

Twilio Account Keys are secure, revocable keys that are used to authenticate API requests. Each key is unique and can be used in place of a primary account SID and auth token. These keys can be created, viewed, and revoked through the Twilio console.

## Table Usage Guide

The `twilio_account_key` table provides insights into Twilio Account Keys. As a developer or system administrator, you can explore detailed information about each key, including its creation date, last update date, and status. This table is particularly useful for managing and auditing the use of API keys in your Twilio account.

**Important Notes**
- You must authenticate using either an [Auth Token](https://www.twilio.com/console) or a [Main API Key](https://www.twilio.com/docs/iam/keys/api-key) to query this table.

## Examples

### Basic info
Determine the areas in which your Twilio account keys were created to manage and track their usage over time. This aids in understanding the account key's lifecycle and ensuring their optimal utilization.

```sql+postgres
select
  sid,
  friendly_name,
  date_created
from
  twilio_account_key;
```

```sql+sqlite
select
  sid,
  friendly_name,
  date_created
from
  twilio_account_key;
```

### List keys older than 90 days
Explore which Twilio account keys are older than 90 days to ensure timely key rotation and enhance security measures. This is useful in maintaining good security practices and avoiding potential vulnerabilities due to outdated keys.

```sql+postgres
select
  sid,
  friendly_name,
  date_created,
  extract(day from current_timestamp - date_created) as age
from
  twilio_account_key
where
  extract(day from current_timestamp - date_created) > 90;
```

```sql+sqlite
select
  sid,
  friendly_name,
  date_created,
  julianday('now') - julianday(date_created) as age
from
  twilio_account_key
where
  julianday('now') - julianday(date_created) > 90;
```