package local

import (
	"sync"

	"github.com/pulumi/terraform/pkg/states"
	"github.com/pulumi/terraform/pkg/states/statemgr"
	"github.com/pulumi/terraform/pkg/terraform"
)

// StateHook is a hook that continuously updates the state by calling
// WriteState on a statemgr.Full.
type StateHook struct {
	terraform.NilHook
	sync.Mutex

	StateMgr statemgr.Writer
}

var _ terraform.Hook = (*StateHook)(nil)

func (h *StateHook) PostStateUpdate(new *states.State) (terraform.HookAction, error) {
	h.Lock()
	defer h.Unlock()

	if h.StateMgr != nil {
		if err := h.StateMgr.WriteState(new); err != nil {
			return terraform.HookActionHalt, err
		}
	}

	return terraform.HookActionContinue, nil
}
