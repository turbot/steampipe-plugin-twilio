# Table: twilio_account_balance

The balance of the current account.

## Examples

### Fetch the balance

```sql
select
  concat(balance, ' ', currency) as balance_formatted
from
  twilio_account_balance;
```
