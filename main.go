package main

import (
	"github.com/turbot/steampipe-plugin-hubspot/hubspot"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hubspot.Plugin})
}