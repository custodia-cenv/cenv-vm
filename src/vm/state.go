package vm

import cenvvm "github.com/custodia-cenv/cenv-vm/src"

// Legt den Core Status fest
func coreSetState(tstate cenvvm.VmState, useMutex bool) {
	// Es wird geprüft ob Mutex verwendet werden sollen
	if useMutex {
		vmMutex.Lock()
		defer vmMutex.Unlock()
	}

	// Der Aktuelle Status wird geschrieben
	vmState = tstate
}
