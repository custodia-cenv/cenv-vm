package vm

import (
	cenvvm "github.com/custodia-cenv/cenv-vm/src"
)

func StartVMAndKeepAlive() error {
	// Der Status wird auf Serving gesetzt
	coreSetState(cenvvm.SERVING, true)
	defer coreSetState(cenvvm.SHUTDOWN, true)

	return nil
}
