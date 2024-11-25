package main

import (
	"runtime"

	"github.com/custodia-cenv/cenv-vm/src/filesys"
	"github.com/custodia-cenv/cenv-vm/src/img"
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

	// Das VM Image wird geladen
	err = img.LoadImgFile("test.img")
	if err != nil {
		panic(err)
	}

	// Das geladene Image wird im Core Registriert
	err = vm.InitVmImgManifestWithCore()
	if err != nil {
		panic(err)
	}

	// Der Prozess wird in Nebendiensten Registriert, bsp: Network Plattform
	err = vm.InitAntoherVmSystemServices()
	if err != nil {
		panic(err)
	}

	// Das Dateisystem wird vorbereitet
	err = filesys.InitVMFileSystem()
	if err != nil {
		panic(err)
	}
}
