package main

import "github.com/custodia-cenv/cenv-vm/src/vm"

func main() {
	err := vm.InitVmProcessInstance()
	if err != nil {
		panic(err)
	}
}
