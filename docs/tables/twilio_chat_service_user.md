# Table: twilio_chat_service_user

The User resource of Programmable Chat represents a single user who is identified by an identity value that you provide when the User resource is created. The User resource's identity must be unique within its Service instance.

The `twilio_chat_service_user` table can be used to query users within a chat service, and **you must specify which chat service** with `where service_sid='IS69cdg66f24de48919638c0a0bfaf2a70'`.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  identity,
  service_sid,
  date_created,
  is_online,
  account_sid
from
  twilio_chat_service_user;
```

### List all online users

```sql
select
  distinct sid,
  friendly_name,
  identity,
  service_sid,
  is_online,
  account_sid
from
  twilio_chat_service_user
where
  is_online;
```

### List channel count per user

```sql
select
  distinct sid,
  friendly_name,
  identity,
  joined_channels_count,
  account_sid
from
  twilio_chat_service_user;
```
