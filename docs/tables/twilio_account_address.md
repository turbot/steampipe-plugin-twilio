# Table: twilio_account_address

An Address instance resource represents your or your customer’s physical location within a country. Around the world, some local authorities require the name and address of the user to be on file with Twilio to purchase and own a phone number.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  customer_name,
  validated,
  street,
  city,
  region,
  postal_code,
  account_sid
from
  twilio_account_address;
```

### List unverified addresses

```sql
select
  sid,
  friendly_name,
  customer_name,
  verified,
  street,
  city,
  region,
  postal_code,
  account_sid
from
  twilio_account_address
where
  not verified;
```

### List addresses used as emergency address

```sql
select
  sid,
  friendly_name,
  customer_name,
  emergency_enabled,
  street,
  city,
  region,
  postal_code,
  account_sid
from
  twilio_account_address
where
  emergency_enabled;
```
