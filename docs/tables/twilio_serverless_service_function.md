---
title: "Steampipe Table: twilio_serverless_service_function - Query Twilio Serverless Service Functions using SQL"
description: "Allows users to query Twilio Serverless Service Functions, specifically the details of each function, providing insights into function configurations, runtime, and status."
---

# Table: twilio_serverless_service_function - Query Twilio Serverless Service Functions using SQL

Twilio Serverless is a service that allows developers to build and run Twilio applications without having to manage servers. Serverless Service Functions are individual units of code that are part of these applications. They are event-driven, meaning they run in response to triggers such as HTTP requests or incoming phone calls.

## Table Usage Guide

The `twilio_serverless_service_function` table provides insights into Serverless Service Functions within Twilio. As a developer or operations engineer, explore function-specific details through this table, including configurations, runtime, and status. Utilize it to uncover information about functions, such as their configurations, the runtime used, and their current status.

## Examples

### Basic info
Discover the segments that have been recently created within your Twilio Serverless service functions. This enables you to analyze the settings to understand which functions are associated with a specific service and account.

```sql+postgres
select
  sid,
  friendly_name,
  date_created,
  service_sid,
  account_sid
from
  twilio_serverless_service_function;
```

```sql+sqlite
select
  sid,
  friendly_name,
  date_created,
  service_sid,
  account_sid
from
  twilio_serverless_service_function;
```