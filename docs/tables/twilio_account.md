# Table: twilio_account

Retrieves the list of account, along with any sub-accounts belonging to it.

**Note:** You must authenticate either using `Authorization Token`, or `Main API Keys` to query this table.

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
