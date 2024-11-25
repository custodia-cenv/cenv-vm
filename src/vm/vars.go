package vm

import (
	"sync"

	"github.com/custodia-cenv/bngsocket-go"
	cenvvm "github.com/custodia-cenv/cenv-vm/src"
)

var (
	clientIpcVmSokcet *bngsocket.BngConn
	vmMutex           = new(sync.Mutex)
	vmState           cenvvm.VmState
)
