---
title: "Steampipe Table: twilio_chat_service_user - Query Twilio Chat Service Users using SQL"
description: "Allows users to query Twilio Chat Service Users, providing detailed information about each user associated with a specific chat service in Twilio."
---

# Table: twilio_chat_service_user - Query Twilio Chat Service Users using SQL

Twilio Chat Service Users are entities within Twilio's Chat Service that represent individual participants in a chat. Each user has a unique identity and can be associated with one or more chat services. The user's identity, role, and other attributes can be managed and queried through Twilio's API.

## Table Usage Guide

The `twilio_chat_service_user` table provides insights into users within Twilio's Chat Service. As an application developer or system administrator, you can explore user-specific details through this table, including their roles, identities, and other associated metadata. Utilize it to uncover information about users, such as their participation in different chat services, their roles within these services, and the management of their identities.

**Important Notes**
- You must specify the `service_sid` in the `where` clause to query this table.

## Examples

### Basic info
Discover the segments that have been created within a specific chat service, including their status and creation date. This can be particularly useful in identifying active users and understanding the overall usage patterns of the service.

```sql
select
  sid,
  friendly_name,
  identity,
  service_sid,
  date_created,
  is_online,
  account_sid
from
  twilio_chat_service_user
where
  service_sid = 'IS69abc66f24de48919638c0a0bfaf2a70';
```

### List online users
Explore which users are currently online in a particular chat service. This can be particularly useful for real-time user engagement or to monitor active participation within a specific chat service.

```sql
select
  distinct sid,
  friendly_name,
  identity,
  service_sid,
  is_online,
  account_sid
from
  twilio_chat_service_user
where
  service_sid = 'IS69abc66f24de48919638c0a0bfaf2a70'
  and is_online;
```

### List channel count per user
Analyze the settings to understand the distribution of chat channels for each user within a specific service. This can be useful for identifying user engagement and managing resource allocation.

```sql
select
  distinct sid,
  friendly_name,
  identity,
  joined_channels_count,
  account_sid
from
  twilio_chat_service_user
where
  service_sid = 'IS69abc66f24de48919638c0a0bfaf2a70';
```