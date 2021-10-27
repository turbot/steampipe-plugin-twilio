# Table: twilio_chat_service

Twilio Programmable Chat makes it easy for you to add chat features into your web and native mobile applications without building or scaling a real-time chat backend.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  default_service_role_sid,
  typing_indicator_timeout,
  limits,
  account_sid
from
  twilio_chat_service;
```

### List services with reachability indicator enabled

```sql
select
  sid,
  friendly_name,
  reachability_enabled,
  account_sid
from
  twilio_chat_service
where
  reachability_enabled;
```

### List services with consumption horizon enabled

```sql
select
  sid,
  friendly_name,
  read_status_enabled,
  account_sid
from
  twilio_chat_service
where
  read_status_enabled;
```

### Get user count by chat service

```sql
select
  s.sid,
  s.friendly_name,
  count(u.sid)
from
  twilio_chat_service as s,
  twilio_chat_service_user as u
where
  u.service_sid = s.sid
group by
  s.sid,
  s.friendly_name;
```
