---
title: "Steampipe Table: twilio_chat_service - Query Twilio Chat Services using SQL"
description: "Allows users to query Twilio Chat Services, specifically providing information about each chat service's unique ID, friendly name, webhook configuration, and more."
---

# Table: twilio_chat_service - Query Twilio Chat Services using SQL

Twilio Chat Services is a feature within Twilio's communication platform that allows the creation and management of chat services for applications. It offers a set of tools to build and scale real-time chat applications, offering features like message history, user identity and status, typing indicators, and more. Twilio Chat Services helps in creating a rich and interactive communication experience in any application.

## Table Usage Guide

The `twilio_chat_service` table provides insights into each chat service within Twilio's platform. As a developer or system administrator, explore service-specific details through this table, including service ID, friendly name, default service role, and webhook configuration. Utilize it to manage and monitor your chat services, ensuring optimal configuration and performance.

## Examples

### Basic info
Gain insights into the specific chat services within your Twilio account, including their unique identifiers and associated roles, to better understand and manage your communication services. This could be particularly useful for auditing purposes or for streamlining your chat services.

```sql+postgres
select
  sid,
  friendly_name,
  default_service_role_sid,
  typing_indicator_timeout,
  limits,
  account_sid
from
  twilio_chat_service;
```

```sql+sqlite
select
  sid,
  friendly_name,
  default_service_role_sid,
  typing_indicator_timeout,
  limits,
  account_sid
from
  twilio_chat_service;
```

### List services with reachability indicator enabled
Explore which chat services have the reachability indicator enabled. This can help you identify services where users are notified of their message's delivery status, enhancing communication efficiency.

```sql+postgres
select
  sid,
  friendly_name,
  reachability_enabled,
  account_sid
from
  twilio_chat_service
where
  reachability_enabled;
```

```sql+sqlite
select
  sid,
  friendly_name,
  reachability_enabled,
  account_sid
from
  twilio_chat_service
where
  reachability_enabled = 1;
```

### List services with consumption horizon enabled
Explore which chat services have the consumption horizon feature enabled. This query is useful in identifying services that allow users to track the last read message, enhancing user experience by keeping track of conversation progress.

```sql+postgres
select
  sid,
  friendly_name,
  read_status_enabled,
  account_sid
from
  twilio_chat_service
where
  read_status_enabled;
```

```sql+sqlite
select
  sid,
  friendly_name,
  read_status_enabled,
  account_sid
from
  twilio_chat_service
where
  read_status_enabled = 1;
```

### Get user count by chat service
Explore which chat services have the most users to understand their popularity and usage trends. This can help in allocating resources effectively or planning targeted promotions.

```sql+postgres
select
  s.sid,
  s.friendly_name,
  count(u.sid)
from
  twilio_chat_service as s,
  twilio_chat_service_user as u
where
  u.service_sid = s.sid
group by
  s.sid,
  s.friendly_name;
```

```sql+sqlite
select
  s.sid,
  s.friendly_name,
  count(u.sid)
from
  twilio_chat_service as s
join
  twilio_chat_service_user as u on u.service_sid = s.sid
group by
  s.sid,
  s.friendly_name;
```