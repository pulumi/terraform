package main

import "github.com/pulumi/terraform/pkg/command/workdir"

func WorkingDir(originalDir string, overrideDataDir string) *workdir.Dir {
	ret := workdir.NewDir(".") // caller should already have used os.Chdir in "-chdir=..." mode
	ret.OverrideOriginalWorkingDir(originalDir)
	if overrideDataDir != "" {
		ret.OverrideDataDir(overrideDataDir)
	}
	return ret
}
