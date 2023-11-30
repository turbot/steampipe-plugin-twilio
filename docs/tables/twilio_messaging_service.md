---
title: "Steampipe Table: twilio_messaging_service - Query Twilio Messaging Services using SQL"
description: "Allows users to query Twilio Messaging Services, providing details on the messaging service's SID, friendly name, status, type, and more."
---

# Table: twilio_messaging_service - Query Twilio Messaging Services using SQL

Twilio Messaging Service is a feature within Twilio that allows developers to send SMS, MMS, and WhatsApp messages. It provides a unified API to send messages and manage message delivery across multiple channels. This service also offers features such as message scheduling, content redaction, and delivery intelligence.

## Table Usage Guide

The `twilio_messaging_service` table provides insights into the Messaging Services within Twilio. As a developer or system administrator, explore specific details through this table, including the service's SID, friendly name, status, and type. Utilize it to manage and monitor your messaging services, ensuring efficient and secure communication.

## Examples

### Basic info
Explore which messaging services have been created on your Twilio account, along with their creation dates and validity periods. This is useful to keep track of your services and their active durations.

```sql
select
  sid,
  friendly_name,
  date_created,
  validity_period,
  account_sid
from
  twilio_messaging_service;
```

### List services with smart encoding enabled
Uncover the details of messaging services that have smart encoding enabled, allowing you to assess whether your communication platform is optimized for efficient data usage. This is particularly useful for managing and optimizing costs associated with data transmission in your messaging services.

```sql
select
  sid,
  friendly_name,
  date_created,
  smart_encoding,
  account_sid
from
  twilio_messaging_service
where
  smart_encoding;
```