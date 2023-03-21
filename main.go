package main

import (
	"github.com/albeego/terraform-provider-influxdb-v2-onboarding/influxdbv2_onboarding"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: influxdbv2_onboarding.Provider})
}
