![image](https://hub.steampipe.io/images/plugins/turbot/twilio-social-graphic.png)

# Twilio Plugin for Steampipe

Use SQL to query calls, messages and other communication functions from your Twilio account.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/twilio)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/twilio/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-twilio/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install twilio
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/twilio#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/twilio#configuration).

Run a query:

```sql
select
  sid,
  friendly_name,
  extract(day from now() - date_created) as age,
  account_sid
from
  twilio_account_key;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-twilio.git
cd steampipe-plugin-twilio
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/twilio.spc
```

Try it!

```shell
steampipe query
> .inspect twilio
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-twilio/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-twilio/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Twilio Plugin](https://github.com/turbot/steampipe-plugin-twilio/labels/help%20wanted)
