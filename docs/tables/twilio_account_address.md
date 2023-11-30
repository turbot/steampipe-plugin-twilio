---
title: "Steampipe Table: twilio_account_address - Query Twilio Account Addresses using SQL"
description: "Allows users to query Twilio Account Addresses, specifically the physical address details associated with a Twilio account."
---

# Table: twilio_account_address - Query Twilio Account Addresses using SQL

Twilio Account Addresses are resources within Twilio that contain the physical address details associated with a Twilio account. These addresses are often used for regulatory compliance and emergency services purposes. They can be customer addresses, or the addresses of the business that owns the Twilio account.

## Table Usage Guide

The `twilio_account_address` table provides insights into the physical addresses associated with a Twilio account. As a DevOps engineer, you can explore address-specific details through this table, including the customer's address or the business's address associated with the Twilio account. Utilize it to uncover information about addresses for regulatory compliance and emergency services purposes.

## Examples

### Basic info
Explore which Twilio account addresses have been validated. This can help identify instances where incorrect or incomplete information has been provided, allowing for necessary corrections and updates to ensure accurate communication.

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
Discover the segments that contain unverified addresses within customer accounts. This is particularly useful for identifying potential inaccuracies or inconsistencies in your customer data.

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

### List emergency addresses
Explore which accounts have emergency addresses enabled to ensure rapid response capabilities in critical situations. This can be particularly beneficial in managing risk and improving safety measures.

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