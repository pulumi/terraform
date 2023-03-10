package getproviders

import (
	"context"

	"github.com/pulumi/terraform/pkg/addrs"
)

// A Source can query a particular source for information about providers
// that are available to install.
type Source interface {
	AvailableVersions(ctx context.Context, provider addrs.Provider) (VersionList, Warnings, error)
	PackageMeta(ctx context.Context, provider addrs.Provider, version Version, target Platform) (PackageMeta, error)
	ForDisplay(provider addrs.Provider) string
}
