package main

import (
	"github.com/edify42/terraform-provider-stackjanitor/stackjanitor"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: stackjanitor.Provider})
}
