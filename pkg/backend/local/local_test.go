package local

import (
	"flag"
	"os"
	"testing"

	_ "github.com/pulumi/terraform/pkg/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
