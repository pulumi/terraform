package main

import (
	"github.com/pulumi/terraform/pkg/grpcwrap"
	"github.com/pulumi/terraform/pkg/plugin"
	simple "github.com/pulumi/terraform/pkg/provider-simple"
	"github.com/pulumi/terraform/pkg/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
