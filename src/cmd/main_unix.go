package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	cenvvm "github.com/custodia-cenv/cenv-vm/src"
	"github.com/custodia-cenv/cenv-vm/src/filesys"
	"github.com/custodia-cenv/cenv-vm/src/host/filesystem"
	"github.com/custodia-cenv/cenv-vm/src/img"
	"github.com/custodia-cenv/cenv-vm/src/vm"
	cenvxcore "github.com/custodia-cenv/cenvx-core/src"
	"github.com/custodia-cenv/cenvx-core/src/cmd"
)

func main() {
	// Set maximum number of CPU cores for the Go runtime
	runtime.GOMAXPROCS(1)

	// Define flags
	image := flag.String("img", "", "Path to the image file (required)")
	workingDir := flag.String("workingdir", "", "Path to the working directory")

	// Parse the flags
	flag.Parse()

	// Check if the required flag is provided
	if *image == "" {
		fmt.Println("Error: -img is required.")
		flag.Usage()
		os.Exit(1)
	}

	// It is checked whether the WorkingDir exists
	var virtualFileSystemPath string
	var virtualFileSystemMethode cenvvm.VFileSystemMethode
	if *workingDir != "" {
		// Es wird geprüft ob es sich um einen Zulässigen Path handelt
		if filesystem.HasUserAccess(*workingDir) {

		}

	} else {
		fmt.Println("Error: -vfs is required.")
		flag.Usage()
		os.Exit(1)
	}

	// The welcome screen is displayed
	cmd.ShowBanner(cenvxcore.VMBanner)

	// It is checked whether it is a supported OS
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
	err = filesys.InitVMFileSystem(virtualFileSystemPath, virtualFileSystemMethode)
	if err != nil {
		panic(err)
	}

	// Der VM Porzess wird gestartet und am leben gehalten
	err = vm.StartVMAndKeepAlive()
	if err != nil {
		panic(err)
	}

	// Bye
	fmt.Println("By")
}
