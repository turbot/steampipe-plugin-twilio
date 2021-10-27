# Table: twilio_account_application

An Application instance resource represents an application that you have created with Twilio. An application inside of Twilio is a set of URLs and other configuration data that tells Twilio how to behave when one of your Twilio numbers receives a call or SMS message.

## Examples

### Basic info

```sql
select
  sid,
  friendly_name,
  sms_url,
  sms_method,
  voice_url,
  voice_method,
  account_sid
from
  twilio_account_application;
```

### List applications with caller ID lookup feature enabled

```sql
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_caller_id_lookup,
  account_sid
from
  twilio_account_application
where
  voice_caller_id_lookup;
```

### List applications with no voice fallback URL configured

```sql
select
  sid,
  friendly_name,
  sms_url,
  voice_url,
  voice_fallback_url,
  account_sid
from
  twilio_account_application
where
  voice_fallback_url is null;
```
