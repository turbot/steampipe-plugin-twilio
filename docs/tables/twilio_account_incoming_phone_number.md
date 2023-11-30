---
title: "Steampipe Table: twilio_account_incoming_phone_number - Query Twilio Incoming Phone Numbers using SQL"
description: "Allows users to query Incoming Phone Numbers in Twilio, providing details such as phone number, capabilities, status, etc."
---

# Table: twilio_account_incoming_phone_number - Query Twilio Incoming Phone Numbers using SQL

Twilio Incoming Phone Numbers are unique phone numbers that are assigned to your Twilio account. These numbers can be used to make and receive calls and text messages. They come with various features and capabilities, such as voice, fax, SMS, and MMS.

## Table Usage Guide

The `twilio_account_incoming_phone_number` table provides insights into the incoming phone numbers associated with your Twilio account. As a developer or a system administrator, you can use this table to retrieve detailed information about each phone number, including its capabilities (voice, fax, SMS, MMS), status, and more. This information can be useful for managing and monitoring your communication capabilities with Twilio.

## Examples

### Basic info
Explore which phone numbers are associated with your Twilio account, identify their current status, and gain insights into when they were created. This is useful for managing and tracking your Twilio resources effectively.

```sql
select
  sid,
  friendly_name,
  phone_number,
  status,
  date_created,
  account_sid
from
  twilio_account_incoming_phone_number;
```

### List phone numbers with no emergency address registered
Identify phone numbers that have not registered an emergency address. This is useful for ensuring that all numbers are properly set up for emergency situations.

```sql
select
  sid,
  friendly_name,
  phone_number,
  status,
  date_created,
  account_sid
from
  twilio_account_incoming_phone_number
where
  emergency_address_status = 'unregistered';
```

### List call logs for a phone number
Gain insights into the completed calls associated with a specific phone number, including details such as the caller, receiver, timing, duration, and cost. This can be particularly useful for tracking communication patterns, monitoring costs, or identifying potential misuse of services.

```sql
select
  c.called_from as caller,
  c.called_to as receiver,
  c.start_time,
  c.end_time,
  c.duration,
  c.price,
  c.price_unit,
  c.account_sid
from
  twilio_account_call as c,
  twilio_account_incoming_phone_number as ph
where
  c.status = 'completed'
  and (
    ph.phone_number = c.called_to
    or ph.phone_number = c.called_from
  );
```