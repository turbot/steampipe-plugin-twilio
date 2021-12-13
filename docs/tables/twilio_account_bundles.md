# Table: twilio_account_bundles

The Regulatory Bundles are Item Assignments of End-Users and Supporting Documents for regulatory compliance.

Depending on the configuration of the bundle, the bundle is being assessed against a Regulation (e.g., Germany local phone numbers for a business). Different Regulations need Item Assignments combinations of End-User Types and Supporting Document Types.

## Examples

### Basic info

```sql
select 
  sid
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle;
```

### List bundles in draft status

```sql
select 
  sid
  account_sid, 
  friendly_name, 
  email, 
  regulation_sid, 
  url, 
  valid_until 
from 
  twilio_account_bundle
where 
  status = 'draft';
```
