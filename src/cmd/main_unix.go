package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/custodia-cenv/cenv-vm/src/host/filesystem"
	"github.com/custodia-cenv/cenv-vm/src/vfimg"
	"github.com/custodia-cenv/cenv-vm/src/vm"
	cenvxcore "github.com/custodia-cenv/cenvx-core/src"
	"github.com/custodia-cenv/cenvx-core/src/cmd"
)

func main() {
	// Set the maximum number of CPU cores for the Go runtime
	runtime.GOMAXPROCS(1)

	// Parse the flags
	flag.Parse()

	// Check if the required -img flag is provided
	if *image == "" {
		fmt.Println("Error: -img is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Display the welcome screen
	cmd.ShowBanner(cenvxcore.VMBanner)

	// Check if the OS is supported
	cmd.OSSupportCheck()

	// Prepare the working directory
	err := vm.SetupWorkingDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Output the parsed flags
	fmt.Printf("Image: %s\n", *image)

	// Load the additional virtual file systems
	for i, pair := range vfsFlags {
		// Es wird geprüft ob die Datei vorhanden ist
		if !filesystem.FileExists(pair.File) {
			fmt.Printf("Error: virtual filesystem image %s not found\n", pair.File)
			os.Exit(1)
		}

		fmt.Printf("VFS %d: Type=%s, File=%s\n", i+1, pair.Type, pair.File)
	}

	// Prepare the VM process
	err = vm.InitVmProcessInstance()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Load the VM image
	err = vfimg.LoadImgFile(*image)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Register the loaded image in the core
	err = vm.InitVmImgManifestWithCore()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Register the process in background services, e.g., Network Platform
	err = vm.InitAntoherVmSystemServices()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Start the VM process and keep it alive
	err = vm.StartVMAndKeepAlive()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Bye
	fmt.Println("Bye")
}
