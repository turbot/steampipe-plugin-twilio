# Table: twilio_messaging_service

Messaging Services is a free Twilio feature designed to make it easier to send
messages at-scale, as well as across multiple countries. In short, a Messaging
Service is a container for multiple Twilio message senders (e.g. phone numbers,
WhatsApp senders).

It offers you additional message sending intelligence and content features,
such as automatic number selection, when you pass your Messaging Service
information in your API requests instead of a specific From number.

## Examples

### Basic info

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
