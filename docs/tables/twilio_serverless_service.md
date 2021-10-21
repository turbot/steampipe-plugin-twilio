# Table: twilio_serverless_service

A Twilio Serverless services is an application container to store all your Functions and Assets, and used to manage deployments and separate environments.

## Examples

### Basic info

```sql
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

```sql
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
