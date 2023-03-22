## v0.4.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#20](https://github.com/turbot/steampipe-plugin-twilio/pull/20))

## v0.3.1 [2022-10-05]

_Bug fixes_

- Fixed all tables not returning more than 50 results due to the max limit being set instead of page size.

## v0.3.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#16](https://github.com/turbot/steampipe-plugin-twilio/pull/16))
- Recompiled plugin with Go version `1.19`. ([#16](https://github.com/turbot/steampipe-plugin-twilio/pull/16))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#12](https://github.com/turbot/steampipe-plugin-twilio/pull/12))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#11](https://github.com/turbot/steampipe-plugin-twilio/pull/11))

## v0.1.0 [2021-12-15]

_What's new?_

- New tables added
  - [twilio_account_bundle](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_bundle) ([#5](https://github.com/turbot/steampipe-plugin-twilio/pull/5))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#8](https://github.com/turbot/steampipe-plugin-twilio/pull/8))
- Recompiled plugin with Go version 1.17 ([#8](https://github.com/turbot/steampipe-plugin-twilio/pull/8))

## v0.0.1 [2021-10-26]

_What's new?_

- New tables added

  - [twilio_account](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account)
  - [twilio_account_address](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_address)
  - [twilio_account_application](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_application)
  - [twilio_account_call](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_call)
  - [twilio_account_incoming_phone_number](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_incoming_phone_number)
  - [twilio_account_key](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_key)
  - [twilio_account_message](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_message)
  - [twilio_chat_service](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_chat_service)
  - [twilio_chat_service_user](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_chat_service_user)
  - [twilio_messaging_service](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_messaging_service)
  - [twilio_serverless_service](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_serverless_service)
  - [twilio_serverless_service_function](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_serverless_service_function)
