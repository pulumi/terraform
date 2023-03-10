package terraform

import (
	backendInit "github.com/pulumi/terraform/pkg/backend/init"
)

func init() {
	// Initialize the backends
	backendInit.Init(nil)
}
