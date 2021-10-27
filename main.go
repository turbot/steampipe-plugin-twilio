package main

import (
	"github.com/turbot/steampipe-plugin-twilio/twilio"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: twilio.Plugin})
}
