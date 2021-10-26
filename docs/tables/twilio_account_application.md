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

### List applications by name

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
  twilio_account_application
where
  friendly_name = 'MyApp';
```
