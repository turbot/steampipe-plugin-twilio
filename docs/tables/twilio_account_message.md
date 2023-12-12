---
title: "Steampipe Table: twilio_account_message - Query Twilio Account Messages using SQL"
description: "Allows users to query Twilio Account Messages, specifically the details of individual messages sent or received on a Twilio account, providing insights into message traffic and potential anomalies."
---

# Table: twilio_account_message - Query Twilio Account Messages using SQL

Twilio Messaging is a service within Twilio that allows you to send and receive SMS, MMS, and other types of messages globally. It provides a simple and powerful API for developers to integrate messaging capabilities into their applications. Twilio Messaging helps you stay informed about the details of individual messages sent or received on a Twilio account.

## Table Usage Guide

The `twilio_account_message` table provides insights into individual messages sent or received on a Twilio account. As a developer or IT professional, explore message-specific details through this table, including status, direction, and associated metadata. Utilize it to uncover information about messages, such as those with delivery issues, the direction of messages (inbound or outbound), and verification of message content.

## Examples

### Basic info
Explore which messages were sent in what direction and their status to gain insights into communication patterns and potential issues. This can be useful in assessing the effectiveness of communication strategies and identifying areas for improvement.

```sql+postgres
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message;
```

```sql+sqlite
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message;
```

### List outgoing messages
Explore the history of outgoing messages sent through your account. This can be useful for tracking communication patterns, identifying potential issues, or auditing message activity.

```sql+postgres
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message
where
  direction = 'outbound-api';
```

```sql+sqlite
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message
where
  direction = 'outbound-api';
```

### List undelivered messages
Discover the segments that contain undelivered messages to understand potential communication gaps or issues within your system. This can be particularly useful in troubleshooting or improving customer communication strategies.

```sql+postgres
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message
where
  status <> 'delivered';
```

```sql+sqlite
select
  sid,
  sent_to,
  direction,
  status,
  date_sent,
  account_sid
from
  twilio_account_message
where
  status <> 'delivered';
```