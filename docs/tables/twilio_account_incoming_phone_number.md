# Table: twilio_account_incoming_phone_number

An incoming phone number instance resource represents a Twilio phone number provisioned from Twilio, ported or hosted to Twilio.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  phone_number,
  status,
  date_created,
  account_sid
from
  twilio_account_incoming_phone_number;
```

### List phone numbers with no emergency address registered

```sql
select
  sid,
  friendly_name,
  phone_number,
  status,
  date_created,
  account_sid
from
  twilio_account_incoming_phone_number
where
  emergency_address_status = 'unregistered';
```

### List all call logs for a phone number

```sql
select
  c.called_from as caller,
  c.called_to as receiver,
  c.start_time,
  c.end_time,
  c.duration,
  c.price,
  c.price_unit,
  c.account_sid
from
  twilio_account_call as c,
  twilio_account_incoming_phone_number as ph
where
  c.status = 'completed'
  and (
    ph.phone_number = c.called_to
    or ph.phone_number = c.called_from
  );
```
