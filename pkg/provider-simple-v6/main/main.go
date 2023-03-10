package main

import (
	"github.com/pulumi/terraform/pkg/grpcwrap"
	plugin "github.com/pulumi/terraform/pkg/plugin6"
	simple "github.com/pulumi/terraform/pkg/provider-simple-v6"
	"github.com/pulumi/terraform/pkg/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
