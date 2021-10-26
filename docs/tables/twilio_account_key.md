# Table: twilio_account_key

The Twilio account key provides a powerful and flexible primitive for managing access to the Twilio API.

**Note:** You must authenticate either using [Auth Token](https://www.twilio.com/console), or [Main API Keys](https://www.twilio.com/docs/iam/keys/api-key) to query this table.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  date_created
from
  twilio_account_key;
```

### List keys older than 90 days

```sql
select
  sid,
  friendly_name,
  date_created,
  extract(day from current_timestamp - date_created) as age
from
  twilio_account_key
where
  extract(day from current_timestamp - date_created) > 90;
```
