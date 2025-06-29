// ================================================================
// main.go - Entry point for the Terraform provider
// ================================================================

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"jxtsaminfra_ec2_change_instance_type/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
