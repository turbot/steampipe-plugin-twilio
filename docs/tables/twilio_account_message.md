# Table: twilio_account_message

Retrieves all message calls for your Twilio account.

## Examples

### Basic info

```sql
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

### List all outgoing messages

```sql
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

```sql
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
