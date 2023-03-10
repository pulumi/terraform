package clistate

import (
	"testing"

	"github.com/pulumi/terraform/pkg/command/arguments"
	"github.com/pulumi/terraform/pkg/command/views"
	"github.com/pulumi/terraform/pkg/states/statemgr"
	"github.com/pulumi/terraform/pkg/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
