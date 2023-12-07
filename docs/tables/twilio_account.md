---
title: "Steampipe Table: twilio_account - Query Twilio Accounts using SQL"
description: "Allows users to query Twilio Accounts, specifically the account details and status, providing insights into the account usage, balance, and activity."
---

# Table: twilio_account - Query Twilio Accounts using SQL

Twilio Accounts is a resource within Twilio that allows users to manage their account details, status, and usage. It provides a centralized way to monitor and manage account balance, activity, and other related details. Twilio Accounts helps users stay informed about their account status and take appropriate actions when necessary.

## Table Usage Guide

The `twilio_account` table provides insights into the Twilio Accounts. As a developer or account manager, explore account-specific details through this table, including status, balance, and activity. Utilize it to uncover information about account usage and manage the account more effectively.

## Examples

### Basic info
Explore which Twilio accounts are sub-accounts and their current status. This is useful to understand the hierarchy and health of your Twilio account setup.

```sql+postgres
select
  friendly_name,
  status,
  is_sub_account
from
  twilio_account;
```

```sql+sqlite
select
  friendly_name,
  status,
  is_sub_account
from
  twilio_account;
```

### List trial accounts
Discover the segments that are utilizing trial accounts in your Twilio services. This can help assess the elements within your business that are testing or experimenting with Twilio's features, aiding in resource allocation and strategic planning.

```sql+postgres
select
  sid,
  friendly_name,
  type,
  status,
  is_sub_account
from
  twilio_account
where
  type = 'Trial';
```

```sql+sqlite
select
  sid,
  friendly_name,
  type,
  status,
  is_sub_account
from
  twilio_account
where
  type = 'Trial';
```