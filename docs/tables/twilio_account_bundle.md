---
title: "Steampipe Table: twilio_account_bundle - Query Twilio Account Bundles using SQL"
description: "Allows users to query Twilio Account Bundles, providing detailed information about each bundle, including its type, status, and the account it is associated with."
---

# Table: twilio_account_bundle - Query Twilio Account Bundles using SQL

Twilio Account Bundles are collections of resources that are grouped together and billed as a single unit. Each bundle is associated with a specific account and can include various types of resources, such as phone numbers, messaging services, and more. Account Bundles provide a consolidated view of resource usage and costs, simplifying billing and resource management for Twilio users.

## Table Usage Guide

The `twilio_account_bundle` table provides insights into account bundles within Twilio. As a DevOps engineer or a cloud accountant, explore bundle-specific details through this table, including the types of resources included in each bundle, their status, and associated costs. Utilize it to monitor resource usage, manage costs, and optimize resource allocation across your Twilio account.

## Examples

### Basic info
Explore which Twilio account bundles are active and when they will expire. This can help in managing and optimizing your Twilio services effectively.

```sql+postgres
select 
  sid
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle;
```

```sql+sqlite
select 
  sid,
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle;
```

### List bundles in draft status
Discover the segments that contain draft status bundles in your Twilio account. This is useful for assessing which bundles are yet to be finalized and may require further editing or approval.

```sql+postgres
select 
  sid
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle
where 
  status = 'draft';
```

```sql+sqlite
select 
  sid,
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle
where 
  status = 'draft';
```