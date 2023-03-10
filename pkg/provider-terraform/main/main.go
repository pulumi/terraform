package main

import (
	"github.com/pulumi/terraform/pkg/builtin/providers/terraform"
	"github.com/pulumi/terraform/pkg/grpcwrap"
	"github.com/pulumi/terraform/pkg/plugin"
	"github.com/pulumi/terraform/pkg/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
