---
title: "Steampipe Table: twilio_serverless_service - Query Twilio Serverless Services using SQL"
description: "Allows users to query Twilio Serverless Services, providing information about the services, including service SID, unique name, date created, date updated, and more."
---

# Table: twilio_serverless_service - Query Twilio Serverless Services using SQL

Twilio Serverless is a development and runtime environment that enables developers to build and run Twilio applications without having to set up and manage servers. It provides an integrated environment for building, testing, and deploying Twilio applications, making it easier and faster to create and maintain Twilio applications. Twilio Serverless supports multiple programming languages and provides a range of tools and features to help developers build and run applications.

## Table Usage Guide

The `twilio_serverless_service` table provides insights into Twilio Serverless Services. As a developer or system administrator, you can explore service-specific details through this table, including service SID, unique name, date created, date updated, and more. This table can be utilized to manage and monitor the services, understand the usage patterns, and troubleshoot any issues.

## Examples

### Basic info
Discover the segments that have been created on the Twilio Serverless Service platform, including their names and creation dates. This query is useful for gaining insights into the services' setup and their associated account IDs without needing to dive into credential details.

```sql+postgres
select
  sid,
  friendly_name,
  date_created,
  include_credentials,
  account_sid
from
  twilio_serverless_service;
```

```sql+sqlite
select
  sid,
  friendly_name,
  date_created,
  include_credentials,
  account_sid
from
  twilio_serverless_service;
```

### List services not editable via UI
Determine the services that cannot be modified through the user interface. This can be useful in identifying services that may require manual intervention or additional permissions for modifications.

```sql+postgres
select
  sid,
  friendly_name,
  date_created,
  include_credentials,
  account_sid
from
  twilio_serverless_service
where
  not ui_editable;
```

```sql+sqlite
select
  sid,
  friendly_name,
  date_created,
  include_credentials,
  account_sid
from
  twilio_serverless_service
where
  ui_editable = 0;
```