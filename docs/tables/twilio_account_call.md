---
title: "Steampipe Table: twilio_account_call - Query Twilio Calls using SQL"
description: "Allows users to query Twilio Calls. This table provides details about individual calls made through the Twilio platform."
---

# Table: twilio_account_call - Query Twilio Calls using SQL

Twilio Calls represent individual voice calls made through the Twilio platform. They include details such as the call's status, start and end times, duration, and the phone numbers involved. Twilio's call resource is a key component of its voice communication services.

## Table Usage Guide

The `twilio_account_call` table provides insights into individual call details within the Twilio platform. As a developer or system administrator, explore call-specific details through this table, including call status, duration, and associated phone numbers. Utilize it to analyze call patterns, troubleshoot call issues, and monitor the usage of Twilio's voice communication services.

## Examples

### Basic info
Gain insights into the direction and status of calls made through a specific account, including when each call started and ended. This can help in tracking call activities and understanding call patterns for customer service improvement.

```sql+postgres
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call;
```

```sql+sqlite
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call;
```

### List outgoing calls
Analyze the details of outgoing calls made via an application's API, providing insights into call statuses and durations, which can be useful for assessing call performance and usage patterns.

```sql+postgres
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  direction = 'outbound-api';
```

```sql+sqlite
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  direction = 'outbound-api';
```

### List unsuccessful calls
Explore which calls were unsuccessful to identify potential issues with your telecommunication system. This can help in troubleshooting and improving the overall call success rate.

```sql+postgres
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  status = 'failed';
```

```sql+sqlite
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  status = 'failed';
```