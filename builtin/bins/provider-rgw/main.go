package main

import (
	"github.com/hashicorp/terraform/builtin/providers/rgw"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rgw.Provider,
	})
}
