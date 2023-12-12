---
title: "Steampipe Table: twilio_account_application - Query Twilio Account Applications using SQL"
description: "Allows users to query Twilio Account Applications, specifically the application details and related metadata, providing insights into application configurations and settings."
---

# Table: twilio_account_application - Query Twilio Account Applications using SQL

Twilio Account Applications represent the settings and configurations for a specific application in Twilio. It encapsulates the behavior you’d like your call to exhibit once it’s connected. This includes things like the URL of the script you’d like to run, whether to record the calls, and whether to record the calls as mono or dual-channel.

## Table Usage Guide

The `twilio_account_application` table provides insights into application settings within Twilio. As a DevOps engineer, explore application-specific details through this table, including URLs of the scripts to run, recording preferences, and channel settings. Utilize it to uncover information about applications, such as those with specific configurations, the behavior of calls once connected, and the verification of application settings.

## Examples

### Basic info
Explore which Twilio applications are currently linked to your account, with a focus on understanding how each application is configured for sending SMS and voice messages. This can be useful for auditing communication settings and ensuring the correct methods are in place.

```sql+postgres
select
  sid,
  friendly_name,
  sms_url,
  sms_method,
  voice_url,
  voice_method,
  account_sid
from
  twilio_account_application;
```

```sql+sqlite
select
  sid,
  friendly_name,
  sms_url,
  sms_method,
  voice_url,
  voice_method,
  account_sid
from
  twilio_account_application;
```

### List applications with caller ID lookup feature enabled
Gain insights into the applications that have the caller ID lookup feature enabled. This can be useful for auditing purposes, ensuring that the feature is only activated for the intended applications.

```sql+postgres
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_caller_id_lookup,
  account_sid
from
  twilio_account_application
where
  voice_caller_id_lookup;
```

```sql+sqlite
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_caller_id_lookup,
  account_sid
from
  twilio_account_application
where
  voice_caller_id_lookup = 1;
```

### List applications with no voice fallback URL configured
Discover which applications lack a configured voice fallback URL. This can help in identifying potential communication gaps in your Twilio account, ensuring that all applications have a fallback option for voice services.

```sql+postgres
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_fallback_url,
  account_sid
from
  twilio_account_application
where
  voice_fallback_url is null;
```

```sql+sqlite
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_fallback_url,
  account_sid
from
  twilio_account_application
where
  voice_fallback_url is null;
```