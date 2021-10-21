# Table: twilio_account

Retrieves the list of account, along with any subaccounts belonging to it.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  type,
  status,
  date_created
from
  twilio_account;
```

### List all trial accounts

```sql
select
  sid,
  friendly_name,
  type,
  status,
  date_created
from
  twilio_account
where
  type = 'Trial';
```

### List inactive accounts

```sql
select
  sid,
  friendly_name,
  type,
  status,
  date_created
from
  twilio_account
where
  status <> 'active';
```

### List all sub-accounts

```sql
select
  sid,
  friendly_name,
  type,
  status,
  date_created
from
  twilio_account
where
  sid <> owner_account_sid;
```
