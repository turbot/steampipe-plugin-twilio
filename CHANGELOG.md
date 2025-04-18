## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#57](https://github.com/turbot/steampipe-plugin-twilio/pull/57))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#57](https://github.com/turbot/steampipe-plugin-twilio/pull/57))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#55](https://github.com/turbot/steampipe-plugin-twilio/pull/55))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#55](https://github.com/turbot/steampipe-plugin-twilio/pull/55))

## v0.7.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#49](https://github.com/turbot/steampipe-plugin-twilio/pull/49))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#49](https://github.com/turbot/steampipe-plugin-twilio/pull/49))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-twilio/blob/main/docs/LICENSE). ([#49](https://github.com/turbot/steampipe-plugin-twilio/pull/49))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#48](https://github.com/turbot/steampipe-plugin-twilio/pull/48))

## v0.6.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#32](https://github.com/turbot/steampipe-plugin-twilio/pull/32))

## v0.6.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#29](https://github.com/turbot/steampipe-plugin-twilio/pull/29))
- Recompiled plugin with Go version `1.21`. ([#29](https://github.com/turbot/steampipe-plugin-twilio/pull/29))

## v0.5.0 [2023-03-31]

_What's new?_

- New tables added
  - [twilio_account_balance](https://hub.steampipe.io/plugins/turbot/twilio/tables/twilio_account_balance) ([#19](https://github.com/turbot/steampipe-plugin-twilio/pull/19)) (Thanks [@reva](https://github.com/reva) for the contribution!)

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
