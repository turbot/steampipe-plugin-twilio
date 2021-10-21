# Table: twilio_account_call

Retrieves all phone calls for your Twilio account.

## Examples

### Basic info

```sql
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call;
```

### List all outgoing calls

```sql
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  direction = 'outbound-api';
```

### List unsuccessful calls

```sql
select
  sid,
  called_to,
  direction,
  status,
  start_time,
  end_time,
  account_sid
from
  twilio_account_call
where
  status = 'failed';
```
