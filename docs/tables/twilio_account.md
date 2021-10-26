# Table: twilio_account

Retrieves the list of account, along with any sub-accounts belonging to it.

**Note:** You must authenticate either using [Auth Token](https://www.twilio.com/console), or [Main API Keys](https://www.twilio.com/docs/iam/keys/api-key) to query this table.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  type,
  status,
  is_sub_account
from
  twilio_account;
```

### List trail accounts

```sql
select
  sid,
  friendly_name,
  type,
  status,
  is_sub_account
from
  twilio_account
where
  type = 'Trial';
```
