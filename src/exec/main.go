package main

import (
	"runtime"

	"github.com/custodia-cenv/cenv-vm/src/vm"
	cenvxcore "github.com/custodia-cenv/cenvx-core/src"
	"github.com/custodia-cenv/cenvx-core/src/cmd"
)

func main() {
	// Maximale Anzahl von CPU-Kernen für die Go-Runtime festlegen
	runtime.GOMAXPROCS(1)

	// Der Willkomensbildschrim wird angezeigt
	cmd.ShowBanner(cenvxcore.VMBanner)

	// Es wird geprüft ob es sich um Unterstützes OS handelt
	cmd.OSSupportCheck()

	// Der VM Prozess wird vorbereitet
	err := vm.InitVmProcessInstance()
	if err != nil {
		panic(err)
	}
}
