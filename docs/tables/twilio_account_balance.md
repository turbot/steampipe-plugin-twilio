---
title: "Steampipe Table: twilio_account_balance - Query Twilio Account Balances using SQL"
description: "Allows users to query Twilio Account Balances, providing insights into the current balance and currency type of Twilio accounts."
---

# Table: twilio_account_balance - Query Twilio Account Balances using SQL

Twilio Account Balance is a feature within Twilio that provides information about the current balance and the type of currency used in a Twilio account. It is a crucial component in understanding the financial status of a Twilio account, helping users to monitor and manage their usage and costs effectively. Twilio Account Balance allows users to stay informed about their account balance and take appropriate actions when necessary.

## Table Usage Guide

The `twilio_account_balance` table provides insights into the current balance and currency type of Twilio accounts. As a financial or operations analyst, explore account-specific details through this table, including the current balance and the type of currency used. Utilize it to uncover information about your Twilio accounts, such as the current financial status and the currency type, aiding in effective financial management and planning.

## Examples

### Fetch the balance
Discover the current balance in your Twilio account, expressed in the appropriate currency. This can be useful for monitoring your usage and expenses in real time.

```sql
select
  concat(balance, ' ', currency) as balance_formatted
from
  twilio_account_balance;
```