package main

import (
	localexec "github.com/pulumi/terraform/pkg/builtin/provisioners/local-exec"
	"github.com/pulumi/terraform/pkg/grpcwrap"
	"github.com/pulumi/terraform/pkg/plugin"
	"github.com/pulumi/terraform/pkg/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
