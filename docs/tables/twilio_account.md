# Table: twilio_account

Retrieves the list of account, along with any sub-accounts belonging to it.

**Note:** You must authenticate either using `Main API Keys`, or `Authorization Token` to query this table.

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
